package acache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hi-iwi/AaGo/ae"
	"log"
	"strconv"
	"time"
)

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
	if ttl > 0 && ttl < expires {
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
	if ttl > 0 && ttl < expires {
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
	if ttl > 0 && ttl < expires {
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

func HMIncr(ctx context.Context, rdb *redis.Client, expires time.Duration, k string, fields []string) ([]int64, *ae.Error) {
	replies := make([]int64, len(fields))
	var err error
	ttl, _ := rdb.TTL(ctx, k).Result()

	_, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		var err1 error
		for i, field := range fields {
			if replies[i], err1 = pipe.HIncrBy(ctx, k, field, 1).Result(); err1 != nil {
				return err1
			}
		}
		if ttl <= 0 {
			err1 = pipe.Expire(ctx, k, expires).Err()
		}
		return err1
	})

	return replies, ae.NewRedisError(err)
}

func HMIncrIds(ctx context.Context, rdb *redis.Client, expires time.Duration, k string, ids []uint64) ([]int64, *ae.Error) {
	fields := make([]string, len(ids))
	for i, id := range ids {
		fields[i] = strconv.FormatUint(id, 10)
	}
	return HMIncr(ctx, rdb, expires, k, fields)
}

// 使当前时段的放到最后
func BatchKeys(n uint8, final uint8, ignoreCurrent bool, f func(uint8) string) []string {
	l := final
	if ignoreCurrent {
		l--
	}
	ks := make([]string, l)
	var i, j uint8
	// 当前时段一定要在最后
	if n+1 < final {
		for i = n + 1; i < final; i++ {
			ks[j] = f(i)
			j++
		}
	}

	for i = 0; i < n; i++ {
		ks[j] = f(i)
		j++
	}
	if ignoreCurrent {
		return ks
	}

	ks[j] = f(n)

	return ks
}

// final 不包括
func Key(n uint8, final uint8, prev bool, f func(uint8) string) string {
	if prev {
		n--
	}
	if n < 0 || n >= final {
		n = final - 1
	}
	return f(n)
}

func hourIdx(interval uint8) (uint8, uint8) {
	max := uint8(24) // 一天24小时，因数：1，2，3，4，6，8，12，24
	if interval > (max/2) || interval == 0 {
		return 1, 1
	}
	if interval > max {
		interval = max
	}
	final := max / interval
	n := uint8(time.Now().Hour()) % final
	return n, final
}

// @param interval uint8  一天内，每interval小时一张表，最大不能超过24，且被24整除，1，2，3，4，6，8，12，24
func HourlyKey(interval uint8, prev bool, f func(uint8) string) string {
	n, final := hourIdx(interval)
	return Key(n, final, prev, f)
}
func HourlyKeys(interval uint8, ignoreCurrent bool, f func(uint8) string) []string {
	n, final := hourIdx(interval)
	return BatchKeys(n, final, ignoreCurrent, f)
}

func dayIdx(interval uint8) (uint8, uint8) {
	max := uint8(30) // 30 因数：1 2 3 5 6 10 15 30  因数越多，能够拆分的可能性越多，所以选择每个月30天
	if interval > (max/2) || interval == 0 {
		return 1, 1
	}
	if interval > max {
		interval = max
	}
	final := max / interval
	n := uint8(time.Now().Day()) % final
	return n, final
}

// @param interval uint8 一个月内，每interval天一张表
func DailyKey(interval uint8, prev bool, f func(uint8) string) string {
	n, final := dayIdx(interval)
	return Key(n, final, prev, f)
}

func DailyKeys(interval uint8, ignoreCurrent bool, f func(uint8) string) []string {
	n, final := dayIdx(interval)
	return BatchKeys(n, final, ignoreCurrent, f)
}
