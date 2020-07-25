package router

import (
	"net/http"
)

// Routes routes application
func Routes(subscribe func(string, func(w http.ResponseWriter, r *http.Request))) {
	subscribe("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Application running..."))
	})
}
