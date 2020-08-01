package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/database"
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
		next.ServeHTTP(w, r)
	}
}
