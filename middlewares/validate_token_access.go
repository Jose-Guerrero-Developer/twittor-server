package middlewares

import (
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"
)

/*ValidateTokenAccess validate access token */
func ValidateTokenAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var JWT models.JWT
		var Galex galex.Driver
		_, _, _, err := JWT.ValidateToken(Galex.Utils().URL.GetHeader("Authorization"))
		if err != nil {
			Galex.Response().Failed("010", "Error validate token", "Impossible to generate access token. "+err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
