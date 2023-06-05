package rabbitadmin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	"github.com/szluyu99/rabbit-admin/internal/models"
)

// default data
func (m *ServerManager) handleDefaultPermissions(c *gin.Context) {
	user := rabbit.CurrentUser(c)
	if !user.IsSuperUser {
		handleError(c, http.StatusForbidden, "only super user can init permissions")
		return
	}

	if err := models.InitDefaultPermissions(m.db); err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, true)
}

func (m *ServerManager) handleDefaultRoles(c *gin.Context) {
	user := rabbit.CurrentUser(c)
	if !user.IsSuperUser {
		handleError(c, http.StatusForbidden, "only super user can init roles")
		return
	}

	roles := []*rabbit.Role{
		{Name: "admin", Label: "Admin"},
		{Name: "user", Label: "User"},
		{Name: "test", Label: "Test"},
	}

	if err := models.InitDefaultRoles(m.db, roles); err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, true)
}

// role
func (m *ServerManager) handleQueryRole(c *gin.Context) {
	var form PageForm
	if err := c.BindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	page, limit := CheckPageLimit(form.Page, form.Limit)
	list, count, err := models.GetRoleList(m.db, page, limit, form.Keyword)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, PageResult[rabbit.Role]{
		Page:       page,
		Limit:      limit,
		Keyword:    form.Keyword,
		Items:      list,
		TotalCount: count,
	})
}

// permission
func (m *ServerManager) handleCreateWebObjectPermissions(c *gin.Context) {
	name := c.Param("name")

	if _, err := models.CreateWebObjectPermissions(m.db, name); err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, true)
}

func (m *ServerManager) handleQueryPermission(c *gin.Context) {
	var form PageForm
	if err := c.BindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	list, err := models.GetPermissionTreeList(m.db, form.Keyword)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": list,
	})
}

// user
func (m *ServerManager) handleQueryUser(c *gin.Context) {
	var form PageForm
	if err := c.BindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	page, limit := CheckPageLimit(form.Page, form.Limit)
	list, count, err := models.GetUserList(m.db, page, limit, form.Keyword)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, PageResult[rabbit.User]{
		Page:       page,
		Limit:      limit,
		Keyword:    form.Keyword,
		Items:      list,
		TotalCount: count,
	})
}

func (m *ServerManager) handleUpdateUserRole(c *gin.Context) {
	var form UserForm
	if err := c.BindJSON(&form); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	user, err := rabbit.UpdateRolesForUser(m.db, form.ID, form.RoleIds)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
