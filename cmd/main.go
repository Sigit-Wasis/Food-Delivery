package main

import (
	"log"

	"food-delivery/config"
	"food-delivery/internal/database"
	"food-delivery/internal/restaurant/handler"
	"food-delivery/internal/restaurant/repository"
	"food-delivery/internal/restaurant/service"
	userHandler "food-delivery/internal/users/handler"
	userRepository "food-delivery/internal/users/repository"
	userService "food-delivery/internal/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq" // Driver untuk PostgreSQL
)

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
	app.Use(logger.New())       // Logging untuk request
	app.Use(recover.New())      // Menangani panic agar server tetap berjalan
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",              // Bisa diatur ke domain tertentu
		AllowMethods: "GET,POST,PUT",   // Metode HTTP yang diizinkan
	}))

	// Middleware Custom (contoh: autentikasi atau header default)
	app.Use(func(c *fiber.Ctx) error {
		// Tambahkan header default untuk semua respons
		c.Set("X-Powered-By", "Fiber")
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	// Rute untuk restoran
	app.Get("/restaurants", restaurantHandler.GetRestaurants)
	app.Post("/restaurants", restaurantHandler.AddRestaurant)

	// Rute untuk users
	app.Get("/users", userHandler.LoginUser)
	app.Post("/users", userHandler.RegisterUser)

	// Jalankan server
	log.Println("Server running on http://localhost:4000")
	log.Fatal(app.Listen(":4000"))
}
