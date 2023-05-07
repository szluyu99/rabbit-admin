package models

import (
	"time"

	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

const (
	ArticleStatusPublic  = "public"
	ArticleStatusPrivate = "private"

	ArticleTypeOriginal  = "original"
	ArticleTypeReprint   = "reprint"
	ArticleTypeTranslate = "translate"
)

// article belongTo category
// article belongTo user
// article many2many tag
type Article struct {
	ID        int           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	GroupID   uint          `json:"group_id"`
	Group     *rabbit.Group `json:"-"`

	Title       string `json:"title" gorm:"type:varchar(100);not null"`
	Desc        string `json:"desc" gorm:"type:varchar(200)"`
	Content     string `json:"content"`
	Img         string `json:"img" gorm:"type:varchar(100)"`
	Type        string `json:"type" gorm:"type:varchar(20)"`   // original | reprint | translate
	Status      string `json:"status" gorm:"type:varchar(20)"` // public | private
	IsTop       bool   `json:"is_top"`
	IsDelete    bool   `json:"is_delete"`
	OriginalUrl string `json:"original_url" gorm:"type:varchar(100)"`

	CategoryId int          `json:"category_id"`
	Category   *Category    `json:"category" gorm:"foreignkey:CategoryId"`
	UserId     int          `json:"user_id"`
	User       *rabbit.User `json:"user" gorm:"foreignkey:UserId"`
	// Tags []*Tag `gorm:"many2many:article_tag;" json:"tags"`
}

func GetArticleList(db *gorm.DB, page, limit int, keyword string) ([]Article, int, error) {
	keys := map[string]string{
		"title": keyword,
		"desc":  keyword,
	}
	return gormpher.ListPageKeyword[Article](db, page, limit, keys)
}
