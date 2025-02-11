package controllers

import (
	"net/http"
	"wuji/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// GetUserProfile 获取用户个人信息
func (uc *UserController) GetUserProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User

	if err := uc.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile 更新用户个人信息
func (uc *UserController) UpdateUserProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var updateData struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "个人信息更新成功"})
}
