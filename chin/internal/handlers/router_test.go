package handlers_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"pjm.dev/chin/internal/crow"
	"pjm.dev/chin/internal/handlers"
	"pjm.dev/chin/internal/nook"
	"pjm.dev/chin/test"
)

// testNook is a seeded nook database that can be used for testing.
//
// It will used many times by GetTestHandler. The first time GetTestHandle is
// called testNook will be seeded, which takes some time. Subsequent calls to
// GetTestHandler will reuse testNook.
var testNook *gorm.DB

// getTestHandler returns a http.Handler that can be used for testing.
//
// The handler...
//   - contains all routes registered by RegisterRoutes.
//   - has a seeded nook database in the context.
//   - operates in a transaction that's rolled back after the request is handled.
func getTestHandler(t *testing.T) http.Handler {
	router := mux.NewRouter()
	handlers.RegisterRoutes(router)
	var handler http.Handler = router

	if testNook == nil { // initialize testNook
		seed := test.OpenFixture(t, "Chinook.sqlite")
		defer seed.Close()

		var err error
		testNook, err = nook.GetSQLite(seed.Name())
		if err != nil {
			t.Fatalf("failed to seed testNook: %v", err)
		}
	}

	handler = wrapInNookTransaction(t, router)
	handler = handlers.WrapWithNookInContext(handler, testNook)
	return handler
}

// wrapInNookTransaction wraps a handler that was previously wrapped with
// handlers.WrapWithNookInContext in a transaction.
//
// The transaction begins before calling handler.ServeHTTP and rolls back after.
//
// This is useful for testing handlers that modify the database, so that
// different tests do not interfere with each other.
func wrapInNookTransaction(t *testing.T, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := handlers.GetNookFromContext(r.Context())
		if err != nil {
			t.Fatalf("failed to get nook from context: %v", err)
		}
		tx := db.Begin()
		ctx := handlers.GetContextWithNook(r.Context(), tx)
		handler.ServeHTTP(w, r.WithContext(ctx))
		tx.Rollback()
	})
}

// unseededTestNook is an unseeded nook database that can be used for testing.
var unseededTestNook *gorm.DB

func getCrowHandler(t *testing.T) (http.Handler, *crow.Crow) {
	router := mux.NewRouter()
	handlers.RegisterRoutes(router)
	var h http.Handler = router

	if unseededTestNook == nil {
		var err error
		unseededTestNook, err = nook.GetSQLite(":memory:")
		if err != nil {
			t.Fatalf("failed to get in-memory sqlite db: %v", err)
		}
	}

	h, c := crow.WrapInCrow(h, func(r *http.Request) *gorm.DB {
		db, err := handlers.GetNookFromContext(r.Context())
		if err != nil {
			t.Fatalf("failed to get nook from context: %v", err)
		}
		return db
	})
	h = wrapInNookTransaction(t, h)
	h = handlers.WrapWithNookInContext(h, unseededTestNook)

	return h, c
}
