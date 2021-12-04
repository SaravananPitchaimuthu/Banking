package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/SaravananPitchaimuthu/banking/Banking/service"
)

type Customer struct {
	Name    string `json:"first_name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Saravanan", "Trichy", "620106"},
	// 	{"Saranya", "Trichy", "620106"},
	// 	{"Saraswathi", "Trichy", "620106"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
