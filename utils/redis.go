package utils

import (
	"starzeng.com/gin-demo/pkg/redis"
	"time"
)

func Set(s string) {
	redis.RDB.Set(redis.Ctx, "s", s, 0)
}

func BlacklistAdd(jti string, ttl int64) {
	redis.RDB.Set(redis.Ctx, "blacklist:"+jti, 1, time.Duration(ttl))
}

func BlacklistCheck(jti string) bool {
	exists, _ := redis.RDB.Exists(redis.Ctx, "blacklist:"+jti).Result()
	return exists > 0
}
