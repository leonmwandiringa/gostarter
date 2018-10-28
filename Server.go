package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type contacts struct {
	email  string
	name   string
	number string
}

func main() {

	router := mux.NewRouter()
	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))

	//handle home route
	router.HandleFunc("/", loggingMiddleware(home)).Methods("GET")
	router.HandleFunc("/things/{name}", loggingMiddleware(route2)).Methods("GET")
	router.HandleFunc("/register", loggingMiddleware(register)).Methods("POST")

	//serve public file
	router.PathPrefix("/").Handler(fs)

	http.ListenAndServe(":80", router)
}

//Home handle request
func home(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Fprintf(w, "requested %s", vals)
}

//logger middleware
func loggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		f(w, r)
	}
}

func register(w http.ResponseWriter, r *http.Request) {

	details := contacts{
		email:  r.FormValue("email"),
		name:   r.FormValue("name"),
		number: r.FormValue("number"),
	}

	fmt.Fprintf(w, "found %s", details)
}
func route2(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Fprintf(w, "found %s", vals["name"])
}
