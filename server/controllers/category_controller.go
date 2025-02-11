package controllers

import (
	"net/http"
	"wuji/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

// CreateCategory 创建分类
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	userID := c.GetUint("user_id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.UserID = userID
	if err := cc.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetCategories 获取用户的所有分类
func (cc *CategoryController) GetCategories(c *gin.Context) {
	userID := c.GetUint("user_id")
	var categories []models.Category

	if err := cc.DB.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取每个分类下的物品数量
	for i := range categories {
		var count int64
		cc.DB.Model(&models.Asset{}).Where("user_id = ? AND category = ?", userID, categories[i].Name).Count(&count)
		categories[i].ItemCount = int(count)
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateCategory 更新分类信息
func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")
	userID := c.GetUint("user_id")
	var category models.Category

	if err := cc.DB.Where("id = ? AND user_id = ?", categoryID, userID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.UserID = userID
	if err := cc.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetCategory 获取单个分类详情
func (cc *CategoryController) GetCategory(c *gin.Context) {
	categoryID := c.Param("id")
	userID := c.GetUint("user_id")
	var category models.Category

	if err := cc.DB.Where("id = ? AND user_id = ?", categoryID, userID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	// 获取该分类下的物品数量
	var count int64
	cc.DB.Model(&models.Asset{}).Where("user_id = ? AND category = ?", userID, category.Name).Count(&count)
	category.ItemCount = int(count)

	c.JSON(http.StatusOK, category)
}

// DeleteCategory 删除分类
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")
	userID := c.GetUint("user_id")

	if err := cc.DB.Where("id = ? AND user_id = ?", categoryID, userID).Delete(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "分类已删除"})
}
