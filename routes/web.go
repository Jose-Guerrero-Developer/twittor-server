package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

/*setRoutesWeb subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesWeb() {
	/* route start */
	var Home controllers.Home
	subscribe("/", middlewares.CheckConnectionStatus(Home.Get)).Methods("GET")
}
