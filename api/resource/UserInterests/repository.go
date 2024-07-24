package userinterests

import (
	mediatype "pop_culture/api/resource/MediaType"
	user "pop_culture/api/resource/User"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInterestRepository struct {
	db *gorm.DB
}

func NewUserInterestRepository(db *gorm.DB) *UserInterestRepository {
	return &UserInterestRepository{
		db: db,
	}

}

//Get list of User Interests
//Add Interests to User
//Remove Interests from User

func (r UserInterestRepository) GetInterestFromUser(userID uuid.UUID) ([]mediatype.TypeMedia, error) {
	user := &user.User{}
	if err := r.db.Where("id = ?", userID).Preload("Interests").First(&user).Error; err != nil {
		return nil, err
	}
	return user.Interests, nil

}

func (r UserInterestRepository) AddInterestToUser(userID uuid.UUID, mediaTypeID uint) (*user.User, error) {

	user, mediaType, err := getUserAndMediaType(r, userID, mediaTypeID)
	if err != nil {
		return nil, err
	}

	user.Interests = append(user.Interests, *mediaType)
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil

}

func (r UserInterestRepository) RemoveInterestFromUser(userID uuid.UUID, mediaTypeID uint) (*user.User, error) {
	myUser, mediaType, err := getUserAndMediaType(r, userID, mediaTypeID)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(myUser).Association("Interests").Delete(&mediaType); err != nil {
		return nil, err

	}

	return myUser, nil

}

func getUserAndMediaType(r UserInterestRepository, userID uuid.UUID, mediaTypeID uint) (*user.User, *mediatype.TypeMedia, error) {
	user := &user.User{}
	if err := r.db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, nil, err
	}

	mediaType := &mediatype.TypeMedia{}
	if err := r.db.Where("id = ?", mediaTypeID).First(mediaType).Error; err != nil {
		return nil, nil, err
	}

	return user, mediaType, nil

}
