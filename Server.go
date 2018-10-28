package main

import (
	"fmt"
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
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/things/{name}", route2).Methods("GET")
	router.HandleFunc("/register", register).Methods("POST")

	//serve public file
	router.PathPrefix("/").Handler(fs)

	http.ListenAndServe(":80", router)
}

//Home handle request
func home(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Fprintf(w, "requested %s", vals)
}

func register(w http.ResponseWriter, r *http.Request) {

	details := contacts{
		email:  r.FormValue("email"),
		name:   r.FormValue("name"),
		number: r.FormValue("number"),
	}

	fmt.Fprintf(w, "found %s", &details)
}
func route2(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Fprintf(w, "found %s", vals["name"])
}
