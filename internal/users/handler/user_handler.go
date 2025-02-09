package handler

import (
	"food-delivery/internal/users/models"
	"food-delivery/internal/users/service"
	"food-delivery/pkg/response"

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

// Index menangani permintaan daftar user
func (h *UserHandler) Index(c *fiber.Ctx) error {
	users, err := h.Service.GetUsers()
	if err != nil {
		return response.SendResponse(c, response.NewErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve users", err.Error()))
	}
	return response.SendResponse(c, response.NewSuccessResponse(fiber.StatusOK, "Users retrieved successfully", users))
}
