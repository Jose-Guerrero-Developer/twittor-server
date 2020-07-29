package controllers

import "net/http"

// Home controller
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application running..."))
}
