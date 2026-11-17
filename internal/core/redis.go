package core

import (
	"ServeGin/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() {
	var conf config.Redis
	// 初始化配置
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	// 建立连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("连接redis失败:%v", err)
		return
	}
	logrus.Infof("redis连接成功")
}
