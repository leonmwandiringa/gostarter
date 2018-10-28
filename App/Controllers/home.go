package home

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Home handle request
func Home(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	fmt.Fprintf(w, "requested %s", vals)
}
