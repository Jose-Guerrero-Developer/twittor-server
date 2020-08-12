package response

import (
	"encoding/json"

	"github.com/devJGuerrero/twittor-server/galex/utils/url"
)

/*Success returns a response structure in json format */
func (Http *Driver) Success(data interface{}, statusHTTP int) {
	var GalexURL url.Driver

	GalexURL.WriteHeader(statusHTTP)
	json.NewEncoder(GalexURL.GetWriter()).Encode(data)
}
