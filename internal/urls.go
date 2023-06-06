package rabbitadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"gorm.io/gorm"
)

const (
	AuthGroupField = "_auth_group_"
)

func (m *ServerManager) RegisterHandlers(r *gin.Engine) error {
	ar := r.Group("/api").Use(rabbit.WithAuthentication(), rabbit.WithAuthorization("/api"))
	// .Use(m.HandleAuth) // TODO: for group
	rabbit.RegisterAuthorizationHandlers(m.db, ar)

	// auth
	{
		ar.POST("/role", m.handleQueryRole)
		ar.POST("/permission", m.handleQueryPermission)

		ar.GET("/role/default", m.handleDefaultRoles)
		ar.GET("/permission/default", m.handleDefaultPermissions)
		ar.GET("/permission/object/:name", m.handleCreateWebObjectPermissions)
	}
	{
		ar.GET("/article/:key", m.handleGetArticle)
		ar.POST("/article", m.handleQueryArticle)
		ar.POST("/article/save_or_update", m.handleSaveOrUpdateArticle)
	}
	{
		ar.POST("/user", m.handleQueryUser)
		ar.PATCH("/user/role", m.handleUpdateUserRole)
	}

	objects := []rabbit.WebObject{
		m.tagObject(),
		m.categoryObject(),
		m.articleObject(),
		{
			Name:         "config",
			Model:        rabbit.Config{},
			Editables:    []string{"Key", "Value", "Desc"},
			Filterables:  []string{"Key", "Value"},
			Searchables:  []string{"Key", "Desc"},
			GetDB:        func(c *gin.Context, isCreate bool) *gorm.DB { return m.db },
			AllowMethods: rabbit.QUERY | rabbit.EDIT | rabbit.DELETE,
		},
	}
	rabbit.RegisterObjects(ar, objects)

	return nil
}

// TODO: use group
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

func handleError(c *gin.Context, code int, err any) {
	switch val := err.(type) {
	case error:
		c.AbortWithStatusJSON(code, gin.H{"error": val.Error()})
	case string:
		c.AbortWithStatusJSON(code, gin.H{"error": val})
	}
}

// func (m *ServerManager) HandleAuth(c *gin.Context) {
// 	group := m.GetAuthGroup(c)
// 	if group == nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "forbidden"})
// 		return
// 	}
// 	c.Next()
// }

// func (m *ServerManager) GetAuthGroup(c *gin.Context) *rabbit.Group {
// 	val, ok := c.Get(AuthGroupField)
// 	if ok {
// 		return val.(*rabbit.Group)
// 	}

// 	group := rabbit.CurrentGroup(c)
// 	if group != nil {
// 		c.Set(AuthGroupField, group)
// 		return group
// 	}

// 	user := rabbit.CurrentUser(c)
// 	if user == nil {
// 		return nil
// 	}

// 	group, err := rabbit.GetFirstGroupByUser(m.db, user.ID)
// 	if err != nil {
// 		log.Println("get user first group fail: ", user.ID, err)
// 		return nil
// 	}

// 	c.Set(AuthGroupField, group)
// 	return group
// }
