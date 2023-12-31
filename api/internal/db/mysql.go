package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetMySQLFromEnvironment returns a *gorm.DB that's connected to a MySQL
// database.
//
// The DSN is constructed from the environment variable MYSQL_ROOT_PASSWORD. All
// other values are set to their defaults. The database name is hardcoded to
// "Chinook".
func GetMySQLFromEnvironment() (*gorm.DB, error) {
	rootPassword, ok := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("environment variable MYSQL_ROOT_PASSWORD not set")
	}

	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf("root:%s@tcp(localhost:3306)/Chinook", rootPassword)),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database\n%w", err)
	}

	return db, nil
}
