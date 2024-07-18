package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) Create(user *User) (*User, error) {

	if err := r.database.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Read(id uuid.UUID) (*User, error) {

	user := &User{}
	if err := r.database.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *User) (int64, error) {
	result := r.database.Model(&User{}).Select("Name", "Email", "UpatedAt").Where("id = ?", user.ID).Updates(user)

	return result.RowsAffected, result.Error

}

func (r *UserRepository) Delete(id uuid.UUID) (int64, error) {
	result := r.database.Where("id = ?", id).Delete(&User{})
	return result.RowsAffected, result.Error

}
