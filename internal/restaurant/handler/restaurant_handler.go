package handler

import (
	"fmt"
	"food-delivery/internal/restaurant/models"
	"food-delivery/internal/restaurant/service"
	"food-delivery/pkg/response"
	"strconv"
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
)

type RestaurantHandler struct {
	Service *service.RestaurantService
}

func NewRestaurantHandler(service *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{Service: service}
}

// Validasi input restoran
func validateRestaurant(restaurant models.Restaurant) error {
	if strings.TrimSpace(restaurant.Name) == "" {
		return fmt.Errorf("Restaurant name cannot be empty")
	}
	if len(restaurant.Name) < 3 {
		return fmt.Errorf("Restaurant name must be at least 3 characters")
	}

	if strings.TrimSpace(restaurant.Address) == "" {
		return fmt.Errorf("Restaurant address cannot be empty")
	}

	// Validasi cuisine_type
    if strings.TrimSpace(restaurant.CuisineType) == "" {
        return fmt.Errorf("Cuisine type cannot be empty")
    }
    for _, c := range restaurant.CuisineType {
        if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
            return fmt.Errorf("Cuisine type must only contain letters and spaces")
        }
    }

	if restaurant.Rating < 1 || restaurant.Rating > 5 {
		return fmt.Errorf("Rating must be between 1 and 5")
	}

	return nil
}

// Handler untuk mendapatkan daftar restoran
func (h *RestaurantHandler) GetRestaurants(c *fiber.Ctx) error {
	restaurants, err := h.Service.ListRestaurants()

	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(fiber.StatusInternalServerError, "Failed to fetch restaurants", err.Error()))
	}

	if len(restaurants) == 0 {
		return response.SendResponse(c, response.NewSuccessResponse(fiber.StatusOK, "No restaurants found", nil))
	}

	return response.SendResponse(c, response.NewSuccessResponse(fiber.StatusOK, "Restaurants retrieved successfully", restaurants))
}

// Handler untuk menambahkan restoran
func (h *RestaurantHandler) AddRestaurant(c *fiber.Ctx) error {
	// Ambil nilai rating sebagai string dari form-data
	ratingStr := c.FormValue("rating")

	// Validasi jika rating kosong
	if ratingStr == "" {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest,
			"Validation error",
			"Rating cannot be empty",
		))
	}

	// Konversi rating ke float64 terlebih dahulu
	rating, err := strconv.ParseFloat(ratingStr, 32)
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest,
			"Validation error",
			"Invalid rating format",
		))
	}

	// Buat objek restoran dari form data
	restaurant := models.Restaurant{
		Name:        c.FormValue("name"),
		Address:     c.FormValue("address"),
		CuisineType: c.FormValue("cuisine_type"),
		Rating:      float32(rating),
	}

	// Validasi input menggunakan fungsi validasi terpisah
	if err := validateRestaurant(restaurant); err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest,
			"Validation error",
			err.Error(),
		))
	}

	// Panggil service untuk menambahkan restoran
	if err := h.Service.CreateRestaurant(restaurant); err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusInternalServerError,
			"Internal Server Error",
			"Failed to add restaurant: "+err.Error(),
		))
	}

	// Kembalikan respons sukses
	return response.SendResponse(c, response.NewSuccessResponse(
		fiber.StatusCreated,
		"Restaurant added successfully",
		nil,
	))
}

// DeleteRestaurant menghapus restoran berdasarkan ID
func (h *RestaurantHandler) DeleteRestaurant(c *fiber.Ctx) error {
	// Ambil ID dari parameter URL
	idStr := c.Params("id")

	// Konversi ID menjadi integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest,
			"Invalid restaurant ID",
			"Restaurant ID must be a number",
		))
	}

	// Panggil service untuk menghapus restoran
	err = h.Service.DeleteRestaurant(id)
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusInternalServerError,
			"Failed to delete restaurant",
			err.Error(),
		))
	}

	// Kirim respons sukses
	return response.SendResponse(c, response.NewSuccessResponse(
		fiber.StatusOK,
		"Restaurant deleted successfully",
		nil,
	))
}