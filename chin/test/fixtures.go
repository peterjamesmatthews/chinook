package test

import (
	"os"
	"path/filepath"
	"testing"

	"pjm.dev/chin/internal/meta"
)

// OpenFixture opens the fixture at the path and returns a file descriptor.
//
// If fixturePath is not absolute, it will be rooted in the fixtures directory.
//
// If an error occurs opening the file, the test will fail fatally.
func OpenFixture(t *testing.T, fixturePath string) *os.File {
	t.Helper()

	if !filepath.IsAbs(fixturePath) {
		fixturePath = filepath.Join(meta.Root, "test", "fixtures", fixturePath)
	}

	file, err := os.Open(fixturePath)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	return file
}
