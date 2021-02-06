package handler

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/dto"
	"github.com/dpjungmin/jellypi-server/service"
	"github.com/gofiber/fiber/v2"
)

type (
	// AuthHandler defines all auth related handlers
	AuthHandler interface {
		Login(*fiber.Ctx) error
	}

	authHandler struct {
		authSrv service.AuthService
	}
)

// NewAuthHandler generates a new auth handler
func NewAuthHandler(authSrv service.AuthService) AuthHandler {
	return &authHandler{authSrv}
}

// Login handler
func (h *authHandler) Login(c *fiber.Ctx) error {
	// Parse body
	d := new(dto.LoginRequest)
	if err := c.BodyParser(d); err != nil {
		return err
	}
	// Generate token
	token, err := h.authSrv.Login(d)
	if err != nil {
		return dto.NewErrorResponse(c, err)
	}
	// Send response
	return c.Status(http.StatusOK).JSON(token)
}
