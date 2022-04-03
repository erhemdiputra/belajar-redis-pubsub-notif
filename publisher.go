package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

const (
	testKey  = "test_key_123"
	dummyKey = "dummy_key_123"
	dummyVal = "true"
	duration = 10
)

func initPublisher() error {
	conn, err := initRedis(address)
	if err != nil {
		log.Printf("initRedis got error: %v - initSubscriber", err)
		return err
	}

	_, err = setNX(conn, testKey, dummyVal)
	if err != nil {
		log.Printf("setNX testKey got error: %v - initSubscriber", err)
		return err
	}

	_, err = expire(conn, testKey, duration)
	if err != nil {
		log.Printf("expire testKey got error: %v - initSubscriber", err)
		return err
	}

	_, err = setNX(conn, dummyKey, dummyVal)
	if err != nil {
		log.Printf("setNX dummyKey got error: %v - initSubscriber", err)
		return err
	}

	_, err = expire(conn, dummyKey, duration)
	if err != nil {
		log.Printf("expire dummyKey got error: %v", err)
		return err
	}

	return nil
}

func setNX(conn redis.Conn, key string, value string) (interface{}, error) {
	return conn.Do("SETNX", key, value)
}

func expire(conn redis.Conn, key string, duration int) (interface{}, error) {
	return conn.Do("EXPIRE", key, duration)
}
