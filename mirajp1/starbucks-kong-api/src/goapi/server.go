package main

import (
	"fmt"
	"log"
	"net/http"
    "io/ioutil"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
    "github.com/gorilla/schema"
	"github.com/unrolled/render"
    "encoding/json"
)

var decoder = schema.NewDecoder()


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
	mx.HandleFunc("/upstream", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api", createNewKongApiHandler(formatter)).Methods("POST")
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        resp, err := http.Get("http://kong-aws:8001" )
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

func getUpstreamURLHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"URL"})
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
        
	    fmt.Printf("[Kong DEBUG] POST: " + "kong-aws:8001/ => %+v", kongApi) 
		formatter.JSON(w, http.StatusOK , kongApi)
	}
}
