package battleData

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

func GetTTL(c *gin.Context, rdb *redis.Client, key string) (time.Duration, error) {
	ctx := context.Background()
	ttl, err := rdb.TTL(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			c.JSON(400, gin.H{"code": 1, "msg": "Key does not exist or has no expiration time"})
			return -1, nil // 键不存在或没有设置过期时间
		}
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to get TTL"})
		return 0, err // 其他错误
	}
	if ttl == -1 {
		c.JSON(400, gin.H{"code": 1, "msg": "Key has no expiration time"})
		return -1, nil // 键没有设置过期时间
	}
	// 将秒转换为分钟
	//ttlInMinutes := ttl / 60
	return ttl, nil
}
