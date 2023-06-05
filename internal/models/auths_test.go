package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szluyu99/rabbit"
)

func TestInitDefaultAuth(t *testing.T) {
	db := initDB()

	// create default roles
	{
		roles := []*rabbit.Role{
			{Name: "admin", Label: "Admin"},
			{Name: "user", Label: "User"},
		}

		err := InitDefaultRoles(db, roles)
		assert.Nil(t, err)

		var count int64
		db.Model(&rabbit.Role{}).Count(&count)
		assert.Equal(t, int64(2), count)
	}

	// with default, should not create again
	{
		roles := []*rabbit.Role{
			{Name: "admin", Label: "Admin"},
			{Name: "user", Label: "User"},
			{Name: "test", Label: "Test"},
		}

		err := InitDefaultRoles(db, roles)
		assert.Nil(t, err)

		var count int64
		db.Model(&rabbit.Role{}).Count(&count)
		assert.Equal(t, int64(3), count)
	}
}

func TestGetRoleList(t *testing.T) {
	db := initDB()
	rabbit.CreateRoleWithPermissions(db, "role1", "ROLE1", []*rabbit.Permission{
		{Name: "p1"},
		{Name: "p2"},
	})

	rabbit.CreateRoleWithPermissions(db, "role2", "ROLE2", []*rabbit.Permission{
		{Name: "p3"},
		{Name: "p4"},
	})

	list, count, err := GetRoleList(db, 1, 10, "")
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
	assert.Equal(t, "role1", list[0].Name)
	assert.Equal(t, "role2", list[1].Name)
}

func TestGetPermissionTreeList(t *testing.T) {
	db := initDB()

	_, err := CreateWebObjectPermissions(db, "role")
	assert.Nil(t, err)

	list, err := GetPermissionTreeList(db, "")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(list))
	assert.Equal(t, 5, len(list[0].Children))
}

func TestCreateWebObjectPermissions(t *testing.T) {
	db := initDB()

	// create
	{
		_, err := CreateWebObjectPermissions(db, "role")
		assert.Nil(t, err)

		var count int64
		db.Model(&rabbit.Permission{}).Count(&count)
		assert.Equal(t, int64(6), count)
	}

	// with default, should not create again
	{
		_, err := CreateWebObjectPermissions(db, "role")
		assert.Nil(t, err)

		var count int64
		db.Model(&rabbit.Permission{}).Count(&count)
		assert.Equal(t, int64(6), count)
	}
}
