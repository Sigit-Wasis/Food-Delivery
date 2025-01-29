package database

import (
	"database/sql"
	"log"

	"food-delivery/config"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// ConnectDatabase membuat koneksi ke database
func ConnectDatabase(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Uji koneksi
	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
