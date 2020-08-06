package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
)

/*Web subscribe the routes to the *mux.Router instance */
func Web() {
	/* route start */
	var Home controllers.Home
	subscribe("GET", "/", Home.Get)
}
