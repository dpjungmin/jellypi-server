package dto

// LoginRequest DTO
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Validate structure
func (dto *LoginRequest) Validate() Errors {
	var errs Errors

	if dto.Username == "" && dto.Email == "" {
		errs = append(errs, "username or email is required")
	}

	if dto.Password == "" {
		errs = append(errs, "password is required")
	}

	return errs
}
