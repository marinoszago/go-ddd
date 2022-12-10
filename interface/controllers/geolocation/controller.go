// Package geolocation contains controller level code for geolocation handling
package geolocation

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"go-ddd/application/geolocation"
)

// Controller is an interface for a controller.
//
//go:generate mockgen -destination ./mock/controller.go -package mock . Controller
type Controller interface {
	GetLocationDataForIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// ControllerImpl is a struct for a controller layer.
type ControllerImpl struct {
	service geolocation.ApiService
}

// NewController creates an instance of the ControllerImpl struct.
func NewController(service geolocation.ApiService) ControllerImpl {
	return ControllerImpl{
		service: service,
	}
}

// GetLocationDataForIP fetches data for an ip address.
func (c ControllerImpl) GetLocationDataForIP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ip := ps.ByName("ip")

	isValidIP, err := ContainsValidFieldForRegex(IPRegex, ip)
	if err != nil || !isValidIP {
		Error(w, http.StatusBadRequest, err, "given IP address is invalid")

		return
	}

	data, err := c.service.GetLocationDataForIP(ip)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "an error occurred")

		return
	}

	JSON(w, http.StatusOK, data)
}
