package controllers

import (
	"geek-pc-gin/initializers"
	"geek-pc-gin/models"
	"geek-pc-gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	//fmt.Println(credentials.Username)
	//fmt.Println(credentials.Password)

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//fmt.Println("开始调用登录服务")
	// 调用登录服务
	token, err := services.Login(initializers.DB, initializers.Redis, credentials.Username, credentials.Password)
	if err != nil {
		//fmt.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println("开始返回token")
	// 返回 Token
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
}

func Register(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 绑定请求 JSON 数据
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 调用服务层的 RegisterUser 函数进行注册
	err := services.RegisterUser(initializers.DB, credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Test(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "You have access", "user": user})
}

func Profile(c *gin.Context) {
	// 从 JWT Token 中获取用户 ID
	username, exists := c.Get("username")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// 根据 userID 查询用户信息
	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"ID":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"photo":    user.Photo, // 假设有这个字段
	})
}
