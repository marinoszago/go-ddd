package geolocation

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	geodomain "go-ddd/domain/geolocation"
	"go-ddd/infrastructure/persistence/mysql"
	"go-ddd/infrastructure/persistence/mysql/geolocation"
)

func TestServiceImpl_GetLocationDataForIP(t *testing.T) {
	t.Run("given a valid IP address return a GeoData entity", func(t *testing.T) {

		dbClient, err := mysql.NewClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			"root", "password", "localhost", "3306"))
		require.NoError(t, err)

		dbClient.CreateDatabase("test_db")
		dbClient.UseDatabase("test_db")

		err = dbClient.RunGormEntityMigrations()
		require.NoError(t, err)

		apiRepository := geolocation.NewRepository(dbClient)
		service := NewService(apiRepository)

		gd := geodomain.GeoData{
			ID:          1,
			IPAddress:   "1.1.1.1",
			CountryCode: "GR",
			Country:     "Greece",
			City:        "Athens",
			Lat:         "12.2222",
			Lng:         "-12.3333",
		}

		err = apiRepository.Upsert(gd)
		require.NoError(t, err)

		result, err := service.GetLocationDataForIP("1.1.1.1")
		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("given a valid IP address return empty results", func(t *testing.T) {

		dbClient, err := mysql.NewClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			"root", "password", "localhost", "3306"))
		require.NoError(t, err)

		dbClient.CreateDatabase("test_db")
		dbClient.UseDatabase("test_db")

		err = dbClient.RunGormEntityMigrations()
		require.NoError(t, err)

		apiRepository := geolocation.NewRepository(dbClient)
		service := NewService(apiRepository)

		gd := geodomain.GeoData{
			ID:          1,
			IPAddress:   "1.2.1.1",
			CountryCode: "GR",
			Country:     "Greece",
			City:        "Athens",
			Lat:         "12.2222",
			Lng:         "-12.3333",
		}

		err = apiRepository.Upsert(gd)
		require.NoError(t, err)

		result, err := service.GetLocationDataForIP("1.1.1.1")
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Len(t, result, 0)
	})

	t.Run("given a valid IP address return an error", func(t *testing.T) {

		dbClient, err := mysql.NewClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			"root", "password", "localhost", "3306"))
		require.NoError(t, err)

		err = dbClient.AddError(errors.New("an error occurred"))
		require.Error(t, err)

		apiRepository := geolocation.NewRepository(dbClient)
		service := NewService(apiRepository)

		result, err := service.GetLocationDataForIP("1.1.1.1")
		require.Error(t, err)
		require.Nil(t, result)
	})
}
