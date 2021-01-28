package service

import (
	d "github.com/dpjungmin/jellypi-server/dto"
	e "github.com/dpjungmin/jellypi-server/entity"
	r "github.com/dpjungmin/jellypi-server/repository"
	"github.com/gofiber/fiber/v2"
)

// IUserService interface
type IUserService interface {
	GetUser(string) (*e.User, error)
	CreateUser(*d.CreateUserRequest) (*e.User, error)
}

// UserService structure
type UserService struct {
	userRepo r.UserRepository
}

// GetUser returns a user by the given id
func (s *UserService) GetUser(userID string) (*e.User, error) {
	return s.userRepo.FindByID(userID)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(dto *d.CreateUserRequest) (*e.User, *d.Error) {
	// Validate DTO
	if err := dto.Validate(); err != nil {
		return nil, d.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Create user entity
	u := &e.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
	// Validate user entity
	if err := u.Validate(); err != nil {
		return nil, d.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Create new user
	return s.userRepo.Create(u)
}

// NewUserService creates a new user service
func NewUserService(userRepo r.UserRepository) UserService {
	return UserService{userRepo}
}
