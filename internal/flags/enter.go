package flags

import (
	"ServeGin/internal/global"
	"flag"
	"fmt"
	"os"
)

/*
命名行参数结构体
*/
type FlagOptions struct {
	File    string // 配置文件
	Version bool   // 版本
	DB      bool   // 迁移数据库
}

var Options FlagOptions

/*
操作命令行,实现读取配置文件,版本，迁移数据库
*/
func Parse() {
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	flag.BoolVar(&Options.Version, "v", false, "打印当前版本")
	flag.BoolVar(&Options.DB, "db", false, "迁移数据库表结构")
	flag.Parse()
}

/*
运行函数,实现命令行，进行表结构迁移以及打印版本号
*/
func Run() {
	if Options.DB {
		MigrateDB()
		os.Exit(0)
	}
	if Options.Version {
		fmt.Println("当前后端的版本号", global.Version)
		os.Exit(0)
	}
}
