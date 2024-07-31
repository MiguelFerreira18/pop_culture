package auth

import (
	"errors"
	user "pop_culture/api/resource/User"
	"pop_culture/util/hash"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(gorm *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: gorm,
	}
}

func (r AuthRepository) LoadUser(userID string, email string) (*user.User, error) {
	userBody := &user.User{}
	if err := r.db.Where("id = ?", userID).Where("email = ?", email).Preload("Role").First(&userBody).Error; err != nil {
		return nil, err
	}
	return userBody, nil

}

func (r AuthRepository) Login(email string, password string) (*user.User, error) {
	userBody := &user.User{}
	if err := r.db.Where("email = ?", email).First(&userBody).Error; err != nil {
		return nil, err
	}
	if hash.CheckPassword(password, userBody.Password) {
		return nil, errors.New("Passwords do not match")
	}

	return userBody, nil

}
