-- 初始化数据库脚本
-- 如果使用 docker-compose，数据库会自动创建
-- 如果需要手动创建，可以执行此脚本

CREATE DATABASE IF NOT EXISTS `bgame` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `bgame`;

-- 注意：表结构会通过 GORM AutoMigrate 自动创建
-- 如果需要手动创建，可以参考以下 SQL（但建议使用 AutoMigrate）

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `nickname` varchar(50) DEFAULT NULL,
  `status` tinyint DEFAULT '1',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_username` (`username`),
  UNIQUE KEY `idx_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 管理员表
CREATE TABLE IF NOT EXISTS `admins` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` tinyint NOT NULL DEFAULT '3',
  `status` tinyint DEFAULT '1',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_admins_username` (`username`),
  KEY `idx_admins_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入一个默认超级管理员（密码：admin123）
-- 注意：实际密码需要经过 bcrypt 加密
-- 可以使用以下 Go 代码生成加密后的密码：
-- password, _ := util.HashPassword("admin123")
INSERT INTO `admins` (`username`, `password`, `role`, `status`, `created_at`, `updated_at`) 
VALUES ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 1, 1, NOW(), NOW())
ON DUPLICATE KEY UPDATE `username`=`username`;

