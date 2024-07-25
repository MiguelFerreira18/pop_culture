package mediatypeattribute

import (
	attribute "pop_culture/api/resource/Attribute"
	mediatype "pop_culture/api/resource/MediaType"

	"gorm.io/gorm"
)

type MediaTypeAttributeRepository struct {
	db *gorm.DB
}

func NewMediaTypeAttributeRepository(db *gorm.DB) *MediaTypeAttributeRepository {
	return &MediaTypeAttributeRepository{
		db: db,
	}
}

//Get Attributes from MediaType
//Add attribute to media type
//Remove attribute from media type

func (r MediaTypeAttributeRepository) GetAttributesFromMediaType(mediaTypeID uint) ([]attribute.Attribute, error) {
	mediaType := mediatype.TypeMedia{}
	if err := r.db.Where("id = ?", mediaTypeID).Preload("Attributes").First(&mediaType).Error; err != nil {
		return nil, err
	}

	return mediaType.Attributes, nil
}

func (r MediaTypeAttributeRepository) AddAttribute(mediaTypeID uint, attributeID uint) (*TypemediaAttribute, error) {
	typemediaAttribute := NewTypeMediaAttribute(mediaTypeID, attributeID)
	if err := r.db.Create(&typemediaAttribute).Error; err != nil {
		return nil, err
	}

	return typemediaAttribute, nil
}

func (r MediaTypeAttributeRepository) RemoveAttribute(mediaTypeID uint, attributeID uint) (*TypemediaAttribute, error) {

	typeMediaAttribute := TypemediaAttribute{}
	if err := r.db.Where("type_media_id = ?", mediaTypeID).Where("attribute_id = ?", attributeID).First(&typeMediaAttribute).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&typeMediaAttribute).Error; err != nil {
		return nil, err
	}

	return &typeMediaAttribute, nil
}

// func getAttributeAndMediaType(r MediaTypeAttributeRepository, attributeID uint, mediaTypeID uint) (*attribute.Attribute, *mediatype.TypeMedia, error) {
// 	attribute := attribute.Attribute{}
// 	if err := r.db.Where("id = ?", attributeID).First(&attribute).Error; err != nil {
// 		return nil, nil, err
// 	}
//
// 	mediaType := mediatype.TypeMedia{}
// 	if err := r.db.Where("id = ? ", mediaTypeID).Preload("Attributes").First(&mediaType).Error; err != nil {
// 		return nil, nil, err
// 	}
// 	return &attribute, &mediaType, nil
// }
