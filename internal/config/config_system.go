package config

import "fmt"

/*
web服务系统配置
*/
type System struct {
	Mode string `yaml:"mode"` // gin运行模式
	IP   string `yaml:"ip"`   // 运行ip
	Port int    `yaml:"port"` // 运行端口
}

/*
获取addr
*/
func (sys *System) Addr() string {
	return fmt.Sprintf("%s:%d", sys.IP, sys.Port)
}
