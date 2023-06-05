package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db := rabbit.InitDatabase("", "", nil)
	MakeMigrate(db)
	return db
}

func TestGetArticleList(t *testing.T) {
	db := initDB()
	db.Create(&Article{Title: "test1", CategoryId: 1})
	db.Create(&Article{Title: "test2", CategoryId: 2, IsTop: true})
	db.Create(&Article{Title: "test3", CategoryId: 3})

	{
		// basic + order
		list, count, err := GetArticleList(db, 1, 10, "", "", "", false, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 3, count)
		assert.Equal(t, "test2", list[0].Title) // is_top desc

		// search category_id
		list, count, err = GetArticleList(db, 1, 10, "", "", "", false, 0, 3)
		assert.Nil(t, err)
		assert.Equal(t, 1, count)
		assert.Equal(t, "test3", list[0].Title)
	}

	{
		// search tag_id
		db.Create(&Tag{Name: "tag1"})
		err := db.Create(&ArticleTag{TagID: 1, ArticleID: 1}).Error
		assert.Nil(t, err)

		list, count, err := GetArticleList(db, 1, 10, "", "", "", false, 1, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, count)
		assert.Equal(t, "test1", list[0].Title)
	}
}

func TestSaveOrUpdateArticle(t *testing.T) {
	db := initDB()
	// save
	{
		err := SaveOrUpdateArticle(db, 0, "title", "content", "cover", "status", "typ", "originalUrl", 1, true, []string{"tag1", "tag2"})
		assert.Nil(t, err)

		_, count, err := GetArticleList(db, 1, 10, "", "", "", false, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, count)

		c, _ := Count[Tag](db)
		assert.Equal(t, 2, c)
	}
	{
		err := SaveOrUpdateArticle(db, 0, "title", "content", "cover", "status", "typ", "originalUrl", 1, true, []string{"tag1", "tag2"})
		assert.Nil(t, err)

		_, count, err := GetArticleList(db, 1, 10, "", "", "", false, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, count)

		c, _ := Count[Tag](db)
		assert.Equal(t, 2, c)
	}
	{
		// 1-1, 1-2, 2-1, 2-2
		c, _ := Count[ArticleTag](db)
		assert.Equal(t, 4, c)
	}
	// update
	{
		SaveOrUpdateArticle(db, 1, "newTitle", "content", "cover", "status", "typ", "originalUrl", 1, true, []string{"tag1", "tag2", "tag3"})

		a, err := rabbit.GetByID[Article](db, 1)
		assert.Nil(t, err)
		assert.Equal(t, "newTitle", a.Title)

		// 1-1, 1-2, 1-3, 2-1, 2-2
		c, _ := Count[ArticleTag](db)
		assert.Equal(t, 5, c)
	}
}
