package handler

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/dto"
	srv "github.com/dpjungmin/jellypi-server/service"
	"github.com/dpjungmin/jellypi-server/tools/logger"
	"github.com/gofiber/fiber/v2"
)

// UserHandler defines all user related handlers
type UserHandler interface {
	GetUser(*fiber.Ctx) error
	CreateUser(*fiber.Ctx) error
}

type userHandler struct {
	userSrv srv.UserService
}

// NewUserHandler generates a new user handler
func NewUserHandler(userSrv srv.UserService) UserHandler {
	return &userHandler{userSrv}
}

// GetUser handler
func (h *userHandler) GetUser(c *fiber.Ctx) error {
	user, err := h.userSrv.GetUser("userID") // fake id
	if err != nil {
		logger.Error("Failed to get user", err)
	}
	return c.JSON(user)
}

// CreateUser handler
func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	// Parse body
	d := new(dto.CreateUserRequest)
	if err := c.BodyParser(d); err != nil {
		return err
	}
	// Create user
	u, err := h.userSrv.CreateUser(d)
	if err != nil {
		return dto.NewErrorResponse(c, err)
	}
	// Send response
	return c.Status(http.StatusCreated).JSON(u)
}
