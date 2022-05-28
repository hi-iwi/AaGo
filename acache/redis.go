package acache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hi-iwi/AaGo/ae"
	"log"
	"strconv"
	"time"
)

// final 不包括
func Key(n int, final int, prev bool, f func(int) string) string {
	if prev {
		n--
	}
	if n < 0 || n >= final {
		n = final - 1
	}
	return f(n)
}
func HourlyKey(prev bool, f func(int) string) string {
	return Key(time.Now().Hour(), 24, prev, f)
}
func WeeklyKey(prev bool, f func(int) string) string {
	day := time.Now().Day()
	x := day / 7 // 0 1 2 3
	return Key(x, 4, prev, f)
}

// 使当前时段的放到最后
func BatchKeys(n int, final int, ignoreCurrent bool, f func(int) string) []string {
	l := final
	if ignoreCurrent {
		l--
	}
	ks := make([]string, l)
	j := 0
	// 当前时段一定要在最后
	if n+1 < final {
		for i := n + 1; i < final; i++ {
			ks[j] = f(i)
			j++
		}
	}
	for i := 0; i < n; i++ {
		ks[j] = f(i)
		j++
	}
	if ignoreCurrent {
		return ks
	}

	ks[j] = f(n)

	return ks
}
func HourlyKeys(ignoreCurrent bool, f func(int) string) []string {
	return BatchKeys(time.Now().Hour(), 24, ignoreCurrent, f)
}
func WeeklyKeys(ignoreCurrent bool, f func(int) string) []string {
	day := time.Now().Day()
	x := day / 7 // 0 1 2 3
	return BatchKeys(x, 4, ignoreCurrent, f)
}

// 一般用于 SUnion
func Uint64s(vs []string, err error) ([]uint64, *ae.Error) {
	if err != nil {
		return nil, ae.NewRedisError(err)
	}
	if len(vs) == 0 {
		return nil, ae.NotFound
	}
	ids := make([]uint64, len(vs))
	for i, v := range vs {
		ids[i], err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Printf("redis uint64s %s is not uint64\n", v)
		}
	}
	return ids, nil
}

func HSet(ctx context.Context, rdb *redis.Client, expires time.Duration, k string, values ...interface{}) *ae.Error {
	var err error
	ttl, _ := rdb.TTL(ctx, k).Result()
	if ttl < expires {
		_, err = rdb.HSet(ctx, k, values...).Result()
	} else {
		_, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			err1 := pipe.HSet(ctx, k, values...).Err()
			err2 := pipe.Expire(ctx, k, expires).Err()
			return ae.CatchError(err1, err2)
		})
	}
	return ae.NewRedisError(err)
}

func HMSet(ctx context.Context, rdb *redis.Client, expires time.Duration, k string, values ...interface{}) *ae.Error {
	var err error
	ttl, _ := rdb.TTL(ctx, k).Result()
	if ttl < expires {
		_, err = rdb.HMSet(ctx, k, values...).Result()
	} else {
		_, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			err1 := pipe.HMSet(ctx, k, values...).Err()
			err2 := pipe.Expire(ctx, k, expires).Err()
			return ae.CatchError(err1, err2)
		})
	}
	return ae.NewRedisError(err)
}

func HIncrBy(ctx context.Context, rdb *redis.Client, expires time.Duration, k string, field string, incr int64) (int64, *ae.Error) {
	var reply int64
	var err error
	ttl, _ := rdb.TTL(ctx, k).Result()
	if ttl < expires {
		reply, err = rdb.HIncrBy(ctx, k, field, incr).Result()
	} else {
		_, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			var err1 error
			reply, err1 = pipe.HIncrBy(ctx, k, field, incr).Result()
			err2 := pipe.Expire(ctx, k, expires).Err()
			return ae.CatchError(err1, err2)
		})
	}
	return reply, ae.NewRedisError(err)
}

func HIncr(ctx context.Context, rdb *redis.Client, ttl time.Duration, k string, field string) (int64, *ae.Error) {
	return HIncrBy(ctx, rdb, ttl, k, field, 1)
}
