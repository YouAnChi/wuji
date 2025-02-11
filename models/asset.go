package models

import (
	"time"

	"gorm.io/gorm"
)

// Asset 资产模型
type Asset struct {
	gorm.Model
	UserID       uint        `json:"user_id" gorm:"index"`
	Name         string      `json:"name"`                                  // 资产名称
	Category     string      `json:"category"`                              // 分类
	Type         string      `json:"type"`                                  // 类型（如：相机、电脑等）
	Price        float64     `json:"price"`                                 // 购买价格
	PurchaseDate time.Time   `json:"purchase_date"`                         // 购买日期
	WarrantyEnd  time.Time   `json:"warranty_end"`                          // 保修截止日期
	DailyPrice   float64     `json:"daily_price"`                           // 每日成本
	Status       string      `json:"status"`                                // 使用状态
	Icon         string      `json:"icon"`                                  // 自定义图标
	Remark       string      `json:"remark"`                                // 备注
	ExtraCosts   []ExtraCost `json:"extra_costs" gorm:"foreignKey:AssetID"` // 额外费用
}

// ExtraCost 额外费用模型
type ExtraCost struct {
	gorm.Model
	AssetID    uint      `json:"asset_id" gorm:"index"`
	Name       string    `json:"name"`        // 费用名称
	Amount     float64   `json:"amount"`      // 费用金额
	ExpireDate time.Time `json:"expire_date"` // 到期时间（如会员）
}
