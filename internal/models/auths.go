package models

import (
	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

// TODO: load from config
func InitDefaultPermissions(db *gorm.DB) error {
	// role
	{
		p := rabbit.Permission{Name: "role", ParentID: 0}
		if err := db.Where("name", p.Name).FirstOrCreate(&p).Error; err != nil {
			return err
		}
		ps := []*rabbit.Permission{
			{Name: "role.query", Uri: "/role", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "role.create", Uri: "/role", Method: "PUT", ParentID: p.ID},
			{Name: "role.delete", Uri: "/role/:key", Method: "DELETE", ParentID: p.ID},
			{Name: "role.update", Uri: "/role/:key", Method: "PATCH", ParentID: p.ID},
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
			{Name: "permission.query", Uri: "/permission", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "permission.create", Uri: "/permission", Method: "PUT", ParentID: p.ID},
			{Name: "permission.delete", Uri: "/permission/:key", Method: "DELETE", ParentID: p.ID},
			{Name: "permission.update", Uri: "/permission/:key", Method: "PATCH", ParentID: p.ID},
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
			{Name: "user.query", Uri: "/user", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "user.role", Uri: "/user/role", Method: "PATCH", ParentID: p.ID},
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
			{Name: "config.query", Uri: "/config", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "config.create", Uri: "/config", Method: "PUT", ParentID: p.ID},
			{Name: "config.delete", Uri: "/config/:key", Method: "DELETE", ParentID: p.ID},
			{Name: "config.update", Uri: "/config/:key", Method: "PATCH", ParentID: p.ID},
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
			{Name: "article.get", Uri: "/article/:key", Method: "GET", ParentID: p.ID, Anonymous: true},
			{Name: "article.query", Uri: "/article", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "article.create", Uri: "/article", Method: "PUT", ParentID: p.ID},
			{Name: "article.save_or_update", Uri: "/article/save_or_update", Method: "POST", ParentID: p.ID},
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
			{Name: "tag.get", Uri: "/tag/:key", Method: "GET", ParentID: p.ID, Anonymous: true},
			{Name: "tag.query", Uri: "/tag", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "tag.create", Uri: "/tag", Method: "PUT", ParentID: p.ID},
			{Name: "tag.update", Uri: "/tag", Method: "PATCH", ParentID: p.ID},
			{Name: "tag.all", Uri: "/tag/all", Method: "GET", ParentID: p.ID, Anonymous: true},
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
			{Name: "category.get", Uri: "/category/:key", Method: "GET", ParentID: p.ID, Anonymous: true},
			{Name: "category.query", Uri: "/category", Method: "POST", ParentID: p.ID, Anonymous: true},
			{Name: "category.create", Uri: "/category", Method: "PUT", ParentID: p.ID},
			{Name: "category.update", Uri: "/category", Method: "PATCH", ParentID: p.ID},
			{Name: "category.all", Uri: "/category/all", Method: "GET", ParentID: p.ID, Anonymous: true},
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
		{Name: name + ".get", Uri: "/" + name, Method: "GET", ParentID: p.ID},
		{Name: name + ".query", Uri: "/" + name, Method: "POST", ParentID: p.ID},
		{Name: name + ".create", Uri: "/" + name, Method: "PUT", ParentID: p.ID},
		{Name: name + ".delete", Uri: "/" + name, Method: "DELETE", ParentID: p.ID},
		{Name: name + ".update", Uri: "/" + name, Method: "PATCH", ParentID: p.ID},
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
