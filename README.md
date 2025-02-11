# Wuji 项目文档

## 项目概述
Wuji是一个基于Go语言开发的资产管理系统，使用Gin框架作为Web服务器，GORM作为ORM框架，MySQL作为数据存储。该系统支持用户管理个人资产、自定义分类、计算资产折旧等功能。

## 技术栈
- 后端框架：Gin v1.10.0
- 数据库：MySQL
- ORM框架：GORM v1.25.5
- 其他依赖：
  - golang-jwt/jwt v5.0.0（认证）
  - go-playground/validator v10.20.0（数据验证）

## 项目结构
```
/wuji
├── controllers/         # 控制器层，处理HTTP请求
│   ├── asset_controller.go    # 资产管理控制器
│   └── category_controller.go # 分类管理控制器
├── models/             # 数据模型层
│   ├── asset.go       # 资产相关模型
│   └── user.go        # 用户相关模型
└── main.go            # 应用入口文件
```

## 数据模型设计

### User 模型
```go
type User struct {
    gorm.Model
    OpenID      string     // 微信OpenID
    Nickname    string     // 用户昵称
    AvatarURL   string     // 头像URL
    MemberUntil *string    // 会员到期时间
    Assets      []Asset    // 用户的资产
    Categories  []Category // 用户自定义分类
}
```

### Category 模型
```go
type Category struct {
    gorm.Model
    UserID      uint   // 用户ID
    Name        string // 分类名称
    Icon        string // 分类图标
    Description string // 分类描述
    ItemCount   int    // 分类下的物品数量（非数据库字段）
}
```

## API接口文档

### 分类管理接口

#### 创建分类
- 请求方法：POST
- 路径：/categories
- 请求体：
```json
{
    "name": "电子产品",
    "icon": "electronics",
    "description": "电子设备分类"
}
```
- 响应：返回创建的分类信息

#### 获取分类列表
- 请求方法：GET
- 路径：/categories
- 响应：返回用户所有分类列表，包含每个分类下的物品数量

#### 更新分类
- 请求方法：PUT
- 路径：/categories/:id
- 请求体：同创建分类
- 响应：返回更新后的分类信息

#### 删除分类
- 请求方法：DELETE
- 路径：/categories/:id
- 响应：返回删除成功消息

## 核心功能实现

### 1. 用户资产管理
- 支持创建、查询、更新和删除资产信息
- 自动计算资产每日成本（基于购买价格和保修期限）
- 支持记录额外费用支出

### 2. 分类管理
- 支持用户自定义资产分类
- 动态统计各分类下的资产数量
- 分类信息支持图标和描述

### 3. 数据安全
- 所有接口都需要用户认证
- 资产和分类操作都会验证用户身份
- 数据访问严格限制在用户自己的范围内

## 部署说明

### 环境要求
- Go 1.21或以上
- MySQL 5.7或以上

### 数据库配置
1. 创建数据库：
```sql
CREATE DATABASE wuji CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 配置数据库连接：
在`main.go`中修改数据库连接信息：
```go
dsn := "root:password@tcp(127.0.0.1:3306)/wuji?charset=utf8mb4&parseTime=True&loc=Local"
```

### 启动服务
1. 安装依赖：
```bash
go mod download
```

2. 运行服务：
```bash
go run main.go
```

服务将在8080端口启动，可以通过 http://localhost:8080 访问。

## 开发计划
1. 添加用户认证中间件
2. 实现资产统计和分析功能
3. 添加图片上传功能
4. 优化性能和安全性
5. 添加单元测试和集成测试