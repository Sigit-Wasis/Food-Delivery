// internal/restaurant/models/restaurant.go
package models

type Restaurant struct {
	ID          int64   `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	CuisineType string  `json:"cuisine_type" validate:"required"`
	Rating      float32 `json:"rating" validate:"required,gte=1,lte=5"`
}
