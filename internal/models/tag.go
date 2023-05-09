package models

import (
	"errors"
	"time"

	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

type ArticleTag struct {
	TagID     uint `gorm:"primaryKey"`
	ArticleID uint `gorm:"primaryKey"`
}

type Tag struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	GroupID   uint         `json:"group_id"`
	Group     rabbit.Group `json:"-"`

	Name string `json:"name" gorm:"type:varchar(20);not null;uniqueIndex"`

	// Articles []*Article `gorm:"many2many:article_tag;" json:"articles"`
}

// check if tag is used by article
func (t *Tag) BeforeDelete(tx *gorm.DB) error {
	var count int64
	result := tx.Model(&ArticleTag{}).Where("tag_id = ?", t.ID).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return errors.New("tag is used by article")
	}
	return nil
}

func CheckTagExistByName(db *gorm.DB, name string) bool {
	var tag Tag
	result := db.Model(&Tag{}).Where("name = ?", name).First(&tag)
	return result.Error == nil
}

// create tags by names, if tag exist, ignore
func CreateTagsByNames(db *gorm.DB, names []string) error {
	for _, name := range names {
		if CheckTagExistByName(db, name) {
			continue
		}
		tag := Tag{Name: name}
		result := db.Model(&Tag{}).Create(&tag)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func GetTagIdsByNames(db *gorm.DB, names []string) ([]uint, error) {
	var ids []uint
	result := db.Model(&Tag{}).Where("name in ?", names).Pluck("id", &ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return ids, nil
}
