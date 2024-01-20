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

	// register all routes
	handlers.RegisterRoutes(router)

	// get dsn that connects to nook
	dsn, err := nook.GetDSN()
	if err != nil {
		log.Fatalf("failed to get dsn\n%v", err)
	}

	// get a db database connection
	db, err := nook.GetMySQL(dsn)
	if err != nil {
		log.Fatalf("failed to get nook database with dsn %s\n%v", dsn, err)
	}

	// wrap router with nook db in context
	handler := handlers.WrapWithNookInContext(router, db)

	// start the server
	http.Handle("/", handler)
	log.Default().Println("Listening on port 3000")
	http.ListenAndServe(":3000", handler)
}
