package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config berisi konfigurasi aplikasi
type Config struct {
	DatabaseURL string
}

// LoadConfig membaca konfigurasi dari file .env atau environment variables
func LoadConfig() *Config {
	// Muat file .env (opsional)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
	}
}

// getEnv membaca variabel lingkungan dengan nilai default
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
