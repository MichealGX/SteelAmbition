package battleData

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"platform/database"
	"strconv"
)

func GetEnergyData(c *gin.Context, rdb *redis.Client, ID uint) (database.UserModule, error) {
	// 获取用户的能量值数据
	var userModule database.UserModule
	ctx := context.Background()
	data, err := rdb.HGet(ctx, "user:"+strconv.Itoa(int(ID)), "modules").Bytes()
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to get battle data from redis"})
		return userModule, err
	}
	// 反序列化能量值数据
	err = json.Unmarshal(data, &userModule)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to unmarshal battle data"})
		return userModule, err
	}
	return userModule, nil
}
