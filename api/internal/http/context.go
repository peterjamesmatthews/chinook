package http

import (
	"context"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type contextKey struct{}

// chinook is the key used to store the chinook database in the context.
var chinook contextKey = contextKey{}

// WrapWithChinookInContext returns handler with chinook set in the context.
//
// chinook should be retrieved from the context using GetChinookFromContext.
//
//	 func handler(w http.ResponseWriter, r *http.Request) {
//		  chinook, err := GetChinookFromContext(r.Context())
//			/* use chinook... */
//		}
func WrapWithChinookInContext(handler http.Handler, chinook *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContextWithChinook(r.Context(), chinook)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetChinookFromContext tries to retrieve the chinook database from ctx.
//
// # Errors
//   - ErrMissingDatabaseInContext: if the database is not present in the context
func GetChinookFromContext(ctx context.Context) (*gorm.DB, error) {
	val := ctx.Value(chinook)
	var nilDB *gorm.DB
	if val == nilDB {
		return nil, fmt.Errorf("missing database in context")
	}

	db, ok := val.(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("expected *gorm.DB in context, got %T", val)
	}

	return db, nil
}

// GetContextWithChinook returns a new context with the chinook database set to
// a key that can be used by GetChinookFromContext later.
func GetContextWithChinook(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, chinook, db)
}
