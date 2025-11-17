package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"bgame/internal/model"
	"bgame/pkg/mysql"
	"bgame/pkg/redis"
)

const (
	adminCachePrefix = "admin:" // 管理员缓存前缀
	adminCacheTTL    = 1800 * time.Second // 管理员缓存时间
)

type AdminDAO struct{}

func NewAdminDAO() *AdminDAO {
	return &AdminDAO{}
}

// Create 创建管理员
func (d *AdminDAO) Create(admin *model.Admin) error {
	return mysql.DB.Create(admin).Error
}

// GetByID 根据ID获取管理员（带缓存）
func (d *AdminDAO) GetByID(id uint) (*model.Admin, error) {
	// 先查缓存
	cacheKey := fmt.Sprintf("%s%d", adminCachePrefix, id)
	cached, err := redis.Client.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var admin model.Admin
		if json.Unmarshal([]byte(cached), &admin) == nil {
			return &admin, nil
		}
	}

	// 查数据库
	var admin model.Admin
	if err := mysql.DB.Where("id = ? AND status = 1", id).First(&admin).Error; err != nil {
		return nil, err
	}

	// 写入缓存
	if adminData, err := json.Marshal(admin); err == nil {
		redis.Client.Set(context.Background(), cacheKey, adminData, adminCacheTTL)
	}

	return &admin, nil
}

// GetByUsername 根据用户名获取管理员
func (d *AdminDAO) GetByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	if err := mysql.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// Update 更新管理员
func (d *AdminDAO) Update(admin *model.Admin) error {
	err := mysql.DB.Save(admin).Error
	if err == nil {
		// 清除缓存
		cacheKey := fmt.Sprintf("%s%d", adminCachePrefix, admin.ID)
		redis.Client.Del(context.Background(), cacheKey)
	}
	return err
}

// DeleteCache 删除管理员缓存
func (d *AdminDAO) DeleteCache(adminID uint) {
	cacheKey := fmt.Sprintf("%s%d", adminCachePrefix, adminID)
	redis.Client.Del(context.Background(), cacheKey)
}

