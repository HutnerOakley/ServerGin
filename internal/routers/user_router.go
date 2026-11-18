package routers

import (
	"ServeGin/internal/api"

	"github.com/gin-gonic/gin"
)

/*
用户路由
*/
func UserRouter(g *gin.RouterGroup) {
	userApi := api.App.UserApi

	g.POST("users/login", userApi.LoginView)
}
