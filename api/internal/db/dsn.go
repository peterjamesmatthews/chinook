package db

import (
	"fmt"
	"os"
)

// GetDSN returns a MySQL DSN that can be used to connect to the Chinook
// database.
//
// # Errors
//   - error: MYSQL_ROOT_PASSWORD not set
func GetDSN() (string, error) {
	password, ok := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	if !ok {
		return "", fmt.Errorf("MYSQL_ROOT_PASSWORD not set")
	}

	return fmt.Sprintf("root:%s@tcp(localhost:3306)/Chinook", password), nil
}
