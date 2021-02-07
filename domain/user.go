package domain

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dpjungmin/jellypi-server/config"
	"github.com/dpjungmin/jellypi-server/utils"
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
	} else if !utils.EmailRegex.MatchString(e.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

// GenerateToken will generate an access token
func (e *User) GenerateToken() (*string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = e.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.JWT.Key))
	if err != nil {
		return nil, errors.New("Failed to generate token")
	}

	return &t, nil
}
