package domain

import (
	"net/http"

	"github.com/dpjungmin/jellypi-server/dto"
	"github.com/dpjungmin/jellypi-server/utils/logger"
	"gorm.io/gorm"
)

// UserRepository defines all data accessing methods
type UserRepository interface {
	Create(*User) (*User, *dto.Error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository generates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create creates a new user
func (r *userRepository) Create(u *User) (*User, *dto.Error) {
	var xUser User

	r.db.Where(&User{Email: u.Email}).Find(&xUser)
	if xUser.Email == u.Email {
		return nil, dto.NewError(http.StatusConflict, "email conflict")
	}

	r.db.Where(&User{Username: u.Username}).Find(&xUser)
	if xUser.Username == u.Username {
		return nil, dto.NewError(http.StatusConflict, "username conflict")
	}

	if err := r.db.Create(&u).Error; err != nil {
		logger.Error("[DATABASE]::[CREATE_ERROR]", err)
		return nil, dto.NewError(http.StatusBadRequest, err.Error())
	}

	return u, nil
}
