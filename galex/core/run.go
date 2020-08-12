package core

import (
	"log"
	"net/http"

	"github.com/devJGuerrero/twittor-server/galex/database"

	"github.com/devJGuerrero/twittor-server/galex/router"
	"github.com/devJGuerrero/twittor-server/routes"

	"github.com/devJGuerrero/twittor-server/configs"
	"github.com/devJGuerrero/twittor-server/galex/configuration"
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
