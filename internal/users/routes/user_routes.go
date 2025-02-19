// internal/users/routes/user_routes.go
package routes

import (
	"food-delivery/internal/users/handler"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes mengelompokkan semua endpoint user
func UserRoutes(router fiber.Router, handler *handler.UserHandler) {
	router.Get("/", handler.Index) // Menampilkan semua user
	router.Post("/register", handler.RegisterUser) // Registrasi user baru
	router.Post("/login", handler.LoginUser) // Login user
	router.Delete("/:id", handler.DeleteUser)
	router.Put("/:id/password", handler.UpdatePassword) // update password
}