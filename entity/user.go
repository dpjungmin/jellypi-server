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
	Email      string `json:"email" gorm:"type:varchar(50); unique; not null"`
	IsVerifeid bool   `json:"is_verified" gorm:"type:boolean; default:false; not null"`
}

// Validate structure
func (e *User) Validate() error {
	if e.Username == "" {
		return errors.New("username is required")
	}

	if len(e.Username) < 3 || len(e.Username) > 55 {
		return errors.New("username must be between 3 to 55 characters")
	}

	if e.Password == "" {
		return errors.New("password is required")
	}

	if e.Email == "" {
		return errors.New("email is required")
	}

	if !tools.EmailRegex.MatchString(e.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
