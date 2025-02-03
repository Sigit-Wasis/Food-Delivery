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

	// Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Rute untuk restoran
	// @Summary Get list of restaurants
	// @Description Mendapatkan daftar restoran
	// @Tags Restaurants
	// @Accept json
	// @Produce json
	// @Success 200 {array} []handler.RestaurantResponse
	// @Router /restaurants [get]
	app.Get("/restaurants", restaurantHandler.GetRestaurants)

	// @Summary Add a new restaurant
	// @Description Menambahkan restoran baru
	// @Tags Restaurants
	// @Accept json
	// @Produce json
	// @Param restaurant body handler.RestaurantRequest true "Restaurant Data"
	// @Success 201 {object} handler.RestaurantResponse
	// @Router /restaurants [post]
	app.Post("/restaurants", restaurantHandler.AddRestaurant)

	// Rute untuk users
	// @Summary Login user
	// @Description Melakukan login user
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param user body handler.LoginRequest true "User Data"
	// @Success 200 {object} handler.LoginResponse
	// @Router /users [get]
	app.Get("/users", userHandler.LoginUser)

	// @Summary Register user
	// @Description Registrasi user baru
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param user body handler.RegisterRequest true "User Data"
	// @Success 201 {object} handler.RegisterResponse
	// @Router /users [post]
	app.Post("/users", userHandler.RegisterUser)

	// Jalankan server
	log.Println("Server running on http://localhost:4000")
	log.Fatal(app.Listen(":4000"))
}
