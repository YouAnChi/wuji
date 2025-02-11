package main

import (
	"log"
	"net/http"
	"wuji/server/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化数据库连接
	dsn := "root:12345678@tcp(127.0.0.1:3306)/wuji?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 初始化控制器
	assetController := &controllers.AssetController{DB: db}
	categoryController := &controllers.CategoryController{DB: db}
	userController := &controllers.UserController{DB: db}

	r := gin.Default()

	// 配置CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// API路由组
	api := r.Group("/api")
	{
		// 资产相关路由
		assets := api.Group("/assets")
		{
			assets.GET("/", assetController.GetAssets)
			assets.POST("/", assetController.CreateAsset)
			assets.GET("/:id", assetController.GetAsset)
			assets.PUT("/:id", assetController.UpdateAsset)
			assets.DELETE("/:id", assetController.DeleteAsset)
		}

		// 分类相关路由
		categories := api.Group("/categories")
		{
			categories.GET("/", categoryController.GetCategories)
			categories.POST("/", categoryController.CreateCategory)
			categories.GET("/:id", categoryController.GetCategory)
			categories.PUT("/:id", categoryController.UpdateCategory)
			categories.DELETE("/:id", categoryController.DeleteCategory)
		}

		// 用户相关路由
		users := api.Group("/users")
		{
			users.GET("/profile", userController.GetUserProfile)
			users.PUT("/profile", userController.UpdateUserProfile)
		}
	}

	// 启动服务器
	log.Printf("Server is running on http://localhost:8080")
	r.Run(":8080")
}
