package geolocation

import (
	"testing"

	"github.com/stretchr/testify/require"

	internalimport "go-ddd/internal/import"
)

func TestImportedRecordValidator_Validate(t *testing.T) {
	t.Run("given a validator, validate a record and return a GeoData entity", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()

		rec := []string{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		gd := irv.Validate()

		require.NotNil(t, gd)
		require.Len(t, irv.Errors, 0)
		require.Equal(t, gd.ID, uint(0))
		require.Equal(t, gd.IPAddress, rec[0])
		require.Equal(t, gd.CountryCode, rec[1])
		require.Equal(t, gd.Country, rec[2])
		require.Equal(t, gd.City, rec[3])
		require.Equal(t, gd.Lat, rec[4])
		require.Equal(t, gd.Lng, rec[5])
		require.Equal(t, irv.Stats["accepted"], 1)
		require.Equal(t, irv.Stats["discarded"], 0)
	})

	t.Run("given a validator, validate a record with empty field(s) and return a discarded record", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()

		rec := []string{"", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		gd := irv.Validate()

		require.NotNil(t, gd)
		require.Equal(t, irv.Stats["accepted"], 0)
		require.Equal(t, irv.Stats["discarded"], 1)
	})

	t.Run("given a validator, validate a record with non empty field(s) but invalid type and return a discarded record", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()

		// contains invalid IP address
		rec := []string{"1000.2.3.4", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		gd := irv.Validate()

		require.NotNil(t, gd)
		require.Equal(t, irv.Stats["accepted"], 0)
		require.Equal(t, irv.Stats["discarded"], 1)

		// contains invalid IP address
		rec = []string{"test", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}

		irv.Rec = internalimport.NewParsedRecord(rec)
		irv.Stats = ImportStats{}

		gd = irv.Validate()

		require.NotNil(t, gd)
		require.Equal(t, irv.Stats["accepted"], 0)
		require.Equal(t, irv.Stats["discarded"], 1)
	})
}
func TestValidateFieldsForValidType(t *testing.T) {
	t.Run("given a validator, validates successfully an IP address, returns true", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.True(t, isValid)
	})

	t.Run("given a validator, validates an invalid IP address, returns false", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"300.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.False(t, isValid)
	})

	t.Run("given a validator, validates a valid country code, returns true", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.True(t, isValid)
	})

	t.Run("given a validator, validates an invalid country code, returns false", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "gr", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.False(t, isValid)
	})

	t.Run("given a validator, validates a valid country, returns true", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.True(t, isValid)
	})

	t.Run("given a validator, validates an invalid country, returns false", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "N", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.False(t, isValid)
	})

	t.Run("given a validator, validates a valid city, returns true", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.True(t, isValid)
	})

	t.Run("given a validator, validates an invalid country, returns false", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "D", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.False(t, isValid)
	})

	t.Run("given a validator, validates a valid latitude, returns true", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.True(t, isValid)
	})

	t.Run("given a validator, validates an invalid latitude, returns false", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "12312312312", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.False(t, isValid)
	})

	t.Run("given a validator, validates a valid longitude, returns true", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.True(t, isValid)
	})

	t.Run("given a validator, validates an invalid longitude, returns false", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()
		rec := []string{"200.106.141.15", "GR", "Nepal", "DuBuquemouth", "-84.87503094689836", "7131312.206435933364332", "7823011346"}
		irv.Rec = internalimport.NewParsedRecord(rec)

		isValid := irv.validateFieldsForValidType()
		require.False(t, isValid)
	})
}
func TestImportedRecordValidator_RemoveDuplicates(t *testing.T) {
	t.Run("given a slice of parsed records, remove duplicates and return a new slice", func(t *testing.T) {
		irv := NewImportedRecordValidatorDefault()

		records := []internalimport.ParsedRecord{
			internalimport.NewParsedRecord([]string{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"200.106.142.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"}),
			internalimport.NewParsedRecord([]string{"1.1.1.1", "GR", "Greece", "Athens", "-84.87503094689836", "7.206435933364332", "7823011346"}),
		}

		unique := irv.RemoveDuplicates(records)

		require.NotNil(t, unique)
		require.Len(t, unique, 3)
		require.NotEqual(t, unique, records)
	})
}
