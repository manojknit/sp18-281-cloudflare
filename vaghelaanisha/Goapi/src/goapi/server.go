/*
	 API in Go (Version 3)
	Uses MongoDB 
	(For use with Kong API Key)
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	//"strconv"
	//"goji.io/pat"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"github.com/go-redis/redis"
)

// MongoDB Config
var mongodb_server = "mongodb://54.183.14.90,52.53.150.53,54.153.37.72,54.67.37.239,54.67.123.195"
var mongodb_database = "admin"
var mongodb_collection = "starbucks"



// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/login", login(formatter)).Methods("POST")
	mx.HandleFunc("/signup", signup(formatter)).Methods("POST")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// --------------------- POST ----------------------------
func login(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var m user
    	_ = json.NewDecoder(req.Body).Decode(&m)		
		session, err := mgo.Dial(mongodb_server)
        fmt.Println(m)
		if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
       	var result user
        err = c.Find(bson.M{"Username" : m.Username}).One(&result)
		fmt.Println(err)
        if err != nil {
			//log.Fatal(err)
			formatter.JSON(w, 404, "Username not registered")
			return
        }

		if result.Password != m.Password {
			formatter.JSON(w, 404, "Password is incorrect")
			return
		}
		
		check:=red_setHandler(result.UserId);
		if(check){
			formatter.JSON(w, http.StatusOK, result)
		}else{
			formatter.JSON(w, 500, "Cannot connect to redis")
		}
	}
}

func signup(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var m user
		uuid,_ := uuid.NewV4()
    	_ = json.NewDecoder(req.Body).Decode(&m)
    	fmt.Println("",m)

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.PrimaryPreferred, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"UserId" : uuid.String(), "Username" : m.Username, "Password" : m.Password}
        err = c.Insert(query)
		fmt.Println(m.Username)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, "User registered Successfully")
		return
			
	}
}

func get_client()(*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:6397",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

    pong, err := client.Ping().Result()
	fmt.Println("ponging",pong,err)
	return client
	
}

 //get key from cache
func red_setHandler(uid string)(bool){
	
	client:=get_client()
	err := client.Set(uid, "UserId", 0).Err()
	if err != nil {
		panic(err)
	}
	_, err = client.Get(uid).Result()
	
	fmt.Println(uid,"my uid")
	return true
		
}




