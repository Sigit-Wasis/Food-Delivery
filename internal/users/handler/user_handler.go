package handler

import (
	"food-delivery/internal/users/models"
	"food-delivery/internal/users/service"
	"food-delivery/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service *service.UserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// RegisterUser menangani pendaftaran user baru
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var user models.User

	// Decode body JSON
	if err := c.BodyParser(&user); err != nil {
		return response.SendResponse(c, response.NewErrorResponse(fiber.StatusBadRequest, "Invalid input", err.Error()))
	}

	// Panggil service untuk registrasi user
	if err := h.Service.RegisterUser(user); err != nil {
		return response.SendResponse(c, response.NewErrorResponse(fiber.StatusInternalServerError, "Failed to register user", err.Error()))
	}

	// Kirim response sukses
	return response.SendResponse(c, response.NewSuccessResponse(fiber.StatusCreated, "User registered successfully", nil))
}

// LoginUser menangani login user
func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode body JSON
	if err := c.BodyParser(&credentials); err != nil {
		return response.SendResponse(c, response.NewErrorResponse(fiber.StatusBadRequest, "Invalid input", err.Error()))
	}

	// Panggil service untuk login
	user, err := h.Service.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(fiber.StatusUnauthorized, "Invalid credentials", err.Error()))
	}

	// Kirim response sukses dengan data user
	return response.SendResponse(c, response.NewSuccessResponse(fiber.StatusOK, "Login successful", user))
}

// GetUsers mendapatkan daftar pengguna
// @Summary Get all users
// @Description Mendapatkan semua data pengguna yang terdaftar
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{} "Berhasil mendapatkan daftar user"
// @Failure 500 {object} map[string]interface{} "Terjadi kesalahan server"
// @Router /users [get]
func (h *UserHandler) Index(c *fiber.Ctx) error {
	users, err := h.Service.GetUsers()
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusInternalServerError,
			"Failed to retrieve users",
			err.Error(),
		))
	}

	return response.SendResponse(c, response.NewSuccessResponse(
		fiber.StatusOK,
		"Users retrieved successfully",
		users,
	))
}

type UpdatePasswordRequest struct {
	NewPassword string `json:"new_password"`
}

// Handler untuk update password user
func (h *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	// Ambil parameter ID
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest, "Invalid ID format", err.Error(),
		))
	}

	// Parse request body
	var req UpdatePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest, "Invalid request body", err.Error(),
		))
	}

	// Validasi password baru
	if req.NewPassword == "" {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest, "New password cannot be empty", "",
		))
	}

	// Panggil service untuk update password
	err = h.Service.UpdatePassword(id, req.NewPassword)
	if err != nil {
		if err.Error() == "user not found" {
			return response.SendResponse(c, response.NewErrorResponse(
				fiber.StatusNotFound, "User not found", "",
			))
		}
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusInternalServerError, "Failed to update password", err.Error(),
		))
	}

	// Berhasil diupdate
	return response.SendResponse(c, response.NewSuccessResponse(
		fiber.StatusOK, "Password updated successfully", nil,
	))
}

// Delete user handler
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	// Ambil parameter ID
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusBadRequest, "Invalid ID format", err.Error(),
		))
	}

	// Panggil service untuk menghapus user
	err = h.Service.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			return response.SendResponse(c, response.NewErrorResponse(
				fiber.StatusNotFound, "User not found", "",
			))
		}
		return response.SendResponse(c, response.NewErrorResponse(
			fiber.StatusInternalServerError, "Failed to delete user", err.Error(),
		))
	}

	// Berhasil dihapus
	return response.SendResponse(c, response.NewSuccessResponse(
		fiber.StatusOK, "User deleted successfully", nil,
	))
}