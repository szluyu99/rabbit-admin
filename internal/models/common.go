package models

import (
	"time"

	"gorm.io/gorm"
)

type Universal struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MakeMigrate(db *gorm.DB) error {
	if err := db.SetupJoinTable(&Article{}, "Tags", &ArticleTag{}); err != nil {
		return err
	}

	return db.AutoMigrate(
		&Tag{},
		&Category{},
		&Article{},
	)
}
