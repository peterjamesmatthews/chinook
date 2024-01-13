package handlers_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"pjm.dev/chinook/internal/db"
	"pjm.dev/chinook/internal/handlers"
	"pjm.dev/chinook/test"
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

	handler = wrapInTransaction(t, router)
	handler = handlers.WrapWithChinookInContext(handler, testChinook)
	return handler
}

// wrapInTransaction wraps a handler that was previously wrapped with
// handlers.WrapWithChinookInContext in a transaction.
//
// The transaction begins before calling handler.ServeHTTP and rolls back after.
//
// This is useful for testing handlers that modify the database, so that
// different tests do not interfere with each other.
func wrapInTransaction(t *testing.T, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chinook, err := handlers.GetChinookFromContext(r.Context())
		if err != nil {
			t.Fatalf("failed to get chinook from context: %v", err)
		}
		tx := chinook.Begin()
		defer tx.Rollback()
		ctx := handlers.GetContextWithChinook(r.Context(), tx)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// unseededTestChinook is an unseeded chinook database that can be used for testing.
var unseededTestChinook *gorm.DB

type foo func(*testing.T, map[any][]any)

func getFooHandler(t *testing.T) (http.Handler, foo, foo) {
	router := mux.NewRouter()
	handlers.RegisterChinookRoutes(router)
	var handler http.Handler = router

	if unseededTestChinook == nil {
		var err error
		unseededTestChinook, err = db.GetSQLite(":memory:")
		if err != nil {
			t.Fatalf("failed to get in-memory sqlite db: %v", err)
		}
	}

	// create function that will seed the database
	seedDB := func(t *testing.T, seed map[any][]any) {
		// migrate schema
		var models []any
		for model := range seed {
			models = append(models, model)
		}

		if err := unseededTestChinook.AutoMigrate(models...); err != nil {
			t.Fatalf("failed to migrate schema: %v", err)
		}

		// insert data
		for _, records := range seed {
			if err := unseededTestChinook.Create(records).Error; err != nil {
				t.Fatalf("failed to insert data: %v", err)
			}
		}
	}

	// create function that will assert the database
	assertDB := func(t *testing.T, want map[any][]any) {

		// for each model in want
		for model, wantRecords := range want {
			// get records from database
			var gotRecords []any // TODO this should be a slice of model's type
			if err := unseededTestChinook.Find(&gotRecords).Error; err != nil {
				t.Fatalf("failed to get records from database: %v", err) // TODO handle no rows
			}

			// assert that the records match want
			// TODO need a way to compare wantRecords to gotRecords
			// ? generic soft unordered equality
		}

		t.Fatal("not implemented")
		return
	}

	handler = wrapInTransaction(t, router)
	handler = handlers.WrapWithChinookInContext(handler, unseededTestChinook)

	return handler, seedDB, assertDB
}
