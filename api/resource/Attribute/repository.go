package attribute

import "gorm.io/gorm"

type AttributeRepository struct {
	db *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) *AttributeRepository {
	return &AttributeRepository{
		db: db,
	}
}

func (r AttributeRepository) List() (Attributes, error) {
	attributes := make([]*Attribute, 0)
	if err := r.db.Find(&attributes).Error; err != nil {
		return nil, err
	}

	return attributes, nil
}

func (r AttributeRepository) Create(a *Attribute) (*Attribute, error) {

	if err := r.db.Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (r AttributeRepository) Read(id uint) (*Attribute, error) {
	attribute := &Attribute{}
	if err := r.db.Where("id = ?", id).First(&attribute).Error; err != nil {
		return nil, err
	}
	return attribute, nil
}

func (r AttributeRepository) Update(a *Attribute) (int64, error) {
	result := r.db.Model(a).Select("Name", "UpdatedAt").Where("id = ?", a.ID).Updates(a)

	return result.RowsAffected, result.Error
}

func (r AttributeRepository) Delete(id uint) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Attribute{})
	return result.RowsAffected, result.Error
}
