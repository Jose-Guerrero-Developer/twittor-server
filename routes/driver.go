package routes

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*_ControllerInstance contains the status of the router instance */
var _ControllerInstance *Driver

/*subscribe records the routes in the *mux instance */
var subscribe func(string, func(http.ResponseWriter, *http.Request)) *mux.Route

/*Driver structure to manage the entire route ecosystem */
type Driver struct {
	router *mux.Router
	routes http.Handler
}

/*LoadDriver instance router */
func (Controller *Driver) LoadDriver() error {
	_ControllerInstance = new(Driver)
	_ControllerInstance.router = mux.NewRouter()
	subscribe = _ControllerInstance.router.HandleFunc
	if err := recover(); err != nil {
		return errors.New("An error occurred in establishing driver routes")
	}
	return nil
}

/*GetRoutes returns the routes registered in the *mux.Router instance */
func (Controller *Driver) GetRoutes() http.Handler {
	_ControllerInstance.setRoutesAPI()
	_ControllerInstance.setRoutesWeb()
	_ControllerInstance.routes = cors.AllowAll().Handler(_ControllerInstance.router)
	return _ControllerInstance.routes
}
