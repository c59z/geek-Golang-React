package controllers

import (
	"fmt"
	"geek-pc-gin/initializers"
	"geek-pc-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	offset := (page - 1) * pageSize

	var articles []models.Article
	result := initializers.DB.Limit(pageSize).Offset(offset).Find(&articles)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "未找到文章",
		})
		return
	}

	var totalCount int64
	initializers.DB.Model(&models.Article{}).Count(&totalCount)

	c.JSON(http.StatusOK, gin.H{
		"results":     articles,
		"total_count": totalCount,
	})
}

func DelArticle(c *gin.Context) {
	// 获取文章的 ID
	articleID := c.Param("id")

	// 转换 ID 为整数类型
	id, err := strconv.Atoi(articleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID错误"})
		return
	}

	// 查找要删除的文章
	var article models.Article
	if err := initializers.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在或者已被删除"})
		return
	}

	// 执行软删除 (也可以改为硬删除)
	if err := initializers.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func AddArticle(c *gin.Context) {
	var requestBody struct {
		ID        uint   `json:"ID"`
		ChannelID uint   `json:"channel_id" binding:"required"`
		Title     string `json:"title" binding:"required"`
		Content   string `json:"content" binding:"required"`
		Author    string `json:"author" binding:"required"`
	}

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断 作品ID传过来了，传过来就是编辑
	// 没有传过来就是新增

	fmt.Println("----------------")
	fmt.Println(requestBody.ID)
	fmt.Println("----------------")

	if requestBody.ID == 0 {
		// 新增文章
		article := models.Article{
			ChannelID: requestBody.ChannelID,
			Title:     requestBody.Title,
			Content:   requestBody.Content,
			Author:    requestBody.Author,
		}
		if err := initializers.DB.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "添加文章失败"})
			return
		}
		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "添加成功!"})
	} else {
		// 修改文章
		var existingArticle models.Article

		// 查找已有文章
		if err := initializers.DB.First(&existingArticle, requestBody.ID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在或已被删除"})
			return
		}

		// 更新文章内容
		existingArticle.ChannelID = requestBody.ChannelID
		existingArticle.Title = requestBody.Title
		existingArticle.Content = requestBody.Content

		if err := initializers.DB.Save(&existingArticle).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "修改文章失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "文章修改成功"})
	}

}

func GetArticleByID(c *gin.Context) {
	// 获取路径参数中的 id
	idParam := c.Param("id")

	// 将 id 转换为整数
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "参数错误"})
	}
	if 0 == id {
		return
	}

	var article models.Article
	if err := initializers.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "获取文章成功", "article": article})
}
