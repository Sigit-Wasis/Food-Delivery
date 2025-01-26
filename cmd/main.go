package main

import (
	"log"

	"food-delivery/internal/database"
	"food-delivery/internal/restaurant/handler"
	"food-delivery/internal/restaurant/repository"
	"food-delivery/internal/restaurant/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq" // Driver untuk PostgreSQL
)

func main() {
	// Inisialisasi database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Jalankan migrasi
	database.MigrateTables(db)

	// Inisialisasi komponen restoran
	restaurantRepo := repository.NewRestaurantRepository(db)
	restaurantService := service.NewRestaurantService(restaurantRepo)
	restaurantHandler := handler.NewRestaurantHandler(restaurantService)

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

	// Tambahkan rute untuk modul users
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

app.Post("/users/register", userHandler.RegisterUser)

	// Jalankan server
	log.Println("Server running on http://localhost:4000")
	log.Fatal(app.Listen(":4000"))
}
