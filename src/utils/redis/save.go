package redis

import (
	"microservice/src/apps"
	"time"
)

func Save(database int, key string, value string, expired time.Duration) error {
	client := apps.RedisConnect(database)
	return client.Set(key, value, expired).Err()
}
