package user

import (
	"bgame/internal/service"
	"bgame/internal/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// Login 用户登录
// @Summary      用户登录
// @Description  用户登录接口，返回JWT token和用户信息
// @Tags         用户接口
// @Accept       json
// @Produce      json
// @Param        request body service.LoginRequest true "登录请求"
// @Success      200  {object}  util.Response{data=service.LoginResponse}
// @Failure      400  {object}  util.Response
// @Router       /api/user/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, "参数错误: "+err.Error())
		return
	}

	resp, err := h.userService.Login(&req)
	if err != nil {
		util.Error(c, err.Error())
		return
	}

	util.Success(c, resp)
}

