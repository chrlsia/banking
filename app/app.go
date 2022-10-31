package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(){

	// mux:=http.NewServeMux()
	router:=mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)


	// start the server
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r) // creates a map
	fmt.Fprint(w,vars["customer_id"])
}
