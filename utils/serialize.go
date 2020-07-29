package utils

import (
	"encoding/json"
	"net/http"
)

// Response response create a resource
func Response(w http.ResponseWriter, statusHTTP int, data map[string]string) {
	response, _ := json.Marshal(data)
	w.WriteHeader(statusHTTP)
	w.Write([]byte(string(response)))
}

// ResponseError response http
func ResponseError(w http.ResponseWriter, statusHTTP int, code string, message string, description string) {
	transform := map[string]string{"code": code, "message": message, "description": description}
	response, _ := json.Marshal(transform)
	http.Error(w, string(response), statusHTTP)
}
