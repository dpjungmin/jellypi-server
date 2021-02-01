package service

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/dto"
	"github.com/dpjungmin/jellypi-server/entity"
	repo "github.com/dpjungmin/jellypi-server/repository"
)

// UserService defines all user related services
type UserService interface {
	GetUser(string) (*entity.User, error)
	CreateUser(*dto.CreateUserRequest) (*entity.User, *dto.Error)
}
type userService struct {
	userRepo repo.UserRepository
}

// NewUserService generates a new user service
func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{userRepo}
}

// GetUser returns a user by the given id
func (s *userService) GetUser(userID string) (*entity.User, error) {
	return s.userRepo.FindByID(userID)
}

// CreateUser creates a new user
func (s *userService) CreateUser(d *dto.CreateUserRequest) (*entity.User, *dto.Error) {
	// Validate DTO
	if errs := d.Validate(); errs != nil {
		return nil, dto.NewErrorWithStack(http.StatusBadRequest, errs, "Request body validation error")
	}
	// Create new user entity
	u := &entity.User{
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
