package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	RedigoPool *redis.Pool
	expiration = 1 * 60
)

func InitializeLocalRedisConnectionPool(address string) {
	fmt.Println("Initializing local redis connection with redigo")

	RedigoPool = newPool(address)

	fmt.Println("Local redis initialized")
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func SetValue(key string, value []byte) error {
	conn := RedigoPool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, expiration, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func GetValue(key string) ([]byte, error) {
	conn := RedigoPool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}
