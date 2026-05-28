#!/bin/bash

# ==============================================================================
# JoyLive Dashboard ARM64 Linux 一键打包构建脚本
# ==============================================================================
# 该脚本在本地机器上运行，执行前端构建、Go 交叉编译，并将所有运行时所需文件
# 打包压缩为适用于 Linux ARM64 的发布包，以便一键上传到 VPS 运行。
# ==============================================================================

set -e

# 获取当前脚本所在目录及项目根目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." &>/dev/null && pwd)"
RELEASE_DIR="$PROJECT_ROOT/release"
PACK_DIR="$RELEASE_DIR/joylive-dashboard"
PACK_NAME="joylive-dashboard-linux-arm64.tar.gz"

echo "========================================="
echo " 开始制作 Linux ARM64 原生部署包 "
echo "========================================="

# 1. 环境依赖校验
echo "--> 1. 正在检查本地构建环境..."
if ! command -v go &>/dev/null; then
    echo "❌ 错误: 未检测到 Go 编译环境，请先安装 Go (1.20+)。"
    exit 1
fi
if ! command -v npm &>/dev/null; then
    echo "❌ 错误: 未检测到 Node.js/NPM，请先安装 Node.js (18+)。"
    exit 1
fi
echo "✅ 本地环境检查通过 (Go & Node.js/NPM 已就绪)"

# 2. 清理历史构建
echo "--> 2. 正在清理历史构建数据..."
cd "$PROJECT_ROOT"
make clean || true
rm -rf "$RELEASE_DIR"
mkdir -p "$PACK_DIR"

# 3. 编译前端静态文件
echo "--> 3. 正在编译 Vue 前端静态文件..."
make build-frontend

# 4. 交叉编译 Go 后端二进制 (Linux ARM64, 禁用 CGO)
echo "--> 4. 正在交叉编译 Linux ARM64 后端二进制..."
make build-linux-arm64

# 验证编译产物是否存在
if [ ! -f "$PROJECT_ROOT/bin/joylivedashboard_linux_arm64" ]; then
    echo "❌ 错误: 未找到编译后的二进制文件 bin/joylivedashboard_linux_arm64"
    exit 1
fi

# 5. 组装发布包目录结构
echo "--> 5. 正在组装发布包文件..."

# 5.1 复制二进制文件并重命名为标准名称
cp "$PROJECT_ROOT/bin/joylivedashboard_linux_arm64" "$PACK_DIR/joylivedashboard"
chmod +x "$PACK_DIR/joylivedashboard"

# 5.2 复制前端静态文件
mkdir -p "$PACK_DIR/dist"
cp -r "$PROJECT_ROOT/frontend/dist/"* "$PACK_DIR/dist/"

# 5.3 复制配置文件
mkdir -p "$PACK_DIR/configs"
cp -r "$PROJECT_ROOT/configs/prod" "$PACK_DIR/configs/prod"
cp "$PROJECT_ROOT/configs/menu.json" "$PACK_DIR/configs/"
cp "$PROJECT_ROOT/configs/menu_cn.json" "$PACK_DIR/configs/"
cp "$PROJECT_ROOT/configs/rbac_model.conf" "$PACK_DIR/configs/"

# 5.4 复制 systemd 和 Nginx 模板
mkdir -p "$PACK_DIR/scripts"
cp "$PROJECT_ROOT/scripts/joylivedashboard.service" "$PACK_DIR/scripts/"
cp "$PROJECT_ROOT/scripts/nginx.conf" "$PACK_DIR/scripts/"

# 5.5 创建空的 data 目录以供 SQLite 数据库使用
mkdir -p "$PACK_DIR/data"

# 5.6 写入快捷运行脚本 start.sh 到发布包中
cat << 'EOF' > "$PACK_DIR/start.sh"
#!/bin/bash
# 进入脚本所在目录，确保相对路径有效
cd "$(dirname "$0")"

if [ -f pid ] && kill -0 $(cat pid) 2>/dev/null; then
    echo "⚠️ JoyLive Dashboard 已经在运行中 (PID: $(cat pid))."
    exit 1
fi

# 创建必要的数据存放目录
mkdir -p data

echo "正在启动 JoyLive Dashboard..."
# 在后台静默启动服务，指定工作路径
nohup ./joylivedashboard start -d configs -c prod -s dist > stdout.log 2>&1 &
echo $! > pid

# 等待 1 秒检查进程是否成功存活
sleep 1
if kill -0 $(cat pid) 2>/dev/null; then
    echo "✅ JoyLive Dashboard 成功启动 (PID: $(cat pid))"
    echo "访问地址: http://localhost:8040 (若在VPS上，请替换为您的VPS IP)"
    echo "实时日志请查看: tail -f stdout.log"
else
    echo "❌ 启动失败，请检查 stdout.log 中的错误日志。"
    cat stdout.log
    rm -f pid
fi
EOF
chmod +x "$PACK_DIR/start.sh"

# 5.7 写入快捷停止脚本 stop.sh 到发布包中
cat << 'EOF' > "$PACK_DIR/stop.sh"
#!/bin/bash
cd "$(dirname "$0")"

if [ -f pid ]; then
    PID=$(cat pid)
    echo "正在停止 JoyLive Dashboard (PID: $PID)..."
    kill $PID
    
    # 等待进程优雅退出
    for i in {1..10}; do
        if kill -0 $PID 2>/dev/null; then
            sleep 0.5
        else
            break
        fi
    done
    
    if kill -0 $PID 2>/dev/null; then
        echo "⚠️ 进程未响应，正在强制结束进程..."
        kill -9 $PID
    fi
    rm -f pid
    echo "✅ JoyLive Dashboard 已停止。"
else
    echo "⚠️ 未发现 pid 文件，尝试使用 pkill 停止进程..."
    pkill joylivedashboard || true
    echo "✅ 操作完成。"
fi
EOF
chmod +x "$PACK_DIR/stop.sh"

# 5.8 写入发布包自述文件 README.md
cat << 'EOF' > "$PACK_DIR/README.md"
# JoyLive Dashboard ARM64 Linux 部署说明书

本部署包专为 Linux ARM64 (aarch64) 架构的 VPS 定制，包含已编译的后端服务及前端静态资源。

## 目录结构说明

- `joylivedashboard`：运行二进制文件
- `configs/`：系统配置文件目录（默认使用 `configs/prod` 的配置）
- `dist/`：前端打包出的静态资源文件
- `data/`：本地数据目录（SQLite 数据库等）
- `start.sh`：一键后台启动脚本
- `stop.sh`：一键停止服务脚本
- `scripts/`：包含 systemd 服务配置及 Nginx 反代配置样例

---

## 部署与运行方法

### 方式一：快捷脚本启动（适合临时测试/开发）

1. 启动服务：
   ```bash
   ./start.sh
   ```
2. 停止服务：
   ```bash
   ./stop.sh
   ```
3. 查看运行日志：
   ```bash
   tail -f stdout.log
   ```

---

### 方式二：使用 Systemd 守护进程运行（推荐生产环境使用）

使用 systemd 可以实现开机自启、崩溃自动拉起，以及标准的系统日志托管。

1. 建议将本文件夹整体移动到 `/opt/joylive-dashboard`：
   ```bash
   sudo mkdir -p /opt
   sudo mv joylive-dashboard /opt/
   ```
2. 将服务描述文件复制到系统服务目录：
   ```bash
   sudo cp /opt/joylive-dashboard/scripts/joylivedashboard.service /etc/systemd/system/
   ```
3. 重载 systemd 配置并启动服务：
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable joylivedashboard
   sudo systemctl start joylivedashboard
   ```
4. 查看服务运行状态：
   ```bash
   sudo systemctl status joylivedashboard
   ```
5. 查看日志：
   ```bash
   journalctl -u joylivedashboard -f
   ```

---

## 域名与反向代理 (Nginx)

如果需要绑定域名并配置 SSL (HTTPS)，可以参考 `scripts/nginx.conf` 样例进行配置。

1. 修改 `/opt/joylive-dashboard/scripts/nginx.conf` 中的 `server_name` 为您的域名。
2. 将该配置内容合入 `/etc/nginx/sites-enabled/` 中。
3. 重新加载 Nginx 服务：`sudo systemctl reload nginx`。
EOF

# 6. 打包发布包
echo "--> 6. 正在打包发布包为 tar.gz..."
cd "$RELEASE_DIR"
tar -czf "$PACK_NAME" joylive-dashboard

# 7. 清理临时生成的组装文件夹
rm -rf "$PACK_DIR"

echo "========================================="
echo "🎉 恭喜！发布包制作成功！"
echo "打包输出路径: $RELEASE_DIR/$PACK_NAME"
echo "文件大小: $(du -sh "$RELEASE_DIR/$PACK_NAME" | cut -f1)"
echo "========================================="
echo "接下来，您只需通过 SCP 将此文件上传至您的 ARM VPS 即可："
echo "  scp $RELEASE_DIR/$PACK_NAME root@your_vps_ip:/tmp/"
echo "========================================="
