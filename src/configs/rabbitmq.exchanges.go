package configs

import "os"

func ExampleExchange() string {
	return os.Getenv("EXAMPLE_EXCHANGE")
}
