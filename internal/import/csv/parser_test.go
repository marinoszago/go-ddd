package csv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var filePath = "../../../resources/data_dump_test.csv"
var invalidFilePath = "../../../resources/data_dump_test_invalid.csv"
var emptyFilePath = "../../../resources/data_dump_test_empty.csv"
var emptyOnlyHeadersFilePath = "../../../resources/data_dump_test_headers.csv"

func TestParser_Parse(t *testing.T) {
	t.Run("given a parser with a valid file, return parsed records", func(t *testing.T) {
		p := NewParser(filePath)

		records, err := p.Parse()
		require.NoError(t, err)
		require.Len(t, records, 2)
	})

	t.Run("given a parser with an invalid file path, return error", func(t *testing.T) {
		p := NewParser(invalidFilePath)

		records, err := p.Parse()
		require.Nil(t, records)
		require.Error(t, err)
	})

	t.Run("given a parser with a valid but empty file, return error", func(t *testing.T) {
		p := NewParser(emptyFilePath)

		records, err := p.Parse()
		require.Nil(t, records)
		require.Error(t, err)
	})

	t.Run("given a parser with a valid file but only containing headers, return nil records and no error",
		func(t *testing.T) {
			p := NewParser(emptyOnlyHeadersFilePath)

			records, err := p.Parse()
			require.Nil(t, records)
			require.NoError(t, err)
		})
}
