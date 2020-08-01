package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

/*setRoutesWeb subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesWeb() {
	/* route start */
	var HomeController controllers.HomeController
	subscribe("/", middlewares.CheckConnectionStatus(HomeController.Get)).Methods("GET")
}
