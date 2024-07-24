package user

import (
	media "pop_culture/api/resource/Media"
	mediatype "pop_culture/api/resource/MediaType"
	"pop_culture/util/validation"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDto struct {
	ID    string
	Name  string
	Email string
}

type User struct {
	ID        uuid.UUID `gorm: "primaryKey"`
	Name      string
	Email     *string `gorm:"unique"`
	Password  string
	Medias    []media.Media         `gorm:"many2many:user_media;"`
	Interests []mediatype.TypeMedia `gorm:"many2many:user_interests;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type FormUser struct {
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

func NewUser(name string, email *string, password string) (*User, error) {
	userName, err := validation.UserNameRules(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := validation.EmailRules(email)
	if err != nil {
		return nil, err
	}
	userPassword, err := validation.PasswordRules(password)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     *userName,
		Email:    userEmail,
		Password: *userPassword,
	}, nil

}

func (u *User) ToDto() *UserDto {
	return &UserDto{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: *u.Email,
	}
}
