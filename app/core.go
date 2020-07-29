package app

import (
	"log"
	"net/http"
	"os"

	"github.com/Jose-Guerrero-Developer/twittorbackend/router"
	"github.com/joho/godotenv"
)

// Run start serve
func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	AppRoutes := router.NewRouter()

	log.Print("Application: localhost:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, AppRoutes))
}
