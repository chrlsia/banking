package main

import (
	"fmt"
	"net/http"
)

type Customer struct{
	Name,City,Zipcode string
}
func main(){
	http.HandleFunc("/greet", greet)

	http.ListenAndServe(":3000",nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello World!")
}