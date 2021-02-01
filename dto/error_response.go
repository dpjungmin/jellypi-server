package dto

import "github.com/gofiber/fiber/v2"

// ErrorResponse is the formated error response
type ErrorResponse struct {
	Error Error `json:"error"`
}

// NewErrorResponse generates a new ErrorResponse instance with an optional message
func NewErrorResponse(c *fiber.Ctx, e *Error) error {
	return c.Status(e.Code).JSON(&ErrorResponse{
		Error: Error{
			Code:    e.Code,
			Message: e.Message,
			Errors:  e.Errors,
		},
	})
}
