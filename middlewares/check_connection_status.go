package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"github.com/Jose-Guerrero-Developer/twittorbackend/database"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/url"
)

/*CheckConnectionStatus check the connection to the database */
func CheckConnectionStatus(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Database database.Driver
		status, error := Database.GetStatus()
		if !status {
			http.Error(w, error, 500)
			return
		}
		var URL url.Driver
		if err := URL.GetDriver(w, r); err != nil {
			var Galex galex.Driver
			Galex.Response().Failed("013", "Error loading driver URL", err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
