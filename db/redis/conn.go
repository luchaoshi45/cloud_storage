package redis

import (
	"cloud_storage/configurator"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

func init() {
	cephConfig := configurator.GetredisConfig()
	redisHost := cephConfig.GetAttr("Host")
	//redisPass := redisConn.GetAttr("Password")
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
}

func GetRedisPool() *redis.Pool {
	return redisPool
}
