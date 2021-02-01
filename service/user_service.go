package service

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/domain"
	"github.com/dpjungmin/jellypi-server/dto"
)

// UserService defines all user related services
type UserService interface {
	CreateUser(*dto.CreateUserRequest) (*domain.User, *dto.Error)
}
type userService struct {
	userRepo domain.UserRepository
}

// NewUserService generates a new user service
func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{userRepo}
}

// CreateUser creates a new user
func (s *userService) CreateUser(d *dto.CreateUserRequest) (*domain.User, *dto.Error) {
	// Validate DTO
	if errs := d.Validate(); errs != nil {
		return nil, dto.NewErrorWithStack(http.StatusBadRequest, errs, "Request body validation error")
	}
	// Create new user entity
	u := &domain.User{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password,
	}
	// Validate entity
	if err := u.Validate(); err != nil {
		return nil, dto.NewError(http.StatusBadRequest, err.Error())
	}
	// Create new user
	return s.userRepo.Create(u)
}
