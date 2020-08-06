package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/router"
	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

var subscribe = router.Subscribe

/*API subscribe the routes to the *mux.Router instance */
func API() {
	/* route start */
	var Auth controllers.Auth
	subscribe("POST", "/api/sign", Auth.Sign)

	/* routes users */
	var User controllers.User
	subscribe("POST", "/api/users", middlewares.ValidateTokenAccess(User.Store))

	/* routes profile */
	var Profile controllers.Profile
	subscribe("GET", "/api/profiles", middlewares.ValidateTokenAccess(Profile.Get))
	subscribe("PUT", "/api/profiles", middlewares.ValidateTokenAccess(Profile.Update))

	/* routes tweets */
	var Tweet controllers.Tweet
	subscribe("GET", "/api/tweets/profile/{id}", middlewares.ValidateTokenAccess(Tweet.GetProfile))
	subscribe("POST", "/api/tweets", middlewares.ValidateTokenAccess(Tweet.Store))
}
