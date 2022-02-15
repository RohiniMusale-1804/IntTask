package services

import (
	"context"
	"encoding/json"
	"fmt"
	"intTask/services/pkg/models"
	repository "intTask/services/pkg/repositories"
	"net/http"
	"sync"
)

// Service describes the service.
type CityService interface {
	CreateCity(context.Context) error
	GetCity(context.Context, string) (*models.City, error)
}

//ServiceImpl **
type CityServiceImpl struct {
	cityRepo repository.CityRepository
}

//NewCityServiceImpl inject depedancies user repositiory
func NewCityServiceImpl(cityRepo repository.CityRepository) CityService {
	return &CityServiceImpl{cityRepo: cityRepo}
}

const (
	APIID = "13c16f611f8c8e675230683764aaf06b"
)

func (cityServiceImpl CityServiceImpl) CreateCity(ctx context.Context) error {

	var wg sync.WaitGroup
	var workers = len(Cities)

	wg.Add(10)
	citiesCh := make(chan string)

	for i := 0; i < workers; i++ {
		go cityServiceImpl.FetchCityData(ctx, citiesCh, &wg) //cityData
	}

	for _, citi := range Cities {
		citiesCh <- citi

	}

	wg.Wait()
	close(citiesCh)

	return nil
}

//FetchCityData ##
func (cityServiceImpl CityServiceImpl) FetchCityData(ctx context.Context,
	cityCh <-chan string, wg *sync.WaitGroup) {

	for {

		cityname, ok := <-cityCh
		if !ok {
			wg.Done()

			return
		}
		//ToDo: This part can be part of env variable
		url := fmt.Sprintf("https://pro.openweathermap.org/data/2.5/forecast/climate?q=%s&appid=%s",
			cityname,
			APIID,
		)
		data, err := http.Get(url)
		if err != nil {

			return
		}

		weatherDesc := models.WeatherDesc{}
		jsonDecoder := json.NewDecoder(data.Body)
		err = jsonDecoder.Decode(&weatherDesc)
		if err != nil {

			return
		}

		cityDetails := models.City{}
		cityDetails.Name = cityname
		cityDetails.Lat = weatherDesc.Coord.Lat
		cityDetails.Lon = weatherDesc.Coord.Lon
		cityDetails.Temp = weatherDesc.Main.Temp
		cityDetails.Pressure = weatherDesc.Main.Pressure
		cityDetails.SeaLevel = weatherDesc.Main.SeaLevel

		err = cityServiceImpl.cityRepo.CreateCity(ctx, cityDetails)
		if err != nil {

			return
		}
	}
}

func (cityServiceImpl CityServiceImpl) GetCity(ctx context.Context, cityName string) (resp *models.City, err error) {

	resp, err = cityServiceImpl.cityRepo.GetCity(ctx, cityName)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
