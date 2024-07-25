package migration

import (
	"log"
	a "pop_culture/api/resource/Attribute"
	m "pop_culture/api/resource/Media"
	mt "pop_culture/api/resource/MediaType"
	mediatypeattribute "pop_culture/api/resource/MediaTypeAttribute"
	u "pop_culture/api/resource/User"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&u.User{}, &mt.TypeMedia{}, &m.Media{}, &a.Attribute{}, &mediatypeattribute.TypemediaAttribute{})
	if err := db.SetupJoinTable(&mt.TypeMedia{}, "Attributes", &mediatypeattribute.TypemediaAttribute{}); err != nil {
		log.Fatal(err)
		return
	}
}
