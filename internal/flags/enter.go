package flags

import (
	"flag"
)

/*
命名行参数结构体
*/
type FlagOptions struct {
	File string
}

var Options FlagOptions

/*
操作命令行,实现读取配置文件
*/
func Run() {
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	flag.Parse()
}
