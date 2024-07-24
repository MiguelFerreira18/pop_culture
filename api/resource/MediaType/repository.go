package mediatype

import "gorm.io/gorm"

type MediaTypeRepository struct {
	database *gorm.DB
}

func NewMediaTypeRepository(database *gorm.DB) *MediaTypeRepository {
	return &MediaTypeRepository{database}
}

//CRUD

func (r MediaTypeRepository) Create(media *TypeMedia) (*TypeMedia, error) {
	if err := r.database.Create(&media).Error; err != nil {
		return nil, err
	}
	return media, nil
}

func (r MediaTypeRepository) Read(id uint) (*TypeMedia, error) {
	newTypeMedia := &TypeMedia{}
	if err := r.database.Preload("Attributes").First(&newTypeMedia, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return newTypeMedia, nil
}

func (r MediaTypeRepository) Update(media *TypeMedia) (int64, error) {
	result := r.database.Model(&media).Select("Name", "UpdatedAt").Where("id = ?", media.ID).Updates(media)

	return result.RowsAffected, result.Error
}

func (r MediaTypeRepository) Delete(id uint) (int64, error) {
	result := r.database.Where("id = ?", id).Delete(&TypeMedia{})
	return result.RowsAffected, result.Error
}
