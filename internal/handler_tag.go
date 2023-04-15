package rabbitadmin

import (
	"github.com/restsend/gormpher"
	"github.com/szluyu99/rabbit-admin/internal/models"
)

func (m *ServerManager) tagObject() gormpher.WebObject {
	return gormpher.WebObject{
		Model:       &models.Tag{},
		GetDB:       m.getDB,
		Editables:   []string{"Name"},
		Searchables: []string{"Name"},
		Filterables: []string{"Name"},
		Orderables:  []string{"CreatedAt"},
	}
}
