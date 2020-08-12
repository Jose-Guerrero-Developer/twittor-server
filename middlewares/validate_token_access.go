package middlewares

import (
	"net/http"

	"github.com/devJGuerrero/twittor-server/galex/utils/url"

	"github.com/devJGuerrero/twittor-server/galex/response"

	"github.com/devJGuerrero/twittor-server/models"
)

/*ValidateTokenAccess validate access token */
func ValidateTokenAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var JWT models.JWT
		var GalexURL url.Driver
		var GalexResponse response.Driver
		_, _, _, err := JWT.ValidateToken(GalexURL.GetHeader("Authorization"))
		if err != nil {
			GalexResponse.Failed("010", "Error validate token", "Impossible to generate access token. "+err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
