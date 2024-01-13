package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"pjm.dev/chinook/internal/db/model"
	"pjm.dev/chinook/internal/handlers"
)

func TestGetGenres(t *testing.T) {
	handler, seed, assert := getFooHandler(t)

	tests := []struct {
		name     string
		seed     map[any][]any
		request  *http.Request
		response *http.Response
		want     map[any][]any
	}{
		{
			name: "three genres",
			seed: map[any][]any{
				model.Genre{}: []model.Genre{
					{Name: "Foo"},
					{Name: "Bar"},
					{Name: "Baz"},
				},
			},
			request: httptest.NewRequest(http.MethodGet, "/genres", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"83"},
				},
				Body: io.NopCloser(strings.NewReader(`[{"GenreId":1,"Name":"Foo"},{"GenreId":2,"Name":"Bar"},{"GenreId":3,"Name":"Baz"}]`)),
			},
			want: map[any][]any{
				model.Genre{}: []model.Genre{
					{GenreID: 1, Name: "Foo"},
					{GenreID: 2, Name: "Bar"},
					{GenreID: 3, Name: "Baz"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			seed(t, test.seed)

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()
			if err := handlers.AssertSoftResponseEquality(t, test.response, response); err != nil {
				t.Errorf("response mismatch\n%v", err)
			}

			assert(t, test.want)
		})
	}
}

func TestGetGenre(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/genres/9", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"27"},
				},
				Body: io.NopCloser(strings.NewReader(`{"GenreId":9,"Name":"Pop"}`)),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if err := handlers.AssertSoftResponseEquality(t, test.response, response); err != nil {
				t.Errorf("response mismatch\n%v", err)
			}
		})
	}
}

func TestCreateGenre(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodPost, "/genres", strings.NewReader(`{"Name":"Junk"}`)),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"29"},
				},
				Body: io.NopCloser(strings.NewReader(`{"GenreId":26,"Name":"Junk"}`)),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if err := handlers.AssertSoftResponseEquality(t, test.response, response); err != nil {
				t.Errorf("response mismatch\n%v", err)
			}
		})
	}
}

func TestPatchGenre(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodPatch, "/genres/22", strings.NewReader(`{"Name":"Junk"}`)),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"29"},
				},
				Body: io.NopCloser(strings.NewReader(`{"GenreId":22,"Name":"Junk"}`)),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if err := handlers.AssertSoftResponseEquality(t, test.response, response); err != nil {
				t.Errorf("response mismatch\n%v", err)
			}
		})
	}
}

func TestDeleteGenre(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodDelete, "/genres/9", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"27"},
				},
				Body: io.NopCloser(strings.NewReader(`{"GenreId":9,"Name":"Pop"}`)),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if err := handlers.AssertSoftResponseEquality(t, test.response, response); err != nil {
				t.Errorf("response mismatch\n%v", err)
			}
		})
	}
}
