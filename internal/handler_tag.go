package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"github.com/szluyu99/rabbit-admin/internal/models"
	"gorm.io/gorm"
)

func (m *ServerManager) tagObject() rabbit.WebObject {
	return rabbit.WebObject{
		Name:        "tag",
		Model:       &models.Tag{},
		GetDB:       m.getDB,
		Editables:   []string{"Name"},
		Searchables: []string{"Name"},
		Filterables: []string{"Name"},
		Orderables:  []string{"CreatedAt"},
		Views: []rabbit.QueryView{
			allTagQueryView(),
		},
		AllowMethods: rabbit.GET | rabbit.CREATE | rabbit.DELETE | rabbit.EDIT | rabbit.BATCH,
	}
}

func allTagQueryView() rabbit.QueryView {
	return rabbit.QueryView{
		Name:   "all",
		Method: http.MethodGet,
		Prepare: func(db *gorm.DB, c *gin.Context) (*gorm.DB, *rabbit.QueryForm, error) {
			return db, &rabbit.QueryForm{Page: 1, Limit: -1}, nil
		},
	}
}

func (m *ServerManager) handleQueryTag(c *gin.Context) {
	var form rabbit.QueryForm
	if err := c.ShouldBindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	list, total, err := models.GetTagList(m.db, form.Page, form.Limit, form.Keyword)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	resultList := make([]TagQueryResult, len(list))
	for i, item := range list {
		resultList[i] = TagQueryResult{
			Tag: models.Tag{
				ID:        item.ID,
				Name:      item.Name,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				Articles:  item.Articles,
			},
			ArticleCount: len(item.Articles),
		}
	}

	c.JSON(http.StatusOK, PageResult[TagQueryResult]{
		Page:       form.Page,
		Limit:      form.Limit,
		Keyword:    form.Keyword,
		TotalCount: total,
		Items:      resultList,
	})
}
