package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var redis_connect = "192.168.99.100:6379"
var mongodb_server1 = "192.168.99.100:27017"
var mongodb_server2 = "192.168.99.100:27018"
var mongodb_server3 = "192.168.99.100:27019"
var mongodb_database = "cmpe281"
var mongodb_collection = "redistest"

var servers = []string{mongodb_server1, mongodb_server2, mongodb_server3}

// Configure and return a server
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
  //define routine here
	
	mx.HandleFunc("/getid/{item_id}", getHandler(formatter)).Methods("GET")
	mx.HandleFunc("/postitem", postHandler(formatter)).Methods("POST")
}


// Helper Functions

func getFromMongo(session *mgo.Session, serialNumber string) User {

	var result User
	//get from mongo
	if session != nil {
		c := session.DB(mongodb_database).C(mongodb_collection)
		err := c.Find(bson.M{"serialnumber": serialNumber}).One(&result)
		if err != nil {
			//could not find in mongo (inserting into mongo for now. TODO: Make proper)
			// c.Insert(bson.M{"SerialNumber": "1", "Name": "Test"})
			fmt.Println("Some Error in Get, maybe data is not present")
		}
	}
	return result

}

func getSession(mongodb_bal_server string) *mgo.Session {
	// Connect to mongo cluster
	//mongodb_bal_server := Balance()
	fmt.Println("mongo connecting to " + mongodb_bal_server)
	s, err := mgo.Dial(mongodb_bal_server)

	if err == nil {
		s.SetMode(mgo.Monotonic, true)
	}

	return s
}


