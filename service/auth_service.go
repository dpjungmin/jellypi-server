package service

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/domain"
	"github.com/dpjungmin/jellypi-server/dto"
)

type (
	// AuthService defines all auth related services
	AuthService interface {
		Login(*dto.LoginRequest) (*dto.Token, *dto.Error)
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
func (s *authService) Login(d *dto.LoginRequest) (*dto.Token, *dto.Error) {
	// Validate DTO
	if errs := d.Validate(); errs != nil {
		return nil, dto.NewErrorWithStack(http.StatusBadRequest, errs, "Request body validation error")
	}

	// u, err := s.userRepo.FindByCredentials(d.Email, d.Password)
	// token, err := u.GenerateToken()

	// return token, nil

	return nil, nil
}
