package middleware

import (
	"ServeGin/internal/utils/res"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

/*
限流器结构体
*/
type Limiter struct {
	limit      int                // 限制的请求数量
	duration   time.Duration      // 时间窗口
	timestamps map[string][]int64 // 请求的时间戳
	mu         sync.Mutex         // 锁
}

/*
创建限流器
*/
func NewLimiter(limiter int, duration time.Duration) *Limiter {
	return &Limiter{
		limit:      limiter,
		duration:   duration,
		timestamps: make(map[string][]int64),
	}
}

/*
限流中间件
*/
func (l *Limiter) LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP() // 获取客户端ip

		l.mu.Lock() // 加锁
		defer l.mu.Unlock()

		// 检查请求时间戳切片是否存在
		if _, ok := l.timestamps[ip]; !ok {
			l.timestamps[ip] = make([]int64, 0)
		}
		now := time.Now().Unix() // 当前时间戳

		// 移除过期的请求时间戳
		for i := 0; i < len(l.timestamps[ip]); i++ {
			if l.timestamps[ip][i] < now-int64(l.duration.Seconds()) {
				l.timestamps[ip] = append(l.timestamps[ip][:i], l.timestamps[ip][i+1:]...)
				i--
			}
		}

		// 检查请求数量是否超过限制
		if len(l.timestamps[ip]) >= l.limit {
			res.FailWithMsg("请求过多", c)
			c.Abort()
			return
		}

		// 添加当前请求时间戳到切片
		l.timestamps[ip] = append(l.timestamps[ip], now)

		// 继续处理请求
		c.Next()
	}
}
