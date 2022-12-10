// Package geolocation contains the repository code for geolocation
package geolocation

import (
	"go-ddd/domain/geolocation"
	"go-ddd/infrastructure/persistence/mysql"
)

// RepositoryImpl Implements geolocation.Repository.
type RepositoryImpl struct {
	dbClient mysql.Client
}

// NewRepository returns initialised RepositoryImpl.
func NewRepository(dbClient mysql.Client) geolocation.Repository {
	return &RepositoryImpl{dbClient: dbClient}
}

// Upsert saves or updates a GeoData entity.
func (r RepositoryImpl) Upsert(ge geolocation.GeoData) error {
	if err := r.dbClient.Save(&ge).Error; err != nil {
		return err
	}

	return nil
}

// GetLocationDataFromIP is the function implementation that returns all rows for a given IP address.
func (r RepositoryImpl) GetLocationDataFromIP(ip string) ([]geolocation.GeoData, error) {
	var ge []geolocation.GeoData

	r.dbClient.Find(&ge, "ip_address LIKE ?", ip)

	if r.dbClient.Error != nil {
		return nil, r.dbClient.Error
	}

	return ge, nil
}
