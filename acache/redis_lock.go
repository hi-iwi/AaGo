package acache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hi-iwi/AaGo/ae"
	"time"
)

// 申请原子性的锁
func ApplyLock(ctx context.Context, rdb *redis.Client, expires time.Duration, k string) *ae.Error {
	err := rdb.SetNX(ctx, k+":lock", 1, expires).Err()
	return ae.NewRedisError(err)
}
