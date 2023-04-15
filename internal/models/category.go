package models

import (
	"time"

	"github.com/restsend/rabbit"
)

type Category struct {
	ID        int          `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	GroupID   uint         `json:"group_id"`
	Group     rabbit.Group `json:"-"`

	Name string `gorm:"type:varchar(20);not null" json:"name"`

	// Articles []Article `gorm:"foreignKey:CategoryId"` // 重写外键
}
