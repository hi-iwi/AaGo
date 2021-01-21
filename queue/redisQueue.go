package queue

import (
	"errors"
	"reflect"
	"time"

	"github.com/go-redis/redis"
	"github.com/hi-iwi/AaGo/cnf"
)

type redisQueue struct {
	redisConf cnf.Redis
}

func NewRedisQueue(cnf cnf.Redis) Queue {
	return &redisQueue{
		redisConf: cnf,
	}
}

func (q *redisQueue) conn() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     q.redisConf.Authority,
		Password: q.redisConf.Auth,
		DB:       0,
	})
	return cli
}
func (q *redisQueue) Pub(channel string, message string, qos Qos) error {
	cli := q.conn()
	defer cli.Close()

	return cli.Publish(channel, message).Err()
}

func (q *redisQueue) Sub(channels interface{}, hanldler func(payload string), timeout time.Duration) error {
	cli := q.conn()
	defer cli.Close()

	var ps *redis.PubSub
	switch reflect.TypeOf(channels).Kind() {
	case reflect.Array:
		ps = cli.Subscribe(channels.([]string)...)
	case reflect.String:
		ps = cli.Subscribe(channels.(string))
	default:
		return errors.New("Invalid sub channels type")
	}

	if timeout != 0 {
		time.AfterFunc(time.Second*timeout, func() {
			ps.Close()
		})
	} else {
		defer ps.Close()
	}
	if _, err := ps.Receive(); err != nil {
		return err
	}
	ch := ps.Channel()

	for msg := range ch {
		hanldler(msg.Payload)
	}
	return nil
}
