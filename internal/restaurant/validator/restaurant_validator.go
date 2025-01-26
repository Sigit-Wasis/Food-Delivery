package validator

import (
	"food-delivery/internal/restaurant/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRestaurant(restaurant models.Restaurant) error {
	// Validasi menggunakan library validator
	return validate.Struct(restaurant)
}
