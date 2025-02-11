# 物纪 - 个人资产管理小程序

## 项目介绍
物纪是一款专注于个人资产管理的微信小程序，帮助用户更好地管理和追踪个人物品资产。通过直观的界面和完善的功能，用户可以轻松记录、分类和管理各类物品，实现资产的数字化管理。

## 技术架构
### 前端技术栈
- 微信小程序原生框架
- WXML + WXSS + JavaScript
- 微信小程序组件库

### 后端技术栈
- Go语言
- RESTful API设计

## 前后端通信

### 通信方式
前端通过微信小程序提供的 `wx.request()` 方法与后端服务器进行通信。所有请求都发送到后端服务器的 8080 端口。

### 接口认证
所有需要认证的接口都通过 HTTP Header 中的 Authorization 字段传递 token：
```javascript
header: {
  'Authorization': wx.getStorageSync('token')
}
```

### 主要API接口

1. 用户认证
```
POST /api/auth/login
请求体：{ code: string }
响应：{ token: string }
```

2. 获取用户信息
```
GET /api/user/profile
响应：用户信息对象
```

3. 资产管理
```
GET /api/assets            // 获取资产列表
GET /api/assets/:id        // 获取资产详情
POST /api/assets           // 创建新资产
```

4. 分类管理
```
GET /api/categories        // 获取分类列表
```

### 错误处理
前端统一通过 `wx.showToast()` 展示错误信息：
```javascript
wx.showToast({
  title: '加载失败',
  icon: 'error'
});
```

## 功能模块

### 1. 首页
- 资产总览
- 最近添加的物品
- 资产统计数据

### 2. 分类管理
- 物品分类列表
- 添加/编辑分类
- 分类统计信息

### 3. 资产管理
- 物品列表
- 添加/编辑物品
- 物品详情查看
- 资产状态追踪

### 4. 用户中心
- 用户信息管理
- 会员功能
- 系统设置

## 数据模型

### 用户模型 (User)
```go
type User struct {
    ID          string
    Nickname    string
    AvatarURL   string
    MemberUntil time.Time
}
```

### 资产模型 (Asset)
```go
type Asset struct {
    ID          string
    Name        string
    CategoryID  string
    Price       float64
    Status      string
    CreateTime  time.Time
}
```

## 项目进度

### 已完成功能
- [x] 基础界面框架搭建
- [x] 用户登录功能
- [x] 分类管理基础功能
- [x] 资产列表展示

### 开发中功能
- [ ] 资产详情页面
- [ ] 数据统计分析
- [ ] 会员系统
- [ ] 数据导出功能

## 本地开发

1. 克隆项目
```bash
git clone [项目地址]
```

2. 安装依赖
```bash
go mod download
```

3. 启动开发服务器
```bash
go run main.go
```

4. 使用微信开发者工具打开项目目录

5. 配置域名
- 登录微信小程序管理后台
- 进入「开发」-「开发设置」
- 在「服务器域名」中添加以下域名到request合法域名列表：
  - https://tcb-api.tencentcloudapi.com

## 部署说明

1. 编译后端服务
```bash
go build -o wuji
```

2. 配置服务器环境
- 确保Go环境已安装
- 配置必要的环境变量
- 设置数据库连接

3. 小程序发布
- 通过微信小程序管理后台上传代码
- 提交审核并发布

## 贡献指南
欢迎提交Issue和Pull Request，一起完善这个项目。

## 开源协议
MIT License