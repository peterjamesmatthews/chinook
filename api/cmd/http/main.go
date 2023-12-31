package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"pjm.dev/chinook/internal/db"
	chinookHTTP "pjm.dev/chinook/internal/http"
)

// main creates a chinook router, registers the "/" path to be handled by it,
// and listens on port 3000.
func main() {
	// create a mux router
	router := mux.NewRouter()

	// register all chinook routes
	chinookHTTP.RegisterChinookRoutes(router)

	// TODO move me to somewhere else
	rootPassword, ok := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	if !ok {
		log.Fatal("environment variable MYSQL_ROOT_PASSWORD not set")
	}

	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/Chinook", rootPassword)

	// get a chinook database connection
	chinook, err := db.GetMySQL(dsn)
	if err != nil {
		log.Fatalf("failed to get chinook database\n%v", err)
	}

	// wrap router with chinook in context
	handler := chinookHTTP.WrapWithChinookInContext(router, chinook)

	// start the server
	http.Handle("/", handler)
	log.Default().Println("Listening on port 3000")
	http.ListenAndServe(":3000", handler)
}
