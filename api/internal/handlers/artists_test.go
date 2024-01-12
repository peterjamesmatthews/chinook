package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"pjm.dev/chinook/internal/handlers"
	"pjm.dev/chinook/testdata"
)

func TestGetArtists(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/artists", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"13012"},
				},
				Body: testdata.OpenElseFatal(t, "/Users/pjm/Repositories/chinook/api/testdata/GetArtists.json"),
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

func TestGetArtist(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/artists/237", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"65"},
				},
				Body: io.NopCloser(strings.NewReader(`{"ArtistId":237,"Name":"Berliner Philharmoniker & Hans Rosbaud"}`)),
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

func TestCreateArtist(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodPost, "/artists", strings.NewReader(`{"Name":"Foo"}`)),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"30"},
				},
				Body: io.NopCloser(strings.NewReader(`{"ArtistId":276,"Name":"Foo"}`)),
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

func TestPatchArtist(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodPatch, "/artists/237", strings.NewReader(`{"Name":"Bar"}`)),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"30"},
				},
				Body: io.NopCloser(strings.NewReader(`{"ArtistId":237,"Name":"Bar"}`)),
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

func TestDeleteArtist(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodDelete, "/artists/237", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"65"},
				},
				Body: io.NopCloser(strings.NewReader(`{"ArtistId":237,"Name":"Berliner Philharmoniker & Hans Rosbaud"}`)),
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
