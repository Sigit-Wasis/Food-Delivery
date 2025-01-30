package utils

import "regexp"

// IsValidEmail memeriksa apakah format email valid
func IsValidEmail(email string) bool {
	// Regular expression untuk format email yang sederhana
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	r, _ := regexp.Compile(regex)
	return r.MatchString(email)
}
