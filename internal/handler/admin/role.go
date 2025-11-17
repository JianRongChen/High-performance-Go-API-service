package admin

import (
	"bgame/internal/model"
	"bgame/internal/service"
	"bgame/internal/util"
	"github.com/gin-gonic/gin"
)

// CreateAdmin 创建管理员（需要超级管理员权限）
// @Summary      创建管理员
// @Description  创建新管理员账户，需要超级管理员权限
// @Tags         管理员接口
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body service.CreateAdminRequest true "创建管理员请求"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Failure      403  {object}  util.Response
// @Router       /api/admin/create [post]
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var req service.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, "参数错误: "+err.Error())
		return
	}

	if err := h.adminService.CreateAdmin(&req); err != nil {
		util.Error(c, err.Error())
		return
	}

	util.SuccessWithMessage(c, "创建管理员成功", nil)
}

// GetAdminInfo 获取管理员信息
// @Summary      获取管理员信息
// @Description  获取当前登录管理员的信息
// @Tags         管理员接口
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  util.Response{data=model.Admin}
// @Failure      401  {object}  util.Response
// @Router       /api/admin/info [get]
func (h *AdminHandler) GetAdminInfo(c *gin.Context) {
	adminID, exists := c.Get("admin_id")
	if !exists {
		util.Unauthorized(c, "未获取到管理员信息")
		return
	}

	admin, err := h.adminService.GetAdminInfo(adminID.(uint))
	if err != nil {
		util.Error(c, err.Error())
		return
	}

	util.Success(c, admin)
}

// GetRoles 获取所有角色列表
// @Summary      获取角色列表
// @Description  获取所有可用的管理员角色列表
// @Tags         管理员接口
// @Accept       json
// @Produce      json
// @Success      200  {object}  util.Response
// @Router       /api/admin/roles [get]
func (h *AdminHandler) GetRoles(c *gin.Context) {
	roles := []map[string]interface{}{
		{"value": int(model.RoleSuperAdmin), "label": model.RoleSuperAdmin.String()},
		{"value": int(model.RoleAdmin), "label": model.RoleAdmin.String()},
		{"value": int(model.RoleOperator), "label": model.RoleOperator.String()},
	}

	util.Success(c, roles)
}

