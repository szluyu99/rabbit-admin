package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"github.com/szluyu99/rabbit-admin/internal/models"
	"gorm.io/gorm"
)

func (m *ServerManager) categoryObject() rabbit.WebObject {
	return rabbit.WebObject{
		Name:        "category",
		Model:       &models.Category{},
		GetDB:       m.getDB,
		Editables:   []string{"Name"},
		Searchables: []string{"Name"},
		Filterables: []string{"Name"},
		Orderables:  []string{"CreatedAt"},
		Views: []rabbit.QueryView{
			allCategoryQueryView(),
		},
		AllowMethods: rabbit.GET | rabbit.CREATE | rabbit.DELETE | rabbit.EDIT | rabbit.BATCH,
	}
}

// GET /category/all
func allCategoryQueryView() rabbit.QueryView {
	return rabbit.QueryView{
		Name:   "all",
		Method: http.MethodGet,
		Prepare: func(db *gorm.DB, c *gin.Context) (*gorm.DB, *rabbit.QueryForm, error) {
			return db, &rabbit.QueryForm{Page: 1, Limit: -1}, nil
		},
	}
}

func (m *ServerManager) handleQueryCategory(c *gin.Context) {
	var form rabbit.QueryForm
	if err := c.ShouldBindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	list, total, err := models.GetCategoryList(m.db, form.Page, form.Limit, form.Keyword)
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
		Page:       form.Page,
		Limit:      form.Limit,
		Keyword:    form.Keyword,
		TotalCount: total,
		Items:      resultList,
	})
}
