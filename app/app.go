package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start(){

	// mux:=http.NewServeMux()
	router:=mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	// start the server
	log.Fatal(http.ListenAndServe(":3000", router))
}
