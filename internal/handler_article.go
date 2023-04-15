package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit-admin/internal/models"
)

func (m *ServerManager) articleObject() gormpher.WebObject {
	return gormpher.WebObject{
		Model:       &models.Article{},
		GetDB:       m.getDB,
		Editables:   []string{"Title", "Desc", "Content", "Img", "Type", "Status", "IsTop", "IsDelete", "OriginalUrl"},
		Searchables: []string{"Name", "Desc"},
		Filterables: []string{"Type", "Status", "IsTop", "IsDelete"},
		Orderables:  []string{"CreatedAt"},
	}
}

func (m *ServerManager) handlerArticleList(c *gin.Context) {
	var form PageForm
	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, limit := CheckPageLimit(form.Page, form.Limit)
	list, count, err := models.GetArticleList(m.db, page, limit, form.Keyword)
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
