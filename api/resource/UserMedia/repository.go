package usermedia

import (
	media "pop_culture/api/resource/Media"
	user "pop_culture/api/resource/User"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMediaRepository struct {
	db *gorm.DB
}

func NewUserMediaRepository(db *gorm.DB) *UserMediaRepository {
	return &UserMediaRepository{
		db: db,
	}

}

//Get medias from User
//Add media to User
//Remove Media from user

func (r UserMediaRepository) GetMediasFromUser(userID uuid.UUID) ([]media.Media, error) {
	user := user.User{}
	if err := r.db.Where("id = ?", userID).Preload("Medias").First(&user).Error; err != nil {
		return nil, err
	}
	return user.Medias, nil

}

func (r UserMediaRepository) AddMediaToUser(mediaID uint, userID uuid.UUID) (*user.User, error) {

	myMedia, myUser, err := getMediaAndUser(r, mediaID, userID)
	if err != nil {
		return nil, err
	}
	myUser.Medias = append(myUser.Medias, *myMedia)
	if err := r.db.Save(myUser).Error; err != nil {
		return nil, err
	}

	return myUser, nil
}
func (r UserMediaRepository) RemoveMediaFromUser(mediaID uint, userID uuid.UUID) (*user.User, error) {

	myMedia, myUser, err := getMediaAndUser(r, mediaID, userID)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(myUser).Association("Medias").Delete(&myMedia); err != nil {
		return nil, err
	}

	return myUser, nil

}

func getMediaAndUser(r UserMediaRepository, mediaID uint, userID uuid.UUID) (*media.Media, *user.User, error) {
	newMedia := media.Media{}
	if err := r.db.Where("id = ?", mediaID).First(&newMedia).Error; err != nil {
		return nil, nil, err
	}
	newUser := user.User{}
	if err := r.db.Where("id = ?", userID).First(&newUser).Error; err != nil {
		return nil, nil, err
	}
	return &newMedia, &newUser, nil

}
