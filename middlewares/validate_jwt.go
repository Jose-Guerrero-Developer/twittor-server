package middlewares

import (
	"encoding/json"
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
			w.WriteHeader(http.StatusUnauthorized)
			response := utils.ResponseErrorJWT{
				Code:        "010",
				Message:     "Error validate token",
				Description: "Impossible to generate access token. " + err.Error(),
			}
			json.NewEncoder(w).Encode(&response)
			return
		}
		next.ServeHTTP(w, r)
	}
}
