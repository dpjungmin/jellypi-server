package dtos

import (
	"errors"

	"github.com/dpjungmin/jellypi-server/tools"
)

// CreateUserRequest DTO
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Validate structure
func (dto *CreateUserRequest) Validate() error {
	if dto.Username == "" {
		return errors.New("username is required")
	}

	if len(dto.Username) < 3 || len(dto.Username) > 55 {
		return errors.New("username must be at between 3 to 55 characters")
	}

	if dto.Password == "" {
		return errors.New("password is required")
	}

	if len(dto.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if dto.Email == "" {
		return errors.New("email is required")
	}

	if !tools.EmailRegex.MatchString(dto.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
