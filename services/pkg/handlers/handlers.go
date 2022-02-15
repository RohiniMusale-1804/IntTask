package handlers

import (
	"encoding/json"
	"net/http"

	service "intTask/services/pkg/service"

	"github.com/gorilla/mux"
)

//CityHandlersImpl for handler Functions
type CityHandlerStruct struct {
	newService service.CityService
}

//NewHandlerImpl inits dependancies for Handlers
func NewCityHandlerImpl(service service.CityService) *CityHandlerStruct {
	return &CityHandlerStruct{newService: service}
}

//CreateCity handler
func (cityHandlers CityHandlerStruct) CreateCity(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	go cityHandlers.newService.CreateCity(ctx)

	resp := "Running goroutines in background."

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (cityHandlers CityHandlerStruct) GetCityDetails(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(req)
	cityName := vars["city"]
	resp, err := cityHandlers.newService.GetCity(ctx, cityName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
