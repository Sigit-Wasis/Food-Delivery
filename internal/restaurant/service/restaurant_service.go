// internal/restaurant/service/restaurant_service.go
package service

import (
	"errors"
	"fmt"
	"food-delivery/internal/restaurant/models"
	"food-delivery/internal/restaurant/repository"
)

type RestaurantService struct {
	Repo *repository.RestaurantRepository
}

func NewRestaurantService(repo *repository.RestaurantRepository) *RestaurantService {
	return &RestaurantService{Repo: repo}
}

func (s *RestaurantService) ListRestaurants() ([]models.Restaurant, error) {
	return s.Repo.GetRestaurants()
}

func (s *RestaurantService) CreateRestaurant(restaurant models.Restaurant) error {
	// Periksa apakah restoran dengan nama yang sama sudah ada
	existingRestaurant, _ := s.Repo.GetByName(restaurant.Name)
	if existingRestaurant != nil {
		return fmt.Errorf("Restaurant with name '%s' already exists", restaurant.Name)
	}
	
	return s.Repo.AddRestaurant(restaurant)
}

func (s *RestaurantService) GetRestaurantByID(id int) (*models.Restaurant, error) {
	restaurant, err := s.Repo.GetRestaurantByID(id)
	if err != nil {
		if err.Error() == "restaurant not found" {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}
	return restaurant, nil
}

func (s *RestaurantService) DeleteRestaurant(id int) error {
	return s.Repo.DeleteRestaurant(id)
}