package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"
)

/*setRoutesAPI subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesAPI() {
	/* routes auth */
	var Auth controllers.Auth
	subscribe("POST", "/api/sign", Auth.Sign, "")
	/* routes users */
	var User controllers.User
	subscribe("POST", "/api/users", User.Store, "ValidateTokenAccess")
	/* routes profile */
	var Profile controllers.Profile
	subscribe("GET", "/api/profiles", Profile.Get, "ValidateTokenAccess")
	subscribe("PUT", "/api/profiles", Profile.Update, "ValidateTokenAccess")
	/* routes tweets */
	var Tweet controllers.Tweet
	subscribe("GET", "/api/tweets/profile/{id}", Tweet.GetProfile, "ValidateTokenAccess")
	subscribe("POST", "/api/tweets", Tweet.Store, "ValidateTokenAccess")
}
