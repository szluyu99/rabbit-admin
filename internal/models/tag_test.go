package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckTagExistByName(t *testing.T) {
	db := initDB()
	db.Create(&Tag{Name: "test"})
	assert.True(t, CheckTagExistByName(db, "test"))
	assert.False(t, CheckTagExistByName(db, "not_exist"))
}

func TestCreateTagsByNames(t *testing.T) {
	db := initDB()
	CreateTagsByNames(db, []string{"test1", "test2", "test3"})
	c, _ := Count[Tag](db)
	assert.Equal(t, 3, c)

	CreateTagsByNames(db, []string{"test1", "test2", "test3", "test4"})
	c, _ = Count[Tag](db)
	assert.Equal(t, 4, c)
}

func TestGetTagIdsByNames(t *testing.T) {
	db := initDB()

	CreateTagsByNames(db, []string{"test1", "test2", "test3"})

	ids, err := GetTagIdsByNames(db, []string{"test1", "test2", "test3"})
	assert.Nil(t, err)
	assert.Equal(t, []uint{1, 2, 3}, ids)
}

func TestGetTagList(t *testing.T) {
	db := initDB()

	db.Create(&Tag{Name: "t1"})
	db.Create(&Article{Title: "a1"})
	db.Create(&Article{Title: "a2"})
	db.Create(&Article{Title: "a3"})
	db.Create(&ArticleTag{TagID: 1, ArticleID: 1})
	db.Create(&ArticleTag{TagID: 1, ArticleID: 2})
	db.Create(&ArticleTag{TagID: 1, ArticleID: 3})

	list, count, err := GetTagList(db, 1, 10, "")
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
	assert.Equal(t, "t1", list[0].Name)
	assert.Equal(t, 3, len(list[0].Articles))
}
