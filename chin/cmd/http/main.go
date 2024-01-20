package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"pjm.dev/chin/internal/handlers"
	"pjm.dev/chin/internal/nook"
)

func main() {
	// create a mux router
	router := mux.NewRouter()

	// register all chinook routes
	handlers.RegisterChinookRoutes(router)

	// get dsn that connects to chinook database
	dsn, err := nook.GetDSN()
	if err != nil {
		log.Fatalf("failed to get dsn\n%v", err)
	}

	// get a chinook database connection
	chinook, err := nook.GetMySQL(dsn)
	if err != nil {
		log.Fatalf("failed to get chinook database\n%v", err)
	}

	// wrap router with chinook in context
	handler := handlers.WrapWithChinookInContext(router, chinook)

	// start the server
	http.Handle("/", handler)
	log.Default().Println("Listening on port 3000")
	http.ListenAndServe(":3000", handler)
}
