package dto

import (
	"github.com/dpjungmin/jellypi-server/tools"
)

// CreateUserRequest DTO
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Validate structure
func (dto *CreateUserRequest) Validate() Errors {
	var errs Errors

	if dto.Username == "" {
		errs = append(errs, "username is required")
	} else if len(dto.Username) < 3 || len(dto.Username) > 55 {
		errs = append(errs, "username must be at between 3 to 55 characters")
	}

	if dto.Password == "" {
		errs = append(errs, "password is required")
	} else if len(dto.Password) < 6 {
		errs = append(errs, "password must be at least 6 characters")
	}

	if dto.Email == "" {
		errs = append(errs, "email is required")
	} else if !tools.EmailRegex.MatchString(dto.Email) {
		errs = append(errs, "invalid email format")
	}

	return errs
}
