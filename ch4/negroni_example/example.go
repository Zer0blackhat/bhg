package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()
	n.UseHandler(r)
	http.ListenAndServe(":8000", n)
}

// Middleware example

type trivial struct {
}
func (t *trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Executing trivial middleware...")
	next(w, r)
}

// Tell negroni to use implementation in middleware chain
n.Use(&trivial{})

// Other ways to use middleware
UseHandler(handler http.Handler)
UseHandlerFunc(handlerFunc func(w http.responseWriter, r *http.Request)) // Not something to use often as you can't stop next middleware from executing (such as authentication)

