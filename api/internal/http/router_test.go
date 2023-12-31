package http_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"pjm.dev/chinook/internal/db"
	chinookHTTP "pjm.dev/chinook/internal/http"
)

// testChinook is a seeded chinook database that can be used for testing.
//
// It will used many times by GetTestHandler. The first time GetTestHandler is
// called testChinook will be seeded, which takes some time. Subsequent calls to
// GetTestHandler will reuse testChinook.
var testChinook *gorm.DB

// GetTestHandler returns a http.Handler that can be used for testing.
//
// The handler...
//   - contains all routes registered by RegisterChinookRoutes.
//   - has a seeded chinook database in the context.
//   - operates in a transaction that's rolled back after the request is handled.
func GetTestHandler(t *testing.T) http.Handler {
	router := mux.NewRouter()
	chinookHTTP.RegisterChinookRoutes(router)
	var handler http.Handler = router

	if testChinook == nil {
		var err error
		testChinook, err = db.GetSQLite("/Users/pjm/Repositories/chinook/api/testdata/Chinook.sqlite") // TODO get path from testdata
		if err != nil {
			t.Fatalf("failed to get sqlite database: %v", err)
		}
	}

	handler = WrapInTransaction(t, router)
	handler = chinookHTTP.WrapWithChinookInContext(handler, testChinook)
	return handler
}

// WrapInTransaction wraps a handler that was previously wrapped with
// chinookHTTP.WrapWithChinookInContext in a transaction.
//
// The transaction begins before calling handler.ServeHTTP and rolls back after.
//
// This is useful for testing handlers that modify the database, so that
// different tests do not interfere with each other.
func WrapInTransaction(t *testing.T, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chinook, err := chinookHTTP.GetChinookFromContext(r.Context())
		if err != nil {
			t.Fatalf("failed to get chinook from context: %v", err)
		}
		tx := chinook.Begin()
		defer tx.Rollback()
		ctx := chinookHTTP.GetContextWithChinook(r.Context(), tx)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
