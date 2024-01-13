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

func TestGetEmployees(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/employees", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"2838"},
				},
				Body: test.OpenElseFatal(t, "/Users/pjm/Repositories/chinook/api/testdata/GetEmployees.json"),
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

func TestGetEmployee(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodGet, "/employees/1", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"358"},
				},
				Body: io.NopCloser(strings.NewReader(`{"EmployeeId":1,"LastName":"Adams","FirstName":"Andrew","Title":"General Manager","ReportsTo":0,"BirthDate":"1962-02-18T00:00:00Z","HireDate":"2002-08-14T00:00:00Z","Address":"11120 Jasper Ave NW","City":"Edmonton","State":"AB","Country":"Canada","PostalCode":"T5K 2N1","Phone":"+1 (780) 428-9482","Fax":"+1 (780) 428-3457","Email":"andrew@chinookcorp.com"}`)),
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

func TestCreateEmployee(t *testing.T) {
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
				"/employees",
				strings.NewReader(`{"LastName":"Bar","FirstName":"Fooby","Title":"Software Engineer","ReportsTo":0,"BirthDate":"1962-02-18T00:00:00Z","HireDate":"2002-08-14T00:00:00Z","Address":"11120 Jasper Ave NW","City":"Edmonton","State":"AB","Country":"Canada","PostalCode":"T5K 2N1","Phone":"+1 (780) 428-9482","Fax":"+1 (780) 428-3457","Email":"andrew@chinookcorp.com"}`),
			),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"357"},
				},
				Body: io.NopCloser(strings.NewReader(`{"EmployeeId":9,"LastName":"Bar","FirstName":"Fooby","Title":"Software Engineer","ReportsTo":0,"BirthDate":"1962-02-18T00:00:00Z","HireDate":"2002-08-14T00:00:00Z","Address":"11120 Jasper Ave NW","City":"Edmonton","State":"AB","Country":"Canada","PostalCode":"T5K 2N1","Phone":"+1 (780) 428-9482","Fax":"+1 (780) 428-3457","Email":"andrew@chinookcorp.com"}`)),
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

func TestPatchEmployee(t *testing.T) {
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
				"/employees/1",
				strings.NewReader(`{"LastName":"Bar","FirstName":"Fooby"}`),
			),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"355"},
				},
				Body: io.NopCloser(strings.NewReader(`{"EmployeeId":1,"LastName":"Bar","FirstName":"Fooby","Title":"General Manager","ReportsTo":0,"BirthDate":"1962-02-18T00:00:00Z","HireDate":"2002-08-14T00:00:00Z","Address":"11120 Jasper Ave NW","City":"Edmonton","State":"AB","Country":"Canada","PostalCode":"T5K 2N1","Phone":"+1 (780) 428-9482","Fax":"+1 (780) 428-3457","Email":"andrew@chinookcorp.com"}`)),
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

func TestDeleteEmployee(t *testing.T) {
	handler := getTestHandler(t)

	tests := []struct {
		name     string
		request  *http.Request
		response *http.Response
	}{
		{
			name:    "happy path",
			request: httptest.NewRequest(http.MethodDelete, "/employees/4", nil),
			response: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type":   []string{"application/json"},
					"Content-Length": []string{"361"},
				},
				Body: io.NopCloser(strings.NewReader(`{"EmployeeId":4,"LastName":"Park","FirstName":"Margaret","Title":"Sales Support Agent","ReportsTo":2,"BirthDate":"1947-09-19T00:00:00Z","HireDate":"2003-05-03T00:00:00Z","Address":"683 10 Street SW","City":"Calgary","State":"AB","Country":"Canada","PostalCode":"T2P 5G3","Phone":"+1 (403) 263-4423","Fax":"+1 (403) 263-4289","Email":"margaret@chinookcorp.com"}`)),
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
