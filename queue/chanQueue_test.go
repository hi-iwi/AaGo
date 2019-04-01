package queue_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/luexu/AaGo/queue"
)

func TestChanQueue(t *testing.T) {
	q := queue.NewChanQueue()

	go func() {
		for i := 0; i < 10000; i++ {
			time.Sleep(time.Second)
			q.Pub("test", fmt.Sprintf("%d", i), queue.Qos0)
			q.Pub("test2", fmt.Sprintf("%d", i), queue.Qos0)
			log.Printf("send: %d", i)
		}
	}()
	go testChanQueueSub(t)

	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
	}
	t.Log("test chan queue done")
}

func testChanQueueSub(t *testing.T) {
	log.Println("subd")
	q := queue.NewChanQueue()
	go q.Sub("test", func(payload string) {
		log.Println(payload)
	}, time.Duration(0))

	go q.Sub("test2", func(payload string) {
		log.Println(payload)
	}, time.Duration(0))
}
