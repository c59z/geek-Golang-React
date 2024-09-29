package middlewares

import (
	"context"
	"geek-pc-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
)

var ctx = context.Background()

// AuthMiddleware 验证请求头中的 token
func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 放行 Login 和 Register 接口
		if strings.HasPrefix(c.Request.URL.Path, "/auth/login") || strings.HasPrefix(c.Request.URL.Path, "/auth/register") {
			c.Next()
			return
		}

		// 从请求头中获取 token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort() // 阻止请求继续处理
			return
		}

		username, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", username)

		// fmt.Println(username)
		// 从 Redis 中检查 token 是否有效
		val, err := rdb.Get(ctx, username).Result()
		if err == redis.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort() // 阻止请求继续处理
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token"})
			c.Abort() // 阻止请求继续处理
			return
		}

		// 如果 token 有效，将其保存到上下文中，供后续处理使用
		c.Set("user", val)

		// 继续处理请求
		c.Next()
	}
}
