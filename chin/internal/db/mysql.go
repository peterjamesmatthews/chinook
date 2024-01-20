package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetMySQL opens a MySQL database with dsn and returns a *gorm.DB.
func GetMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database\n%w", err)
	}

	return db, nil
}
