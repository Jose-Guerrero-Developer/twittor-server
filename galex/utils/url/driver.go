package url

import (
	"errors"
	"net/http"
)

/*_Context stores the status of the URL driver */
var _Context *Driver

/*Driver structure manages pagination */
type Driver struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

/*GetDriver Set initial settings for URLs */
func (Controller *Driver) GetDriver(w http.ResponseWriter, r *http.Request) error {
	_Context = new(Driver)
	_Context.Writer = w
	_Context.Request = r
	Controller.payload()
	if err := recover(); err != nil {
		return errors.New("An error occurred when reading URL input data")
	}
	return nil
}
