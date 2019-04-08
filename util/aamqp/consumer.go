package aamqp

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/manucorporat/try"
	"github.com/streadway/amqp"
)

const RECOVER_INTERVAL_TIME = 6 * 60

type Consumer struct {
	cc      ConnectionConfig
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error

	Tag      string
	uri      string
	exchange Exchange

	lastRecoverTime int64
	currentStatus   atomic.Value
}

func NewConsumer(tag, uri string, ex Exchange, ccs ...ConnectionConfig) *Consumer {
	name, err := os.Hostname()
	if err != nil {
		name = ex.Name + ex.Kind
	}
	var cc ConnectionConfig
	if len(ccs) > 0 {
		cc = ccs[0]
	}

	consumer := &Consumer{
		cc:              cc.withDefault(),
		Tag:             fmt.Sprintf("%s-%s", tag, name),
		uri:             uri,
		exchange:        ex,
		done:            make(chan error),
		lastRecoverTime: time.Now().Unix(),
	}
	consumer.currentStatus.Store(true)
	return consumer
}

func (c *Consumer) Connect() (err error) {
	ac := amqp.Config{
		Properties: amqp.Table{
			"product": defaultProduct,
			"version": defaultVersion,
		},
		Heartbeat: c.cc.Heartbeat,
		Locale:    c.cc.Locale,
		Dial:      amqp.DefaultDial(c.cc.Timeout),
	}
	c.conn, err = amqp.DialConfig(c.uri, ac)

	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}

	go func() {
		log.Println("closing: ", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		c.done <- errors.New("Channel Closed")
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	err = c.channel.ExchangeDeclare(
		c.exchange.Name,
		c.exchange.Kind,
		c.exchange.Durable,
		c.exchange.AutoDelete,
		c.exchange.Internal,
		c.exchange.NoWait,
		c.exchange.Args,
	)
	if err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	return
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
		c.channel = nil
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}

func (c *Consumer) Reconnect(cp ConsumeParams, que Queue, qos *BasicQos, bindings ...QueueBinding) (<-chan amqp.Delivery, error) {
	c.Close()

	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.ConsumeQueue(cp, que, qos, bindings...)

}

// ConsumeQueue
// 同一个连接，可以进行多个 ConsumeQueue。对于相同订阅的， 同一条消息，只有其中一个 ConsumeQueue 可以接收到消息。
func (c *Consumer) ConsumeQueue(cp ConsumeParams, que Queue, qos *BasicQos, bindings ...QueueBinding) (<-chan amqp.Delivery, error) {
	if c.channel == nil {
		return nil, fmt.Errorf("no connected channel")
	}
	queue, err := c.channel.QueueDeclare(
		que.Name,
		que.Durable,
		que.AutoDelete,
		que.Exclusive,
		que.NoWait,
		que.Args,
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	if qos != nil {
		if err = c.channel.Qos(qos.PrefetchCount, qos.PrefetchSize, qos.Global); err != nil {
			return nil, fmt.Errorf("Qos Setting: %s", err)
		}
	}

	for _, bind := range bindings {
		if err = c.channel.QueueBind(queue.Name, bind.Key, bind.Exchange, bind.NoWait, bind.Args); err != nil {
			return nil, fmt.Errorf("Queue Bind: %s", err)
		}
	}

	return c.channel.Consume(
		queue.Name,
		c.Tag,
		false, // cp.AutoAck,
		cp.Exclusive,
		cp.NoLocal,
		cp.NoWait,
		cp.Args,
	)
}

func (c *Consumer) Handle(deliveries <-chan amqp.Delivery, fn func([]byte) bool, threads int, cp ConsumeParams, que Queue, qos *BasicQos, bindings ...QueueBinding) {

	var err error
	for {

		// 当 deliveries 置空后，这些协程将会被全部自动回收
		for i := 0; i < threads; i++ {
			go func() {
				for msg := range deliveries {
					ret := false
					try.This(func() {
						body := msg.Body[:]
						ret = fn(body)
					}).Finally(func() {
						if ret == true {
							msg.Ack(false)
							currentTime := time.Now().Unix()
							if currentTime-c.lastRecoverTime > RECOVER_INTERVAL_TIME && !c.currentStatus.Load().(bool) {
								c.currentStatus.Store(true)
								c.lastRecoverTime = currentTime
								c.channel.Recover(true)
							}
						} else {
							// this really a litter dangerous. if the worker is panic very quickly,
							// it will ddos our sentry server......plz, add [retry-ttl] in header.
							msg.Nack(false, true)
							c.currentStatus.Store(false)
						}
					}).Catch(func(e try.E) {
						log.Printf("delivery failed: %s\n", e)
					})
				}
				log.Println("Consumer coroutine destroyed")
			}()
		}

		runtime.Gosched()

		// Go into reconnect loop when c.done is passed non nil values
		if <-c.done != nil {
			c.currentStatus.Store(false)
			retryTime := 1
			for {
				log.Printf("reconnecting %dth time\n", retryTime)
				deliveries, err = c.Reconnect(cp, que, qos, bindings...)
				if err != nil {
					log.Printf("reconnecting failed: %s\n", err)
					retryTime++
				} else {
					log.Println("reconnecting success")
					break
				}
				time.Sleep(time.Duration(15+rand.Intn(60)+2*retryTime) * time.Second)
			}
		}

		time.Sleep(time.Second)
		log.Println("aamqp consumer re-handling...")
	}
}
