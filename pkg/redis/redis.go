package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"starzeng.com/gin-demo/config"
	"strconv"
	"time"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	cfg := config.AppConfig.Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		retryCount := 0
		maxRetries := 3 // 设置最大重试次数
		for retryCount < maxRetries {
			fmt.Printf("Redis连接失败: %v, 正在尝试重新连接...\n", err)
			_, err = RDB.Ping(Ctx).Result()
			if err == nil {
				break
			}
			retryCount++
			time.Sleep(2 * time.Second) // 添加延迟再重试
		}
		if err != nil {
			log.Fatalf("Redis连接失败，已达到最大重试次数: %v", err)
		}
	}

	fmt.Println("Redis连接成功")
}

// Set 设置 key，带过期时间（秒）
func Set(key string, value interface{}, expiration time.Duration) error {
	return RDB.Set(Ctx, key, value, expiration).Err()
}

// Get 获取 key 对应的值，返回字符串和错误
func Get(key string) (string, error) {
	return RDB.Get(Ctx, key).Result()
}

// Del 删除 key
func Del(keys ...string) (int64, error) {
	return RDB.Del(Ctx, keys...).Result()
}

// Exists 判断 key 是否存在，存在返回 true
func Exists(key string) (bool, error) {
	cnt, err := RDB.Exists(Ctx, key).Result()
	return cnt > 0, err
}

// Expire 设置 key 过期时间，expiration 单位秒
func Expire(key string, expiration time.Duration) (bool, error) {
	return RDB.Expire(Ctx, key, expiration).Result()
}

func BlacklistAdd(jti string, ttl int64) {
	RDB.Set(Ctx, "blacklist:"+jti, 1, time.Duration(ttl))
}

func BlacklistCheck(jti string) bool {
	exists, _ := RDB.Exists(Ctx, "blacklist:"+jti).Result()
	return exists > 0
}
