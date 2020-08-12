package configs

import (
	"os"

	"github.com/devJGuerrero/twittor-server/galex/configuration"
)

/*Driver Source of the package */
type Driver struct{}

var subscribe = configuration.Subscribe

/*Payload Register configuration files */
func Payload() {
	var Configs Driver
	Configs.App()
	Configs.Database()
}

/*App Configs app */
func (Configs *Driver) App() {
	/*Name APP NAME */
	subscribe("APP_NAME", os.Getenv("APP_NAME"), "Twittor")
	/*Application PORT */
	subscribe("PORT", os.Getenv("PORT"), "8080")
	/*Application APP SECRET */
	subscribe("APP_SECRET", os.Getenv("APP_SECRET"), "123456")
}
