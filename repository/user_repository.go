package repository

import (
	"github.com/dpjungmin/jellypi-server/dto"
	"github.com/dpjungmin/jellypi-server/entity"
	"github.com/dpjungmin/jellypi-server/tools/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserRepository defines all data accessing methods
type UserRepository interface {
	FindByID(string) (*entity.User, error)
	Create(*entity.User) (*entity.User, *dto.Error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository generates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// FindByID finds a user by id
func (r *userRepository) FindByID(id string) (*entity.User, error) {
	return nil, nil
}

// Create creates a new user
func (r *userRepository) Create(u *entity.User) (*entity.User, *dto.Error) {
	var existingUser entity.User

	r.db.Where(&entity.User{Email: u.Email}).Find(&existingUser)
	if existingUser.Email == u.Email {
		return nil, dto.NewError(fiber.StatusConflict, "email conflict")
	}

	r.db.Where(&entity.User{Username: u.Username}).Find(&existingUser)
	if existingUser.Username == u.Username {
		return nil, dto.NewError(fiber.StatusConflict, "username conflict")
	}

	if err := r.db.Create(&u).Error; err != nil {
		logger.Error("[DATABASE]::[CREATE_ERROR]", err)
		return nil, dto.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return u, nil
}
