package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string	`json:"full_name" xml:"name"`
	City    string	`json:"city" xml:"city"`
	Zipcode string	`json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashish", City: "New Delhi", Zipcode: "110075"},
		{Name: "Rod", City: "New Delhi", Zipcode: "110075"},
	}
	if r.Header.Get("Content-Type")=="application/xml"{
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	
}

func getCustomer(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r) // creates a map
	fmt.Fprint(w,vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Post request received")
}
