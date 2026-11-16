package core

import (
	"ServeGin/internal/config"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

/*
读取配置项
*/
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile("settings.yaml")
	if err != nil {
		fmt.Printf("配置文件读取错误 %s", err)
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		fmt.Printf("配置文件格式错误 %s", err)
		return
	}
	return
}
