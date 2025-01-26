// internal/users/repository/user_repository.go
package repository

import (
	"database/sql"
	"errors"
	"food-delivery/internal/users/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password, role FROM users WHERE email = $1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.Role)
	return err
}
