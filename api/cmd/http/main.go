package main

import (
	"log"
	"net/http"
)

// main creates a chinook router, registers the "/" path to be handled by it,
// and listens on port 3000.
func main() {
	router := NewChinookRouter()

	http.Handle("/", router)

	log.Default().Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
