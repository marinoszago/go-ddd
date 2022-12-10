package geolocation

// Repository is the interface for a GeoData Entity.
//
//go:generate mockgen -destination ./mock/repository.go -package mock . Repository
type Repository interface {
	Upsert(ge GeoData) error
	GetLocationDataFromIP(ip string) ([]GeoData, error)
}
