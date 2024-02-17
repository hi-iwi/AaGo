package ae

import "github.com/go-redis/redis/v8"

func NewRedisError(err error) *Error {
	switch err {
	case redis.Nil:
		return NotFound
	case nil:
		return nil
	default:
		msg, pos := CallerMsg(err.Error(), 1)
		return NewError(500, pos+" redis: "+msg)
	}
}
