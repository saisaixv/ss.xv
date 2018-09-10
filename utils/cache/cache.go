package cache

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func newPool(server, password string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedis(server, password string) error {

	if pool != nil {
		return fmt.Errorf("pool is not nil")
	}
	pool = newPool(server, password)
}
