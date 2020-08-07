package url

import (
	"net/http"
)

/*GetWriter  Write HTTP request */
func (Url *Driver) GetWriter() http.ResponseWriter {
	return storage.Writer
}

/*AddHeader Add header */
func (Url *Driver) AddHeader(name string, value string) {
	storage.Writer.Header().Set(name, value)
}

/*WriteHeader Overwrite HTTP response status */
func (Url *Driver) WriteHeader(statusCodeHTTP int) {
	storage.Writer.WriteHeader(statusCodeHTTP)
}
