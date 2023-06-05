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
	}
}

func allCategoryQueryView() rabbit.QueryView {
	return rabbit.QueryView{
		Name:   "all",
		Method: http.MethodGet,
		Prepare: func(db *gorm.DB, c *gin.Context) (*gorm.DB, *rabbit.QueryForm, error) {
			return db, &rabbit.QueryForm{Page: 1, Limit: -1}, nil
		},
	}
}
