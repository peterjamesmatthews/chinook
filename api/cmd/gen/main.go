package main

import (
	"log"

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

	db, err := db.GetMySQLFromEnvironment()
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
