package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/url"
)

/*InitiateLifeCycle start the life cycle of the Galex */
func (Controller *Driver) InitiateLifeCycle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var URL url.Driver
		if err := URL.GetDriver(w, r); err != nil {
			var Response response.Driver
			Response.Failed("013", "Error loading driver URL", err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
