package configs

import "os"

func GrpcHost() string {
	return os.Getenv("GRPC_HOST")
}
func GrpcPort() string {
	return os.Getenv("GRPC_PORT")
}
