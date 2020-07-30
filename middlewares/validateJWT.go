package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/authentication"
)

// ValidateJWT validate authorization token access
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := authentication.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error validate token. "+err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
