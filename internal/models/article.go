package models

import (
	"time"

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
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// GroupID   uint      `json:"group_id"`

	Title       string `json:"title" gorm:"type:varchar(100);not null"`
	Desc        string `json:"desc" gorm:"type:varchar(200)"`
	Content     string `json:"content"`
	Cover       string `json:"cover" gorm:"type:varchar(100)"`
	Type        string `json:"type" gorm:"type:varchar(20)"`   // original | reprint | translate
	Status      string `json:"status" gorm:"type:varchar(20)"` // public | private
	IsTop       bool   `json:"is_top"`
	IsDelete    bool   `json:"is_delete"`
	OriginalUrl string `json:"original_url" gorm:"type:varchar(100)"`

	CategoryId uint `json:"category_id"`
	// UserId     uint `json:"user_id"`

	// User *rabbit.User `json:"user" gorm:"foreignkey:UserId"`
	// Group    *rabbit.Group `json:"-"`
	Category *Category `json:"category" gorm:"foreignKey:CategoryId"`
	Tags     []*Tag    `json:"tags" gorm:"many2many:article_tag;"`
}

// delete article_tag
func (a *Article) BeforeDelete(tx *gorm.DB) error {
	return tx.Where("article_id = ?", a.ID).Delete(&ArticleTag{}).Error
}

// 1. Create tags by names
// 2. Create/Update article
// 3. Clear article_tag by article_id
// 4. Create article_tag by tag_id and article_id
func SaveOrUpdateArticle(db *gorm.DB, id uint, title, content, cover, typ, status, originalUrl string, categoryID uint, isTop bool, tagNames []string) error {
	// 1
	if err := CreateTagsByNames(db, tagNames); err != nil {
		return err
	}

	// 2
	article := Article{
		ID:          id,
		Title:       title,
		Content:     content,
		Cover:       cover,
		Type:        typ,
		Status:      status,
		OriginalUrl: originalUrl,
		CategoryId:  categoryID,
		IsTop:       isTop,
	}

	if article.ID != 0 {
		if err := db.Model(&article).
			Select("*").
			Omit("user_id", "created_at").
			Updates(&article).Error; err != nil {
			return err
		}
	} else {
		if err := db.Create(&article).Error; err != nil {
			return err
		}
	}

	// 3
	if err := db.Delete(&ArticleTag{}, "article_id", article.ID).Error; err != nil {
		return err
	}

	// 4
	ids, err := GetTagIdsByNames(db, tagNames)
	if err != nil {
		return err
	}
	for _, id := range ids {
		result := db.Create(&ArticleTag{
			TagID:     id,
			ArticleID: article.ID,
		})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func GetArticleList(db *gorm.DB, page, limit int, keyword, typ, status string, isDelete bool, tagId, categoryId int) ([]Article, int, error) {
	var list = make([]Article, 0)

	db = db.Model(&Article{})
	if keyword != "" {
		db = db.Where("title LIKE ? OR desc LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if typ != "" {
		db = db.Where("type = ?", typ)
	}
	if status != "" {
		db = db.Where("status = ?", status)
	}
	if categoryId > 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	db.Preload("Category").Preload("Tags")
	db.Joins("LEFT JOIN article_tags ON article_tags.article_id = articles.id").Group("articles.id")

	if tagId > 0 {
		db = db.Where("tag_id", tagId)
	}

	db = db.Where("is_delete = ?", isDelete)

	var count int64
	db.Count(&count)
	db.Offset((page - 1) * limit).Limit(limit)
	db.Order("is_top DESC, created_at DESC")
	result := db.Find(&list)

	if result.Error != nil {
		return list, 0, result.Error
	}

	return list, int(count), nil
}
