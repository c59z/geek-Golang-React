package routes

import (
	"geek-pc-gin/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// 创建 Auth 相关的路由组
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/test", controllers.Test)
		authGroup.GET("/profile", controllers.Profile)
		// authGroup.POST("/register", controllers.Register)
	}

	channelGroup := router.Group("/channel")
	{
		channelGroup.GET("/get", controllers.GetChannels)
	}

	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/get", controllers.GetArticles)
		articleGroup.DELETE("/del/:id", controllers.DelArticle)
		articleGroup.POST("/add", controllers.AddArticle)
		articleGroup.GET("/findOne/:id", controllers.GetArticleByID)
	}

}
