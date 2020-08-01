package app

import (
	"log"
	"net/http"
)

/*Run initiates settings (DotEnv, Router, ListenAndServe) and web server */
func Run() {

	if err := Configs.LoadDriver(); err != nil {
		log.Fatal("Failed to load configurations driver. " + err.Error())
	}

	if err := Database.LoadDriver("twittor"); err != nil {
		log.Fatal("Failed to load database driver. " + err.Error())
	}

	if err := Router.LoadDriver(); err != nil {
		log.Fatal("Failed to load driver route system. " + err.Error())
	}

	log.Print("Application: localhost:", Configs.Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+Configs.Get("PORT"), Router.GetRoutes()))
}
