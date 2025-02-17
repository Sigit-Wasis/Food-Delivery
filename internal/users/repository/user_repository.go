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
	query := "SELECT id, username, email, password, role FROM users WHERE email = $1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.Role)
	return err
}

// GetAllUsers mengambil semua daftar user dari database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, username, email, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Cek apakah terjadi error saat iterasi rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Delete user by ID
func (r *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	res, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
