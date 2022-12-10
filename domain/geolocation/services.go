package geolocation

import (
	internalimport "go-ddd/internal/import"
)

// ImportService is an interface defining how to import geolocation data
//
//go:generate mockgen -destination ./mock/import_service.go -package mock . ImportService
type ImportService interface {
	Save(importedRecords []internalimport.ParsedRecord) (ImportStats, ImportErrors)
}

// ApiService is an interface defining how to implement an api geolocation service
//
//go:generate mockgen -destination ./mock/api_service.go -package mock . ApiService
type ApiService interface {
	GetLocationDataForIP(ip string) ([]GeoData, error)
}
