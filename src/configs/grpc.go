package configs

import "os"

func GetPortGrpcServer() string {
	return os.Getenv("GRPC_PORT")
}
