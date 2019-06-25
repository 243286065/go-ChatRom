package redis

import (
	"fmt"
	cfg "go-ChatRom/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			// 1. 打开连接
			c, err := redis.Dial("tcp", cfg.RedisHost)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// 2. 访问认证
			if _, err = c.Do("AUTH", cfg.RedisPass); err != nil {
				fmt.Println(err)
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
	data, err := pool.Get().Do("KEYS", "*")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func RedisPool() *redis.Pool {
	return pool
}
