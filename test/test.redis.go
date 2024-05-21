package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 创建一个新的Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 没有设置密码
		DB:       0,                // 使用默认DB
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("无法连接到Redis:", err)
		return
	}
	fmt.Println("连接成功:", pong)

	// 设置一个键值对
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("设置键值对失败:", err)
		return
	}

	// 获取键的值
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("获取键值失败:", err)
		return
	}
	fmt.Println("key的值是:", val)

	// 删除键
	err = rdb.Del(ctx, "key").Err()
	if err != nil {
		fmt.Println("删除键失败:", err)
		return
	}
	fmt.Println("键已删除")
}
