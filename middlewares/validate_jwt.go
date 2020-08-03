package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/authentication"
	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"
)

/*ValidateAccessToken validate access token */
func ValidateAccessToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := authentication.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			utils.ResponseFailed(w, "010", "Error validate token", "Impossible to generate access token. "+err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
