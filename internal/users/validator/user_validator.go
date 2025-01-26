package validator

import (
	"food-delivery/internal/users/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateUser(user models.User) error {
	// Validasi menggunakan library validator
	return validate.Struct(user)
}
