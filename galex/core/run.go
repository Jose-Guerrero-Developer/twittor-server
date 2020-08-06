package core

import (
	"log"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/router"
	"github.com/Jose-Guerrero-Developer/twittorbackend/routes"

	"github.com/Jose-Guerrero-Developer/twittorbackend/configs"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/configuration"
)

/*Run Turn on the server and start the application life cycle */
func Run() error {
	/* Load configuration driver */
	var GalexConfigs configuration.Driver
	if err := GalexConfigs.EstablishDriver(); err != nil {
		return err
	}
	configs.Payload()

	/* Load database driver */
	var GalexDatabase database.Driver
	if err := GalexDatabase.EstablishDriver(GalexConfigs.Get("DATABASE_NAME")); err != nil {
		return err
	}

	/* Load route driver */
	var GalexRouter router.Driver
	GalexRouter.EstablishDriver()
	routes.API()
	routes.Web()

	/* Set up web server */
	log.Fatal(http.ListenAndServe(":"+GalexConfigs.Get("PORT"), GalexRouter.GetRoutes()))
	return nil
}
