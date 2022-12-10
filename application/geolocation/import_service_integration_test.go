package geolocation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"go-ddd/infrastructure/persistence/mysql"
	"go-ddd/infrastructure/persistence/mysql/geolocation"
	internalimport "go-ddd/internal/import"
)

func TestImportService_Save(t *testing.T) {
	t.Run("given a slice of parsed records, successfully save them", func(t *testing.T) {
		records := []internalimport.ParsedRecord{
			internalimport.NewParsedRecord([]string{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"200.106.142.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"1.1.1.1", "GR", "Greece", "Athens", "-84.87503094689836", "7.206435933364332", "7823011346"}),
		}

		dbClient, err := mysql.NewClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			"root", "password", "localhost", "3306"))
		require.NoError(t, err)

		dbClient.CreateDatabase("test_db")
		dbClient.UseDatabase("test_db")

		err = dbClient.RunGormEntityMigrations()
		require.NoError(t, err)

		repository := geolocation.NewRepository(dbClient)
		service := NewImportService(repository)

		stats, errors := service.Save(records)

		require.Equal(t, stats["accepted"], 3)
		require.Equal(t, stats["discarded"], 0)
		require.NotNil(t, errors)
		require.Empty(t, errors)

		apiService := NewService(repository)

		gd, err := apiService.GetLocationDataForIP("1.1.1.1")
		require.NoError(t, err)

		require.NotNil(t, gd)
		require.Len(t, gd, 1)
	})

	t.Run("given a slice of invalid parsed records, unsuccessfully save them", func(t *testing.T) {
		records := []internalimport.ParsedRecord{
			internalimport.NewParsedRecord([]string{"500.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"200.106.142.15", "assssdd", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"test", "GR", "Greece", "Athens", "-84.87503094689836", "7.206435933364332", "7823011346"}),
		}

		dbClient, err := mysql.NewClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			"root", "password", "localhost", "3306"))
		require.NoError(t, err)

		dbClient.CreateDatabase("test_db")
		dbClient.UseDatabase("test_db")

		err = dbClient.RunGormEntityMigrations()
		require.NoError(t, err)

		repository := geolocation.NewRepository(dbClient)
		service := NewImportService(repository)

		stats, errors := service.Save(records)

		require.Equal(t, stats["accepted"], 0)
		require.Equal(t, stats["discarded"], 3)
		require.NotNil(t, errors)
		require.Empty(t, errors)

		apiService := NewService(repository)

		gd, err := apiService.GetLocationDataForIP("1.1.1.1")
		require.NoError(t, err)

		require.NotNil(t, gd)
		require.Len(t, gd, 0)
	})
}
