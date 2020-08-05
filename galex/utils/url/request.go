package url

import "net/http"

/*GetRequest get request */
func (Controller *Driver) GetRequest() *http.Request {
	return _Context.Request
}
