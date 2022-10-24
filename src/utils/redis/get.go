package redis

import (
	"microservice/src/apps"
)

func Get(database int, key string) (string, error) {
	client := apps.RedisConnect(database)
	return client.Get(key).Result()
}
