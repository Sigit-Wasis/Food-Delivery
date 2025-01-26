package database

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Driver untuk PostgreSQL
)

func InitDB() (*sql.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	// Membaca variabel koneksi database dari .env
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Membuat koneksi ke database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Mengecek apakah koneksi berhasil
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	return db, nil
}
