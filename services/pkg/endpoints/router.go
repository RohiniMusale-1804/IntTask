package endpoints

import (
	handler "intTask/services/pkg/handlers"

	"github.com/gorilla/mux"
)

//NewCityRoute Collection of all routes
func NewCityRoute(router *mux.Router, handler *handler.CityHandlerStruct) {
	router.HandleFunc("/cities", handler.CreateCity).Methods("GET")
	router.HandleFunc("/cities/{city}", handler.GetCityDetails).Methods("GET")
}
