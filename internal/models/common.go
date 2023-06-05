package models

import (
	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

func MakeMigrate(db *gorm.DB) error {
	if err := db.SetupJoinTable(&Article{}, "Tags", &ArticleTag{}); err != nil {
		return err
	}

	return db.AutoMigrate(
		&Tag{},
		&Category{},
		&Article{},
		&rabbit.Permission{},
		&rabbit.Role{},
		&rabbit.Group{},
	)
}

func FindWithCount[T any](db *gorm.DB, list []T, pos, limit int) ([]T, int, error) {
	var count int64
	result := db.Count(&count)
	if result.Error != nil {
		return list, 0, result.Error
	}

	result = db.Offset((pos - 1) * limit).Limit(limit).Find(&list)
	if result.Error != nil {
		return list, 0, result.Error
	}

	return list, int(count), nil
}

func Count[T any](db *gorm.DB, where ...any) (int, error) {
	var count int64
	if len(where) > 0 {
		db = db.Where(where[0], where[1:]...)
	}
	result := db.Model(new(T)).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}
