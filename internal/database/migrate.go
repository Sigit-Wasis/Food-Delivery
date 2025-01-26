package database

import (
	"database/sql"
	"log"
)

// MigrateTables menjalankan migrasi tabel secara manual dengan SQL
func MigrateTables(db *sql.DB) {
	log.Println("Running migrations...")

	// Skrip SQL untuk membuat tabel `restaurants`
	createRestaurantsTable := `
	CREATE TABLE IF NOT EXISTS restaurants (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		address VARCHAR(255) NOT NULL,
		cuisine_type VARCHAR(100) NOT NULL,
		rating FLOAT DEFAULT 0
	);`

	// Jalankan skrip migrasi
	_, err := db.Exec(createRestaurantsTable)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migration completed successfully.")
}
