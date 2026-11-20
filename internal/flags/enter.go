package flags

import (
	"ServeGin/internal/global"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

/*
命名行参数结构体
*/
type FlagOptions struct {
	File    string // 配置文件
	Version bool   // 版本
	DB      bool   // 迁移数据库
	Menu    string // 菜单
	Type    string // 类型 create list remove
}

var Options FlagOptions

/*
操作命令行,实现读取配置文件,版本，迁移数据库
*/
func Parse() {
	flag.StringVar(&Options.File, "f", "settings-prod.yaml", "配置文件路径")
	flag.StringVar(&Options.Menu, "m", "", "菜单 user")
	flag.StringVar(&Options.Type, "t", "", "操作类型 create list")
	flag.BoolVar(&Options.Version, "v", false, "打印当前版本")
	flag.BoolVar(&Options.DB, "db", false, "迁移数据库表结构")
	flag.Parse()
}

/*
运行函数,实现命令行，进行表结构迁移、打印版本号、用户创建
*/
func Run() {
	if Options.DB {
		err := MigrateDB()
		if err != nil {
			logrus.Fatalf("数据库迁移失败，程序退出: %v", err)
		}
		logrus.Infof("表结构迁移成功")
		os.Exit(0)
	}
	if Options.Version {
		fmt.Println("当前后端的版本号", global.Version)
		os.Exit(0)
	}
	if Options.Menu == "user" {
		var user User
		switch Options.Type {
		case "create":
			user.Create()
		case "list":
			user.List()
		}
		os.Exit(0)
	}
}
