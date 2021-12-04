package app

import (
	"net/http"

	"github.com/SaravananPitchaimuthu/banking/Banking/domain"
	"github.com/SaravananPitchaimuthu/banking/Banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	// create router
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// starting server
	http.ListenAndServe("localhost:8000", router)
}
