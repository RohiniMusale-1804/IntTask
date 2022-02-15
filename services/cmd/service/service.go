package cmd

import (
	"fmt"
	db "intTask/db"
	endpoints "intTask/services/pkg/endpoints"
	handler "intTask/services/pkg/handlers"
	repository "intTask/services/pkg/repositories"
	service "intTask/services/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	conn := db.NewDbConnection()
	cityRepository := repository.NewCityRepositoryImpl(conn)
	cityService := service.NewCityServiceImpl(cityRepository)
	cityHandler := handler.NewCityHandlerImpl(cityService)

	endpoints.NewCityRoute(r, cityHandler)
	//The approach can be converted into more asynch log using zerolog package.
	fmt.Println("Apllication Running!")
	fmt.Println(http.ListenAndServe(":9001", r))
}
