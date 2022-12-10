package geolocation

import (
	"go-ddd/domain/geolocation"
)

// ApiService is an interface for the api_service.
//
//go:generate mockgen -destination ./mock/api_service.go -package mock . ApiService
type ApiService interface {
	GetLocationDataForIP(ip string) ([]geolocation.GeoData, error)
}

// ServiceImpl is a struct for the application layer service for an Entity.
type ServiceImpl struct {
	repository geolocation.Repository
}

// NewService creates a new ServiceImpl struct.
func NewService(repository geolocation.Repository) ServiceImpl {
	return ServiceImpl{
		repository: repository,
	}
}

// GetLocationDataForIP fetches all the records for a specific IP.
func (s ServiceImpl) GetLocationDataForIP(ip string) ([]geolocation.GeoData, error) {
	return s.repository.GetLocationDataFromIP(ip)
}
