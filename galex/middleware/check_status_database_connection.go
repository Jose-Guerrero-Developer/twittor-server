package middleware

import (
	"net/http"

	"github.com/devJGuerrero/twittor-server/galex/database"
)

/*CheckStatusDatabaseConnection Check the status of the database connection */
func (Controller *Driver) CheckStatusDatabaseConnection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var GalexDatabase database.Driver
		if status, error := GalexDatabase.GetStatus(); !status {
			http.Error(w, error, 500)
			return
		}
		next.ServeHTTP(w, r)
	})
}
