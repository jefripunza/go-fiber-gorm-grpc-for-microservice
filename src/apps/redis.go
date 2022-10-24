package apps

import (
	"fmt"
	"microservice/src/configs"
	"time"

	"github.com/go-redis/redis"
)

func RedisConnect(database int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.RedisURL(),
		Password: configs.RedisPassword(),
		DB:       database,
	})
}

func RedisTryConnection() {
	go func() {
		time.Sleep(22)
		client := RedisConnect(0)
		pong, err := client.Ping().Result()
		if err != nil {
			panic("Redis not connect !!")
		}
		fmt.Println("Redis is connect,", pong, "!!!")
	}()
}
