package main

import (
	"ServeGin/internal/core"
	"ServeGin/internal/flags"
	"ServeGin/internal/global"
	"ServeGin/internal/routers"
)

/*
入口
*/
func main() {
	// 日志初始化
	core.InitLogger()
	// 解析命令行
	flags.Parse()
	// 配置文件
	global.Config = core.ReadConfig()
	// 数据库
	global.DB = core.InitGorm()
	// redis
	global.Redis = core.InitRedis()
	// 执行迁移
	flags.Run()

	// 运行web服务
	routers.Run()

}
