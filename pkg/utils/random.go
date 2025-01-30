package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomString menghasilkan string acak dengan panjang tertentu
func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
