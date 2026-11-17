package core

import (
	"ServeGin/internal/global"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

/*
redis初始化
*/
func InitRedis() (client *redis.Client) {
	cfg := global.Config.Redis
	if cfg.Addr == "" {
		logrus.Warnf("未配置redis连接")
		return
	}
	// 初始化配置
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	// 建立连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("连接redis失败:%v", err)
		return
	}
	logrus.Infof("redis连接成功")
	return rdb
}
