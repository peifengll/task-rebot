package goredis

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var rdb *redis.Client
var Once sync.Once

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "118.25.23.154:8089",
		Password: "123456", // 没有密码，默认值
		DB:       0,        // 默认DB 0
	})

}

func GetRedisClient() *redis.Client {
	Once.Do(func() {
		Init()
	})
	return rdb
}
