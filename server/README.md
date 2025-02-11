# 物纪后端服务

## 项目架构

```
server/
├── controllers/     # 控制器层，处理HTTP请求和业务逻辑
│   ├── asset_controller.go     # 资产相关接口
│   ├── category_controller.go  # 分类相关接口
│   └── user_controller.go      # 用户相关接口
├── models/         # 数据模型层，定义数据结构和数据库操作
│   ├── asset.go    # 资产模型
│   └── user.go     # 用户模型
├── main.go        # 程序入口，服务器配置和路由设置
└── README.md      # 项目说明文档
```

## 技术栈

- Go 1.16+
- RESTful API设计
- MySQL数据库
- JWT认证

## API接口说明

服务器默认运行在8080端口，前端通过微信小程序的`wx.request()`方法进行通信。

### 主要接口列表

#### 1. 用户认证
```
POST /api/auth/login
请求体：{ code: string }  // 微信登录code
响应：{ token: string }   // JWT认证token
```

#### 2. 用户信息
```
GET /api/user/profile
请求头：Authorization: Bearer <token>
响应：{
    id: string,
    nickname: string,
    avatar: string
}
```

#### 3. 资产管理
```
GET /api/assets            // 获取资产列表
GET /api/assets/:id        // 获取资产详情
POST /api/assets           // 创建新资产
PUT /api/assets/:id        // 更新资产信息
DELETE /api/assets/:id     // 删除资产
```

#### 4. 分类管理
```
GET /api/categories        // 获取分类列表
POST /api/categories       // 创建新分类
PUT /api/categories/:id    // 更新分类
DELETE /api/categories/:id // 删除分类
```

## 数据模型

### User模型
```go
type User struct {
    ID        string    `json:"id"`
    Nickname  string    `json:"nickname"`
    Avatar    string    `json:"avatar"`
    OpenID    string    `json:"open_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Asset模型
```go
type Asset struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CategoryID  string    `json:"category_id"`
    UserID      string    `json:"user_id"`
    Value       float64   `json:"value"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 开发环境配置

1. 安装Go 1.16或更高版本
2. 克隆项目代码
3. 安装依赖：
```bash
go mod download
```
4. 配置MySQL数据库
5. 创建配置文件`.env`：
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=wuji
JWT_SECRET=your_jwt_secret
```

## 本地开发

1. 启动服务器：
```bash
go run main.go
```

2. 服务器默认运行在`http://localhost:8080`

3. API测试：
```bash
# 测试服务器是否正常运行
curl http://localhost:8080/api/health
```

## 部署说明

1. 编译项目：
```bash
go build -o wuji
```

2. 配置生产环境：
- 设置必要的环境变量
- 配置MySQL数据库连接
- 配置域名和SSL证书（如需要）

3. 运行服务：
```bash
./wuji
```

建议使用PM2或Supervisor等进程管理工具来管理服务进程。

## 注意事项

1. 所有API请求（除登录接口外）都需要在请求头中携带JWT token
2. 请确保数据库配置正确且有足够的连接数
3. 建议在生产环境中启用HTTPS
4. 定期备份数据库
5. 监控服务器状态和日志