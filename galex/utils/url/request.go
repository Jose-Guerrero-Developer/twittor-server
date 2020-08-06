package url

import "net/http"

/*GetRequest get request */
func (Url *Driver) GetRequest() *http.Request {
	return storage.Request
}

/*GetHeader Returns one sent parameter per header */
func (Url *Driver) GetHeader(key string) string {
	return storage.Request.Header.Get(key)
}
