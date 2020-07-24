package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

// Init start serve
func Init() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Print("PORT: ", PORT)
	router := mux.NewRouter()
	router.HandleFunc("/", test)
	handler := cors.AllowAll().Handler(router)
	log.Print("Application: localhost:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
