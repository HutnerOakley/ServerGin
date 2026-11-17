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
	flags.Run()
	global.Config = core.ReadConfig()
}
