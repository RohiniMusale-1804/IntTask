package repositories_test

import (
	"context"
	"intTask/mocks"
	"intTask/services/pkg/models"
	service "intTask/services/pkg/service"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCityRepositoryImpl_GetCity(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCityRepository(ctrl)
	service := service.NewCityServiceImpl(mockRepo)

	type args struct {
		ctx      context.Context
		cityName string
	}
	tests := []struct {
		name    string
		args    args
		fn      func()
		want    *models.City
		wantErr bool
	}{
		{
			"Test 1 :Get cities",
			args{
				context.Background(),
				"Agra",
			},
			func() {
				mockRepo.EXPECT().GetCity(gomock.Any(), gomock.Any()).Return(&models.City{}, nil)
			},
			&models.City{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			_, err := service.GetCity(tt.args.ctx, tt.args.cityName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
