package user

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
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
	UpatedAt  time.Time
	DeletedAt time.Time
}

func nameRules(name string) (*string, error) {
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`
	if len(name) < 5 && len(name) > 20 {
		return nil, errors.New("Name is not the correct size: " + strconv.Itoa(len(name)))
	}
	if matchPattern(name, pattern) {
		return nil, errors.New("Name is ")
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

}
func NewUser(name string, email *string, password string) {

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
