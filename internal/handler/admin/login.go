package admin

import (
	"bgame/internal/service"
	"bgame/internal/util"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *service.AdminService
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{
		adminService: service.NewAdminService(),
	}
}

// Login 管理员登录
// @Summary      管理员登录
// @Description  管理员登录接口，返回JWT token和管理员信息
// @Tags         管理员接口
// @Accept       json
// @Produce      json
// @Param        request body service.AdminLoginRequest true "登录请求"
// @Success      200  {object}  util.Response{data=service.AdminLoginResponse}
// @Failure      400  {object}  util.Response
// @Router       /api/admin/login [post]
func (h *AdminHandler) Login(c *gin.Context) {
	var req service.AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, "参数错误: "+err.Error())
		return
	}

	resp, err := h.adminService.Login(&req)
	if err != nil {
		util.Error(c, err.Error())
		return
	}

	util.Success(c, resp)
}

