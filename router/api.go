package router

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/controllers"

	"github.com/Jose-Guerrero-Developer/twittorbackend/middlewares"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Router struct
type Router struct {
	instance *mux.Router
	routes   http.Handler
}

// NewRouter instance mux
func NewRouter() http.Handler {
	router := new(Router)
	router.instance = mux.NewRouter()
	router.getRoutes()
	router.routes = cors.AllowAll().Handler(router.instance)
	return router.routes
}

func (router *Router) getRoutes() {
	var subscribe = router.instance.HandleFunc
	// START
	subscribe("/", middlewares.
		CheckDatabaseConnectionStatus(controllers.Home)).Methods("GET")
	// USERS
	subscribe("/users", middlewares.
		CheckDatabaseConnectionStatus(controllers.UserStore)).Methods("POST")
	// LOGIN
	subscribe("/login", middlewares.
		CheckDatabaseConnectionStatus(controllers.AuthLogin)).Methods("POST")
	// PROFILE
	subscribe("/profile", middlewares.
		CheckDatabaseConnectionStatus(
			middlewares.
				ValidateJWT(controllers.Profile))).Methods("GET")
}
