package mediatype

import (
	"pop_culture/util/validation"

	"gorm.io/gorm"
)

type TypeMedia struct {
	gorm.Model
	Name string
}

type TypeMediaForm struct {
	Name string `json:"name"`
}
type TypeMediaDTO struct {
	ID   uint
	Name string
}

func NewTypeMedia(name string) (*TypeMedia, error) {
	mediaName, err := validation.MediaNameRules(name)
	if err != nil {
		return nil, err
	}
	return &TypeMedia{
		Name: *mediaName,
	}, nil
}

func (tm TypeMedia) ToDTO() *TypeMediaDTO {
	return &TypeMediaDTO{
		ID:   tm.ID,
		Name: tm.Name,
	}
}
