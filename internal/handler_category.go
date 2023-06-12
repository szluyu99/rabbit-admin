package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit-admin/internal/models"
	"gorm.io/gorm"
)

func (m *ServerManager) categoryObject() gormpher.WebObject {
	return gormpher.WebObject{
		Name:         "category",
		Model:        &models.Category{},
		GetDB:        m.getDB,
		Pagination:   true,
		EditFields:   []string{"Name"},
		SearchFields: []string{"Name"},
		FilterFields: []string{"Name"},
		OrderFields:  []string{"CreatedAt"},
		Views: []gormpher.QueryView{
			allCategoryQueryView(),
		},
		AllowMethods: gormpher.GET | gormpher.CREATE | gormpher.DELETE | gormpher.EDIT | gormpher.BATCH,
	}
}

// GET /category/all
func allCategoryQueryView() gormpher.QueryView {
	return gormpher.QueryView{
		Name:   "all",
		Method: http.MethodGet,
		Prepare: func(db *gorm.DB, c *gin.Context, pagination bool) (*gorm.DB, *gormpher.QueryForm, error) {
			return db, &gormpher.QueryForm{Pos: 1, Limit: -1}, nil
		},
	}
}

func (m *ServerManager) handleQueryCategory(c *gin.Context) {
	var form gormpher.QueryForm
	if err := c.ShouldBindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	list, total, err := models.GetCategoryList(m.db, form.Pos, form.Limit, form.Keyword)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	resultList := make([]CategoryQueryResult, len(list))
	for i, item := range list {
		resultList[i] = CategoryQueryResult{
			Category: models.Category{
				ID:        item.ID,
				Name:      item.Name,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				Articles:  item.Articles,
			},
			ArticleCount: len(item.Articles),
		}
	}

	c.JSON(http.StatusOK, PageResult[CategoryQueryResult]{
		Page:       form.Pos,
		Limit:      form.Limit,
		Keyword:    form.Keyword,
		TotalCount: total,
		Items:      resultList,
	})
}
