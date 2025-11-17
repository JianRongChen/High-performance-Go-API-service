package service

import (
	"errors"

	"bgame/internal/dao"
	"bgame/internal/model"
	"bgame/internal/util"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		userDAO: dao.NewUserDAO(),
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Nickname string `json:"nickname" binding:"max=50"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token    string       `json:"token"`
	UserInfo *model.User  `json:"user_info"`
}

// Register 用户注册
func (s *UserService) Register(req *RegisterRequest) error {
	// 检查用户名是否已存在
	_, err := s.userDAO.GetByUsername(req.Username)
	if err == nil {
		return errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	if req.Email != "" {
		_, err := s.userDAO.GetByEmail(req.Email)
		if err == nil {
			return errors.New("邮箱已被注册")
		}
	}

	// 加密密码
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Nickname: req.Nickname,
		Status:   1,
	}

	if err := s.userDAO.Create(user); err != nil {
		return errors.New("创建用户失败")
	}

	return nil
}

// Login 用户登录
func (s *UserService) Login(req *LoginRequest) (*LoginResponse, error) {
	// 获取用户
	user, err := s.userDAO.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, errors.New("用户已被禁用")
	}

	// 验证密码
	if !util.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := util.GenerateUserToken(user.ID, user.Username)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	// 清除敏感信息
	user.Password = ""

	return &LoginResponse{
		Token:    token,
		UserInfo: user,
	}, nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID uint) (*model.User, error) {
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if user.Status != 1 {
		return nil, errors.New("用户已被禁用")
	}

	// 清除敏感信息
	user.Password = ""
	return user, nil
}

