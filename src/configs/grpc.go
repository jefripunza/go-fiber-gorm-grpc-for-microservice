package configs

import "os"

func GetHostGrpcServer() string {
	return os.Getenv("GRPC_HOST")
}

func GetPortGrpcServer() string {
	return os.Getenv("GRPC_PORT")
}
