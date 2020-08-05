package url

import "net/http"

/*GetRequest get request */
func (Controller *Driver) GetRequest() *http.Request {
	return _Context.Request
}

/*GetHeader Returns one sent parameter per header */
func (Controller *Driver) GetHeader(key string) string {
	return _Context.Request.Header.Get(key)
}
