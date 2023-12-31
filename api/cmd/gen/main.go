package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gen"
	"pjm.dev/chinook/internal/db"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/db/dao",
		Mode: gen.WithoutContext |
			gen.WithDefaultQuery |
			gen.WithQueryInterface,
	})

	// TODO move me to somewhere else
	rootPassword, ok := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	if !ok {
		log.Fatal("environment variable MYSQL_ROOT_PASSWORD not set")
	}

	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/Chinook", rootPassword)

	db, err := db.GetMySQL(dsn)
	if err != nil {
		log.Fatal(err)
	}

	g.UseDB(db)

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
