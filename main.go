package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"first_name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe("localhost:8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello world !!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Saravanan", "Trichy", "620106"},
		{"Saranya", "Trichy", "620106"},
		{"Saraswathi", "Trichy", "620106"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
