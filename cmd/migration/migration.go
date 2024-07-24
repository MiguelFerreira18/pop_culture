package migration

import (
	a "pop_culture/api/resource/Attribute"
	m "pop_culture/api/resource/Media"
	mt "pop_culture/api/resource/MediaType"
	u "pop_culture/api/resource/User"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&u.User{}, &mt.TypeMedia{}, &m.Media{}, &a.Attribute{})
}
