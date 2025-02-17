// internal/users/service/user_service.go
package service

import (
	"errors"
	"food-delivery/internal/users/models"
	"food-delivery/internal/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(user models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.Repo.CreateUser(user)
}

func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetUsers mengambil semua daftar user
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

// Delete user service
func (s *UserService) DeleteUser(id int) error {
	// Cek apakah user ada sebelum menghapus
	err := s.Repo.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}