package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	chinookHTTP "pjm.dev/chinook/internal/http"
)

// main creates a chinook router, registers the "/" path to be handled by it,
// and listens on port 3000.
func main() {
	r := mux.NewRouter()
	chinookHTTP.RegisterChinookRoutes(r)
	http.Handle("/", r)
	log.Default().Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
