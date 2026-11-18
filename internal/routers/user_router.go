package routers

import (
	"ServeGin/internal/api"
	"ServeGin/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

/*
用户路由
*/
func UserRouter(g *gin.RouterGroup) {
	userApi := api.App.UserApi
	loginLimiter := middleware.NewLimiter(5, time.Minute)

	g.POST("users/login", loginLimiter.LimitMiddleware(), userApi.LoginView)
	g.GET("users", middleware.JWTAuthMiddleware(), userApi.UserList)
}
