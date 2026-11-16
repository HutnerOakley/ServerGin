package core

import (
	"ServeGin/internal/config"
	"ServeGin/internal/flags"
	"ServeGin/internal/global"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

/*
读取配置项
*/
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		logrus.Fatalf("配置文件读取错误 %s", err)
		return
	}
	// 读取yaml文件
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		logrus.Fatalf("配置文件格式错误 %s", err)
		return
	}
	logrus.Infof("%s 配置文件读取成功", flags.Options.File)
	return
}

/*
修改配置文件
*/
func DumpConfig() {
	// 将配置转化为yaml字节流
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		logrus.Errorf("配置文件转换错误 %s", err)
		return
	}
	// 把yaml字节流写入到文件中
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("文件写入错误 %s", err)
		return
	}
	logrus.Infof("配置文件写入成功")
}
