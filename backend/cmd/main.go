package main

import (
	"log"

	"github.com/ryank157/babyTracker/backend/db"
)

func main() {

	postgresDb, err := db.NewPostgresDB("postgres://username:password@localhost:5432/database_name")
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	defer postgresDb.Close()

	queries := postgresDb.Queries()

}
