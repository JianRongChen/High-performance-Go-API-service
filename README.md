# 高性能 Go API 服务

一个基于 Golang 构建的高性能服务端 API 项目，单机 QPS > 20000。

## 功能特性

### 核心功能
- ✅ **用户模块**：注册、登录、用户信息查询
- ✅ **管理员模块**：管理员登录、角色管理、权限控制
  - 支持三种角色：超级管理员、普通管理员、操作员
  - 基于角色的权限控制（RBAC）
- ✅ **JWT 认证机制**：用户和管理员分离的 Token 认证
- ✅ **Redis 缓存支持**：用户和管理员信息缓存，提升查询性能

### 中间件
- ✅ **限流中间件**：基于 Redis 的滑动窗口限流算法，防止接口滥用
- ✅ **认证中间件**：用户认证和管理员认证分离
- ✅ **CORS 跨域支持**：支持跨域请求
- ✅ **Panic 恢复**：自动捕获和记录 panic 错误
- ✅ **请求日志**：自动记录所有 HTTP 请求

### 日志系统
- ✅ **按日期分割**：每天自动创建新的日志文件
- ✅ **按级别分离**：Info 和 Error 日志分别存储
  - `logs/YYYY-MM-DD.info.log` - 信息日志
  - `logs/YYYY-MM-DD.error.log` - 错误日志
- ✅ **自动轮转**：跨日期自动切换日志文件
- ✅ **控制台输出**：同时输出到文件和控制台

### 开发工具
- ✅ **Swagger API 文档**：自动生成交互式 API 文档，支持在线测试
- ✅ **热重载开发**：使用 Air 实现代码变更自动重启
- ✅ **Docker Compose**：一键启动 MySQL 和 Redis 开发环境

### 性能优化
- ✅ **数据库连接池**：MySQL 连接池优化配置
- ✅ **Redis 连接池**：Redis 连接池优化配置
- ✅ **优雅关闭**：支持优雅关闭服务
- ✅ **Gin Release 模式**：生产环境性能优化

## 技术栈

- **Web框架**: Gin
- **数据库**: MySQL 8.x + GORM
- **缓存**: Redis 7.x
- **认证**: JWT
- **密码加密**: bcrypt

## 项目结构

```
bgame/
├── cmd/
│   └── server/
│       └── main.go                 # 程序入口
├── internal/
│   ├── config/                     # 配置加载
│   ├── handler/                    # Gin 控制器
│   │   ├── user/                   # 用户接口
│   │   │   ├── login.go
│   │   │   └── register.go
│   │   └── admin/                  # 管理员接口
│   │       ├── login.go
│   │       └── role.go
│   ├── middleware/                 # 中间件
│   │   ├── auth.go                 # JWT 认证
│   │   ├── cors.go                 # 跨域支持
│   │   ├── logger.go               # 请求日志
│   │   ├── ratelimit.go            # 限流
│   │   └── recovery.go             # Panic 恢复
│   ├── model/                      # GORM 模型
│   │   ├── user.go
│   │   └── admin.go
│   ├── service/                    # 业务逻辑层
│   │   ├── user_service.go
│   │   └── admin_service.go
│   ├── dao/                        # 数据访问层（含缓存）
│   │   ├── user_dao.go
│   │   └── admin_dao.go
│   ├── util/                       # 工具函数
│   │   ├── jwt.go                  # JWT 工具
│   │   ├── logger.go               # 日志工具
│   │   ├── password.go             # 密码加密
│   │   └── response.go             # 响应格式化
│   └── router/                     # 路由配置
│       ├── router.go
│       ├── user_router.go
│       └── admin_router.go
├── pkg/
│   ├── redis/                      # Redis 客户端封装
│   └── mysql/                      # GORM 封装
├── scripts/                        # 工具脚本
│   ├── check_swagger.sh           # Swagger 配置检查
│   ├── check_tools.sh             # 工具安装检查
│   ├── create_admin.go            # 生成管理员密码
│   ├── init_db.sql                # 数据库初始化
│   ├── run_air.sh                 # Air 启动脚本
│   ├── run_air.bat                # Air 启动脚本（Windows）
│   ├── setup_dev.sh               # 开发环境设置
│   ├── setup_dev.bat              # 开发环境设置（Windows）
│   └── test_api.sh                # API 测试脚本
├── docs/                           # 文档目录
│   ├── docs.go                    # Swagger 生成代码
│   ├── swagger.json               # Swagger JSON
│   ├── swagger.yaml               # Swagger YAML
│   └── LOGGING.md                 # 日志系统文档
├── deployments/
│   └── docker-compose.yml          # Docker 部署配置
├── logs/                           # 日志目录（自动创建）
│   ├── YYYY-MM-DD.info.log        # Info 日志
│   └── YYYY-MM-DD.error.log       # Error 日志
├── .air.toml                       # Air 热重载配置
├── .gitignore
├── go.mod
├── go.sum
├── config.yaml                     # 配置文件
└── README.md
```

## 快速开始

### 1. 环境要求

- Go 1.21+
- MySQL 8.x
- Redis 7.x

### 2. 使用 Docker Compose 启动数据库（推荐）

```bash
# 启动 MySQL 和 Redis
cd deployments
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

### 3. 配置

编辑 `config.yaml` 文件，配置数据库连接信息：

```yaml
mysql:
  host: "localhost"
  port: 3306
  user: "root"
  password: "root123"
  database: "bgame"

redis:
  host: "localhost"
  port: 6379
```

### 4. 安装依赖和开发工具

**所有平台：**
```bash
# 安装项目依赖
go mod download
go mod tidy

# 安装开发工具（Air 热重载 + Swag Swagger 文档生成）
go install github.com/air-verse/air@latest
go install github.com/swaggo/swag/cmd/swag@latest

```

**或使用批处理脚本（Windows）：**
```bash
make.bat deps          # 安装项目依赖
make.bat install-tools # 安装开发工具
```

**检查工具安装状态：**
```bash
bash scripts/check_tools.sh  # 检查工具是否已安装并配置正确
```

**或使用开发环境设置脚本（一键安装）：**
```bash
bash scripts/setup_dev.sh    # Linux/Mac
scripts\setup_dev.bat        # Windows
```

### 5. 生成 Swagger 文档

**所有平台：**
```bash
# 方法1: 直接使用 swag 命令
swag init -g cmd/server/main.go -o docs

# 方法2: 使用脚本（推荐）
bash scripts/swagger.sh      # Linux/Mac/Git Bash
scripts\swagger.bat          # Windows CMD
```

**或使用开发环境设置脚本（会自动生成）：**
```bash
bash scripts/setup_dev.sh
```

**注意**：必须使用 `-g cmd/server/main.go` 参数指定入口文件路径，不能直接运行 `swag init`。

文档生成后，启动服务并访问 `http://localhost:8080/swagger/index.html` 查看 API 文档

### 6. 运行服务

**开发模式（推荐，支持热重载）：**

所有平台:
```bash
# 方法1: 直接运行（如果 PATH 已配置）
air

# 方法2: 使用完整路径（推荐，避免 PATH 问题）
$(go env GOPATH)/bin/air

# 方法3: 使用启动脚本
bash scripts/run_air.sh  # Linux/Mac
scripts\run_air.bat      # Windows

# 生成文档并启动热重载
swag init -g cmd/server/main.go -o docs && air
```

**使用启动脚本（推荐）：**
```bash
bash scripts/run_air.sh      # Linux/Mac/Git Bash
scripts\run_air.bat          # Windows CMD
```

**生产模式：**

所有平台:
```bash
go run cmd/server/main.go              # 直接运行
go build -o bin/server cmd/server/main.go  # 编译
./bin/server                           # 运行编译后的程序（Linux/Mac）
bin\server.exe                         # 运行编译后的程序（Windows）
```

**Windows 用户：**
```bash
go run cmd/server/main.go
go build -o bin/server.exe cmd/server/main.go
bin\server.exe
```

### 7. 访问 Swagger 文档

启动服务后，访问以下地址查看交互式 API 文档：

```
http://localhost:8080/swagger/index.html
```

### 8. 测试 API

```bash
# 使用测试脚本
bash scripts/test_api.sh

# 或手动测试
curl http://localhost:8080/health
```

## API 接口

### 用户接口

#### 用户注册
```http
POST /api/user/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com",
  "nickname": "测试用户"
}
```

#### 用户登录
```http
POST /api/user/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "123456"
}
```

响应：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_info": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "nickname": "测试用户"
    }
  }
}
```

#### 获取用户信息（需要认证）
```http
GET /api/user/info
Authorization: Bearer {token}
```

### 管理员接口

#### 管理员登录
```http
POST /api/admin/login
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

#### 获取角色列表
```http
GET /api/admin/roles
```

#### 获取管理员信息（需要认证）
```http
GET /api/admin/info
Authorization: Bearer {token}
```

#### 创建管理员（需要超级管理员权限）
```http
POST /api/admin/create
Authorization: Bearer {token}
Content-Type: application/json

{
  "username": "newadmin",
  "password": "admin123",
  "role": 2
}
```

角色说明：
- 1: 超级管理员
- 2: 普通管理员
- 3: 操作员

## 性能优化

1. **数据库连接池**: 配置了合理的连接池大小，减少连接开销
2. **Redis 缓存**: 用户和管理员信息缓存，减少数据库查询
3. **限流中间件**: 基于 Redis 的滑动窗口限流，防止接口被滥用
4. **Gin 性能模式**: 使用 Release 模式，关闭调试信息
5. **连接复用**: HTTP Keep-Alive 和数据库连接复用

## 配置说明

主要配置项在 `config.yaml`：

### Server 配置
```yaml
server:
  host: "0.0.0.0"        # 监听地址
  port: 8080             # 监听端口
  mode: "release"        # 运行模式：debug, release, test
  read_timeout: 30       # 读取超时（秒）
  write_timeout: 30      # 写入超时（秒）
```

### MySQL 配置
```yaml
mysql:
  host: "localhost"
  port: 3306
  user: "bgame"
  password: "123456"
  database: "bgame"
  charset: "utf8mb4"
  max_open_conns: 100    # 最大打开连接数
  max_idle_conns: 10     # 最大空闲连接数
  conn_max_lifetime: 3600 # 连接最大生存时间（秒）
```

### Redis 配置
```yaml
redis:
  host: "localhost"
  port: 6379
  password: ""           # Redis 密码（如有）
  database: 0
  pool_size: 100         # 连接池大小
  min_idle_conns: 10     # 最小空闲连接数
```

### JWT 配置
```yaml
jwt:
  secret: "your-secret-key-change-in-production"  # ⚠️ 生产环境必须修改
  user_expire: 7200      # 用户 Token 过期时间（秒，2小时）
  admin_expire: 3600     # 管理员 Token 过期时间（秒，1小时）
```

### 限流配置
```yaml
rate_limit:
  enabled: true          # 是否启用限流
  rps: 10000            # 每秒请求数限制
  burst: 20000          # 突发请求数限制
```

### 日志配置
```yaml
log:
  level: "info"          # 日志级别：debug, info, warn, error
  dir: "logs"            # 日志目录
```

**日志文件格式**：
- `logs/YYYY-MM-DD.info.log` - Info 级别日志（包含 Info、Warn、Debug）
- `logs/YYYY-MM-DD.error.log` - Error 级别日志

日志文件按日期自动创建和切换，无需手动管理。

## 开发建议

1. **首次运行前**：
   - 确保 MySQL 和 Redis 已启动
   - 检查 `config.yaml` 中的数据库连接配置
   - 运行 `bash scripts/setup_dev.sh` 一键设置开发环境

2. **数据库**：
   - 数据库表会自动创建（通过 GORM AutoMigrate）
   - 如需手动初始化，可参考 `scripts/init_db.sql`

3. **生产环境**：
   - ⚠️ **必须修改** JWT Secret（`config.yaml` 中的 `jwt.secret`）
   - 根据实际负载调整连接池大小和限流参数
   - 设置合适的日志级别（建议 `info` 或 `warn`）

4. **管理员账户**：
   - 首次创建管理员可使用 `scripts/create_admin.go` 生成密码哈希
   - 或通过 Swagger 文档的创建管理员接口（需要超级管理员权限）

5. **日志管理**：
   - 日志文件会自动按日期创建，建议定期清理旧日志
   - 查看日志：`tail -f logs/$(date +%Y-%m-%d).info.log`
   - 查看错误：`tail -f logs/$(date +%Y-%m-%d).error.log`

## 性能测试

项目已针对高性能进行优化：

- **连接池优化**: MySQL 和 Redis 都配置了合理的连接池
- **缓存策略**: 用户和管理员信息使用 Redis 缓存
- **限流保护**: 基于 Redis 的滑动窗口限流算法
- **Gin 优化**: 使用 Release 模式，关闭调试信息

预期性能指标：
- 单机 QPS > 20000
- 响应时间 < 10ms（缓存命中）
- 支持高并发请求

## 常用命令

**直接使用 Go 命令（所有平台）：**
```bash
# 安装开发工具
go install github.com/air-verse/air@latest
go install github.com/swaggo/swag/cmd/swag@latest

# 项目依赖管理
go mod download
go mod tidy

# 生成 Swagger 文档
swag init -g cmd/server/main.go -o docs

# 开发模式（热重载）
air

# 构建项目
go build -o bin/server cmd/server/main.go

# 运行项目
go run cmd/server/main.go

# Docker 服务管理
cd deployments && docker-compose up -d
cd deployments && docker-compose down
cd deployments && docker-compose logs -f

# 代码格式化
go fmt ./...

# 清理构建文件
rm -rf bin/ tmp/ docs/  # Linux/Mac
rmdir /s /q bin tmp docs  # Windows
```

**使用脚本工具：**
```bash
# 开发环境设置（一键安装所有工具）
bash scripts/setup_dev.sh

# 检查工具安装状态
bash scripts/check_tools.sh

# 检查 Swagger 配置
bash scripts/check_swagger.sh

# 启动热重载
bash scripts/run_air.sh

# API 测试
bash scripts/test_api.sh
```

## Swagger API 文档

项目集成了 Swagger 自动文档生成功能：

### 生成文档
```bash
# 直接使用命令
swag init -g cmd/server/main.go -o docs

# 或使用脚本
bash scripts/swagger.sh      # Linux/Mac/Git Bash
scripts\swagger.bat          # Windows CMD
```

**重要**：必须指定 `-g cmd/server/main.go` 参数，不能直接运行 `swag init`。

### 访问文档
启动服务后访问：`http://localhost:8080/swagger/index.html`

### 文档特性
- ✅ 交互式 API 测试界面
- ✅ 自动生成请求/响应示例
- ✅ 支持 JWT 认证测试（点击右上角 "Authorize" 按钮）
- ✅ 完整的参数说明和错误码
- ✅ 按模块分类（用户接口、管理员接口、系统接口）

### 检查配置
```bash
bash scripts/check_swagger.sh
```

## 热重载开发

项目使用 [Air](https://github.com/air-verse/air) 实现热重载功能：

### 安装 Air
```bash
go install github.com/air-verse/air@latest
```

### 启动热重载
```bash
# 方法1: 直接运行（如果 PATH 已配置）
air

# 方法2: 使用完整路径
$(go env GOPATH)/bin/air

# 方法3: 使用启动脚本（推荐）
bash scripts/run_air.sh      # Linux/Mac/Git Bash
scripts\run_air.bat          # Windows CMD
```

### 工作原理
- 监控 `.go`、`.yaml`、`.yml` 文件变更
- 自动重新编译和重启服务
- 排除 `tmp/`、`vendor/`、`docs/`、`scripts/` 等目录

### Windows 用户注意
- `.air.toml` 配置文件已设置为使用 `main.exe`（Windows 需要 .exe 扩展名）
- `scripts/run_air.sh` 脚本会自动检测操作系统并修复配置
- 如果遇到执行错误，检查 `.air.toml` 中的 `bin` 和 `cmd` 配置

配置文件：`.air.toml`（可根据需要自定义）

## 日志系统

项目实现了完整的日志系统，支持按日期和级别自动分割日志文件。

### 日志文件格式
- `logs/YYYY-MM-DD.info.log` - Info 级别日志（包含 Info、Warn、Debug）
- `logs/YYYY-MM-DD.error.log` - Error 级别日志

### 使用示例
```go
import "bgame/internal/util"

util.Info("这是一条信息日志")
util.LogError("这是一条错误日志: %v", err)
util.Warn("这是一条警告日志")
util.Debug("这是一条调试日志")  // 仅在 level=debug 时记录
```

### 自动日志记录
- HTTP 请求自动记录（状态码 < 400 → info.log，>= 400 → error.log）
- 系统事件自动记录（启动、关闭、数据库连接等）
- Panic 错误自动记录到 error.log
- 每天自动创建新的日志文件

### 查看日志
```bash
# 查看今天的 Info 日志
tail -f logs/$(date +%Y-%m-%d).info.log

# 查看今天的 Error 日志
tail -f logs/$(date +%Y-%m-%d).error.log

# Windows PowerShell
Get-Content logs/2024-01-01.info.log -Wait
```

详细文档请参考：[docs/LOGGING.md](docs/LOGGING.md)

## 项目特性总结

### 架构设计
- **分层架构**：Handler → Service → DAO → Model
- **依赖注入**：清晰的模块依赖关系
- **接口分离**：用户和管理员模块完全分离

### 安全特性
- JWT Token 认证
- 密码 bcrypt 加密
- 基于角色的权限控制（RBAC）
- 限流保护防止接口滥用

### 性能特性
- Redis 缓存减少数据库查询
- 数据库连接池优化
- HTTP Keep-Alive 连接复用
- Gin Release 模式优化

### 开发体验
- Swagger 自动文档生成
- Air 热重载开发
- 完整的日志系统
- Docker Compose 一键启动环境

## 故障排查

### 常见问题

1. **服务启动失败**
   - 检查 MySQL 和 Redis 是否已启动
   - 检查 `config.yaml` 中的连接配置
   - 查看日志文件：`logs/YYYY-MM-DD.error.log`

2. **Swagger 文档无法访问**
   - 确保已运行 `swag init -g cmd/server/main.go -o docs`
   - 检查 `docs/` 目录是否存在
   - 运行 `bash scripts/check_swagger.sh` 检查配置

3. **Air 热重载不工作**
   - 检查 Air 是否已安装：`which air` 或 `air -v`
   - Windows 用户检查 `.air.toml` 中的 `.exe` 配置
   - 查看 `tmp/build-errors.log` 了解编译错误

4. **数据库连接失败**
   - 检查 MySQL 服务是否运行
   - 验证用户名、密码、数据库名是否正确
   - 检查防火墙和网络连接

5. **Redis 连接失败**
   - 检查 Redis 服务是否运行
   - 验证端口和密码配置
   - 检查 Redis 连接：`redis-cli ping`

## License

MIT

