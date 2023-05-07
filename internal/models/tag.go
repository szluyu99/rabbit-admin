package models

import (
	"time"

	"github.com/szluyu99/rabbit"
)

type Tag struct {
	ID        int          `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	GroupID   uint         `json:"group_id"`
	Group     rabbit.Group `json:"-"`

	Name string `json:"name" gorm:"type:varchar(20);not null"`
	// Articles []*Article `gorm:"many2many:article_tag;" json:"articles"`
}
