package galex

import "github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"

/*Controller structure manages controller galex */
type Controller struct {
	*_Packages
}

/*Response Http response package */
func (Controller *Controller) Response() *response.Driver {
	return _Context._IoC.Response
}
