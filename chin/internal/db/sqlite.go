package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetSQLite opens a SQLite database with dsn and returns a *gorm.DB.
//
// gorm logs are disabled.
func GetSQLite(dsn string) (*gorm.DB, error) {
	chinook, err := gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // disable logging
		},
	)
	if err != nil {
		return nil, err
	}

	return chinook, nil
}
