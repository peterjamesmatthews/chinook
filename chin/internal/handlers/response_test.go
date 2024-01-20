package handlers_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"pjm.dev/chin/internal/handlers"
)

func TestWriteJSONToResponse(t *testing.T) {
	tests := []struct {
		name    string
		payload any
		want    *http.Response
		err     error
	}{
		{
			name:    "happy path",
			payload: "hello world",
			want: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Length": []string{"14"},
					"Content-Type":   []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`"hello world"`)),
			},
		},
		{
			name:    "doesn't escape html",
			payload: "Foo & Bar",
			want: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Length": []string{"12"},
					"Content-Type":   []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`"Foo & Bar"`)),
			},
		},
		{
			name:    "error marshalling payload",
			payload: func() {},
			err:     fmt.Errorf("failed to marshal payload\njson: unsupported type: func()"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := handlers.WriteJSONToResponse(w, test.payload)
			got := w.Result()

			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("unexpected error\n got = %v\nwant %v", err, test.err)
			}
			err = handlers.AssertSoftResponseEquality(t, test.want, got)
			if err != nil {
				t.Error(fmt.Errorf("response mismatch\n%w", err))
			}
		})
	}
}

func TestHandleWritingJSONToResponse(t *testing.T) {
	tests := []struct {
		name    string
		payload any
		want    *http.Response
		err     error
	}{
		{
			name:    "happy path",
			payload: "hello world",
			want: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Length": []string{"14"},
					"Content-Type":   []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`"hello world"`)),
			},
		},
		{
			name:    "error writing payload to response",
			payload: func() {},
			want: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header:     http.Header{},
				Body:       io.NopCloser(strings.NewReader("failed to write payload to response\nfailed to marshal payload\njson: unsupported type: func()")),
			},
			err: fmt.Errorf("failed to marshal payload\njson: unsupported type: func()"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			gotErr := handlers.HandleWritingJSONToResponse(w, test.payload)
			got := w.Result()

			err := handlers.AssertNilParity(t, test.err, gotErr)
			if err != nil {
				t.Error(err)
			}

			if gotErr != nil && gotErr.Error() != test.err.Error() {
				t.Errorf("\nunexpected error\n got = %v\nwant %v", gotErr, test.err)
			}

			gotErr = handlers.AssertSoftResponseEquality(t, test.want, got)
			if gotErr != nil {
				t.Error(fmt.Errorf("response mismatch\n%w", gotErr))
			}
		})
	}
}
