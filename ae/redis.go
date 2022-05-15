package ae

import "github.com/go-redis/redis/v8"

func NewRedisError(err error) *Error {
	switch err {
	case redis.Nil:
		return NotFound
	case nil:
		return nil
	default:
		pos := Caller(1)
		return NewError(500, pos+" redis: "+err.Error())
	}
}

func Redis(cmd *redis.StatusCmd) *Error {
	return NewRedisError(cmd.Err())
}
