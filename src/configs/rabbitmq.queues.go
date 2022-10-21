package configs

import "os"

func ExampleQueue() string {
	return os.Getenv("EXAMPLE_QUEUE")
}
