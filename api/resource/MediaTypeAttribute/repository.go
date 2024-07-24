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

func (r MediaTypeAttributeRepository) AddAttribute(mediaTypeID uint, attributeID uint) (*mediatype.TypeMedia, error) {

	attribute, mediaType, err := getAttributeAndMediaType(r, attributeID, mediaTypeID)
	if err != nil {
		return nil, err
	}
	mediaType.Attributes = append(mediaType.Attributes, *attribute)

	if err := r.db.Save(&mediaType).Error; err != nil {
		return nil, err
	}
	return mediaType, nil
}

func (r MediaTypeAttributeRepository) RemoveAttribute(mediaTypeID uint, attributeID uint) (*mediatype.TypeMedia, error) {

	attribute, mediaType, err := getAttributeAndMediaType(r, attributeID, mediaTypeID)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&mediaType).Association("Attributes").Delete(&attribute); err != nil {
		return nil, err
	}
	return mediaType, nil
}

func getAttributeAndMediaType(r MediaTypeAttributeRepository, attributeID uint, mediaTypeID uint) (*attribute.Attribute, *mediatype.TypeMedia, error) {
	attribute := attribute.Attribute{}
	if err := r.db.Where("id = ?", attributeID).First(&attribute).Error; err != nil {
		return nil, nil, err
	}

	mediaType := mediatype.TypeMedia{}
	if err := r.db.Where("id = ? ", mediaTypeID).Preload("Attributes").First(&mediaType).Error; err != nil {
		return nil, nil, err
	}
	return &attribute, &mediaType, nil
}
