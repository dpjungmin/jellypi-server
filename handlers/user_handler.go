package handlers

import (
	"github.com/dpjungmin/jellypi-server/dtos"
	"github.com/dpjungmin/jellypi-server/services"
	"github.com/dpjungmin/jellypi-server/tools/logger"
	"github.com/gofiber/fiber/v2"
)

// IUserHandler interface
type IUserHandler interface {
	GetUser(*fiber.Ctx) error
	CreateUser(*fiber.Ctx) error
}

// UserHandler structure
type UserHandler struct {
	s services.UserService
}

// GetUser handler
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	user, err := h.s.GetUser("userID") // fake id
	if err != nil {
		logger.Error("Failed to get user", err)
	}
	return c.JSON(user)
}

// CreateUser handler
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Parse body
	dto := new(dtos.CreateUserRequest)
	if err := c.BodyParser(dto); err != nil {
		return err
	}
	// Create user
	u, err := h.s.CreateUser(dto)
	if err != nil {
		return c.Status(err.Code).JSON(dtos.NewErrorResponse(err.Code, err.Message))
	}
	// Send response
	return c.Status(fiber.StatusCreated).JSON(u)
}

// NewUserHandler creates a new user handler
func NewUserHandler(s services.UserService) UserHandler {
	return UserHandler{s}
}
