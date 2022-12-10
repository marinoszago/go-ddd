package mysql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("given a database, run migrations", func(t *testing.T) {
		dbClient, err := NewClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			"root", "password", "localhost", "3306"))
		require.NoError(t, err)

		dbClient.CreateDatabase("test_db")
		dbClient.UseDatabase("test_db")

		err = dbClient.RunGormEntityMigrations()
		require.NoError(t, err)
	})
}
