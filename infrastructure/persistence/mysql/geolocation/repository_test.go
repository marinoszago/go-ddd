package geolocation

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"go-ddd/domain/geolocation"
	"go-ddd/domain/geolocation/mock"
)

func TestRepositoryImpl_GetLocationDataFromIP(t *testing.T) {
	t.Run("given a valid ip address, return GeoData results", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := "1.1.1.1"
		expected := geolocation.GeoData{
			ID:          1,
			IPAddress:   "1.1.1.1",
			CountryCode: "GR",
			Country:     "Greece",
			City:        "Athens",
			Lat:         "12.2312312",
			Lng:         "2.33232",
		}

		mockRepository := mock.NewMockRepository(ctrl)
		mockRepository.EXPECT().GetLocationDataFromIP(ip).Times(1).Return([]geolocation.GeoData{expected}, nil)

		ge, err := mockRepository.GetLocationDataFromIP(ip)

		require.NoError(t, err)
		require.Equal(t, ge, []geolocation.GeoData{expected})
	})

	t.Run("given a valid ip address, return error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ip := "1.1.1.1"

		mockRepository := mock.NewMockRepository(ctrl)
		mockRepository.EXPECT().GetLocationDataFromIP(ip).Times(1).Return(nil, errors.New("an error occurred"))

		ge, err := mockRepository.GetLocationDataFromIP(ip)

		require.Error(t, err)
		require.Len(t, ge, 0)
	})
}

func TestRepositoryImpl_Upsert(t *testing.T) {
	t.Run("given a GeoData entity, successfully save it", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := geolocation.GeoData{
			ID:          1,
			IPAddress:   "1.1.1.1",
			CountryCode: "GR",
			Country:     "Greece",
			City:        "Athens",
			Lat:         "12.2312312",
			Lng:         "2.33232",
		}

		mockRepository := mock.NewMockRepository(ctrl)
		mockRepository.EXPECT().Upsert(expected).Return(nil)

		err := mockRepository.Upsert(expected)
		require.NoError(t, err)
	})

	t.Run("given a GeoData entity, successfully save it", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := geolocation.GeoData{
			ID:          1,
			IPAddress:   "1.1.1.1",
			CountryCode: "GR",
			Country:     "Greece",
			City:        "Athens",
			Lat:         "12.2312312",
			Lng:         "2.33232",
		}

		mockRepository := mock.NewMockRepository(ctrl)
		mockRepository.EXPECT().Upsert(expected).Return(errors.New("an error occurred"))

		err := mockRepository.Upsert(expected)
		require.Error(t, err)
	})
}
