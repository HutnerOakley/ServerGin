package global

import (
	"ServeGin/internal/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

/*
全局变量
*/
var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
