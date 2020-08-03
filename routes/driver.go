package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*_Context contains the status of the router instance */
var _Context *Driver

/*Driver structure to manage the entire route ecosystem */
type Driver struct {
	router    *mux.Router
	routes    http.Handler
	subscribe func(string, func(http.ResponseWriter, *http.Request)) *mux.Route
}

/*LoadDriver instance router */
func (Controller *Driver) LoadDriver() error {
	_Context = new(Driver)
	_Context.router = mux.NewRouter()
	_Context.subscribe = _Context.router.HandleFunc
	if err := recover(); err != nil {
		return errors.New("An error occurred in establishing driver routes")
	}
	return nil
}

/*GetRoutes returns the routes registered in the *mux.Router instance */
func (Controller *Driver) GetRoutes() http.Handler {
	_Context.setRoutesAPI()
	_Context.setRoutesWeb()
	_Context.routes = cors.AllowAll().Handler(_Context.router)
	return _Context.routes
}

/*subscribe records the routes in the *mux instance */
func subscribe(method string, route string, resource func(http.ResponseWriter, *http.Request), validations string) {
	switch validations {
	case "CheckConnectionStatus":
		_Context.subscribe(route, middlewares.CheckConnectionStatus(resource)).Methods(method)
	case "ValidateAccessToken":
		_Context.subscribe(route, middlewares.ValidateAccessToken(resource)).Methods(method)
	case "CheckConnectionStatus, ValidateAccessToken":
		_Context.subscribe(route, middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(resource))).Methods(method)
	default:
		if validations == string("") {
			_Context.subscribe(route, resource).Methods(method)
		} else {
			log.Println("Impossible to process middlewares on request. Route: ", route, " => middlewares: ", validations)
		}
	}
}
