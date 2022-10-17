package utils

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
)

func GrpcConnectTo(service_name string) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", os.Getenv("GRPC_HOST"), os.Getenv(service_name)), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}
