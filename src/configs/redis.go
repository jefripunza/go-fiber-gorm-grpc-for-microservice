package configs

import "os"

func RedisURL() string {
	return os.Getenv("REDIS_URL")
}

func RedisPassword() string {
	return os.Getenv("REDIS_PASS")
}
