package apps

import (
	"context"
	"fmt"
	"log"
	"main-service/proto"
	"main-service/src/configs"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcHooks struct {
	proto.UnimplementedMainServiceServer // server gRPC sekarang
}

func GrpcServer() {

	go func() {
		time.Sleep(22)
		grpc_port := configs.GrpcPort()

		// Run gRPC
		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpc_port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		svc := grpc.NewServer()
		proto.RegisterMainServiceServer(svc, &GrpcHooks{}) // register this gRPC with real name service
		reflection.Register(svc)

		fmt.Printf("gRPC Started at http://%v:%v ...\n", configs.GrpcHost(), grpc_port)
		if err := svc.Serve(lis); err != nil {
			log.Fatalf("failed to start gRPC serve: %v", err)
		}
	}()

}

// ===========================================================
//-> Titipan Function

func (s *GrpcHooks) Detail(_ context.Context, request *proto.ResponseMain) (*proto.ResponseMain, error) {
	// your logic from gRPC request
	return &proto.ResponseMain{
		Id:          request.GetId(),
		Name:        "Example Product",
		Price:       1922,
		Description: "manusia tidak akan bisa di hancurkan selama dia masih setia pada hatinya...",
	}, nil
}
