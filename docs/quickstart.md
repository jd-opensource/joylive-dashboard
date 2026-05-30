# Quick Start Guide

> [中文版本](./quickstart-zh.md)

This guide will help you quickly set up and run the JoyLive Dashboard in different environments.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Local Development](#local-development)
- [Docker Deployment](#docker-deployment)
- [Configuration](#configuration)
- [Verification](#verification)
- [Troubleshooting](#troubleshooting)

## Prerequisites

### For Local Development

- **Go**: 1.19 or higher
- **Node.js**: 18.x or higher
- **npm**: 8.x or higher
- **MySQL**: 5.7 or higher (or use SQLite for testing)
- **Redis**: 6.0 or higher (optional, can use memory cache)
- **Git**: For cloning the repository

### For Docker Deployment

- **Docker**: 20.10 or higher
- **Docker Compose**: 1.29 or higher (optional, for easier deployment)

## Local Development

### Step 1: Clone the Repository

```bash
git clone https://github.com/jd-opensource/joylive-dashboard.git
cd joylive-dashboard
```

### Step 2: Database Setup

#### Option A: Using MySQL (Recommended for Production)

1. Create the database:

```sql
CREATE DATABASE joylive CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
```

2. Initialize the database schema:

```bash
mysql -u root -p joylive < scripts/init.sql
```

3. Update the database configuration in `configs/dev/server.toml`:

```toml
[Storage.DB]
Type = "mysql"
DSN = "username:password@tcp(localhost:3306)/joylive?charset=utf8mb4&parseTime=True&loc=Local"
```

#### Option B: Using SQLite (Quick Testing)

SQLite is enabled by default in development mode. The database file will be created automatically at `data/joylivedashboard.db`.

To use SQLite, update `configs/dev/server.toml`:

```toml
[Storage.DB]
Type = "sqlite3"
DSN = "data/joylivedashboard.db"
```

### Step 3: Configure Redis (Optional)

If you want to use Redis for caching, update `configs/dev/server.toml`:

```toml
[Storage.Cache]
Type = "redis"  # Change from "memory" to "redis"

[Storage.Cache.Redis]
Addr = "127.0.0.1:6379"
Password = ""
DB = 1
```

For development, you can keep the default memory cache.

### Step 4: Install Dependencies

#### Backend Dependencies

```bash
# Go dependencies are managed by go.mod
go mod download
```

#### Frontend Dependencies

```bash
cd frontend
npm install
cd ..
```

### Step 5: Build and Run

#### Option A: Build and Run Together

```bash
# Build frontend and backend, then start the server
make build-all
make serve
```

#### Option B: Development Mode (Separate Frontend and Backend)

**Terminal 1 - Backend:**

```bash
make start
```

**Terminal 2 - Frontend:**

```bash
cd frontend
npm run dev
```

In development mode:
- Backend runs on `http://localhost:8040`
- Frontend dev server runs on `http://localhost:5173` (with hot reload)
- Frontend proxies API requests to the backend

### Step 6: Access the Application

Open your browser and navigate to:

- **Production Build**: `http://localhost:8040`
- **Development Mode**: `http://localhost:5173`

**Default Credentials:**
- Username: `admin`
- Password: `admin`

## Docker Deployment

### Quick Start with Docker Compose (Recommended)

#### Step 1: Clone the Repository

```bash
git clone https://github.com/jd-opensource/joylive-dashboard.git
cd joylive-dashboard
```

#### Step 2: Configure (Optional)

Edit `configs/prod/server.toml` if you need to customize settings. The default configuration uses SQLite and is ready to use.

#### Step 3: Start with Docker Compose

```bash
docker-compose up -d
```

This will:
- Build the Docker image (if not already built)
- Start the container in detached mode
- Mount `./data` for persistent storage
- Expose port 8040

#### Step 4: Check Status

```bash
# View logs
docker-compose logs -f joylive

# Check container status
docker-compose ps
```

#### Step 5: Access the Application

Open your browser and navigate to `http://localhost:8040`

**Default Credentials:**
- Username: `admin`
- Password: `admin`

### Manual Docker Deployment

#### Step 1: Build the Docker Image

```bash
# Using Makefile
make docker-build

# Or using Docker directly
docker build -t joylivedashboard:latest .
```

The Docker build process:
1. Builds the frontend (Vue 3 + Vite)
2. Builds the backend (Go)
3. Creates a single optimized image with both components

#### Step 2: Run the Container

```bash
# Basic run
docker run -d \
  --name joylive-dashboard \
  -p 8040:8040 \
  -v $(pwd)/data:/app/data \
  joylivedashboard:latest

# With custom configuration
docker run -d \
  --name joylive-dashboard \
  -p 8040:8040 \
  -v $(pwd)/data:/app/data \
  -v $(pwd)/configs/prod:/app/configs/prod \
  -e GIN_MODE=release \
  joylivedashboard:latest
```

#### Step 3: Verify the Deployment

```bash
# Check container logs
docker logs -f joylive-dashboard

# Check container status
docker ps | grep joylive-dashboard

# Test health endpoint
curl http://localhost:8040/health
```

### Using External MySQL with Docker

If you want to use an external MySQL database:

1. Update `configs/prod/server.toml`:

```toml
[Storage.DB]
Type = "mysql"
DSN = "username:password@tcp(mysql-host:3306)/joylive?charset=utf8mb4&parseTime=True&loc=Local"
```

2. Initialize the database:

```bash
mysql -h mysql-host -u username -p joylive < scripts/init.sql
```

3. Restart the container:

```bash
docker-compose restart joylive
```

## Configuration

### Key Configuration Files

- `configs/dev/server.toml` - Development environment configuration
- `configs/prod/server.toml` - Production environment configuration
- `configs/menu.json` - English menu configuration
- `configs/menu_cn.json` - Chinese menu configuration

### Important Configuration Options

#### General Settings

```toml
[General]
AppName = "joylivedashboard"
Debug = true  # Set to false in production
DefaultLoginPwd = "21232f297a57a5a743894a0e4a801fc3"  # MD5("admin")
MenuFile = "menu.json"  # Or "menu_cn.json" for Chinese

[General.HTTP]
Addr = ":8040"  # Server port
```

#### Database Configuration

```toml
[Storage.DB]
Type = "mysql"  # Options: sqlite3, mysql, postgres
DSN = "root:password@tcp(localhost:3306)/joylive?charset=utf8mb4&parseTime=True&loc=Local"
AutoMigrate = true  # Automatically create/update tables
```

#### Cache Configuration

```toml
[Storage.Cache]
Type = "memory"  # Options: memory, redis, badger

[Storage.Cache.Redis]
Addr = "127.0.0.1:6379"
Password = ""
DB = 1
```

### Frontend Environment Variables

- `frontend/.env.dev` - Development environment
- `frontend/.env.prod` - Production environment

```env
# API endpoint
VITE_API_HTTP=http://localhost:8040

# Application title
VITE_APP_TITLE=JoyLive Dashboard
```

## Verification

### 1. Check Service Status

```bash
# For local development
curl http://localhost:8040/health

# For Docker deployment
docker exec joylive-dashboard wget -qO- http://localhost:8040/health
```

Expected response: `{"status":"ok"}`

### 2. Access the Web Interface

1. Open `http://localhost:8040` in your browser
2. Login with default credentials (`admin` / `admin`)
3. You should see the dashboard homepage

### 3. Verify Database Connection

Check the logs for successful database connection:

```bash
# Local development
# Check terminal output

# Docker deployment
docker logs joylive-dashboard | grep -i "database"
```

### 4. Test API Endpoints

```bash
# Get API version
curl http://localhost:8040/api/v1/version

# Login test
curl -X POST http://localhost:8040/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'
```

## Troubleshooting

### Common Issues

#### 1. Port Already in Use

**Error:** `bind: address already in use`

**Solution:**
```bash
# Find process using port 8040
lsof -i :8040

# Kill the process
kill -9 <PID>

# Or change the port in configs/dev/server.toml
[General.HTTP]
Addr = ":8041"
```

#### 2. Database Connection Failed

**Error:** `Error connecting to database`

**Solution:**
- Verify MySQL is running: `mysql -u root -p`
- Check database exists: `SHOW DATABASES;`
- Verify credentials in `configs/dev/server.toml`
- Ensure database is initialized: `mysql -u root -p joylive < scripts/init.sql`

#### 3. Frontend Build Failed

**Error:** `npm run build failed`

**Solution:**
```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
npm run build:prod
```

#### 4. Docker Build Failed

**Error:** `Docker build failed`

**Solution:**
```bash
# Clean Docker cache
docker system prune -a

# Rebuild with no cache
docker build --no-cache -t joylivedashboard:latest .
```

#### 5. Cannot Login

**Problem:** Login fails with correct credentials

**Solution:**
- Check if database is initialized properly
- Verify the default admin user exists in the database
- Check logs for authentication errors
- Try resetting the password in the database:
  ```sql
  UPDATE sys_user SET password = '21232f297a57a5a743894a0e4a801fc3' WHERE username = 'admin';
  ```

### Viewing Logs

#### Local Development

Logs are printed to the console where you ran `make start` or `make serve`.

#### Docker Deployment

```bash
# View all logs
docker logs joylive-dashboard

# Follow logs in real-time
docker logs -f joylive-dashboard

# View last 100 lines
docker logs --tail 100 joylive-dashboard

# With Docker Compose
docker-compose logs -f joylive
```

### Performance Optimization

#### For Production Deployment

1. **Enable Redis for caching:**
   ```toml
   [Storage.Cache]
   Type = "redis"
   ```

2. **Disable debug mode:**
   ```toml
   [General]
   Debug = false
   ```

3. **Optimize database connections:**
   ```toml
   [Storage.DB]
   MaxOpenConns = 100
   MaxIdleConns = 50
   ```

4. **Use production build:**
   ```bash
   make build-all
   ```

## Next Steps

After successfully starting the application:

1. **Change Default Password**: Go to User Settings and change the default admin password
2. **Configure Spaces**: Create microservice spaces for your services
3. **Register Services**: Add your microservices to the dashboard
4. **Set Up Policies**: Configure routing, rate limiting, and circuit breaker policies
5. **Monitor Services**: Use the dashboard to monitor service health and metrics

## Additional Resources

- [Project README](../README.md)
- [Deployment Guide](../DEPLOY.md)
- [GitHub Repository](https://github.com/jd-opensource/joylive-dashboard)
- [Related Projects](../README.md#related-projects)

## Getting Help

If you encounter issues not covered in this guide:

1. Check the [GitHub Issues](https://github.com/jd-opensource/joylive-dashboard/issues)
2. Review the logs for error messages
3. Consult the configuration files for available options
4. Open a new issue with detailed information about your problem
