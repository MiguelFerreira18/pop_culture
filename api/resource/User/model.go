package user

import (
	"errors"
	"pop_culture/util/hash"
	"regexp"
	"strconv"
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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type FormUser struct {
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

func nameRules(name string) (*string, error) {
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`
	if len(name) < 5 && len(name) > 20 {
		return nil, errors.New("Name is not the correct size: " + strconv.Itoa(len(name)))
	}
	if matchPattern(name, pattern) {
		return nil, errors.New("Name is not conformed")
	}
	return &name, nil
}
func EmailRules(email *string) (*string, error) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if matchPattern(*email, pattern) {
		return email, nil
	}

	return nil, errors.New("Email does not follow the rules")

}
func passwordRules(password string) (*string, error) {
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`

	if matchPattern(password, pattern) {
		return nil, errors.New("The password does not follow the rules")
	}

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return &hashedPassword, nil
}
func NewUser(name string, email *string, password string) (*User, error) {
	userName, err := nameRules(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := EmailRules(email)
	if err != nil {
		return nil, err
	}
	userPassword, err := passwordRules(password)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     *userName,
		Email:    userEmail,
		Password: *userPassword,
	}, nil

}

func matchPattern(input string, pattern string) bool {

	re := regexp.MustCompile(pattern)

	return re.MatchString(input)

}

func (u *User) ToDto() *UserDto {
	return &UserDto{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: *u.Email,
	}
}
