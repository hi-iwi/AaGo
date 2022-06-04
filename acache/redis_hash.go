package acache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
)

func HScan(ctx context.Context, rdb *redis.Client, dest interface{}, k string, fields ...string) *ae.Error {
	c := rdb.HMGet(ctx, k, fields...)
	v, err := c.Result()
	if err != nil {
		return ae.NewRedisError(err)
	}
	if len(v) != len(fields) {
		return ae.NotFound
	}
	for _, x := range v {
		if atype.IsNil(x) {
			return ae.NotFound
		}
	}
	e := ae.NewRedisError(c.Scan(dest))
	return e
}

func HGetAll(ctx context.Context, rdb *redis.Client, k string, dest interface{}) *ae.Error {
	c := rdb.HGetAll(ctx, k)
	result, err := c.Result()
	if err != nil {
		return ae.NewRedisError(err)
	}
	if len(result) == 0 {
		return ae.NotFound
	}
	e := ae.NewRedisError(c.Scan(dest))
	return e
}

func HGetAllInt(ctx context.Context, rdb *redis.Client, k string, restrict bool) (map[string]int, *ae.Error) {
	c := rdb.HGetAll(ctx, k)
	result, err := c.Result()
	if err != nil {
		return nil, ae.NewRedisError(err)
	}
	n := len(result)
	if n == 0 {
		return nil, ae.NotFound
	}
	value := make(map[string]int, n)
	var x int64
	for k, v := range result {
		if x, err = strconv.ParseInt(v, 10, 32); err != nil {
			if restrict {
				return nil, ae.New(err)
			}
			continue
		}
		value[k] = int(x)
	}
	return value, nil
}

// 只要存在一个，就不报错；全是nil，返回 ae.NotFound
func TryHMGet(ctx context.Context, rdb *redis.Client, k string, fields ...string) ([]interface{}, bool, *ae.Error) {
	v, err := rdb.HMGet(ctx, k, fields...).Result()
	if err != nil {
		return nil, false, ae.NewRedisError(err)
	}
	ok := true
	e := ae.NotFound
	for _, x := range v {
		if !atype.IsNil(x) {
			e = nil // 只要存在一个不是nil，都正确
			if !ok {
				break
			}
		} else {
			ok = false
			if e == nil {
				break
			}
		}
	}
	return v, ok, e
}
func TryHMGetUint64(ctx context.Context, rdb *redis.Client, k string, fields []string, defaultValue uint64) ([]uint64, bool, *ae.Error) {
	iv, ok, e := TryHMGet(ctx, rdb, k, fields...)
	if e != nil {
		return nil, ok, e
	}
	v := make([]uint64, len(fields))
	for i, x := range iv {
		if atype.IsNil(x) {
			v[i] = defaultValue
		} else {
			v[i] = atype.New(x).DefaultUint64(0)
		}
	}
	return v, ok, nil
}
func TryHMGetInt64(ctx context.Context, rdb *redis.Client, k string, fields []string, defaultValue int64) ([]int64, bool, *ae.Error) {
	iv, ok, e := TryHMGet(ctx, rdb, k, fields...)
	if e != nil {
		return nil, false, e
	}
	v := make([]int64, len(fields))
	for i, x := range iv {
		if atype.IsNil(x) {
			v[i] = defaultValue
		} else {
			v[i] = atype.New(x).DefaultInt64(0)
		}
	}
	return v, ok, nil
}
func TryHMGetInt(ctx context.Context, rdb *redis.Client, k string, fields []string, defaultValue int) ([]int, bool, *ae.Error) {
	iv, ok, e := TryHMGet(ctx, rdb, k, fields...)
	if e != nil {
		return nil, false, e
	}
	v := make([]int, len(fields))
	for i, x := range iv {
		if atype.IsNil(x) {
			v[i] = defaultValue
		} else {
			v[i] = atype.New(x).DefaultInt(0)
		}
	}
	return v, ok, nil
}

// 不能有一个是nil
func MustHMGet(ctx context.Context, rdb *redis.Client, k string, fields ...string) ([]interface{}, *ae.Error) {
	v, err := rdb.HMGet(ctx, k, fields...).Result()
	if err != nil {
		return nil, ae.NewRedisError(err)
	}
	if len(v) != len(fields) {
		return nil, ae.NotFound
	}
	for _, x := range v {
		if atype.IsNil(x) {
			return v, ae.NotFound
		}
	}
	return v, nil
}
func MustHMGetUint64(ctx context.Context, rdb *redis.Client, k string, fields ...string) ([]uint64, *ae.Error) {
	iv, e := MustHMGet(ctx, rdb, k, fields...)
	if e != nil {
		return nil, e
	}
	v := make([]uint64, len(fields))
	for i, a := range iv {
		v[i] = atype.New(a).DefaultUint64(0)
	}
	return v, nil
}
func MustHMGetInt64(ctx context.Context, rdb *redis.Client, k string, fields ...string) ([]int64, *ae.Error) {
	iv, e := MustHMGet(ctx, rdb, k, fields...)
	if e != nil {
		return nil, e
	}
	v := make([]int64, len(fields))
	for i, a := range iv {
		v[i] = atype.New(a).DefaultInt64(0)
	}
	return v, nil
}
