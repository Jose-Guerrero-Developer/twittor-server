package routes

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"

	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"
)

/*setRoutesAPI subscribe the routes to the *mux.Router instance */
func (Controller *Driver) setRoutesAPI() {
	/* route login */
	var AuthController controllers.AuthController
	subscribe("/api/sign", middlewares.CheckConnectionStatus(AuthController.Sign)).Methods("POST")
	/* route users */
	var UserController controllers.UserController
	subscribe("/api/users", middlewares.CheckConnectionStatus(UserController.Store)).Methods("POST")
	/* route profile */
	var ProfileController controllers.ProfileController
	subscribe("/api/profile", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(ProfileController.Get))).Methods("GET")
	subscribe("/api/profile", middlewares.CheckConnectionStatus(middlewares.ValidateAccessToken(ProfileController.Update))).Methods("PUT")
}
