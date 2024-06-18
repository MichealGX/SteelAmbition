package database

import "github.com/go-redis/redis/v8"

func RedisLink(rdb **redis.Client) {
	// 链接redis
	*rdb = redis.NewClient(&redis.Options{
		Addr:     "121.36.4.215:6379", // Redis 服务器地址
		Password: "",                  // Redis 密码
		DB:       0,                   // Redis 数据库索引（0 表示默认数据库）
	})
}
