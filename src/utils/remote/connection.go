package remote

import (
	"fmt"
	"microservice/src/configs"
	"os"

	"google.golang.org/grpc"
)

func ConnectTo(service_name string) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", configs.GrpcHost(), os.Getenv(service_name)), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}
