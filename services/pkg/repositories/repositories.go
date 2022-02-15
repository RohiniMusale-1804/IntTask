package repositories

import (
	"context"
	"intTask/services/pkg/models"

	"github.com/jinzhu/gorm"
)

//Repository implimets all methods in Repository
type CityRepository interface {
	CreateCity(context.Context, models.City) error
	GetCity(context.Context, string) (*models.City, error)
}

//CityRepositoryImpl **
type CityRepositoryImpl struct {
	dbConn *gorm.DB
}

func NewCityRepositoryImpl(dbConn *gorm.DB) CityRepository {
	return &CityRepositoryImpl{dbConn: dbConn}
}

func (cityRepositoryImpl CityRepositoryImpl) CreateCity(ctx context.Context, city models.City) (err error) {
	dbConn := cityRepositoryImpl.dbConn
	err = dbConn.Table("cities").Create(&city).Error
	if err != nil {
		return
	}
	return nil
}

func (cityRepositoryImpl CityRepositoryImpl) GetCity(ctx context.Context, cityName string) (*models.City, error) {
	dbConn := cityRepositoryImpl.dbConn
	city := models.City{}
	err := dbConn.Where("name=?", cityName).First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}
