package apps

import (
	"context"
	"fmt"
	"log"
	"main-service/proto"
	"main-service/src/configs"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedMainServiceServer // server gRPC sekarang
}

func GrpcServer() {

	// Run gRPC
	grpc_host := configs.GetHostGrpcServer()
	grpc_port := configs.GetPortGrpcServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := grpc.NewServer()
	proto.RegisterMainServiceServer(svc, &server{}) // register this gRPC with real name service
	reflection.Register(svc)

	fmt.Printf("gRPC Started at http://%v:%v ...\n", grpc_host, grpc_port)
	if err := svc.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC serve: %v", err)
	}

}

// ===========================================================
//-> Titipan Function

func (s *server) Detail(_ context.Context, request *proto.ResponseMain) (*proto.ResponseMain, error) {
	// your logic from gRPC request
	return &proto.ResponseMain{
		Id:          request.GetId(),
		Name:        "Example Product",
		Price:       1922,
		Description: "manusia tidak akan bisa di hancurkan selama dia masih setia pada hatinya...",
	}, nil
}
