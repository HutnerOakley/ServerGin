package middleware

import (
	"ServeGin/internal/utils/jwts"
	"ServeGin/internal/utils/res"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
认证中间件
*/
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取到请求头中的token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.FailWithMsg("认证失败", c)
			c.Abort() // 阻断
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			res.FailWithMsg("Authorization 格式错误", c)
			c.Abort()
			return
		}

		// parts[1]获取到的token，然后进行解析验证
		claims, err := jwts.ParseToken(parts[1])
		if err != nil {
			res.FailWithMsg("访问失败，无效的token，请登录!", c)
			c.Abort()
			return
		}

		// 将用户信息存入到上下文，方便后续获取
		c.Set("userID", claims.UserID)
		c.Set("roleID", claims.RoleID)
		c.Next() // 后续的处理函数可以通过c.Get()来获取当前请求用户的信息
	}
}
