package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	OpenID      string     `json:"open_id" gorm:"uniqueIndex"`          // 微信OpenID
	Nickname    string     `json:"nickname"`                            // 用户昵称
	AvatarURL   string     `json:"avatar_url"`                          // 头像URL
	MemberUntil *string    `json:"member_until"`                        // 会员到期时间
	Assets      []Asset    `json:"assets" gorm:"foreignKey:UserID"`     // 用户的资产
	Categories  []Category `json:"categories" gorm:"foreignKey:UserID"` // 用户自定义分类
}

// Category 自定义分类模型
type Category struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"index"`
	Name        string `json:"name"`                // 分类名称
	Icon        string `json:"icon"`                // 分类图标
	Description string `json:"description"`         // 分类描述
	ItemCount   int    `json:"item_count" gorm:"-"` // 分类下的物品数量，不存储到数据库
}
