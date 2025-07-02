package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	cfg := AppConfig.Redis
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
