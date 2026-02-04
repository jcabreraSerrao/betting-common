package main

import (
	"log"

	"github.com/jcabreraSerrao/betting-common/pkg/db"
	"github.com/jcabreraSerrao/betting-common/pkg/utils"
)

func main() {
	// Initialize config
	config := utils.GetConfig()
	if config.Database.URL == "" {
		log.Fatal("DATABASE_URL is not set in environment or .env file")
	}

	// Connect to database
	gormDB, err := db.NewConnectionDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	log.Println("Database connection established")

	// Run migrations
	if err := db.RunMigrations(gormDB); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("All migrations and seeds completed successfully")
}
