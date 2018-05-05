package main

import (
	"fmt"
	"log"
	"net/http"
    "io/ioutil"
    "bytes"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
    "github.com/gorilla/schema"
	"github.com/unrolled/render"
    "encoding/json"
    "os"
)

var decoder = schema.NewDecoder()
var kong_server string = os.Getenv("KONG_SERVER")

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

//apis for setting up kong
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/apis", getKongApisHandler(formatter)).Methods("GET")
	mx.HandleFunc("/apis", createNewKongApiHandler(formatter)).Methods("POST")
    mx.HandleFunc("/apis", createNewKongApiHandler(formatter)).Methods("DELETE")
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        resp, err := http.Get("http://"+kong_server+":8001" )
	    if err != nil {
		    fmt.Println("[Kong DEBUG] " + err.Error())
		    return
	    }
	    defer resp.Body.Close()
	    body, err := ioutil.ReadAll(resp.Body)
	    fmt.Println("[Kong DEBUG] GET: " + "kong-aws:8001/ => " + string(body)) 
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(body)
		//formatter.JSON(w, http.StatusOK, body)
	}
}

func getKongApisHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        resp, err := http.Get("http://"+kong_server+":8001/apis" )
	    if err != nil {
		    fmt.Println("[Kong DEBUG] " + err.Error())
		    return
	    }
	    defer resp.Body.Close()
	    body, err := ioutil.ReadAll(resp.Body)
	    fmt.Println("[Kong DEBUG] GET: " + "kong-aws:8001/apis => " + string(body)) 
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(body)
		//formatter.JSON(w, http.StatusOK, body)
	}
}

func createNewKongApiHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        err := req.ParseForm()
        if err != nil {
    		formatter.JSON(w, http.StatusBadRequest,"wrong data")
            return
        }

        var kongApi KongApi

        w.Header().Set("Content-Type", "application/json") 
        body, _ := ioutil.ReadAll(req.Body)
        json.Unmarshal(body, &kongApi)
        
	    fmt.Printf("[Kong DEBUG] POST: " + "kong-aws:8001/apis => %+v", kongApi)
        //jsonValue, _ := json.Marshal(values)

        resp, err := http.Post("http://"+kong_server+":8001/apis", 
					        "application/json",  bytes.NewBuffer(body))
        if err != nil {
	        formatter.JSON(w, http.StatusBadRequest,"wrong data")
        }
        defer resp.Body.Close()
        body1, err := ioutil.ReadAll(resp.Body) 
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(body1)
		//formatter.JSON(w, http.StatusOK , kongApi)
	}
}
