package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/router"
	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

var subscribe = router.Subscribe

/*API subscribe the routes to the *mux.Router instance */
func API() {
	/* Route Start */
	var Auth controllers.Auth
	subscribe("POST", "/api/sign", Auth.Sign)
	/* Routes Users */
	var User controllers.User
	subscribe("POST", "/api/users", middlewares.ValidateTokenAccess(User.Store))
	/* Routes Profile */
	var Profile controllers.Profile
	subscribe("GET", "/api/profiles", middlewares.ValidateTokenAccess(Profile.Get))
	subscribe("PUT", "/api/profiles/{id}", middlewares.ValidateTokenAccess(Profile.Update))
	/* Routes Tweets */
	var Tweet controllers.Tweet
	subscribe("GET", "/api/tweets", middlewares.ValidateTokenAccess(Tweet.Get))
	subscribe("GET", "/api/tweets/{id}", middlewares.ValidateTokenAccess(Tweet.GetID))
	subscribe("GET", "/api/tweets/profile/{id}", middlewares.ValidateTokenAccess(Tweet.GetProfile))
	subscribe("POST", "/api/tweets", middlewares.ValidateTokenAccess(Tweet.Store))
	subscribe("PUT", "/api/tweets/{id}", middlewares.ValidateTokenAccess(Tweet.Update))
	subscribe("DELETE", "/api/tweets/{id}", middlewares.ValidateTokenAccess(Tweet.Delete))
	/* Routes Uploads */
	var Upload controllers.Upload
	subscribe("POST", "/api/profiles/{id}/upload/avatar", middlewares.ValidateTokenAccess(Upload.Avatar))
}
