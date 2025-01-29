package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

// RunMigrations menjalankan semua migrasi dari folder migrations
func RunMigrations(db *sql.DB) {
	migrationFolder := "migrations"

	files, err := ioutil.ReadDir(migrationFolder)
	if err != nil {
		log.Fatalf("Failed to read migration folder: %v", err)
	}

	// Urutkan berdasarkan nama file
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Eksekusi tiap file
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(migrationFolder, file.Name())
		log.Printf("Executing migration: %s", filePath)

		sqlContent, err := ioutil.ReadFile(filepath.Clean(filePath))
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", file.Name(), err)
		}

		_, err = db.Exec(string(sqlContent))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", file.Name(), err)
		}

		log.Printf("Migration executed successfully: %s", file.Name())
	}

	log.Println("All migrations completed successfully.")
}
