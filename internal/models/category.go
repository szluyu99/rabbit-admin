package models

import (
	"errors"
	"time"

	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

type Category struct {
	ID        int          `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	GroupID   uint         `json:"group_id"`
	Group     rabbit.Group `json:"-"`

	Name string `gorm:"type:varchar(20);not null" json:"name"`

	// Articles []Article `gorm:"foreignKey:CategoryId"`
}

// check if category is used by article
func (c *Category) BeforeDelete(tx *gorm.DB) error {
	var count int64
	result := tx.Model(&Article{}).Where("category_id = ?", c.ID).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return errors.New("category is used by article")
	}
	return nil
}
