package controllers

import (
	"geek-pc-gin/initializers"
	"geek-pc-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetChannels(c *gin.Context) {

	// 根据 userID 查询用户信息
	var channel []models.Channel
	if err := initializers.DB.Unscoped().Find(&channel).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "无频道"})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"channels": channel,
	})
}
