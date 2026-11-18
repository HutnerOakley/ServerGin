package routers

import (
	"ServeGin/internal/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
运行gin服务
*/
func Run() {
	// 设置gin的运行模式
	gin.SetMode(global.Config.System.Mode)
	r := gin.Default()
	// 静态目录代理
	r.Static("/uploads", "uploads")

	// 路由分组
	g := r.Group("api")
	UserRouter(g)

	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		logrus.Infof("后端服务运行在 %s", addr)
	}

	r.Run(addr)
}
