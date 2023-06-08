package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// GroupID   uint         `json:"group_id"`

	Name string `gorm:"type:varchar(20);not null" json:"name"`

	// for association
	Articles []Article `gorm:"foreignKey:CategoryId"`
	// Group     rabbit.Group `json:"-"`
}

// check if category is used by article
func (c *Category) BeforeDelete(tx *gorm.DB) error {
	var count int64
	result := tx.Model(&Article{}).Where("category_id", c.ID).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return errors.New("the category is used by article")
	}
	return nil
}

func GetCategoryList(db *gorm.DB, page, limit int, keyword string) ([]Category, int, error) {
	list := make([]Category, 0)
	db = db.Model(&Category{}).Preload("Articles")
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	return FindWithCount(db, list, (page-1)*limit, limit)
}
