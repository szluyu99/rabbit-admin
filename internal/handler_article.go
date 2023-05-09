package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"github.com/szluyu99/rabbit-admin/internal/models"
)

func (m *ServerManager) articleObject() rabbit.WebObject {
	return rabbit.WebObject{
		Name:        "article",
		Model:       &models.Article{},
		GetDB:       m.getDB,
		Editables:   []string{"Title", "Desc", "Content", "Img", "Type", "Status", "IsTop", "IsDelete", "OriginalUrl"},
		Searchables: []string{"Name", "Desc"},
		Filterables: []string{"Type", "Status", "IsTop", "IsDelete"},
		OnCreate: func(c *gin.Context, vptr any) error {
			v := vptr.(*models.Article)
			user := rabbit.CurrentUser(c)
			v.UserId = user.ID
			return nil
		},
		AllowMethods: rabbit.BATCH | rabbit.DELETE | rabbit.EDIT,
	}
}

func (m *ServerManager) handleGetArticle(c *gin.Context) {
	key := c.Param("id")
	article, err := models.GetArticle(m.db, key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (m *ServerManager) handleArticleList(c *gin.Context) {
	var form ArticleQueryForm
	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, limit := CheckPageLimit(form.Page, form.Limit)
	list, count, err := models.GetArticleList(m.db, page, limit, form.Keyword, form.Type, form.Status, form.TagId, form.CategoryId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
