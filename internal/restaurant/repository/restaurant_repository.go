// internal/restaurant/repository/restaurant_repository.go
package repository

import (
	"database/sql"
	"errors"

	"food-delivery/internal/restaurant/models"
)

type RestaurantRepository struct {
	DB *sql.DB
}

func NewRestaurantRepository(db *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{DB: db}
}

func (r *RestaurantRepository) GetRestaurants() ([]models.Restaurant, error) {
	rows, err := r.DB.Query("SELECT id, name, address, cuisine_type, rating FROM restaurants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []models.Restaurant
	for rows.Next() {
		var rest models.Restaurant
		if err := rows.Scan(&rest.ID, &rest.Name, &rest.Address, &rest.CuisineType, &rest.Rating); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, rest)
	}
	return restaurants, nil
}

// GetByName mencari restoran berdasarkan nama
func (r *RestaurantRepository) GetByName(name string) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := r.DB.QueryRow("SELECT id, name, address, cuisine_type, rating FROM restaurants WHERE name = $1", name).Scan(
		&restaurant.ID, &restaurant.Name, &restaurant.Address, &restaurant.CuisineType, &restaurant.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tidak ditemukan
		}
		return nil, err // Terjadi error lain
	}
	return &restaurant, nil
}

func (r *RestaurantRepository) AddRestaurant(restaurant models.Restaurant) error {
	query := "INSERT INTO restaurants (name, address, cuisine_type, rating) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, restaurant.Name, restaurant.Address, restaurant.CuisineType, restaurant.Rating)
	if err != nil {
		return err
	}
	return nil
}

func (r *RestaurantRepository) GetRestaurantByID(id int) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	query := "SELECT id, name, address, cuisine_type, rating FROM restaurants WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&restaurant.ID, &restaurant.Name, &restaurant.Address, &restaurant.CuisineType, &restaurant.Rating)
	if err == sql.ErrNoRows {
		return nil, errors.New("restaurant not found")
	} else if err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (r *RestaurantRepository) DeleteRestaurant(id int) error {
	query := "DELETE FROM restaurants WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
