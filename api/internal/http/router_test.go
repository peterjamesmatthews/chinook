package http_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"pjm.dev/chinook/internal/db"
	chinookHTTP "pjm.dev/chinook/internal/http"
)

// GetTestHandler returns a http.Handler that can be used for testing.
//
// The handler contains all routes registered by RegisterChinookRoutes.
//
// The handler contains a seeded chinook database in the context.
//
// If the database cannot be seeded, the test will fail fatally.
func GetTestHandler(t *testing.T) http.Handler {
	router := mux.NewRouter()
	chinookHTTP.RegisterChinookRoutes(router)
	chinook, err := db.GetSQLiteInMemory()
	if err != nil {
		t.Fatalf("failed to get sqlite database: %v", err)
	}
	err = db.SeedTestDatabaseWithChinook(t, chinook)
	if err != nil {
		t.Fatalf("failed to seed test database: %v", err)
	}
	return chinookHTTP.WrapWithChinookInContext(router, chinook)
}
