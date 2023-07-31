package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type badAuth struct { // Simulate authentication for demonstration purposes only
	Username string
	Password string
}

func (b *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) { // 3 parameters for Negroni.Handler
	username := r.URL.Query().Get("username") // Grab username and password from request
	password := r.URL.Query().Get("password")
	if username != b.Username || password != b.Password {
		http.Error(w, "Unauthorized", 401)
		return // Put before next to prevent execution of middleware chain
	}
	ctx := context.WithValue(r.Context(), "username", username) // Initialize from request, common use in web applications
	r = r.WithContext(ctx) // Use new context
	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string) // Use switch routine if you can't ensure the type
	fmt.Fprintf(w, "Hi %s\n", username)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")
	n := negroni.Classic()
	n.Use(&badAuth{
		Username:"admin",
		Password:"password",
	})
	n.UseHandler(r)
	http.ListenAndServe(":8000", n)
}