package main

import (
	"ServeGin/internal/core"
	"ServeGin/internal/flags"
	"ServeGin/internal/global"
)

/*
入口
*/
func main() {
	// 日志初始化
	core.InitLogger()
	// 命令行
	flags.Run()
	// 配置文件
	global.Config = core.ReadConfig()
	// 数据库
	global.DB = core.InitGorm()
	// redis
	global.Redis = core.InitRedis()
}
