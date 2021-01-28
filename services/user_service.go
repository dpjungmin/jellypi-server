package services

import (
	"github.com/dpjungmin/jellypi-server/dtos"
	e "github.com/dpjungmin/jellypi-server/entities"
	"github.com/dpjungmin/jellypi-server/repositories"
	"github.com/gofiber/fiber/v2"
)

// IUserService interface
type IUserService interface {
	GetUser(string) (*e.User, error)
	CreateUser(*e.User) (*e.User, error)
}

// UserService structure
type UserService struct {
	userRepo repositories.UserRepository
}

// GetUser returns a user by the given id
func (s *UserService) GetUser(userID string) (*e.User, error) {
	return s.userRepo.FindByID(userID)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(dto *dtos.CreateUserRequest) (*e.User, *dtos.Error) {
	// Validate DTO
	if err := dto.Validate(); err != nil {
		return nil, dtos.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Create user entity
	u := &e.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
	// Validate user entity
	if err := u.Validate(); err != nil {
		return nil, dtos.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Create new user
	return s.userRepo.Create(u)
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository) UserService {
	return UserService{userRepo}
}
