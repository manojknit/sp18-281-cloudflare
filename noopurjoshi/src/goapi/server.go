package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"github.com/satori/go.uuid"
    "gopkg.in/mgo.v2/bson"
	"github.com/go-redis/redis"
	"hash/fnv"
	"os"
)

// MongoDB Config
var mongodb_server string
// = "mongodb://54.193.41.243:27017,52.53.151.228:27017,54.219.131.204:27017,54.67.31.38:27017,54.215.243.94:27017"
var mongodb_server1 string
//= "mongodb://54.183.14.90,52.53.150.53,54.153.37.72,54.67.37.239,54.67.123.195"
var mongodb_server2 string
//"mongodb://54.67.13.87:27017,54.67.106.101:27017,13.57.39.192:27017,54.153.26.217:27017,52.53.154.42:27017"
var mongodb_database string
//= "admin"
var mongodb_collection string
//= "starbucks"
var redis_server string
//= "192.168.99.100:6379"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	mongodb_server = os.Getenv("MONGO1")
	mongodb_server1 = os.Getenv("MONGO2")
	mongodb_server2 = os.Getenv("MONGO3")
	mongodb_database = os.Getenv("MONGO_DB")
	mongodb_collection = os.Getenv("MONGO_COLLECTION")
	redis_server = os.Getenv("REDIS")
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order/{id}", getCartDetails(formatter)).Methods("GET")
	mx.HandleFunc("/order/{id}", updateOrderInCart(formatter)).Methods("PUT")
	mx.HandleFunc("/order", addOrderToCart(formatter)).Methods("POST")
	mx.HandleFunc("/history/{id}", orderHistory(formatter)).Methods("GET")
}

// sharding
func hash(s string) string {
        h := fnv.New32a()
        h.Write([]byte(s))
		node := h.Sum32()%3
		if node == 0 {
			return mongodb_server
		} else if node == 1 {
			return mongodb_server1
		} else if node == 2 {
			return mongodb_server2
		} else {
			return mongodb_server
		}
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//Redis connection
var client *redis.Client


func get_client()(*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     redis_server,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

    pong, err := client.Ping().Result()
	fmt.Println("ponging",pong,err)
	return client
	
}

 //get key from cache
func red_getHandler(uid string)(bool){
	client:=get_client()	
	_, err := client.Get(uid).Result()
	fmt.Println(uid,"my uid")	
	if err == redis.Nil {
		return false
	}else if err != nil{
		panic(err)
	}else{
		return true
	}		
}

// set key in cache
	func red_setHandler(formatter *render.Render) http.HandlerFunc {  
	return func(w http.ResponseWriter, req *http.Request) {
	var m user
	_ = json.NewDecoder(req.Body).Decode(&m)
	client:=get_client()
	err := client.Set(m.UserId, "dhan_value", 0).Err()
	if err != nil {
		panic(err)
	}
	formatter.JSON(w, http.StatusOK,struct{ Test string }{"Value pushed"})
	}
	}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API  Handler --------------- GET ------------------
func getCartDetails(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var result bson.M
		vars:=mux.Vars(req)
		id := vars["id"]
		check := red_getHandler(id)
		if check {			
			session, err := mgo.Dial(hash(id))
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        
        err = c.Find( bson.M{ "$and": []bson.M{ bson.M{"UserId":id}, bson.M{"Status": "PENDING"} } } ).One(&result)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, result)
		}
	}
}

func orderHistory(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var result []bson.M
		vars:=mux.Vars(req)
		id := vars["id"]
		check := red_getHandler(id)
		if check {			
			session, err := mgo.Dial(hash(id))
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        
        err = c.Find( bson.M{ "$and": []bson.M{ bson.M{"UserId":id}, bson.M{"Status": "PROCESSED"} } } ).All(&result)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, result)
		}
	}
}

// API Update Inventory ----------- PUT --------------
func updateOrderInCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
    	var m orderDetails
    	_ = json.NewDecoder(req.Body).Decode(&m)
		vars:=mux.Vars(req)
		id := vars["id"]
    	fmt.Println("Edit Order")
		check := red_getHandler(m.UserId)		
		if(check){			
			session, err := mgo.Dial(hash(m.UserId))
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
		err = c.Update(bson.M{"UserId": id}, bson.M{"$pull": bson.M{"Orders": bson.M{"OrderId":m.OrderId}}})
		var resp bson.M
        err = c.Find( bson.M{ "$and": []bson.M{ bson.M{"UserId":id}, bson.M{"Status": "PENDING"} } } ).One(&resp)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, resp)
		}
	}
}

// Create
func addOrderToCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var m orderDetails
		var result cart
    	_ = json.NewDecoder(req.Body).Decode(&m)		
    	fmt.Println("Add to cart")
		check := red_getHandler(m.UserId)		
		if check {			
			session, err := mgo.Dial(hash(m.UserId))
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
		err = c.Find( bson.M{ "$and": []bson.M{ bson.M{"UserId":m.UserId}, bson.M{"Status": "PENDING"} } } ).One(&result)
		id := uuid.NewV4()
		fmt.Println("result", result)
		if result.CartId != "" {
			m.OrderId = uuid.NewV4().String()
			err = c.Update(bson.M{"CartId": result.CartId}, bson.M{"$push": bson.M{"Orders": m}})
			if err != nil {
                log.Fatal(err)
			}
		} else {
			m.OrderId = uuid.NewV4().String()
			query := bson.M{"CartId": id.String(), "Orders" :[]orderDetails{m} , "Status":"PENDING", "UserId":m.UserId}
			err = c.Insert(query)
			if err != nil {
                log.Fatal(err)
			}
		}
       	var resp bson.M
        err = c.Find( bson.M{ "$and": []bson.M{ bson.M{"UserId":m.UserId}, bson.M{"Status": "PENDING"} } } ).One(&resp)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, resp)
		}
	}
}
