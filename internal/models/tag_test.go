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
