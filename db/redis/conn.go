package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var (
	redisHost = "127.0.0.1:6379"
	redisPass = ""
)

var redisPool *redis.Pool
var once sync.Once

func RedisConn() *redis.Pool {
	once.Do(func() {
		redisPool = &redis.Pool{
			MaxIdle:     50,
			MaxActive:   30,
			IdleTimeout: 300 * time.Second,
			Dial: func() (redis.Conn, error) {
				// 1 打开连接
				c, err := redis.Dial("tcp", redisHost)
				if err != nil {
					fmt.Print(err)
					return nil, err
				}

				// 2 访问认证
				//if _, err = c.Do("AUTH", redisPass); err != nil {
				//	c.Close()
				//	return nil, err
				//}
				return c, nil
			},

			TestOnBorrow: func(c redis.Conn, lastUsed time.Time) error {
				if time.Since(lastUsed) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	})
	return redisPool
}
