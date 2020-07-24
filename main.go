package main

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/handlers"
)

func main() {
	// status, error := database.CheckConnectionStatus()
	// if !status {
	// 	log.Fatal(error)
	// }
	handlers.Init()
}
