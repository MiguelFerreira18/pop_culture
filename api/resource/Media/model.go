package media

import (
	mediatype "pop_culture/api/resource/MediaType"
	"pop_culture/util/validation"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	Name        string
	MediaTypeID uint
	MediaType   mediatype.TypeMedia `gorm:"foreignKey:MediaTypeID"`
}

type MediaForm struct {
	Name string `json:"name"`
}

type MediaDTO struct {
	ID   uint
	Name string
}

func NewMedia(name string) (*Media, error) {
	mediaName, err := validation.MediaNameRules(name)
	if err != nil {
		return nil, err
	}
	return &Media{
		Name: *mediaName,
	}, nil

}

func (m Media) ToDTO() *MediaDTO {
	return &MediaDTO{
		ID:   m.ID,
		Name: m.Name,
	}
}
