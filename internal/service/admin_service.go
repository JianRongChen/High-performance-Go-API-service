package service

import (
	"errors"

	"bgame/internal/dao"
	"bgame/internal/model"
	"bgame/internal/util"
)

type AdminService struct {
	adminDAO *dao.AdminDAO
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminDAO: dao.NewAdminDAO(),
	}
}

type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminLoginResponse struct {
	Token     string       `json:"token"`
	AdminInfo *model.Admin `json:"admin_info"`
}

type CreateAdminRequest struct {
	Username string           `json:"username" binding:"required,min=3,max=50"`
	Password string           `json:"password" binding:"required,min=6,max=50"`
	Role     model.AdminRole  `json:"role" binding:"required"`
}

// Login 管理员登录
func (s *AdminService) Login(req *AdminLoginRequest) (*AdminLoginResponse, error) {
	// 获取管理员
	admin, err := s.adminDAO.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查管理员状态
	if admin.Status != 1 {
		return nil, errors.New("管理员已被禁用")
	}

	// 验证密码
	if !util.CheckPassword(req.Password, admin.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := util.GenerateAdminToken(admin.ID, admin.Username, int(admin.Role))
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	// 清除敏感信息
	admin.Password = ""

	return &AdminLoginResponse{
		Token:     token,
		AdminInfo: admin,
	}, nil
}

// CreateAdmin 创建管理员（需要超级管理员权限）
func (s *AdminService) CreateAdmin(req *CreateAdminRequest) error {
	// 检查用户名是否已存在
	_, err := s.adminDAO.GetByUsername(req.Username)
	if err == nil {
		return errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 创建管理员
	admin := &model.Admin{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
		Status:   1,
	}

	if err := s.adminDAO.Create(admin); err != nil {
		return errors.New("创建管理员失败")
	}

	return nil
}

// GetAdminInfo 获取管理员信息
func (s *AdminService) GetAdminInfo(adminID uint) (*model.Admin, error) {
	admin, err := s.adminDAO.GetByID(adminID)
	if err != nil {
		return nil, errors.New("管理员不存在")
	}

	if admin.Status != 1 {
		return nil, errors.New("管理员已被禁用")
	}

	// 清除敏感信息
	admin.Password = ""
	return admin, nil
}

// CheckRole 检查角色权限
func (s *AdminService) CheckRole(adminRole model.AdminRole, requiredRole model.AdminRole) bool {
	// 超级管理员拥有所有权限
	if adminRole == model.RoleSuperAdmin {
		return true
	}

	// 角色权限等级：超级管理员 > 管理员 > 操作员
	roleLevel := map[model.AdminRole]int{
		model.RoleSuperAdmin: 1,
		model.RoleAdmin:      2,
		model.RoleOperator:   3,
	}

	return roleLevel[adminRole] <= roleLevel[requiredRole]
}

