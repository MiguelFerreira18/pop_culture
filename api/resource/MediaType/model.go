package mediatype

import (
	attribute "pop_culture/api/resource/Attribute"
	"pop_culture/util/validation"

	"gorm.io/gorm"
)

type TypeMedia struct {
	gorm.Model
	Name       string
	Attributes []attribute.Attribute `gorm:"many2many:typemedia_attribute;"`
}

type TypeMediaForm struct {
	Name string `json:"name"`
}
type TypeMediaDTO struct {
	ID         uint
	Name       string
	Attributes []attribute.AttributeDTO
}

type TypeMedias []*TypeMedia

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
	att := make([]attribute.AttributeDTO, 0, len(tm.Attributes))
	for _, a := range tm.Attributes {
		att = append(att, *a.ToDTO())
	}

	return &TypeMediaDTO{
		ID:         tm.ID,
		Name:       tm.Name,
		Attributes: att,
	}
}

func (tms TypeMedias) ToDTO() []*TypeMediaDTO {
	tpms := make([]*TypeMediaDTO, len(tms))
	for i, t := range tms {
		tpms[i] = t.ToDTO()
	}
	return tpms
}
