package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func initSubscriber() error {
	conn, err := initRedis(address)
	if err != nil {
		log.Printf("initRedis got error: %v - initSubscriber", err)
		return err
	}

	subscriber := Subscriber{
		client: redis.PubSubConn{
			Conn: conn,
		},
		cbMap: map[string]SubscribeCallback{
			patternTestKey:  testCallback,
			patternDummyKey: testCallback,
		},
	}
	listChannel := []string{patternTestKey, patternDummyKey}

	for _, channel := range listChannel {
		err = subscriber.client.PSubscribe(channel)
		if err != nil {
			log.Printf("c.client.PSubscribe() got error: %v - InitSubscriber", err)
			return err
		}
	}

	go func() {
		for {
			switch res := subscriber.client.Receive().(type) {
			case redis.Message:
				subscriber.cbMap[res.Pattern](res.Pattern, res.Channel, string(res.Data))
			case redis.Subscription:
				log.Printf("c.client.Receive() got redis.Subscription:"+
					" channel=%v, kind=%v, count=%v - InitSubscriber",
					res.Channel, res.Kind, res.Count)
			case error:
				log.Printf("c.client.Receive() got error: %v - InitSubscriber", res)
				return
			}
		}
	}()

	return nil
}

func testCallback(pattern string, channel string, message string) {
	log.Printf("testCallback: pattern=%s, channel=%s, message=%s",
		pattern, channel, message)
}
