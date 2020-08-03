package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"

	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

/*setRoutesAPI subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesAPI() {
	/* routes auth */
	var AuthController controllers.AuthController
	subscribe("/api/sign", middlewares.CheckConnectionStatus(AuthController.Sign)).Methods("POST")
	/* routes users */
	var UserController controllers.UserController
	subscribe("/api/users", middlewares.CheckConnectionStatus(UserController.Store)).Methods("POST")
	/* routes profile */
	var ProfileController controllers.ProfileController
	subscribe("/api/profile", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(ProfileController.Get))).Methods("GET")
	subscribe("/api/profile", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(ProfileController.Update))).Methods("PUT")
	/* routes tweet */
	var TweetController controllers.TweetController
	subscribe("/api/tweet", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(TweetController.Store))).Methods("POST")
}
