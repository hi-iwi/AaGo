package queue_test

import (
	"fmt"
	"log"
	"service/aa-oas/conf"
	"testing"
	"time"

	"github.com/hi-iwi/AaGo/queue"
)

func TestRedisQueue(t *testing.T) {
	q := queue.NewRedisQueue(conf.App.MainMasterRedis)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			if err := q.Pub("test", fmt.Sprintf("test go pub %d", i), queue.Qos0); err != nil {
				t.Error(err)
				return
			}
			if err := q.Pub("test2", fmt.Sprintf("test go pub %d", i), queue.Qos0); err != nil {
				t.Error(err)
				return
			}
		}
	}()
	go testRedisQueueSub(t)
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
	}
	t.Log("test redis queue finished")
}
func testRedisQueueSub(t *testing.T) {
	q := queue.NewRedisQueue(conf.App.MainMasterRedis)

	go q.Sub("test", func(payload string) {
		log.Println(payload)
		t.Log(payload)
	}, time.Duration(10))
	go q.Sub("test2", func(payload string) {
		log.Println(payload)
		t.Log(payload)
	}, time.Duration(10))
	t.Log("test sub finished")
}
