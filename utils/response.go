package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

/*_ResponseBasic basic data output */
type _ResponseBasic struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

/*Response returns a response structure in json format */
func Response(w http.ResponseWriter, data bson.M, statusHTTP int) {
	w.WriteHeader(statusHTTP)
	json.NewEncoder(w).Encode(data)
}

/*ResponseFailed respond to a faulty transaction */
func ResponseFailed(w http.ResponseWriter, code string, message string, description string, statusHTTP int) {
	w.WriteHeader(statusHTTP)
	response := _ResponseBasic{
		Code:        code,
		Message:     message,
		Description: description,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Impossible to transform error response data for failed transaction")
	}
}
