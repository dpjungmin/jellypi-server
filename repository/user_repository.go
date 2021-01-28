package repository

import (
	d "github.com/dpjungmin/jellypi-server/dto"
	e "github.com/dpjungmin/jellypi-server/entity"
	"github.com/dpjungmin/jellypi-server/tools/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// IUserRepository interface
type IUserRepository interface {
	FindByID(string) (*e.User, error)
	Create(*e.User) (*e.User, error)
}

// UserRepository structure
type UserRepository struct {
	db *gorm.DB
}

// FindByID finds a user by id
func (r *UserRepository) FindByID(id string) (*e.User, error) {
	return nil, nil
}

// Create creates a new user
func (r *UserRepository) Create(u *e.User) (*e.User, *d.Error) {
	var existingUser e.User

	r.db.Where(&e.User{Email: u.Email}).Find(&existingUser)
	if existingUser.Email == u.Email {
		return nil, d.NewError(fiber.StatusConflict, "email conflict")
	}

	r.db.Where(&e.User{Username: u.Username}).Find(&existingUser)
	if existingUser.Username == u.Username {
		return nil, d.NewError(fiber.StatusConflict, "username conflict")
	}

	if err := r.db.Create(&u).Error; err != nil {
		logger.Error("[DATABASE]::[CREATE_ERROR]", err)
		return nil, d.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return u, nil
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}
