package handler

import (
	d "github.com/dpjungmin/jellypi-server/dto"
	s "github.com/dpjungmin/jellypi-server/service"
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
	s s.UserService
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
	dto := new(d.CreateUserRequest)
	if err := c.BodyParser(dto); err != nil {
		return err
	}
	// Create user
	u, err := h.s.CreateUser(dto)
	if err != nil {
		return c.Status(err.Code).JSON(d.NewErrorResponse(err.Code, err.Message))
	}
	// Send response
	return c.Status(fiber.StatusCreated).JSON(u)
}

// NewUserHandler creates a new user handler
func NewUserHandler(s s.UserService) UserHandler {
	return UserHandler{s}
}
