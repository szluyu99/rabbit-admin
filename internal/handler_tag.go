package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit-admin/internal/models"
	"gorm.io/gorm"
)

func (m *ServerManager) tagObject() gormpher.WebObject {
	return gormpher.WebObject{
		Name:         "tag",
		Model:        &models.Tag{},
		GetDB:        m.getDB,
		Pagination:   true,
		EditFields:   []string{"Name"},
		SearchFields: []string{"Name"},
		FilterFields: []string{"Name"},
		OrderFields:  []string{"CreatedAt"},
		Views: []gormpher.QueryView{
			allTagQueryView(),
		},
		AllowMethods: gormpher.GET | gormpher.CREATE | gormpher.DELETE | gormpher.EDIT | gormpher.BATCH,
	}
}

func allTagQueryView() gormpher.QueryView {
	return gormpher.QueryView{
		Name:   "all",
		Method: http.MethodGet,
		Prepare: func(db *gorm.DB, c *gin.Context, pagination bool) (*gorm.DB, *gormpher.QueryForm, error) {
			return db, &gormpher.QueryForm{Pos: 1, Limit: -1}, nil
		},
	}
}

func (m *ServerManager) handleQueryTag(c *gin.Context) {
	var form gormpher.QueryForm
	if err := c.ShouldBindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	list, total, err := models.GetTagList(m.db, form.Pos, form.Limit, form.Keyword)
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
		Page:       form.Pos,
		Limit:      form.Limit,
		Keyword:    form.Keyword,
		TotalCount: total,
		Items:      resultList,
	})
}
