package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gomodule/redigo/redis"
)

const (
	address             = "localhost:6379"
	patternTestKey      = "__keyspace*:test_key_*"
	patternDummyKey     = "__keyspace*:dummy_key_*"
	patternExpiredEvent = "__keyevent*:expired"
)

type SubscribeCallback func(pattern string, channel string, message string)

type Subscriber struct {
	client redis.PubSubConn
	cbMap  map[string]SubscribeCallback
}

func initRedis(address string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		log.Printf("redis.Dial got error: %v - initRedis", err)
		return nil, err
	}

	return conn, nil
}

func main() {
	err := initSubscriber()
	if err != nil {
		log.Printf("initSubscriber got error: %v - main", err)
		return
	}

	err = initPublisher()
	if err != nil {
		log.Printf("initPublisher got error: %v - main", err)
		return
	}

	waitingSignal()
}

func waitingSignal() {
	// waiting signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println("Signal =", sig)
		done <- true
	}()
	<-done
}
