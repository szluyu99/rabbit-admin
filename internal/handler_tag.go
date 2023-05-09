package rabbitadmin

import (
	"errors"
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
			{
				Name:   "all",
				Method: http.MethodGet,
				Prepare: func(db *gorm.DB, c *gin.Context) (*gorm.DB, *rabbit.QueryForm, error) {
					return db, &rabbit.QueryForm{Page: 1, Limit: -1}, nil
				},
			},
		},
		OnBatchDelete: func(ctx *gin.Context, ids []string) error {
			db := ctx.MustGet(rabbit.DbField).(*gorm.DB)
			var count int64
			result := db.Model(&models.ArticleTag{}).Where("tag_id IN ?", ids).Count(&count)
			if result.Error != nil {
				return result.Error
			}
			if count > 0 {
				return errors.New("a tag is used by article")
			}
			return nil
		},
	}
}
