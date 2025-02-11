# 物纪服务端部署指南

## 环境要求

- Go 1.16 或更高版本
- MySQL 5.7 或更高版本
- Nginx（可选，用于反向代理）

## 部署步骤

### 1. 准备服务器环境

```bash
# 安装 Go
wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# 安装 MySQL
sudo apt update
sudo apt install mysql-server
sudo mysql_secure_installation
```

### 2. 配置数据库

```sql
# 创建数据库和用户
CREATE DATABASE wuji;
CREATE USER 'wuji'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON wuji.* TO 'wuji'@'localhost';
FLUSH PRIVILEGES;
```

### 3. 编译和部署服务

```bash
# 克隆代码到服务器
git clone <repository_url>
cd wuji/server

# 编译
go build -o wuji_server

# 创建配置文件
cat > config.yaml << EOL
database:
  host: localhost
  port: 3306
  user: wuji
  password: your_password
  dbname: wuji

server:
  port: 8080
  jwt_secret: your_jwt_secret
EOL
```

### 4. 创建系统服务

```bash
# 创建服务文件
sudo nano /etc/systemd/system/wuji.service

[Unit]
Description=Wuji Server
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/path/to/wuji/server
ExecStart=/path/to/wuji/server/wuji_server
Restart=always

[Install]
WantedBy=multi-user.target

# 启动服务
sudo systemctl enable wuji
sudo systemctl start wuji
```

### 5. 配置 Nginx 反向代理（可选）

```nginx
server {
    listen 80;
    server_name your_domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### 6. 配置 SSL 证书（推荐）

1. 使用 Let's Encrypt 获取免费的 SSL 证书
2. 配置 Nginx SSL 设置

```nginx
server {
    listen 443 ssl;
    server_name your_domain.com;

    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 维护说明

1. 日志查看：`journalctl -u wuji`
2. 服务状态检查：`systemctl status wuji`
3. 定期备份数据库：`mysqldump -u wuji -p wuji > backup.sql`

## 安全建议

1. 使用强密码
2. 定期更新系统和依赖包
3. 配置防火墙只开放必要端口
4. 启用 SSL 证书实现 HTTPS 访问
5. 定期备份数据

## 故障排查

1. 查看服务日志：`journalctl -u wuji -f`
2. 检查端口占用：`lsof -i :8080`
3. 检查数据库连接：`mysql -u wuji -p`