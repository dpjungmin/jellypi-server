package handler

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/dto"
	"github.com/dpjungmin/jellypi-server/service"
	"github.com/gofiber/fiber/v2"
)

// UserHandler defines all user related handlers
type UserHandler interface {
	CreateUser(*fiber.Ctx) error
}

type userHandler struct {
	userSrv service.UserService
}

// NewUserHandler generates a new user handler
func NewUserHandler(userSrv service.UserService) UserHandler {
	return &userHandler{userSrv}
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
