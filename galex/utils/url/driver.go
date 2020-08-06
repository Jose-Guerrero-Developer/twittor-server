package url

import (
	"net/http"
)

/*Driver Source of the package */
type Driver struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

/*storage Store package status */
var storage *Driver

/*EstablishDriver Set package driver */
func (Url *Driver) EstablishDriver(w http.ResponseWriter, r *http.Request) {
	storage = new(Driver)
	storage.Writer = w
	storage.Request = r
	Url.payload()
}
