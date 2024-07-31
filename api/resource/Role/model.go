package role

import (
	"pop_culture/util/validation"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string
	Description string
}

type RoleDTO struct {
	ID   uint
	Name string
}

type RoleForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewRole(name string, description string) (*Role, error) {
	validatedName, err := validation.RoleNameRules(name)
	if err != nil {
		return nil, err
	}
	validatedDescription, err := validation.RoleDescriptionRules(description)
	if err != nil {
		return nil, err
	}

	return &Role{
		Name:        *validatedName,
		Description: *validatedDescription,
	}, nil

}

func (r Role) ToDTO() *RoleDTO {
	return &RoleDTO{
		ID:   r.ID,
		Name: r.Name,
	}
}
