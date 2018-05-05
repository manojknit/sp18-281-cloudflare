/*
	Starbucks API in Go (Version 3)
	Uses MongoDB and Redis 
	(For use with Kong API Key)
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"github.com/go-redis/redis"
	"os"
	"hash/fnv"
)

// MongoDB Config
var mongodb_server string
var mongodb_server1 string
var mongodb_server2 string
var redis_server string
//="mongodb://54.67.13.87:27017,54.67.106.101:27017,13.57.39.192:27017,54.153.26.217://27017,52.53.154.42:27017"

var mongodb_database string
var mongodb_collection string



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
	mx.HandleFunc("/starbucks/{id}", starbucksHandler(formatter)).Methods("GET")
	mx.HandleFunc("/starbucks", starbucksUpdateHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/starbucks", starbucksNewOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/starbucks/{id}", starbucksDeleteHandler(formatter)).Methods("DELETE")
	//api route to gopayment handler
    //mx.HandleFunc("/redisget/{key}",
	//red_getHandler(formatter)).Methods("GET")
	mx.HandleFunc("/redisSet",
	red_setHandler(formatter)).Methods("POST")
    mx.HandleFunc("/checkoutCart",
	checkoutHandler(formatter)).Methods("POST")
	
}

// sharding function
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


//Redis connection
var client *redis.Client


func get_client()(*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:  redis_server,
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
	
	 _,err := client.Get(uid).Result()
	
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
	var m gumballMachine
	_ = json.NewDecoder(req.Body).Decode(&m)
	client:=get_client()
	err := client.Set(m.UserId, "my_value", 0).Err()
	if err != nil {
		panic(err)
	}
	formatter.JSON(w, http.StatusOK,struct{ Test string }{"Value pushed"})
	}
	}

	
	//CheckoutHandler
	 
	 
	 func checkoutHandler(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request) {
		uuid := uuid.NewV4()
		var payment paymentDetails
		var response string
		
		_ = json.NewDecoder(req.Body).Decode(&payment)		
    	fmt.Println("Insert Payment details in db ")
				
		
		check := red_getHandler(payment.UserId)
		
		if(check){
			
			session, err := mgo.Dial(hash(payment.UserId))
        if err != nil {
                panic(err)
        }
		defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
		
		//Push payment details
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"PaymentId":uuid.String(),"UserId" : payment.UserId, "PaymentType" : payment.PaymentType, "TotalPrice" : payment.TotalPrice, "CartId" : payment.CartId}
        err = c.Insert(query)
		
		//Update Order Status
		err = c.Update(bson.M{"UserId": payment.UserId}, bson.M{"$set": bson.M{"Status": "PROCESSED"}})
		var resp bson.M
        err = c.Find( bson.M{ "$and": []bson.M{ bson.M{"UserId":payment.UserId}, bson.M{"Status": "PROCESSED"} } } ).One(&resp)
		
        if err != nil {
                log.Fatal(err)
        }
		   response = "Your order is processed successfully"
		}else{
		   response= "Session invalid"
		}
				
		formatter.JSON(w, http.StatusOK,struct{ Test string }{response})
		
		}
	 }
// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}



// API  Handler --------------- GET ------------------
func starbucksHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	    
		//var m test
		vars:=mux.Vars(req)
		id,err1 := strconv.Atoi(vars["id"])
	
		if err1 != nil {
			fmt.Println(err1)
		}
				
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
		
		
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"Id" : id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Result :", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Update Inventory ----------- PUT --------------
func starbucksUpdateHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
    	var m test
		
  	_ = json.NewDecoder(req.Body).Decode(&m)		
    	fmt.Println("Update Inventory To: ", m.Count)
		
			
		
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
		
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"Id" : m.Id}
        change := bson.M{"$set": bson.M{ "Count" : m.Count, "SerialNumber":m.SerialNumber,"ModelNumber":m.ModelNumber}}
        err = c.Update(query, change)
        if err != nil {
                log.Fatal(err)
        }
       	var result bson.M
        err = c.Find(bson.M{"Id" : m.Id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }        
        fmt.Println("Result:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}


// --------------------- POST ----------------------------
func starbucksNewOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
		var m test
    	_ = json.NewDecoder(req.Body).Decode(&m)		
    	fmt.Println("Insert into Inventory ")
				
		session, err := mgo.Dial(hash(m.UserId))
        if err != nil {
                panic(err)
        }
		
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"Id" : m.Id, "Count" : m.Count, "ModelNumber" : m.ModelNumber, "SerialNumber" : m.SerialNumber}
        err = c.Insert(query)
        if err != nil {
                log.Fatal(err)
        }
       	var result bson.M
        err = c.Find(bson.M{"Id" : m.Id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, result)
	}
}

// ------------------ Delete ---------------------
func starbucksDeleteHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	
		var m test
    	_ = json.NewDecoder(req.Body).Decode(&m)
		vars := mux.Vars(req)
		id, err1 := strconv.Atoi(vars["id"])
		if err1 != nil {
			fmt.Println(err1)
		}
				
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
		
		fmt.Println("User id:", id)
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Remove(bson.M{"Id" : id})
        if err != nil {
				panic(err)
                log.Fatal(err)
        }
        fmt.Println("Result:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}
