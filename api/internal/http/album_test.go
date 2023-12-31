package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/gorm"
	chinookHTTP "pjm.dev/chinook/internal/http"
	"pjm.dev/chinook/testdata"

	"github.com/gorilla/mux"
)

func TestAlbumRoutes(t *testing.T) {
	r := mux.NewRouter()
	chinookHTTP.RegisterChinookRoutes(r)
	tests := []struct {
		name    string
		request *http.Request
	}{
		{
			name:    "GET /albums",
			request: httptest.NewRequest(http.MethodGet, "/albums", nil),
		},
		{
			name:    "GET /albums/1",
			request: httptest.NewRequest(http.MethodGet, "/albums/1", nil),
		},
		{
			name:    "POST /albums",
			request: httptest.NewRequest(http.MethodPost, "/albums", nil),
		},
		{
			name:    "PATCH /albums/1",
			request: httptest.NewRequest(http.MethodPatch, "/albums/1", nil),
		},
		{
			name:    "DELETE /albums/1",
			request: httptest.NewRequest(http.MethodDelete, "/albums/1", nil),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			match := mux.RouteMatch{}
			if !r.Match(test.request, &match) {
				t.Errorf("failed to match request %v %v", test.request, match.MatchErr)
			}
		})
	}
}

func TestGetAlbums(t *testing.T) {
	handler := GetTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		seed     *gorm.DB
		response *http.Response
		db       struct{}
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
				Body: testdata.OpenElseFatal(t, "/Users/pjm/Repositories/chinook/api/testdata/GetAlbums.json"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if err := AssertSoftResponseEquality(t, test.response, response); err != nil {
				t.Errorf("response mismatch\n%v", err)
			}
		})
	}
}
