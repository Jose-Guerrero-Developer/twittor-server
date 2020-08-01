package utils

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

/*ResponseErrorJWT error response structure */
type ResponseErrorJWT struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

/*Response returns a response structure in json format */
func Response(w http.ResponseWriter, statusHTTP int, data map[string]string) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusHTTP)
	w.Write([]byte(string(response)))
}

/*ResponseError returns an error response structure in json format */
func ResponseError(w http.ResponseWriter, statusHTTP int, code string, message string, description string) {
	transform := map[string]string{"code": code, "message": message, "description": description}
	response, _ := json.Marshal(transform)
	http.Error(w, string(response), statusHTTP)
}

/*EncryptPassword encrypt password */
func EncryptPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
