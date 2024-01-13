package handlers_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"pjm.dev/chinook/internal/handlers"
	"pjm.dev/chinook/test"
)

func TestGetAlbums(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/albums", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"22196"},
				},
				Body: test.OpenFixture(t, "GetAlbums.json"),
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

func TestGetAlbum(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/albums/243", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"71"},
				},
				Body: io.NopCloser(bytes.NewReader([]byte(
					`{"AlbumId":243,"Title":"The Best Of Van Halen, Vol. I","ArtistId":152}`,
				))),
			},
		},
		{
			name:    "not found",
			request: httptest.NewRequest(http.MethodGet, "/albums/9999", nil),
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     http.Header{},
				Body:       io.NopCloser(bytes.NewReader([]byte(`album 9999 not found`))),
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

func TestCreateAlbum(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name: "happy path",
			request: httptest.NewRequest(
				http.MethodPost,
				"/albums",
				bytes.NewReader([]byte(`{"Title":"My Cool Album","ArtistId":152}`)),
			),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"55"},
				},
				Body: io.NopCloser(bytes.NewReader([]byte(
					`{"AlbumId":348,"Title":"My Cool Album","ArtistId":152}`,
				))),
			},
		},
		{
			name:    "invalid body",
			request: httptest.NewRequest(http.MethodPost, "/albums", bytes.NewReader([]byte(`Foobar`))),
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Header:     http.Header{},
				Body:       io.NopCloser(bytes.NewReader([]byte(`failed to decode album: invalid character 'F' looking for beginning of value`))),
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

func TestPatchAlbum(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name: "happy path",
			request: httptest.NewRequest(
				http.MethodPatch,
				"/albums/243",
				bytes.NewReader([]byte(`{"Title":"My Cool Album"}`)),
			),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"55"},
				},
				Body: io.NopCloser(bytes.NewReader([]byte(
					`{"AlbumId":243,"Title":"My Cool Album","ArtistId":152}`,
				))),
			},
		},
		{
			name:    "not found",
			request: httptest.NewRequest(http.MethodPatch, "/albums/9999", nil),
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     http.Header{},
				Body:       io.NopCloser(bytes.NewReader([]byte(`album 9999 not found`))),
			},
		},
		{
			name:    "invalid body",
			request: httptest.NewRequest(http.MethodPatch, "/albums/243", bytes.NewReader([]byte(`Foobar`))),
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Header:     http.Header{},
				Body:       io.NopCloser(bytes.NewReader([]byte(`failed to decode album: invalid character 'F' looking for beginning of value`))),
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

func TestDeleteAlbum(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodDelete, "/albums/243", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"71"},
				},
				Body: io.NopCloser(bytes.NewReader([]byte(
					`{"AlbumId":243,"Title":"The Best Of Van Halen, Vol. I","ArtistId":152}`,
				))),
			},
		},
		{
			name:    "not found",
			request: httptest.NewRequest(http.MethodDelete, "/albums/9999", nil),
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     http.Header{},
				Body:       io.NopCloser(bytes.NewReader([]byte(`album 9999 not found`))),
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
