package handlers_test

import (
	"bytes"
	"errors"
	"intTask/mocks"
	"intTask/services/pkg/handlers"
	service "intTask/services/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCityHandlerStruct_GetCityDetails(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockCityService(ctrl)
	cityHttpHandlers := handlers.NewCityHandlerImpl(mockService)

	type fields struct {
		newService service.CityService
	}
	type args struct {
		Method string
		URL    string
		Body   []byte
	}
	tests := []struct {
		name             string
		cityHttpHandlers *handlers.CityHandlerStruct
		args             args
		fn               func()
		wantStatus       int
	}{
		{
			"Test 1 : Status Sucess",
			cityHttpHandlers,
			args{
				"GET",
				"cities/agra",
				nil,
			},
			func() {
				mockService.EXPECT().GetCity(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
			},
			200,
		},
		{
			"Test 2 : Status Not Found",
			cityHttpHandlers,
			args{
				"GET",
				"cities/nasik",
				nil,
			},
			func() {
				mockService.EXPECT().GetCity(gomock.Any(), gomock.Any()).Return(nil, errors.New("Not Found")).Times(1)
			},
			404,
		},
	}
	for _, tt := range tests {

		//call to  expect mocks
		tt.fn()
		req, err := http.NewRequest(tt.args.Method, tt.args.URL, bytes.NewBuffer(tt.args.Body))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.cityHttpHandlers.GetCityDetails)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tt.wantStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tt.wantStatus)
		}

	}
}
