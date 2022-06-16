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
