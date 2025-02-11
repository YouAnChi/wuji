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

## 当前项目进度

### 已完成功能
1. 基础框架搭建
   - 完成项目结构设计和依赖配置
   - 实现数据库连接和配置
   - 集成Gin框架和路由系统

2. 数据模型实现
   - 设计并实现用户模型
   - 设计并实现分类模型
   - 设计并实现资产模型

3. 分类管理功能
   - 实现分类的创建、查询、更新和删除
   - 支持分类信息的完整管理
   - 实现分类下物品数量的统计

4. 基础安全措施
   - 实现基本的数据验证
   - 添加用户数据隔离机制

### 待开发功能

1. 用户认证系统（优先级：高）
   - 实现JWT认证中间件
   - 添加用户登录和注册接口
   - 实现微信小程序登录集成
   - 添加会员权限控制

2. 资产管理功能（优先级：高）
   - 实现资产CRUD接口
   - 添加资产图片上传功能
   - 实现资产折旧计算
   - 支持批量导入导出

3. 数据统计分析（优先级：中）
   - 实现资产总值统计
   - 添加分类占比分析
   - 实现折旧趋势分析
   - 支持自定义统计报表

4. 系统优化（优先级：中）
   - 添加缓存机制
   - 优化数据库查询
   - 实现数据库索引优化
   - 添加请求限流

5. 测试与文档（优先级：低）
   - 编写单元测试
   - 添加集成测试
   - 完善API文档
   - 添加代码注释

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