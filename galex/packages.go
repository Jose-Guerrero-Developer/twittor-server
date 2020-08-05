package galex

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/configs"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils"
)

/*_Packages Contains the germicidal packages */
type _Packages struct{}

func (Controller *_Packages) Configs() *configs.Driver {
	return _Context._IoC.Configs
}

func (Controller *_Packages) Utils() *utils.Driver {
	return _Context._IoC.Utils
}
