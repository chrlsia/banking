package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start(){

	// mux:=http.NewServeMux()
	mux:=mux.NewRouter()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// start the server
	log.Fatal(http.ListenAndServe(":3000", mux))
}
