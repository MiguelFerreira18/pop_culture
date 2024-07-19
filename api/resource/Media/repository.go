package media

import (
	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{
		db: db,
	}
}

func (r MediaRepository) Create(media *Media) (*Media, error) {

	if err := r.db.Create(&media).Error; err != nil {
		return nil, err
	}
	return media, nil
}

func (r MediaRepository) Read(id uint) (*Media, error) {
	media := &Media{}
	if err := r.db.First(media, "id = ?").Error; err != nil {
		return nil, err
	}
	return media, nil
}

func (r MediaRepository) Update(media *Media) (int64, error) {
	result := r.db.Model(media).Select("Name", "UpdatedAt").Where("id = ?", media.ID).Updates(media)

	return result.RowsAffected, result.Error
}

func (r MediaRepository) Delete(id uint) (int64, error) {
	result := r.db.Select("id = ?", id).Delete(&Media{})
	return result.RowsAffected, result.Error
}
