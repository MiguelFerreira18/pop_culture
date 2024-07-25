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
	Name        string `json:"name"`
	MediaTypeID uint   `json:"mediatypeID"`
}

type MediaDTO struct {
	ID        uint
	Name      string
	TypeMedia *mediatype.TypeMediaDTO
}

type Medias []*Media

func NewMedia(name string, TypeMediaID uint) (*Media, error) {
	mediaName, err := validation.MediaNameRules(name)
	if err != nil {
		return nil, err
	}
	return &Media{
		Name:        *mediaName,
		MediaTypeID: TypeMediaID,
	}, nil

}

func (m Media) ToDTO() *MediaDTO {
	return &MediaDTO{
		ID:        m.ID,
		Name:      m.Name,
		TypeMedia: m.MediaType.ToDTO(),
	}
}

func (ms Medias) ToDTO() []*MediaDTO {
	dtos := make([]*MediaDTO, len(ms))
	for i, m := range ms {
		dtos[i] = m.ToDTO()
	}
	return dtos
}
