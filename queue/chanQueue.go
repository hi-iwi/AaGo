package queue

import (
	"errors"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type chanQueue struct {
	chans map[string]chan string
}

var (
	newChanQueueOnce  sync.Once
	chanQueueInstance *chanQueue
	chansCoMaplock    sync.RWMutex
)

func NewChanQueue() Queue {

	newChanQueueOnce.Do(func() {
		chans := make(map[string]chan string, 10)
		chanQueueInstance = &chanQueue{
			chans: chans,
		}
	})

	return chanQueueInstance
}
func (q *chanQueue) Pub(channel string, message string, qos Qos) error {
	chansCoMaplock.RLock()
	c := q.chans[channel]
	chansCoMaplock.RUnlock()
	if cap(c) == 0 {
		chansCoMaplock.Lock()
		q.chans[channel] = make(chan string, 100)
		chansCoMaplock.Unlock()
	}
	chansCoMaplock.Lock()
	q.chans[channel] <- message
	chansCoMaplock.Unlock()
	return nil
}
func (q *chanQueue) Sub(channels interface{}, hanldler func(payload string), timeout time.Duration) error {

	var chls []string
	switch reflect.TypeOf(channels).Kind() {
	case reflect.Array:
		chls = channels.([]string)
	case reflect.String:
		chls = []string{channels.(string)}
	default:
		return errors.New("Invalid sub channels type")
	}

	var payload string
	for {
		chansCoMaplock.RLock()
		chs := q.chans
		chansCoMaplock.RUnlock()

		for ch := range chs {
			for i := 0; i < len(chls); i++ {
				if chls[i] == ch {
					select {
					case payload = <-chs[ch]:
						hanldler("recv: " + payload)
					default:
					}
				}
			}
		}
		runtime.Gosched()
	}
	return nil
}
