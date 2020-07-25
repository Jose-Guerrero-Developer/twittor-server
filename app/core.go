package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Jose-Guerrero-Developer/twittorbackend/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var routerInstance = mux.NewRouter()

// Run start serve
func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	router.Routes(subscribeRoutes)
	routes := cors.AllowAll().Handler(routerInstance)

	log.Print("Application: localhost:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, routes))
}

func subscribeRoutes(name string, handler func(w http.ResponseWriter, r *http.Request)) {
	routerInstance.HandleFunc(name, handler)
}
