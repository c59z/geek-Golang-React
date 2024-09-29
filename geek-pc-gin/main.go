package main

import (
	"geek-pc-gin/config"
	"geek-pc-gin/initializers"
	"geek-pc-gin/middlewares"
	"geek-pc-gin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	initializers.InitDB()
	initializers.InitRedis()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.Use(middlewares.AuthMiddleware(initializers.Redis))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
