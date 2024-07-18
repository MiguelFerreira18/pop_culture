package migration

import (
	u "pop_culture/api/resource/User"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&u.User{})
}
