package controllers

import "net/http"

/*Home home controller */
type Home struct{}

/*Get returns information about the application */
func (Controller *Home) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application running..."))
}
