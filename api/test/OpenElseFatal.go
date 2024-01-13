package test

import (
	"os"
	"testing"
)

// OpenElseFatal gets a *os.File with os.Open and returns it.
//
// If an error occurs, the test will fail fatally.
func OpenElseFatal(t *testing.T, path string) *os.File {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	return file
}
