package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/database"
)

/*CheckStatusDatabaseConnection Check the status of the database connection */
func (Controller *Driver) CheckStatusDatabaseConnection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var Database database.Driver
		if status, error := Database.GetStatus(); !status {
			http.Error(w, error, 500)
			return
		}
		next.ServeHTTP(w, r)
	})
}
