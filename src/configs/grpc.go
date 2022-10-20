package configs

import "os"

var (
	GrpcHost = os.Getenv("GRPC_HOST")
	GrpcPort = os.Getenv("GRPC_PORT")
)
