package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
)

/*setRoutesWeb subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesWeb() {
	/* route start */
	var Home controllers.Home
	subscribe("GET", "/", Home.Get, "")
}
