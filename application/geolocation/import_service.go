// Package geolocation contains the ServiceImpl layer details for import
package geolocation

import (
	"fmt"

	"go-ddd/domain/geolocation"
	internalimport "go-ddd/internal/import"
)

// ImportService is a struct responsible for the import.
type ImportService struct {
	repository geolocation.Repository
}

// NewImportService constructs a new ImportService struct.
func NewImportService(repository geolocation.Repository) ImportService {
	return ImportService{
		repository: repository,
	}
}

// Save is the import service function responsible for saving a record coming from the imported file.
func (s ImportService) Save(importedRecords []internalimport.ParsedRecord) (geolocation.ImportStats, geolocation.ImportErrors) {
	var importedValidator = geolocation.NewImportedRecordValidatorDefault()

	uniqueImported := importedValidator.RemoveDuplicates(importedRecords)

	fmt.Println("[RUN] >>> will process", len(uniqueImported), "unique entries")
	fmt.Println("[RUN] >>> waiting...")
	for _, importedRecord := range uniqueImported {
		importedValidator.Rec = importedRecord

		entity := importedValidator.Validate()

		if entity != (geolocation.GeoData{}) {
			err := s.repository.Upsert(entity)
			if err != nil {
				importedValidator.Errors[entity] = err
			}
		}
	}

	return importedValidator.Stats, importedValidator.Errors
}
