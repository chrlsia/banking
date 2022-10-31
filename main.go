package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct{
	Name string
	City string
	Zipcode string
}

func main(){
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe(":3000",nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request ){
	customers:=[]Customer{
		{Name:"Ashish",City: "New Delhi", Zipcode:"110075"},
		{Name:"Rod",City: "New Delhi", Zipcode:"110075"},
	}
	json.NewEncoder(w).Encode(customers)
}