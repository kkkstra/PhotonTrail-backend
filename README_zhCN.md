# AI 职达 - backend

## 项目概览

PhotonTrail 是一个专注于摄影作品分享与互动交流的社区。

此后端系统基于 Go 和 Gin Web 框架构建，使用 MySQL 作为数据库，并采用 JWT 进行认证。系统集成了 OSS STS Token 服务。

## 主要功能

- 用户认证：基于 JWT 的登录和注册功能。
- 帖子管理：创建、更新、删除和搜索帖子。
- OSS 集成：适用于 OSS 的 STS Token 服务。
- RESTful API：支持前后端的顺畅交互。

## 项目结构

```
├── Dockerfile            # Dockerfile
├── LICENSE               # 许可证文件
├── README.md             # 项目概览和说明
├── cmd                   # 应用入口
├── configs               # 配置文件
├── data                  # 数据存储
├── docker-compose.yml    # 多服务设置的 Docker Compose 配置
├── internal              # 内部代码，仅限于项目内部使用
├── pkg                   # 可复用的包
└── scripts               # 常用脚本
```

## 技术栈

- 语言：Go
- Web 框架：Gin
- 数据库：MySQL
- 认证：JWT

## 环境要求

- Go 1.16+
- MySQL 5.7+
- Python 3.8+（用于 OSS STS Token 服务）
- Docker（可选，部署用）

## 安装指南

### 1. 克隆仓库

```bash
git clone https://github.com/your-repo/PhotonTrail-backend.git
cd PhotonTrail-backend
```

### 2. 安装 Go 依赖

确保已安装 Go，运行以下命令：

```bash
go mod tidy
```

### 3. 设置 MySQL 数据库

确保 MySQL 运行并创建数据库：

```sql
CREATE DATABASE photon_trail;
```

### 4. 运行数据库迁移

运行所需的迁移以创建数据库结构。

### 5. 启动应用程序

在本地启动后端：

```bash
go run main.go
```

服务器将运行在 http://localhost:8001。

### 7. 启动 OSS STS TOKEN 服务

进入 oss-uploader 目录并启动：

```bash
cd oss-uploader
python main.py
```

服务将在 http://localhost:8000 运行。

## 部署

### Docker 部署

使用 Docker 部署后端：

1.	构建 Docker 镜像：

```bash
docker build -t PhotonTrail-backend .
```

2.	运行 Docker 容器：

```bash
docker run -p 8001:8000 PhotonTrail-backend
```

### Docker Compose 部署

```bash
docker compose up -d --build
```

## 贡献

欢迎贡献！如果您想要参与：

1.	Fork 此仓库。
2.	创建一个新分支（git checkout -b feature-xyz）。
3.	完成更改并提交（git commit -m 'Add new feature'）。
4.	将更改推送到您的 Fork 仓库（git push origin feature-xyz）。
5.	向主仓库提交 Pull Request。

## 许可证

此项目遵循 MIT 许可证。详情见 LICENSE 文件。