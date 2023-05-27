package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {
	fmt.Println("Starting server on port 8000")
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	http.ListenAndServe(":8000", &l)
}