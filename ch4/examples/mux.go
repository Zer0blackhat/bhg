package main

import (
	"github.com/gorilla/mux"
)

r := mux.NewRouter()

// New route to handle GET requests
f.HandleFun("/foo", func(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hi foo")
}).Methods("GET").Host("www.foo.com")

// Match and pass in parameters
r.HandleFunc("/users/{user}", func(w http.ResponseWriter, req *http.Request){
	user := mux.Vars(req)["user"]
	fmt.Fprintf(w, "hi %s\n", user)
}).Methods("GET")

// Match with Regex 
r.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request){
	user := mux.Vars(req)["user"]
	fmt.Fprintf(w, "hi %s\n", user)
}).Methods("GET")parameters
