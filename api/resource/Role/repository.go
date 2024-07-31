package role

import "gorm.io/gorm"

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r RoleRepository) Create(role *Role) (*Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r RoleRepository) Read(id uint) (*Role, error) {
	role := Role{}
	if err := r.db.Where("id = ?", id).First(role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r RoleRepository) Update(role *Role) (int64, error) {
	result := r.db.Model(role).Select("Name", "Description", "UpdatedAt").Where("id = ?", role.ID).Updates(role)
	return result.RowsAffected, result.Error
}

func (r RoleRepository) Delete(id uint) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Role{})
	return result.RowsAffected, result.Error
}
