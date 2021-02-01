package entity

import (
	"errors"
	"time"

	"github.com/dpjungmin/jellypi-server/tools"
)

// User entity
type User struct {
	ID        string    `json:"id" gorm:"type:uuid; default:uuid_generate_v4(); primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp with time zone; default:now(); autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp with time zone; default:now(); autoUpdateTime:milli"`

	Username   string `json:"username" gorm:"type:varchar(55); unique; not null"`
	Password   string `json:"-" gorm:"type:text; not null"`
	Email      string `json:"email" gorm:"type:varchar(55); unique; not null"`
	IsVerified bool   `json:"is_verified" gorm:"type:boolean; default:false; not null"`
}

// Validate validates the entity
func (e *User) Validate() error {
	if e.Username == "" {
		return errors.New("username is required")
	} else if len(e.Username) < 3 || len(e.Username) > 55 {
		return errors.New("username must be at between 3 to 55 characters")
	}

	if e.Password == "" {
		return errors.New("password is required")
	} else if len(e.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if e.Email == "" {
		return errors.New("email is required")
	} else if !tools.EmailRegex.MatchString(e.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
