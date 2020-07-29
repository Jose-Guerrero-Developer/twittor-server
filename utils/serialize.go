package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseErrorJWT response erro json
type ResponseErrorJWT struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

// Response response create a resource
func Response(w http.ResponseWriter, statusHTTP int, data map[string]string) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusHTTP)
	w.Write([]byte(string(response)))
}

// ResponseError response http
func ResponseError(w http.ResponseWriter, statusHTTP int, code string, message string, description string) {
	transform := map[string]string{"code": code, "message": message, "description": description}
	response, _ := json.Marshal(transform)
	http.Error(w, string(response), statusHTTP)
}
