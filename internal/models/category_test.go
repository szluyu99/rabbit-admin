package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategoryList(t *testing.T) {
	db := initDB()

	db.Create(&Category{Name: "c1"})
	db.Create(&Article{Title: "a1", CategoryId: 1})
	db.Create(&Article{Title: "a2", CategoryId: 1})
	db.Create(&Article{Title: "a3", CategoryId: 1})

	list, count, err := GetCategoryList(db, 1, 10, "")
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
	assert.Equal(t, "c1", list[0].Name)
	assert.Equal(t, 3, len(list[0].Articles))
}
