package rabbitadmin

import "github.com/szluyu99/rabbit-admin/internal/models"

func CheckPageLimit(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit <= 1 || limit > 200 {
		limit = 50
	}
	return page, limit
}

type PageForm struct {
	Page    int    `json:"page"`  // current page
	Limit   int    `json:"limit"` // page size
	Keyword string `json:"keyword,omitempty"`
}

type PageResult[T any] struct {
	Page       int    `json:"page,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Keyword    string `json:"keyword,omitempty"`
	TotalCount int    `json:"total,omitempty"`
	Items      []T    `json:"items"`
}

type ArticleQueryForm struct {
	PageForm
	TagId      int    `json:"tag_id"`
	CategoryId int    `json:"category_id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	IsDelete   bool   `json:"is_delete"`
}

type ArticleForm struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Desc        string   `json:"desc"`
	Content     string   `json:"content"`
	Cover       string   `json:"cover"`
	Type        string   `json:"type"`   // original | reprint | translate
	Status      string   `json:"status"` // public | private | draft
	IsTop       bool     `json:"is_top"`
	IsDelete    bool     `json:"is_delete"`
	OriginalUrl string   `json:"original_url"`
	CategoryId  uint     `json:"category_id"`
	TagNames    []string `json:"tag_names"`
}

type RoleForm struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Label         string `json:"label"`
	PermissionIds []uint `json:"permission_ids"`
}

type UserForm struct {
	ID      uint   `json:"id"`
	RoleIds []uint `json:"role_ids"`
}

type CategoryQueryResult struct {
	models.Category
	ArticleCount int `json:"article_count"`
}

type TagQueryResult struct {
	models.Tag
	ArticleCount int `json:"article_count"`
}
