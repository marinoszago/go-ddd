package geolocation

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"go-ddd/application/geolocation/mock"
	"go-ddd/domain/geolocation"
)

func TestControllerImpl_GetLocationDataForIP(t *testing.T) {
	t.Run("given a controller and a request return successful Json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := "1.1.1.1"
		expectedGeoData := []geolocation.GeoData{
			{
				ID:          0,
				IPAddress:   "1.1.1.1",
				CountryCode: "GR",
				Country:     "Greece",
				City:        "Athens",
				Lat:         "13.87656789",
				Lng:         "-10.3312332",
			},
		}

		apiServiceMock := mock.NewMockApiService(ctrl)
		apiServiceMock.EXPECT().GetLocationDataForIP(ip).Return(expectedGeoData, nil)

		writer := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		params := httprouter.Params{httprouter.Param{Key: "ip", Value: ip}}

		gc := NewController(apiServiceMock)

		gc.GetLocationDataForIP(writer, req, params)

		require.Equal(t, writer.Result().StatusCode, 200)

		bodyBytes, err := io.ReadAll(writer.Result().Body)
		require.NoError(t, err)

		bodyString := string(bodyBytes)

		expectedResponse := "[{\"ID\":0,\"ip_address\":\"1.1.1.1\",\"country_code\":\"GR\",\"country\":\"Greece\",\"city\":\"Athens\",\"latitude\":\"13.87656789\",\"longitude\":\"-10.3312332\"}]"

		require.Equal(t, bodyString, expectedResponse)
	})
	t.Run("given a controller and a request return error Json, with code 400", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := "test"

		apiServiceMock := mock.NewMockApiService(ctrl)

		writer := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		params := httprouter.Params{httprouter.Param{Key: "ip", Value: ip}}

		gc := NewController(apiServiceMock)

		gc.GetLocationDataForIP(writer, req, params)

		require.Equal(t, writer.Result().StatusCode, 400)

	})
	t.Run("given a controller and a request return error Json, with code 500", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := "1.1.1.1"
		expectedGeoData := []geolocation.GeoData{
			{
				ID:          0,
				IPAddress:   "1.1.1.1",
				CountryCode: "GR",
				Country:     "Greece",
				City:        "Athens",
				Lat:         "13.87656789",
				Lng:         "-10.3312332",
			},
		}

		apiServiceMock := mock.NewMockApiService(ctrl)
		apiServiceMock.EXPECT().GetLocationDataForIP(ip).Return(expectedGeoData, errors.New("an error occurred"))

		writer := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
		params := httprouter.Params{httprouter.Param{Key: "ip", Value: ip}}

		gc := NewController(apiServiceMock)

		gc.GetLocationDataForIP(writer, req, params)

		require.Equal(t, writer.Result().StatusCode, 500)

		bodyBytes, err := io.ReadAll(writer.Result().Body)
		require.NoError(t, err)

		bodyString := string(bodyBytes)

		expectedResponse := "{\"reason\":\"an error occurred\"}"
		require.Equal(t, bodyString, expectedResponse)
	})
}
