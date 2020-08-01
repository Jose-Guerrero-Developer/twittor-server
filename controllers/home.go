package controllers

import "net/http"

/*HomeController home controller */
type HomeController struct{}

/*Get returns information about the application */
func (Controller *HomeController) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application running..."))
}
