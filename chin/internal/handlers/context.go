package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type contextKey struct{}

// nook is the key used to store the nook database in the context.
var nook contextKey = contextKey{}

// WrapWithNookInContext returns handler with nook set in the context.
//
// nook should be retrieved from the context using GetNookFromContext.
//
//	 func handler(w http.ResponseWriter, r *http.Request) {
//		  db, err := GetNookFromContext(r.Context())
//			/* use db... */
//		}
func WrapWithNookInContext(handler http.Handler, db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContextWithNook(r.Context(), db)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

var ErrMissingDatabaseInContext = errors.New("missing database in context")

type ErrUnexpectedNookType struct{ unexpectedValue any }

func (e ErrUnexpectedNookType) Error() string {
	return fmt.Sprintf("expected *gorm.DB in context, got %T", e.unexpectedValue)
}

// GetNookFromContext tries to retrieve the nook database from ctx.
//
// # Errors
//   - ErrMissingDatabaseInContext: if the database is not present in the context
func GetNookFromContext(ctx context.Context) (*gorm.DB, error) {
	val := ctx.Value(nook)
	if val == nil {
		return nil, ErrMissingDatabaseInContext
	}

	db, ok := val.(*gorm.DB)
	if !ok {
		return nil, &ErrUnexpectedNookType{val}
	}

	return db, nil
}

// GetContextWithNook returns a new context with the nook database set to
// a key that can be used by GetNookFromContext later.
func GetContextWithNook(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, nook, db)
}

// handleGettingNookFromContext is a helper function that gets the nook
// database from the context and writes any errors to the response.
//
//	db, err := handleGettingNookFromContext(w, r)
//	if err != nil {
//		return
//	}
func handleGettingNookFromContext(w http.ResponseWriter, r *http.Request) (*gorm.DB, error) {
	db, err := GetNookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database\n%w", err).Error()))
		return nil, err
	}
	return db, nil
}
