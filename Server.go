package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))

	//handle home route
	router.HandleFunc("/", home)

	//serve public file
	router.PathPrefix("/").Handler(fs)

	http.ListenAndServe(":80", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Fprintf(w, "requested %s", vals)
}
