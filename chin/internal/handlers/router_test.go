package handlers_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"pjm.dev/chin/internal/crow"
	"pjm.dev/chin/internal/db"
	"pjm.dev/chin/internal/handlers"
	"pjm.dev/chin/test"
)

// testChinook is a seeded chinook database that can be used for testing.
//
// It will used many times by GetTestHandler. The first time GetTestHandler is
// called testChinook will be seeded, which takes some time. Subsequent calls to
// GetTestHandler will reuse testChinook.
var testChinook *gorm.DB

// getTestHandler returns a http.Handler that can be used for testing.
//
// The handler...
//   - contains all routes registered by RegisterChinookRoutes.
//   - has a seeded chinook database in the context.
//   - operates in a transaction that's rolled back after the request is handled.
func getTestHandler(t *testing.T) http.Handler {
	router := mux.NewRouter()
	handlers.RegisterChinookRoutes(router)
	var handler http.Handler = router

	if testChinook == nil { // initialize testChinook
		seed := test.OpenFixture(t, "Chinook.sqlite")
		defer seed.Close()

		var err error
		testChinook, err = db.GetSQLite(seed.Name())
		if err != nil {
			t.Fatalf("failed to seed testChinook: %v", err)
		}
	}

	handler = wrapInChinookTransaction(t, router)
	handler = handlers.WrapWithChinookInContext(handler, testChinook)
	return handler
}

// wrapInChinookTransaction wraps a handler that was previously wrapped with
// handlers.WrapWithChinookInContext in a transaction.
//
// The transaction begins before calling handler.ServeHTTP and rolls back after.
//
// This is useful for testing handlers that modify the database, so that
// different tests do not interfere with each other.
func wrapInChinookTransaction(t *testing.T, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chinook, err := handlers.GetChinookFromContext(r.Context())
		if err != nil {
			t.Fatalf("failed to get chinook from context: %v", err)
		}
		tx := chinook.Begin()
		ctx := handlers.GetContextWithChinook(r.Context(), tx)
		handler.ServeHTTP(w, r.WithContext(ctx))
		tx.Rollback()
	})
}

// unseededTestChinook is an unseeded chinook database that can be used for testing.
var unseededTestChinook *gorm.DB

func getCrowHandler(t *testing.T) (http.Handler, *crow.Crow) {
	router := mux.NewRouter()
	handlers.RegisterChinookRoutes(router)
	var h http.Handler = router

	if unseededTestChinook == nil {
		var err error
		unseededTestChinook, err = db.GetSQLite(":memory:")
		if err != nil {
			t.Fatalf("failed to get in-memory sqlite db: %v", err)
		}
	}

	h, c := crow.WrapInCrow(h, func(r *http.Request) *gorm.DB {
		db, err := handlers.GetChinookFromContext(r.Context())
		if err != nil {
			t.Fatalf("failed to get chinook from context: %v", err)
		}
		return db
	})
	h = wrapInChinookTransaction(t, h)
	h = handlers.WrapWithChinookInContext(h, unseededTestChinook)

	return h, c
}
