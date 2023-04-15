package rabbitadmin

import (
	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit-admin/internal/models"
)

func (m *ServerManager) categoryObject() gormpher.WebObject {
	return gormpher.WebObject{
		Model:       &models.Category{},
		GetDB:       m.getDB,
		Editables:   []string{"Name"},
		Searchables: []string{"Name"},
		Filterables: []string{"Name"},
		Orderables:  []string{"CreatedAt"},
	}
}
