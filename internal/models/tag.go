package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ArticleTag struct {
	TagID     uint `gorm:"primaryKey"`
	ArticleID uint `gorm:"primaryKey"`
}

type Tag struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// GroupID   uint         `json:"group_id"`

	Name string `json:"name" gorm:"type:varchar(20);not null;uniqueIndex"`

	// for association
	Articles []*Article `json:"articles" gorm:"many2many:article_tags;"`
	// Group     rabbit.Group `json:"-"`
}

// check if tag is used by article
func (t *Tag) BeforeDelete(tx *gorm.DB) error {
	var count int64
	result := tx.Model(&ArticleTag{}).Where("tag_id", t.ID).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return errors.New("the tag is used by article")
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

func GetTagList(db *gorm.DB, page, limit int, keyword string) ([]Tag, int, error) {
	list := make([]Tag, 0)
	db = db.Model(&Tag{}).Preload("Articles")
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	return FindWithCount(db, list, (page-1)*limit, limit)
}
