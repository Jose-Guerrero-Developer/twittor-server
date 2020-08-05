package app

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/configs"
	"github.com/Jose-Guerrero-Developer/twittorbackend/database"
	"github.com/Jose-Guerrero-Developer/twittorbackend/routes"
)

/*Configs configuration instance */
var Configs configs.Driver

/*Database database instance */
var Database database.Driver

/*Router router instance */
var Router routes.Driver
