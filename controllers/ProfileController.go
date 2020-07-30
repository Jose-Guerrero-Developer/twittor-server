package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

// Profile user controller
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ID := r.URL.Query().Get("id")
	profile := new(models.Profile)
	err := profile.Get(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := utils.ResponseErrorJWT{
			Code:        "008",
			Message:     "Error obtaining data",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(profile)
}
