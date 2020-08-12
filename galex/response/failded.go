package response

import (
	"encoding/json"

	"github.com/devJGuerrero/twittor-server/galex/utils/url"
)

/*Failed respond to a faulty transaction */
func (Http *Driver) Failed(code string, message string, description string, statusHTTP int) {
	var GalexURL url.Driver
	GalexURL.WriteHeader(statusHTTP)
	json.NewEncoder(GalexURL.GetWriter()).Encode(responseBasic{
		Code:        code,
		Message:     message,
		Description: description,
	})
}
