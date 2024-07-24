package attribute

import (
	"pop_culture/util/validation"

	"gorm.io/gorm"
)

type Attribute struct {
	gorm.Model
	Name string
}

type AttributeDTO struct {
	ID   uint
	Name string
}

type AttributeForm struct {
	Name string `json:"name"`
}

func NewAttribute(name string) (*Attribute, error) {

	validName, err := validation.AttributeNameRules(name)
	if err != nil {
		return nil, err
	}

	return &Attribute{
		Name: *validName,
	}, nil

}

func (a Attribute) ToDTO() *AttributeDTO {

	return &AttributeDTO{
		ID:   a.ID,
		Name: a.Name,
	}

}
