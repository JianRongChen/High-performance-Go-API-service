package user

import (
	"bgame/internal/service"
	"bgame/internal/util"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
// @Summary      用户注册
// @Description  用户注册接口
// @Tags         用户接口
// @Accept       json
// @Produce      json
// @Param        request body service.RegisterRequest true "注册请求"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Router       /api/user/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, "参数错误: "+err.Error())
		return
	}

	if err := h.userService.Register(&req); err != nil {
		util.Error(c, err.Error())
		return
	}

	util.SuccessWithMessage(c, "注册成功", nil)
}

// GetUserInfo 获取用户信息
// @Summary      获取用户信息
// @Description  获取当前登录用户的信息
// @Tags         用户接口
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  util.Response{data=model.User}
// @Failure      401  {object}  util.Response
// @Router       /api/user/info [get]
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		util.Unauthorized(c, "未获取到用户信息")
		return
	}

	user, err := h.userService.GetUserInfo(userID.(uint))
	if err != nil {
		util.Error(c, err.Error())
		return
	}

	util.Success(c, user)
}

