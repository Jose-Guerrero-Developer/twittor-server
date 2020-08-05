package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"github.com/Jose-Guerrero-Developer/twittorbackend/authentication"
)

/*ValidateTokenAccess validate access token */
func ValidateTokenAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := authentication.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			var Galex galex.Driver
			Galex.Response().Failed("010", "Error validate token", "Impossible to generate access token. "+err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
