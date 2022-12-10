package geolocation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsValidFieldForRegex(t *testing.T) {
	t.Run("given a valid IP address, return true", func(t *testing.T) {
		ip := "1.1.1.1"

		isValid, err := ContainsValidFieldForRegex(IPRegex, ip)

		require.True(t, isValid)
		require.NoError(t, err)
	})

	t.Run("given an invalid IP address, return false", func(t *testing.T) {
		ip := "300.1.1.1"

		isValid, err := ContainsValidFieldForRegex(IPRegex, ip)

		require.False(t, isValid)
		require.NoError(t, err)
	})

	t.Run("given an invalid IP, return false", func(t *testing.T) {
		ip := "test"

		isValid, err := ContainsValidFieldForRegex(IPRegex, ip)

		require.False(t, isValid)
		require.NoError(t, err)
	})

	t.Run("given an invalid RegEx pattern, return false and error", func(t *testing.T) {
		ip := "1.1.1.1"
		pattern := "\\"
		isValid, err := ContainsValidFieldForRegex(pattern, ip)

		require.False(t, isValid)
		require.Error(t, err)
	})
}
