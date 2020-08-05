package response

import (
	"encoding/json"
	"log"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/url"
)

/*Failed respond to a faulty transaction */
func (Controller *Driver) Failed(code string, message string, description string, statusHTTP int) {
	var URL = new(url.Driver).GetContext()
	URL.Writer.WriteHeader(statusHTTP)
	response := _ResponseBasic{
		Code:        code,
		Message:     message,
		Description: description,
	}
	err := json.NewEncoder(URL.Writer).Encode(response)
	if err != nil {
		log.Println("Impossible to transform error response data for failed transaction")
	}
}
