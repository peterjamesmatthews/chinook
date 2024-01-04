package http

import (
	"context"
	"errors"
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

var ErrMissingDatabaseInContext = errors.New("missing database in context")

type ErrUnexpectedChinookType struct {
	val any
}

func (e ErrUnexpectedChinookType) Error() string {
	return fmt.Sprintf("expected *gorm.DB in context, got %T", e.val)
}

// GetChinookFromContext tries to retrieve the chinook database from ctx.
//
// # Errors
//   - ErrMissingDatabaseInContext: if the database is not present in the context
func GetChinookFromContext(ctx context.Context) (*gorm.DB, error) {
	val := ctx.Value(chinook)
	if val == nil {
		return nil, ErrMissingDatabaseInContext
	}

	db, ok := val.(*gorm.DB)
	if !ok {
		return nil, &ErrUnexpectedChinookType{val}
	}

	return db, nil
}

// GetContextWithChinook returns a new context with the chinook database set to
// a key that can be used by GetChinookFromContext later.
func GetContextWithChinook(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, chinook, db)
}

// handleGettingChinookFromContext is a helper function that gets the chinook
// database from the context and writes any errors to the response.
//
//	chinook, err := handleGettingChinookFromContext(w, r)
//	if err != nil {
//		return
//	}
func handleGettingChinookFromContext(w http.ResponseWriter, r *http.Request) (*gorm.DB, error) {
	chinook, err := GetChinookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database\n%w", err).Error()))
		return nil, err
	}

	// do something with chinook
	return chinook, nil
}
