package galex

import (
	"errors"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/middlewares"

	"github.com/Jose-Guerrero-Developer/twittorbackend/configs"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils"
)

/*_Context Stores instance context */
var _Context *Driver

/*Driver Package structure */
type Driver struct {
	*_Packages
	*Controller
	_IoC struct {
		Utils       *utils.Driver
		Configs     *configs.Driver
		Response    *response.Driver
		Middlewares *middlewares.Driver
	}
}

/*GetDriver Returns the instance of the package */
func GetDriver() (*Driver, error) {
	_Context = new(Driver)
	_Context._IoC.Utils = utils.LoadDriver()
	_Context._IoC.Configs = configs.GetDriver()
	_Context._IoC.Response = response.GetDriver()
	_Context._IoC.Middlewares = middlewares.GetDriver()
	if err := recover(); err != nil {
		return _Context, errors.New("Error loading Galex component instances")
	}
	return _Context, nil
}
