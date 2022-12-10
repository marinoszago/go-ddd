// Package geolocation contains the domain layer logic for a GeoData entity
package geolocation

import (
	"regexp"

	internalimport "go-ddd/internal/import"
)

const (
	// IPRegex contains the regex for a valid IP address.
	IPRegex = "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"

	// CcRegex contains the regex for a valid country code.
	CcRegex = "[A-Z]\\d{0,3}"

	// CountryRegex contains the regex for a valid country.
	CountryRegex = "[a-zA-Z]{2,}"

	// CityRegex contains the regex for a valid city.
	CityRegex = "[a-zA-Z]{2,}"

	// LatRegex contains the regex for a valid latitude.
	LatRegex = "^-?[0-9]{1,3}(?:\\.[0-9]{1,20})?$"

	// LngRegex contains the regex for a valid longitude.
	LngRegex = "^-?[0-9]{1,3}(?:\\.[0-9]{1,20})?$"
)

// GeoData is a struct defining a geolocation entity.
type GeoData struct {
	ID          uint   `gorm:"primaryKey, autoIncrement"`
	IPAddress   string `json:"ip_address"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Lat         string `json:"latitude"`
	Lng         string `json:"longitude"`
}

// Validator is an interface for a validator.
//
//go:generate mockgen -destination ./mock/validator.go -package mock . Validator
type Validator interface {
	Validate() (GeoData, error)
}

// ImportedRecordValidator is a struct implementing the Validator interface.
type ImportedRecordValidator struct {
	Rec    internalimport.ParsedRecord
	Stats  ImportStats
	Errors ImportErrors
}

// ImportStats is a custom type for stats during import.
type ImportStats map[string]int

// ImportErrors is a custom type for errors during import.
type ImportErrors map[GeoData]error

// NewImportedRecordValidatorDefault creates a new ImportRecordValidator.
func NewImportedRecordValidatorDefault() ImportedRecordValidator {
	return ImportedRecordValidator{
		Rec:    internalimport.ParsedRecord{},
		Stats:  make(map[string]int),
		Errors: make(map[GeoData]error),
	}
}

// Validate validates a record.
func (irv ImportedRecordValidator) Validate() GeoData {
	var entity GeoData

	if irv.Rec.ContainsEmptyFields() {
		irv.Stats["discarded"]++

		return GeoData{}
	}

	if !irv.validateFieldsForValidType() {
		irv.Stats["discarded"]++

		return GeoData{}
	}

	irv.Stats["accepted"]++

	entity.IPAddress = irv.Rec.Record[0]
	entity.CountryCode = irv.Rec.Record[1]
	entity.Country = irv.Rec.Record[2]
	entity.City = irv.Rec.Record[3]
	entity.Lat = irv.Rec.Record[4]
	entity.Lng = irv.Rec.Record[5]

	return entity
}

// ContainsValidFieldForRegex checks a field based on a regex.
func ContainsValidFieldForRegex(regex, field string) (bool, error) {
	match, err := regexp.MatchString(regex, field)
	if err != nil {
		return false, err
	}

	return match, nil
}

func (irv ImportedRecordValidator) validateFieldsForValidType() bool {
	isValidIP, err := ContainsValidFieldForRegex(IPRegex, irv.Rec.Record[0])
	if err != nil {
		return false
	}

	isValidCC, err := ContainsValidFieldForRegex(CcRegex, irv.Rec.Record[1])
	if err != nil {
		return false
	}

	isValidCountry, err := ContainsValidFieldForRegex(CountryRegex, irv.Rec.Record[2])
	if err != nil {
		return false
	}

	isValidCity, err := ContainsValidFieldForRegex(CityRegex, irv.Rec.Record[3])
	if err != nil {
		return false
	}

	isValidLat, err := ContainsValidFieldForRegex(LatRegex, irv.Rec.Record[4])
	if err != nil {
		return false
	}

	isValidLng, err := ContainsValidFieldForRegex(LngRegex, irv.Rec.Record[5])
	if err != nil {
		return false
	}

	return isValidIP && isValidCC && isValidCountry && isValidCity && isValidLat && isValidLng
}

// RemoveDuplicates removes duplicates for GeoData.
func (irv ImportedRecordValidator) RemoveDuplicates(pr []internalimport.ParsedRecord) []internalimport.ParsedRecord {
	var unique []internalimport.ParsedRecord

	m := make(map[GeoData]int)

	for _, values := range pr {
		ipAddress := values.Record[0]
		countryCode := values.Record[1]
		country := values.Record[2]
		city := values.Record[3]
		lat := values.Record[4]
		lng := values.Record[5]

		k := GeoData{
			IPAddress:   ipAddress,
			CountryCode: countryCode,
			Country:     country,
			City:        city,
			Lat:         lat,
			Lng:         lng,
		}

		if i, ok := m[k]; ok {
			unique[i] = values
		} else {
			m[k] = len(unique)
			unique = append(unique, values)
		}
	}

	return unique
}
