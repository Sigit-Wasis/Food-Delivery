// internal/users/models/user.go
package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Password harus di-hash
	Role     string `json:"role"`    // e.g., "customer", "admin", "delivery"
}