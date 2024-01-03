package http_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

// assertSoftResponseEquality compares two responses, want and got, and returns
// and error if any field of got is not equal to the corresponding field of
// want.
//
// The comparison is "soft" by the fact that if a field of want is the zero
// value for that type, then the corresponding field of got is allowed to be
// any value.
func assertSoftResponseEquality(t *testing.T, want, got *http.Response) error {
	t.Helper()

	if err := assertNilParity(t, want, got); err != nil {
		return fmt.Errorf("want & got do not have nil parity\n%w", err)
	}

	if want == nil { // got is nil since we already checked for nil parity
		return nil
	}

	if want.StatusCode != 0 {
		if err := assertDeepEquality(t, want.StatusCode, got.StatusCode); err != nil {
			return fmt.Errorf("status code mismatch\n%w", err)
		}
	}

	if want.Header != nil {
		if err := assertDeepEquality(t, want.Header, got.Header); err != nil {
			return fmt.Errorf("header mismatch\n%w", err)
		}
	}

	if want.Body != nil {
		if err := assertTrimmedReaderEquality(t, want.Body, got.Body); err != nil {
			return fmt.Errorf("body mismatch\n%w", err)
		}
	}

	// TODO check other fields of http.Response

	return nil
}

// assertNilParity compares two values, want and got, and returns an error if
// one is nil and the other is not.
func assertNilParity(t *testing.T, want, got any) error {
	t.Helper()

	if want == nil && got != nil {
		return errors.New("want is nil, but got is not")
	}
	if got == nil && want != nil {
		return errors.New("got is nil, but want is not")
	}

	return nil
}

func assertDeepEquality(t *testing.T, want, got any) error {
	t.Helper()

	equal := reflect.DeepEqual(want, got)
	if !equal {
		return fmt.Errorf("want %v\n got %v", want, got)
	}

	return nil
}

func assertTrimmedReaderEquality(t *testing.T, wantReader, gotReader io.Reader) error {
	t.Helper()

	wantBytes, err := io.ReadAll(wantReader)
	if err != nil {
		t.Fatalf("failed to read want: %v", err)
	}

	gotBytes, err := io.ReadAll(gotReader)
	if err != nil {
		t.Fatalf("failed to read got: %v", err)
	}

	want := strings.TrimSpace(string(wantBytes))
	got := strings.TrimSpace(string(gotBytes))

	if want != got {
		t.Fatalf("want %v\n got %v", want, got)
	}

	return nil
}

func assertSoftEquality(t *testing.T, want, got any) error {
	t.Helper()

	// TODO want == nil isn't safe, figure this out

	if reflect.ValueOf(want) == reflect.Zero(reflect.TypeOf(want)).Interface() {
		return nil
	}

	equal := reflect.DeepEqual(want, got)
	if !equal {
		return fmt.Errorf("want %v\n got %v", want, got)
	}

	return nil
}
