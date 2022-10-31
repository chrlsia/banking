package app

import (
	"net/http"
)

func Start(){
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe(":3000", nil)
}
