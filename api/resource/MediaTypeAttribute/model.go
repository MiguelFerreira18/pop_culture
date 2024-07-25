package mediatypeattribute

type TypemediaAttribute struct {
	ID          uint `gorm:"primaryKey;autoIncrement:true"`
	TypeMediaID uint `gorm:"primaryKey;uniqueIndex:idx_typemediaattribute;autoIncrement:false"`
	AttributeID uint `gorm:"primaryKey;uniqueIndex:idx_typemediaattribute;autoIncrement:false"`
}

func NewTypeMediaAttribute(typeMediaID uint, attributeID uint) *TypemediaAttribute {
	return &TypemediaAttribute{
		TypeMediaID: typeMediaID,
		AttributeID: attributeID,
	}
}
