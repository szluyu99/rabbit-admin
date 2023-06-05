package rabbitadmin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"github.com/szluyu99/rabbit-admin/internal/models"
	"gorm.io/gorm"
)

const KEY_CREATE_DEFAULT_ROLE = "CREATE_DEFAULT_ROLE"

type ServerManager struct {
	db *gorm.DB
}

func InitServer(db *gorm.DB, r *gin.Engine) {
	m := NewManager(db)

	if err := models.MakeMigrate(db); err != nil {
		log.Fatal("migrate error: ", err)
	}

	if err := m.Prepare(); err != nil {
		log.Fatal("prepare error: ", err)
	}

	if err := m.RegisterHandlers(r); err != nil {
		log.Fatal("register handlers error: ", err)
	}
}

func NewManager(db *gorm.DB) *ServerManager {
	m := ServerManager{db: db}

	rabbit.Sig().Connect(rabbit.SigUserLogin, m.onUserLogin)
	rabbit.Sig().Connect(rabbit.SigUserCreate, m.onUserCreate)

	return &m
}

// set config key value
func (m *ServerManager) Prepare() error {
	rabbit.CheckValue(m.db, KEY_CREATE_DEFAULT_ROLE, "user")
	return nil
}

func (m *ServerManager) onUserLogin(sender any, params ...any) {
	log.Println("onUserLogin", sender, params)
}

// set default role for user
func (m *ServerManager) onUserCreate(sender any, params ...any) {
	u := sender.(*rabbit.User)
	role, err := rabbit.GetRoleByName(m.db, rabbit.GetValue(m.db, KEY_CREATE_DEFAULT_ROLE))
	if err != nil {
		log.Println("onUserCreate error: ", err)
		return
	}
	err = m.db.Create(rabbit.UserRole{UserID: u.ID, RoleID: role.ID}).Error
	if err != nil {
		log.Println("onUserCreate error: ", err)
		return
	}
}
