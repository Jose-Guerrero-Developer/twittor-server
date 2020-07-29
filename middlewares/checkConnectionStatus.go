package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/database"
)

// CheckDatabaseConnectionStatus check connection status database
func CheckDatabaseConnectionStatus(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, error := database.CheckConnectionStatus()
		if !status {
			http.Error(w, error, 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
