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
	"strconv"
	//"goji.io/pat"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "mongodb://54.193.23.61:27017/"
var mongodb_database = "admin"
var mongodb_collection = "admin"



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
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/gumball/{id}", gumballHandler(formatter)).Methods("GET")
	mx.HandleFunc("/gumball", gumballUpdateHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/gumball", gumballNewOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/gumball/{id}", gumballDeleteHandler(formatter)).Methods("DELETE")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API  Handler --------------- GET ------------------
func gumballHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	    //var m gumballMachine
		
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
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"Id" : id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Gumball Machine:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Update Inventory ----------- PUT --------------
func gumballUpdateHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
    	var m gumballMachine
    	_ = json.NewDecoder(req.Body).Decode(&m)		
    	fmt.Println("Update Gumball Inventory To: ", m.CountGumballs)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"Id" : m.Id}
        change := bson.M{"$set": bson.M{ "CountGumballs" : m.CountGumballs, "SerialNumber":m.SerialNumber,"ModelNumber":m.ModelNumber}}
        err = c.Update(query, change)
        if err != nil {
                log.Fatal(err)
        }
       	var result bson.M
        err = c.Find(bson.M{"Id" : m.Id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }        
        fmt.Println("Gumball Machine:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

// --------------------- POST ----------------------------
func gumballNewOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var m gumballMachine
    	_ = json.NewDecoder(req.Body).Decode(&m)		
    	fmt.Println("Insert into Inventory ")
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"Id" : m.Id, "CountGumballs" : m.CountGumballs, "ModelNumber" : m.ModelNumber, "SerialNumber" : m.SerialNumber}
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
func gumballDeleteHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var m gumballMachine
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
		fmt.Println("Gumball Machine id:", id)
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Remove(bson.M{"Id" : id})
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Gumball Machine:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}


/*

	-- RabbitMQ Setup

	http://localhost:8080

	-- RabbitMQ Create Queue:  

		Queue Name: gumball
		Durable:	no

	-- Gumball MongoDB Create Database

		Database Name: cmpe281
		Collection Name: gumball

  	-- Gumball MongoDB Collection (Create Document) --

    db.gumball.insert(
	    { 
	      Id: 1,
	      CountGumballs: NumberInt(202),
	      ModelNumber: 'M102988',
	      SerialNumber: '1234998871109' 
	    }
	) ;

    -- Gumball MongoDB Collection - Find Gumball Document --

    db.gumball.find( { Id: 1 } ) ;

    {
        "_id" : ObjectId("54741c01fa0bd1f1cdf71312"),
        "Id" : 1,
        "CountGumballs" : 202,
        "ModelNumber" : "M102988",
        "SerialNumber" : "1234998871109"
    }

    -- Gumball MongoDB Collection - Update Gumball Document --

    db.gumball.update( 
        { Dd: 1 }, 
        { $set : { CountGumballs : NumberInt(10) } },
        { multi : false } 
    )

    -- Gumball Delete Documents

    db.gumball.remove({})

 */
