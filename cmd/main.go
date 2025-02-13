package main

import (
	"log"

	"food-delivery/config"
	"food-delivery/internal/database"
	"food-delivery/internal/restaurant/handler"
	"food-delivery/internal/restaurant/repository"
	restaurantRoutes "food-delivery/internal/restaurant/routes"
	"food-delivery/internal/restaurant/service"
	userHandler "food-delivery/internal/users/handler"
	userRepository "food-delivery/internal/users/repository"
	userRoutes "food-delivery/internal/users/routes"
	userService "food-delivery/internal/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq" // Driver untuk PostgreSQL

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Food Delivery API
// @version 1.0
// @description API untuk layanan food delivery
// @host localhost:4000
// @BasePath /
func main() {
	// Muat konfigurasi
	cfg := config.LoadConfig()

	// Hubungkan ke database
	db := database.ConnectDatabase(cfg)
	defer db.Close()

	// Jalankan migrasi
	database.RunMigrations(db)

	// Inisialisasi komponen restoran
	restaurantRepo := repository.NewRestaurantRepository(db)
	restaurantService := service.NewRestaurantService(restaurantRepo)
	restaurantHandler := handler.NewRestaurantHandler(restaurantService)

	// Inisialisasi komponen user
	userRepo := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	// Inisialisasi Fiber
	app := fiber.New()

	// Tambahkan Middleware
	app.Use(logger.New())      // Logging untuk request
	app.Use(recover.New())     // Menangani panic agar server tetap berjalan
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",             // Bisa diatur ke domain tertentu
		AllowMethods: "GET,POST,PUT",  // Metode HTTP yang diizinkan
	}))

	// Middleware Custom
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Powered-By", "Fiber")
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	// Swagger Documentation
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Grouping Routes
	api := app.Group("/api")

	// Grouping untuk Users
	userGroup := api.Group("/users")
	userRoutes.UserRoutes(userGroup, userHandler)

	// Grouping untuk Restaurants
	restaurantGroup := api.Group("/restaurants")
	restaurantRoutes.RestaurantRoutes(restaurantGroup, restaurantHandler)

	log.Println("Server running on http://localhost:4000")
	log.Fatal(app.Listen(":4000"))
}
