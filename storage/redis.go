package storage

import (
	"fmt"
	"time"

	"round_robin_with_weight/config"
	"round_robin_with_weight/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// RedisSession :
type RedisSession struct {
	Client *redis.Client
}

// Init : init the redis client
func (r *RedisSession) Init() error {
	redisConf := config.GlobalConfigurations.RedisConf

	r.Client = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%v:%v", redisConf.Host, redisConf.Port),
		Password:    redisConf.Password,
		DB:          redisConf.DefaultDB,
		DialTimeout: time.Duration(redisConf.MaxConnTimeout) * time.Second,
		ReadTimeout: time.Duration(redisConf.ReadTimeout) * time.Second,
		PoolSize:    redisConf.MaxPoolSize,
		IdleTimeout: time.Duration(redisConf.IdleTimeout),
	})
	return nil
}

// Close : close redis session
func (r *RedisSession) Close() {
	r.Client.Close()
}

// GetRedisKey : get redis key
func (r *RedisSession) GetRedisKey(c *gin.Context, key string) (val string, ok bool) {
	redisClient := r.Client
	val, err := redisClient.Get(key).Result()
	if err != nil {
		message := fmt.Sprintf("Get redis key error, detail:%v", err)
		utils.CommonResponse(c, 400, message, gin.H{})
		return "", false
	}
	return val, true
}

// SetRedisKey : set a key
func (r *RedisSession) SetRedisKey(c *gin.Context, key string, val string) bool {
	redisClient := r.Client
	// set key not expired
	err := redisClient.Set(key, val, 0).Err()
	if err != nil {
		message := fmt.Sprintf("Set redis key error, detail:%v", err)
		utils.CommonResponse(c, 400, message, gin.H{})
		return false
	}
	return true
}

// HGetRedisKey : hash get redis key
func (r *RedisSession) HGetRedisKey(c *gin.Context, key string, field string) (val string, ok bool) {
	redisClient := r.Client
	val, err := redisClient.HGet(key, field).Result()
	if err != nil {
		message := fmt.Sprintf("Hash get redis key error, detail:%v", err)
		utils.CommonResponse(c, 400, message, gin.H{})
		return "", false
	}
	return val, true
}

// HSetRedisKey : hash set a key
func (r *RedisSession) HSetRedisKey(c *gin.Context, key string, field string, val string) bool {
	redisClient := r.Client
	// set key not expired
	err := redisClient.HSet(key, field, val).Err()
	if err != nil {
		message := fmt.Sprintf("Set redis key error, detail:%v", err)
		utils.CommonResponse(c, 400, message, gin.H{})
		return false
	}
	return true
}
