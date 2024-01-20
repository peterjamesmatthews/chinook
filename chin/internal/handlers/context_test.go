package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB = &gorm.DB{}

func TestGetNookFromContext(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want *gorm.DB
		err  error
	}{
		{
			name: "happy path",
			ctx:  context.WithValue(context.Background(), nook, testDB),
			want: testDB,
			err:  nil,
		},
		{
			name: "missing database in context",
			ctx:  context.Background(),
			want: nil,
			err:  ErrMissingDatabaseInContext,
		},
		{
			name: "unexpected nook type",
			ctx:  context.WithValue(context.Background(), nook, "not a *gorm.DB"),
			want: nil,
			err:  &ErrUnexpectedNookType{"not a *gorm.DB"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetNookFromContext(test.ctx)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}

			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("got %v, want %v", err, test.err)
			}
		})
	}
}

func TestHandleGettingNookFromContext(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		response http.Response
		want     *gorm.DB
		err      error
	}{
		{
			name:     "happy path",
			ctx:      context.WithValue(context.Background(), nook, testDB),
			want:     testDB,
			response: http.Response{},
			err:      nil,
		},
		{
			name: "failed to get database",
			ctx:  context.Background(),
			response: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("failed to get database\nmissing database in context")),
			},
			want: nil,
			err:  ErrMissingDatabaseInContext,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequestWithContext(test.ctx, http.MethodGet, "/", nil)
			got, err := handleGettingNookFromContext(w, r)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("got %v, want %v", err, test.err)
			}
			if err = AssertSoftResponseEquality(t, &test.response, w.Result()); err != nil {
				t.Error(err)
			}
		})
	}
}
