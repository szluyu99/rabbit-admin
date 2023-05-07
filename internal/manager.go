package rabbitadmin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"github.com/szluyu99/rabbit-admin/internal/models"
	"gorm.io/gorm"
)

type ServerManager struct {
	db *gorm.DB
}

func InitServer(db *gorm.DB, r *gin.Engine) {
	m := NewManager(db)
	if err := m.Migrate(db); err != nil {
		panic(err)
	}
	if err := m.RegisterHandler(r); err != nil {
		panic(err)
	}
}

func NewManager(db *gorm.DB) *ServerManager {
	m := ServerManager{db: db}

	rabbit.Sig().Connect(rabbit.SigUserLogin, m.onUserLogin)
	rabbit.Sig().Connect(rabbit.SigUserLogin, m.onUserCreate)

	return &m
}

func (m *ServerManager) Migrate(db *gorm.DB) error {
	return models.MakeMigrate(db)
}

func (m *ServerManager) onUserLogin(sender any, params ...any) {
	log.Println("onUserLogin", sender, params)
}

func (m *ServerManager) onUserCreate(sender any, params ...any) {
	log.Println("onUserCreate", sender, params)
}
