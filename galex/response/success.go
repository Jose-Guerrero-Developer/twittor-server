package response

import (
	"encoding/json"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/url"
)

/*Success returns a response structure in json format */
func (Controller *Driver) Success(data interface{}, statusHTTP int) {
	var URL = new(url.Driver).GetContext()
	URL.Writer.WriteHeader(statusHTTP)
	json.NewEncoder(URL.Writer).Encode(data)
}
