package routes

import (
	"github.com/devJGuerrero/twittor-server/controllers"
)

/*Web subscribe the routes to the *mux.Router instance */
func Web() {
	/* route start */
	var Home controllers.Home
	subscribe("GET", "/", Home.Get)
}
