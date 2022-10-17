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
	grpc_port := configs.GetPortGrpcServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	svc := grpc.NewServer()
	proto.RegisterMainServiceServer(svc, &server{})
	reflection.Register(svc)
	fmt.Printf("gRPC Started at http://localhost:%v ...\n", grpc_port)
	if err := svc.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC serve: %v", err)
	}

}

// ===========================================================
//-> Titipan Function

func (s *server) Detail(_ context.Context, request *proto.ResponseMain) (*proto.ResponseMain, error) {
	id := request.GetId()
	return &proto.ResponseMain{Id: id}, nil
}
