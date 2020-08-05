package controllers

import (
	"errors"
	"net/http"
)

/*_Context stores the status of the controllers driver */
var _Context *Driver

/*Driver structure manages controllers */
type Driver struct{}

/*LoadDriver Set initial settings for controllers */
func (Controller *Driver) LoadDriver(w http.ResponseWriter, r *http.Request) error {
	_Context = new(Driver)
	if err := recover(); err != nil {
		return errors.New("An error occurred when reading URL input data")
	}
	return nil
}
