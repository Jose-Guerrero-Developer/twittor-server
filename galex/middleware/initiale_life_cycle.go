package middleware

import (
	"net/http"

	"github.com/devJGuerrero/twittor-server/galex/utils/url"
)

/*InitialeLifeCycle start the life cycle of the Galex */
func (Controller *Driver) InitialeLifeCycle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var GalexURL url.Driver
		GalexURL.EstablishDriver(w, r)
		next.ServeHTTP(w, r)
	})
}
