package router

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Driver Source of the package */
type Driver struct {
	router    *mux.Router
	routes    http.Handler
	subscribe func(string, func(http.ResponseWriter, *http.Request)) *mux.Route
}

/*storage Store package status */
var storage *Driver

/*EstablishDriver Set package driver */
func (Router *Driver) EstablishDriver() {
	storage = new(Driver)
	storage.router = mux.NewRouter()
	storage.subscribe = storage.router.HandleFunc
}

/*GetRoutes Returns the routes registered in the *mux.Router instance */
func (Router *Driver) GetRoutes() http.Handler {
	var GalexMiddleware middleware.Driver
	storage.router.Use(GalexMiddleware.InitialeLifeCycle)
	storage.router.Use(GalexMiddleware.CheckStatusDatabaseConnection)
	storage.routes = cors.AllowAll().Handler(storage.router)
	return storage.routes
}

/*Subscribe Records the routes in the *mux instance */
func Subscribe(method string, route string, resource func(http.ResponseWriter, *http.Request)) {
	storage.subscribe(route, resource).Methods(method)
}
