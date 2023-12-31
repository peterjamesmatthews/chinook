package db

import (
	"io"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pjm.dev/chinook/testdata"
)

// GetSQLite returns a *gorm.DB that's connected to an in-memory SQLite
// database.
//
// gorm logs are disabled.
func GetSQLiteInMemory() (*gorm.DB, error) {
	chinook, err := gorm.Open(
		sqlite.Open("file::memory:"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // disable logging
		},
	)
	if err != nil {
		return nil, err
	}

	return chinook, nil
}

// SeedTestDatabaseWithChinook executes a SQL script on chinook that seeds it
// with the Chinook database.
func SeedTestDatabaseWithChinook(t *testing.T, chinook *gorm.DB) error {
	t.Helper()

	chinookFile := testdata.OpenElseFatal(t, "/Users/pjm/Repositories/chinook/api/testdata/Chinook.sql") // TODO construct using environment variable
	chinookBytes, err := io.ReadAll(chinookFile)
	if err != nil {
		return err
	}

	err = chinook.Exec(string(chinookBytes)).Error
	if err != nil {
		return err
	}

	return nil
}
