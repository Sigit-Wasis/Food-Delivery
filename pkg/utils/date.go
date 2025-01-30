package utils

import "time"

// GetCurrentTimestamp mengembalikan timestamp saat ini
func GetCurrentTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

// ParseDate mengubah string menjadi waktu
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateStr)
}
