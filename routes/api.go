package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
)

/*setRoutesAPI subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesAPI() {
	/* routes auth */
	var Auth controllers.Auth
	subscribe("POST", "/api/sign", Auth.Sign, "CheckConnectionStatus")
	/* routes users */
	var User controllers.User
	subscribe("POST", "/api/users", User.Store, "CheckConnectionStatus, ValidateAccessToken")
	/* routes profile */
	var Profile controllers.Profile
	subscribe("GET", "/api/profiles", Profile.Get, "CheckConnectionStatus, ValidateAccessToken")
	subscribe("PUT", "/api/profiles", Profile.Update, "CheckConnectionStatus, ValidateAccessToken")
	/* routes tweets */
	var Tweet controllers.Tweet
	subscribe("POST", "/api/tweets", Tweet.Store, "CheckConnectionStatus, ValidateAccessToken")
}
