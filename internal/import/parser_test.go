package internalimport

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParsedRecord(t *testing.T) {
	t.Run("creates a new parsed record", func(t *testing.T) {
		sample := []string{"one", "two", "three"}

		pr := NewParsedRecord(sample)

		require.NotNil(t, pr)
		require.Equal(t, pr.Record, sample)
	})

	t.Run("given a parsed record, check if it contains empty fields and return false", func(t *testing.T) {
		sample := []string{"one", "two", "three"}

		pr := NewParsedRecord(sample)

		hasEmptyFields := pr.ContainsEmptyFields()
		require.False(t, hasEmptyFields)
	})

	t.Run("given a parsed record, check if it contains empty fields and return true", func(t *testing.T) {
		sample := []string{""}

		pr := NewParsedRecord(sample)

		hasEmptyFields := pr.ContainsEmptyFields()
		require.True(t, hasEmptyFields)
	})

	t.Run("given a parsed record, check if it the record is nil and return true", func(t *testing.T) {
		sample := []string{""}

		pr := NewParsedRecord(sample)

		hasEmptyFields := pr.ContainsEmptyFields()
		require.True(t, hasEmptyFields)
	})
}
