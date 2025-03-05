// internal/restaurant/routes/restaurant_routes.go
package routes

import (
	"food-delivery/internal/restaurant/handler"

	"github.com/gofiber/fiber/v2"
)

// RestaurantRoutes mengelompokkan semua endpoint restoran
func RestaurantRoutes(router fiber.Router, handler *handler.RestaurantHandler) {
	router.Get("/", handler.GetRestaurants) // Mendapatkan daftar restoran
	router.Post("/", handler.AddRestaurant) // Menambahkan restoran baru
	router.Delete("/:id", handler.DeleteRestaurant) // Menghapus restoran berdasarkan ID
	router.Get("/:id", handler.GetRestaurantByID) // Mendapatkan restoran berdasarkan ID
}
