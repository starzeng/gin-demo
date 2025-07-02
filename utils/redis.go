package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	Ctx = context.Background()
)

func Set(s string) {
	Rdb.Set(Ctx, "s", s, 0)
}

func BlacklistAdd(jti string, ttl int64) {
	Rdb.Set(Ctx, "blacklist:"+jti, 1, time.Duration(ttl))
}

func BlacklistCheck(jti string) bool {
	exists, _ := Rdb.Exists(Ctx, "blacklist:"+jti).Result()
	return exists > 0
}
