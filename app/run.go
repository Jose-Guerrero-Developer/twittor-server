package app

import (
	"log"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"
)

/*Run initiates settings (DotEnv, Router, ListenAndServe) and web server */
func Run() {
	Galex, err := galex.GetDriver()
	if err != nil {
		Galex.Response().Failed("013", "Error loading driver Galex", err.Error(), http.StatusInternalServerError)
		return
	}
	if err := Database.LoadDriver("twittor"); err != nil {
		log.Fatal("Failed to load database driver. " + err.Error())
	}
	if err := Router.LoadDriver(); err != nil {
		log.Fatal("Failed to load driver route system. " + err.Error())
	}
	log.Print("Application: localhost:", Galex.Configs().Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+Galex.Configs().Get("PORT"), Router.GetRoutes()))
}
