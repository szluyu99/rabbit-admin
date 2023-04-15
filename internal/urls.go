package rabbitadmin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restsend/gormpher"
	"github.com/restsend/rabbit"
	"gorm.io/gorm"
)

const (
	AuthGroupField = "_auth_group_"
)

func (m *ServerManager) RegisterHandler(r *gin.Engine) error {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	ar := r.Group("/api") // .Use(m.HandleAuth)

	ar.POST("/article/list", m.handlerArticleList)

	objects := []gormpher.WebObject{
		m.tagObject(),
		m.categoryObject(),
		m.articleObject(),
		{
			Model:     &rabbit.Permission{},
			GetDB:     m.getDB,
			Editables: []string{"Name", "Code"},
		},
		{
			Model:     &rabbit.Group{},
			GetDB:     m.getDB,
			Editables: []string{"Name"},
		},
	}

	gormpher.RegisterObjects(ar, objects)
	// gormpher.RegisterObjectsWithAdmin(r.Group("/admin"), objects)

	return nil
}

func (m *ServerManager) HandleAuth(c *gin.Context) {
	group := m.GetAuthGroup(c)
	if group == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "forbidden"})
		return
	}
	c.Next()
}

func (m *ServerManager) getDB(ctx *gin.Context, isCreate bool) *gorm.DB {
	if isCreate {
		return m.db
	}

	val, ok := ctx.Get(AuthGroupField)
	if !ok {
		return m.db
	}
	group := val.(*rabbit.Group)
	return m.db.Where("group_id", group.ID)
}

func (m *ServerManager) GetAuthGroup(c *gin.Context) *rabbit.Group {
	val, ok := c.Get(AuthGroupField)
	if ok {
		return val.(*rabbit.Group)
	}

	group := rabbit.CurrentGroup(c)
	if group != nil {
		c.Set(AuthGroupField, group)
		return group
	}

	user := rabbit.CurrentUser(c)
	if user == nil {
		return nil
	}

	group, err := rabbit.GetFirstGroupByUser(m.db, user.ID)
	if err != nil {
		log.Println("get user first group fail: ", user.ID, err)
		return nil
	}

	c.Set(AuthGroupField, group)
	return group
}
