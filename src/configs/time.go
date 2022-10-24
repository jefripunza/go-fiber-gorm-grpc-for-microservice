package configs

import "os"

func GetTimezone() string {
	return os.Getenv("TZ")
}
