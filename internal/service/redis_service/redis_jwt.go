package redis_service

import (
	"ServeGin/internal/global"
	"context"
	"fmt"
	"time"
)

var (
	reKey = "logout_%s" // redis中黑名单key前缀
)

/*
将token放入到黑名单中
*/
func AddTokenBlacklist(ctx context.Context, token string, expiration time.Duration) error {
	if expiration <= 0 {
		return nil
	}
	// 放入redis黑名单
	key := fmt.Sprintf(reKey, token)
	err := global.Redis.Set(ctx, key, 1, expiration).Err()
	return err
}

/*
检查token是否在黑名单中
*/
func HasTokenBlacklist(ctx context.Context, token string) (bool, error) {
	key := fmt.Sprintf(reKey, token)
	count, err := global.Redis.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
