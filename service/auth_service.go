package service

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/domain"
	"github.com/dpjungmin/jellypi-server/dto"
)

type (
	// AuthService defines all auth related services
	AuthService interface {
		Login(*dto.LoginRequest) (*string, *dto.Error)
	}

	authService struct {
		userRepo domain.UserRepository
	}
)

// NewAuthService generates a new auth service
func NewAuthService(userRepo domain.UserRepository) AuthService {
	return &authService{userRepo}
}

// TODO: Login
func (s *authService) Login(d *dto.LoginRequest) (*string, *dto.Error) {
	// Validate DTO
	if errs := d.Validate(); errs != nil {
		return nil, dto.NewErrorWithStack(http.StatusBadRequest, errs, "Request body validation error")
	}

	// Find user
	u, err := s.userRepo.FindByCredentials(d.Email, d.Password)
	if err != nil {
		return nil, dto.NewError(http.StatusInternalServerError, err.Error())
	}

	// Generate access token
	token, err := u.GenerateToken()
	if err != nil {
		return nil, dto.NewError(http.StatusInternalServerError, err.Error())
	}

	return token, nil
}
