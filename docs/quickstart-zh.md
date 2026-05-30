# 快速开始指南

> [English Version](./quickstart.md)

本指南将帮助您在不同环境中快速设置和运行 JoyLive Dashboard。

## 目录

- [前置要求](#前置要求)
- [本地开发](#本地开发)
- [Docker 部署](#docker-部署)
- [配置说明](#配置说明)
- [验证部署](#验证部署)
- [故障排查](#故障排查)

## 前置要求

### 本地开发环境

- **Go**: 1.19 或更高版本
- **Node.js**: 18.x 或更高版本
- **npm**: 8.x 或更高版本
- **MySQL**: 5.7 或更高版本（或使用 SQLite 进行测试）
- **Redis**: 6.0 或更高版本（可选，可使用内存缓存）
- **Git**: 用于克隆代码仓库

### Docker 部署环境

- **Docker**: 20.10 或更高版本
- **Docker Compose**: 1.29 或更高版本（可选，用于简化部署）

## 本地开发

### 步骤 1: 克隆代码仓库

```bash
git clone https://github.com/jd-opensource/joylive-dashboard.git
cd joylive-dashboard
```

### 步骤 2: 数据库设置

#### 方案 A: 使用 MySQL（生产环境推荐）

1. 创建数据库：

```sql
CREATE DATABASE joylive CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
```

2. 初始化数据库表结构：

```bash
mysql -u root -p joylive < scripts/init.sql
```

3. 更新 `configs/dev/server.toml` 中的数据库配置：

```toml
[Storage.DB]
Type = "mysql"
DSN = "用户名:密码@tcp(localhost:3306)/joylive?charset=utf8mb4&parseTime=True&loc=Local"
```

#### 方案 B: 使用 SQLite（快速测试）

开发模式下默认启用 SQLite。数据库文件将自动创建在 `data/joylivedashboard.db`。

要使用 SQLite，更新 `configs/dev/server.toml`：

```toml
[Storage.DB]
Type = "sqlite3"
DSN = "data/joylivedashboard.db"
```

### 步骤 3: 配置 Redis（可选）

如果要使用 Redis 作为缓存，更新 `configs/dev/server.toml`：

```toml
[Storage.Cache]
Type = "redis"  # 从 "memory" 改为 "redis"

[Storage.Cache.Redis]
Addr = "127.0.0.1:6379"
Password = ""
DB = 1
```

开发环境可以保持默认的内存缓存。

### 步骤 4: 安装依赖

#### 后端依赖

```bash
# Go 依赖由 go.mod 管理
go mod download
```

#### 前端依赖

```bash
cd frontend
npm install
cd ..
```

### 步骤 5: 构建和运行

#### 方案 A: 一键构建并运行

```bash
# 构建前端和后端，然后启动服务
make build-all
make serve
```

#### 方案 B: 开发模式（前后端分离）

**终端 1 - 后端：**

```bash
make start
```

**终端 2 - 前端：**

```bash
cd frontend
npm run dev
```

开发模式下：
- 后端运行在 `http://localhost:8040`
- 前端开发服务器运行在 `http://localhost:5173`（支持热重载）
- 前端会将 API 请求代理到后端

### 步骤 6: 访问应用

在浏览器中打开：

- **生产构建**: `http://localhost:8040`
- **开发模式**: `http://localhost:5173`

**默认登录凭证：**
- 用户名: `admin`
- 密码: `admin`

## Docker 部署

### 使用 Docker Compose 快速启动（推荐）

#### 步骤 1: 克隆代码仓库

```bash
git clone https://github.com/jd-opensource/joylive-dashboard.git
cd joylive-dashboard
```

#### 步骤 2: 配置（可选）

如需自定义设置，可编辑 `configs/prod/server.toml`。默认配置使用 SQLite，开箱即用。

#### 步骤 3: 使用 Docker Compose 启动

```bash
docker-compose up -d
```

这将会：
- 构建 Docker 镜像（如果尚未构建）
- 以后台模式启动容器
- 挂载 `./data` 目录用于持久化存储
- 暴露 8040 端口

#### 步骤 4: 检查状态

```bash
# 查看日志
docker-compose logs -f joylive

# 检查容器状态
docker-compose ps
```

#### 步骤 5: 访问应用

在浏览器中打开 `http://localhost:8040`

**默认登录凭证：**
- 用户名: `admin`
- 密码: `admin`

### 手动 Docker 部署

#### 步骤 1: 构建 Docker 镜像

```bash
# 使用 Makefile
make docker-build

# 或直接使用 Docker
docker build -t joylivedashboard:latest .
```

Docker 构建过程：
1. 构建前端（Vue 3 + Vite）
2. 构建后端（Go）
3. 创建包含前后端的单一优化镜像

#### 步骤 2: 运行容器

```bash
# 基础运行
docker run -d \
  --name joylive-dashboard \
  -p 8040:8040 \
  -v $(pwd)/data:/app/data \
  joylivedashboard:latest

# 使用自定义配置
docker run -d \
  --name joylive-dashboard \
  -p 8040:8040 \
  -v $(pwd)/data:/app/data \
  -v $(pwd)/configs/prod:/app/configs/prod \
  -e GIN_MODE=release \
  joylivedashboard:latest
```

#### 步骤 3: 验证部署

```bash
# 查看容器日志
docker logs -f joylive-dashboard

# 检查容器状态
docker ps | grep joylive-dashboard

# 测试健康检查端点
curl http://localhost:8040/health
```

### 使用外部 MySQL 数据库

如果要使用外部 MySQL 数据库：

1. 更新 `configs/prod/server.toml`：

```toml
[Storage.DB]
Type = "mysql"
DSN = "用户名:密码@tcp(mysql主机:3306)/joylive?charset=utf8mb4&parseTime=True&loc=Local"
```

2. 初始化数据库：

```bash
mysql -h mysql主机 -u 用户名 -p joylive < scripts/init.sql
```

3. 重启容器：

```bash
docker-compose restart joylive
```

## 配置说明

### 主要配置文件

- `configs/dev/server.toml` - 开发环境配置
- `configs/prod/server.toml` - 生产环境配置
- `configs/menu.json` - 英文菜单配置
- `configs/menu_cn.json` - 中文菜单配置

### 重要配置选项

#### 通用设置

```toml
[General]
AppName = "joylivedashboard"
Debug = true  # 生产环境设置为 false
DefaultLoginPwd = "21232f297a57a5a743894a0e4a801fc3"  # MD5("admin")
MenuFile = "menu.json"  # 中文界面使用 "menu_cn.json"

[General.HTTP]
Addr = ":8040"  # 服务端口
```

#### 数据库配置

```toml
[Storage.DB]
Type = "mysql"  # 选项: sqlite3, mysql, postgres
DSN = "root:密码@tcp(localhost:3306)/joylive?charset=utf8mb4&parseTime=True&loc=Local"
AutoMigrate = true  # 自动创建/更新表结构
```

#### 缓存配置

```toml
[Storage.Cache]
Type = "memory"  # 选项: memory, redis, badger

[Storage.Cache.Redis]
Addr = "127.0.0.1:6379"
Password = ""
DB = 1
```

### 前端环境变量

- `frontend/.env.dev` - 开发环境
- `frontend/.env.prod` - 生产环境

```env
# API 端点
VITE_API_HTTP=http://localhost:8040

# 应用标题
VITE_APP_TITLE=JoyLive Dashboard
```

## 验证部署

### 1. 检查服务状态

```bash
# 本地开发
curl http://localhost:8040/health

# Docker 部署
docker exec joylive-dashboard wget -qO- http://localhost:8040/health
```

预期响应: `{"status":"ok"}`

### 2. 访问 Web 界面

1. 在浏览器中打开 `http://localhost:8040`
2. 使用默认凭证登录（`admin` / `admin`）
3. 应该能看到仪表板首页

### 3. 验证数据库连接

检查日志中的数据库连接信息：

```bash
# 本地开发
# 查看终端输出

# Docker 部署
docker logs joylive-dashboard | grep -i "database"
```

### 4. 测试 API 端点

```bash
# 获取 API 版本
curl http://localhost:8040/api/v1/version

# 登录测试
curl -X POST http://localhost:8040/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'
```

## 故障排查

### 常见问题

#### 1. 端口已被占用

**错误:** `bind: address already in use`

**解决方案:**
```bash
# 查找占用 8040 端口的进程
lsof -i :8040

# 终止该进程
kill -9 <PID>

# 或在 configs/dev/server.toml 中更改端口
[General.HTTP]
Addr = ":8041"
```

#### 2. 数据库连接失败

**错误:** `Error connecting to database`

**解决方案:**
- 验证 MySQL 是否运行: `mysql -u root -p`
- 检查数据库是否存在: `SHOW DATABASES;`
- 验证 `configs/dev/server.toml` 中的凭证
- 确保数据库已初始化: `mysql -u root -p joylive < scripts/init.sql`

#### 3. 前端构建失败

**错误:** `npm run build failed`

**解决方案:**
```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
npm run build:prod
```

#### 4. Docker 构建失败

**错误:** `Docker build failed`

**解决方案:**
```bash
# 清理 Docker 缓存
docker system prune -a

# 无缓存重新构建
docker build --no-cache -t joylivedashboard:latest .
```

#### 5. 无法登录

**问题:** 使用正确凭证登录失败

**解决方案:**
- 检查数据库是否正确初始化
- 验证数据库中是否存在默认管理员用户
- 查看日志中的认证错误
- 尝试在数据库中重置密码:
  ```sql
  UPDATE sys_user SET password = '21232f297a57a5a743894a0e4a801fc3' WHERE username = 'admin';
  ```

### 查看日志

#### 本地开发

日志会打印到运行 `make start` 或 `make serve` 的控制台。

#### Docker 部署

```bash
# 查看所有日志
docker logs joylive-dashboard

# 实时跟踪日志
docker logs -f joylive-dashboard

# 查看最后 100 行
docker logs --tail 100 joylive-dashboard

# 使用 Docker Compose
docker-compose logs -f joylive
```

### 性能优化

#### 生产环境部署建议

1. **启用 Redis 缓存:**
   ```toml
   [Storage.Cache]
   Type = "redis"
   ```

2. **禁用调试模式:**
   ```toml
   [General]
   Debug = false
   ```

3. **优化数据库连接:**
   ```toml
   [Storage.DB]
   MaxOpenConns = 100
   MaxIdleConns = 50
   ```

4. **使用生产构建:**
   ```bash
   make build-all
   ```

## 下一步

成功启动应用后：

1. **修改默认密码**: 进入用户设置修改默认管理员密码
2. **配置空间**: 为您的服务创建微服务空间
3. **注册服务**: 将您的微服务添加到仪表板
4. **设置策略**: 配置路由、限流和熔断策略
5. **监控服务**: 使用仪表板监控服务健康状态和指标

## 其他资源

- [项目 README](../README-zh.md)
- [部署指南](../DEPLOY.md)
- [GitHub 仓库](https://github.com/jd-opensource/joylive-dashboard)
- [相关项目](../README-zh.md#关联项目)

## 获取帮助

如果遇到本指南未涵盖的问题：

1. 查看 [GitHub Issues](https://github.com/jd-opensource/joylive-dashboard/issues)
2. 检查日志中的错误信息
3. 查阅配置文件了解可用选项
4. 提交新的 issue 并提供详细的问题描述
