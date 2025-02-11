package controllers

import (
	"net/http"
	"strconv"
	"wuji/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AssetController struct {
	DB *gorm.DB
}

// CreateAsset 创建资产
func (ac *AssetController) CreateAsset(c *gin.Context) {
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 计算每日成本
	days := asset.WarrantyEnd.Sub(asset.PurchaseDate).Hours() / 24
	asset.DailyPrice = asset.Price / float64(days)

	if err := ac.DB.Create(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
}

// GetAssets 获取用户的所有资产
func (ac *AssetController) GetAssets(c *gin.Context) {
	userID := c.GetUint("user_id")
	var assets []models.Asset

	if err := ac.DB.Where("user_id = ?", userID).Preload("ExtraCosts").Find(&assets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assets)
}

// UpdateAsset 更新资产信息
func (ac *AssetController) UpdateAsset(c *gin.Context) {
	assetID := c.Param("id")
	var asset models.Asset

	if err := ac.DB.First(&asset, assetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "资产不存在"})
		return
	}

	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 重新计算每日成本
	days := asset.WarrantyEnd.Sub(asset.PurchaseDate).Hours() / 24
	asset.DailyPrice = asset.Price / float64(days)

	if err := ac.DB.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
}

// DeleteAsset 删除资产
func (ac *AssetController) DeleteAsset(c *gin.Context) {
	assetID := c.Param("id")

	if err := ac.DB.Delete(&models.Asset{}, assetID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "资产已删除"})
}

// AddExtraCost 添加额外费用
func (ac *AssetController) AddExtraCost(c *gin.Context) {
	assetID := c.Param("id")
	var extraCost models.ExtraCost

	if err := c.ShouldBindJSON(&extraCost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 安全地将字符串ID转换为uint
	parsedID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的资产ID"})
		return
	}

	extraCost.AssetID = uint(parsedID)

	if err := ac.DB.Create(&extraCost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, extraCost)
}
