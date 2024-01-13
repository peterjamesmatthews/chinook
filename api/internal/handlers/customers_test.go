package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"pjm.dev/chinook/internal/handlers"
	"pjm.dev/chinook/test"
)

func TestGetCustomers(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/customers", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"15749"},
				},
				Body: test.OpenElseFatal(t, "/Users/pjm/Repositories/chinook/api/testdata/GetCustomers.json"),
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

func TestGetCustomer(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/customers/42", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"262"},
				},
				Body: io.NopCloser(strings.NewReader(`{"CustomerId":42,"FirstName":"Wyatt","LastName":"Girard","Company":"","Address":"9, Place Louis Barthou","City":"Bordeaux","State":"","Country":"France","PostalCode":"33000","Phone":"+33 05 56 96 96 96","Fax":"","Email":"wyatt.girard@yahoo.fr","SupportRepId":3}`)),
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

func TestCreateCustomer(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"FirstName":"Foo","LastName":"Bar","Email":"foobar@example.com"}`)),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"195"},
				},
				Body: io.NopCloser(strings.NewReader(`{"CustomerId":60,"FirstName":"Foo","LastName":"Bar","Company":"","Address":"","City":"","State":"","Country":"","PostalCode":"","Phone":"","Fax":"","Email":"foobar@example.com","SupportRepId":0}`)),
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

func TestPatchCustomer(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodPatch, "/customers/42", strings.NewReader(`{"FirstName":"Foo"}`)),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"260"},
				},
				Body: io.NopCloser(strings.NewReader(`{"CustomerId":42,"FirstName":"Foo","LastName":"Girard","Company":"","Address":"9, Place Louis Barthou","City":"Bordeaux","State":"","Country":"France","PostalCode":"33000","Phone":"+33 05 56 96 96 96","Fax":"","Email":"wyatt.girard@yahoo.fr","SupportRepId":3}`)),
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

func TestDeleteCustomer(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodDelete, "/customers/42", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"CustomerId":42,"FirstName":"Wyatt","LastName":"Girard","Company":"","Address":"9, Place Louis Barthou","City":"Bordeaux","State":"","Country":"France","PostalCode":"33000","Phone":"+33 05 56 96 96 96","Fax":"","Email":"wyatt.girard@yahoo.fr","SupportRepId":3}`)),
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
