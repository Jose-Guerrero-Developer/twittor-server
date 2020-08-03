package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"

	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

/*setRoutesAPI subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesAPI() {
	/* routes auth */
	var Auth controllers.Auth
	subscribe("/api/sign", middlewares.CheckConnectionStatus(Auth.Sign)).Methods("POST")
	/* routes users */
	var User controllers.User
	subscribe("/api/users", middlewares.CheckConnectionStatus(User.Store)).Methods("POST")
	/* routes profile */
	var Profile controllers.Profile
	subscribe("/api/profiles", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(Profile.Get))).Methods("GET")
	subscribe("/api/profiles", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(Profile.Update))).Methods("PUT")
	/* routes tweets */
	var Tweet controllers.Tweet
	subscribe("/api/tweets", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(Tweet.Store))).Methods("POST")
}
