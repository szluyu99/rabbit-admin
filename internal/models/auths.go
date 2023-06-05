package models

import (
	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

func InitDefaultPermissions(db *gorm.DB) error {
	// role
	{
		p := rabbit.Permission{Name: "role", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "role.query", P1: "/role", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "role.create", P1: "/role", P2: "PUT", ParentID: p.ID},
			{Name: "role.delete", P1: "/role/:key", P2: "DELETE", ParentID: p.ID},
			{Name: "role.update", P1: "/role/:key", P2: "PATCH", ParentID: p.ID},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}
	// permission
	{
		p := rabbit.Permission{Name: "permission", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "permission.query", P1: "/permission", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "permission.create", P1: "/permission", P2: "PUT", ParentID: p.ID},
			{Name: "permission.delete", P1: "/permission/:key", P2: "DELETE", ParentID: p.ID},
			{Name: "permission.update", P1: "/permission/:key", P2: "PATCH", ParentID: p.ID},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}
	// user
	{
		p := rabbit.Permission{Name: "user", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "user.query", P1: "/user", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "user.role", P1: "/user/role", P2: "PATCH", ParentID: p.ID},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}
	// config
	{
		p := rabbit.Permission{Name: "config", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "config.query", P1: "/config", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "config.create", P1: "/config", P2: "PUT", ParentID: p.ID},
			{Name: "config.delete", P1: "/config/:key", P2: "DELETE", ParentID: p.ID},
			{Name: "config.update", P1: "/config/:key", P2: "PATCH", ParentID: p.ID},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}
	// article
	{
		p := rabbit.Permission{Name: "article", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "article.get", P1: "/article/:key", P2: "GET", ParentID: p.ID, Anonymous: true},
			{Name: "article.query", P1: "/article", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "article.create", P1: "/article", P2: "PUT", ParentID: p.ID},
			{Name: "article.save_or_update", P1: "/article/save_or_update", P2: "POST", ParentID: p.ID},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}
	// tag
	{
		p := rabbit.Permission{Name: "tag", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "tag.get", P1: "/tag/:key", P2: "GET", ParentID: p.ID, Anonymous: true},
			{Name: "tag.query", P1: "/tag", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "tag.create", P1: "/tag", P2: "PUT", ParentID: p.ID},
			{Name: "tag.update", P1: "/tag", P2: "PATCH", ParentID: p.ID},
			{Name: "tag.all", P1: "/tag/all", P2: "GET", ParentID: p.ID, Anonymous: true},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}
	// category
	{
		p := rabbit.Permission{Name: "category", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}

		ps := []*rabbit.Permission{
			{Name: "category.get", P1: "/category/:key", P2: "GET", ParentID: p.ID, Anonymous: true},
			{Name: "category.query", P1: "/category", P2: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "category.create", P1: "/category", P2: "PUT", ParentID: p.ID},
			{Name: "category.update", P1: "/category", P2: "PATCH", ParentID: p.ID},
			{Name: "category.all", P1: "/category/all", P2: "GET", ParentID: p.ID, Anonymous: true},
		}
		for _, v := range ps {
			if err := db.Where("name", v.Name).FirstOrCreate(v).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func InitDefaultRoles(db *gorm.DB, roles []*rabbit.Role) error {
	for _, role := range roles {
		if err := db.Where("name", role.Name).FirstOrCreate(role).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateWebObjectPermissions(db *gorm.DB, name string) (*rabbit.Permission, error) {
	p := rabbit.Permission{Name: name, ParentID: 0}
	result := db.Where("name", p.Name).FirstOrCreate(&p)
	if result.Error != nil {
		return nil, result.Error
	}

	ps := []*rabbit.Permission{
		{Name: name + ".get", P1: "/" + name, P2: "GET", ParentID: p.ID},
		{Name: name + ".query", P1: "/" + name, P2: "POST", ParentID: p.ID},
		{Name: name + ".create", P1: "/" + name, P2: "PUT", ParentID: p.ID},
		{Name: name + ".delete", P1: "/" + name, P2: "DELETE", ParentID: p.ID},
		{Name: name + ".update", P1: "/" + name, P2: "PATCH", ParentID: p.ID},
	}

	for _, v := range ps {
		result := db.Where("name", v.Name).FirstOrCreate(v)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return &p, nil
}

func GetRoleList(db *gorm.DB, page, limit int, keyword string) ([]rabbit.Role, int, error) {
	var list = make([]rabbit.Role, 0)

	db = db.Model(&rabbit.Role{})
	db = db.Preload("Users").Preload("Permissions")
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	return FindWithCount(db, list, (page-1)*limit, limit)
}

func GetUserList(db *gorm.DB, page, limit int, keyword string) ([]rabbit.User, int, error) {
	var list = make([]rabbit.User, 0)

	db = db.Model(&rabbit.User{})
	db = db.Preload("Roles")

	if keyword != "" {
		db = db.Where("email LIKE ?", "%"+keyword+"%")
	}

	return FindWithCount(db, list, (page-1)*limit, limit)
}

func GetPermissionTreeList(db *gorm.DB, keyword string) ([]rabbit.Permission, error) {
	var list = make([]rabbit.Permission, 0)

	db = db.Model(&rabbit.Permission{})
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	result := db.Find(&list)
	if result.Error != nil {
		return list, result.Error
	}

	// build tree
	var treeList = make([]rabbit.Permission, 0)
	for _, v := range list {
		if v.ParentID == 0 {
			treeList = append(treeList, v)
		}
	}

	for i := 0; i < len(treeList); i++ {
		for j := 0; j < len(list); j++ {
			if list[j].ParentID == treeList[i].ID {
				treeList[i].Children = append(treeList[i].Children, &list[j])
			}
		}
	}

	return treeList, nil
}
