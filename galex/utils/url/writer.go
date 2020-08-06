package url

import (
	"net/http"
)

/*GetWriter  Write HTTP request */
func (Url *Driver) GetWriter() http.ResponseWriter {
	return storage.Writer
}

/*WriteHeader Overwrite HTTP response status */
func (Url *Driver) WriteHeader(statusCodeHTTP int) {
	storage.Writer.WriteHeader(statusCodeHTTP)
}
