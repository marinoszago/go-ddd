// Package geolocation contains the routes
package geolocation

import (
	"github.com/julienschmidt/httprouter"

	"go-ddd/interface/controllers/geolocation"
)

// Router is a struct for httprouter.
type Router struct {
	controller geolocation.Controller
}

// NewGeolocationRouter creates a new Router.
func NewGeolocationRouter(controller geolocation.Controller) Router {
	return Router{controller: controller}
}

// Routes returns the initialised router.
func (gr Router) Routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/get_location/:ip", gr.controller.GetLocationDataForIP)

	return r
}
