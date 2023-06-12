package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit-admin/internal/models"
)

func (m *ServerManager) articleObject() gormpher.WebObject {
	return gormpher.WebObject{
		Name:         "article",
		Model:        &models.Article{},
		GetDB:        m.getDB,
		Pagination:   true,
		EditFields:   []string{"Title", "Desc", "Content", "Img", "Type", "Status", "IsTop", "IsDelete", "OriginalUrl"},
		SearchFields: []string{"Name", "Desc"},
		FilterFields: []string{"Type", "Status", "IsTop", "IsDelete"},
		AllowMethods: gormpher.BATCH | gormpher.DELETE | gormpher.EDIT,
	}
}

func (m *ServerManager) handleGetArticle(c *gin.Context) {
	db := m.db.Preload("Tags").Preload("Category")
	gormpher.HandleGet[models.Article](c, db, nil)
}

func (m *ServerManager) handleQueryArticle(c *gin.Context) {
	var form ArticleQueryForm
	if err := c.BindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	page, limit := CheckPageLimit(form.Page, form.Limit)
	list, count, err := models.GetArticleList(m.db, page, limit, form.Keyword, form.Type, form.Status, form.IsDelete, form.TagId, form.CategoryId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, PageResult[models.Article]{
		Page:       page,
		Limit:      limit,
		Keyword:    form.Keyword,
		Items:      list,
		TotalCount: count,
	})
}

func (m *ServerManager) handleSaveOrUpdateArticle(c *gin.Context) {
	var form ArticleForm
	if err := c.BindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	if err := models.SaveOrUpdateArticle(m.db,
		form.ID,
		form.Title,
		form.Content,
		form.Cover,
		form.Type,
		form.Status,
		form.OriginalUrl,
		form.CategoryId,
		form.IsTop,
		form.TagNames,
	); err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, true)
}
