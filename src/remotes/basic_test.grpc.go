package remotes

import (
	"context"
	"microservice/proto"
	"microservice/src/utils/remote"
)

func BasicAdd(a uint64, b uint64) (int64, error) {
	conn := remote.ConnectTo("BASIC_SERVICE")
	client := proto.NewBasicServiceClient(conn)

	var response int64
	var err_msg error
	req := &proto.RequestBasic{A: int64(a), B: int64(b)}
	if res, err := client.Add(context.Background(), req); err == nil {
		response = res.Result
	} else {
		err_msg = err
	}
	return response, err_msg
}

func BasicMultiply(a uint64, b uint64) (int64, error) {
	conn := remote.ConnectTo("BASIC_SERVICE")
	client := proto.NewBasicServiceClient(conn)

	var response int64
	var err_msg error
	req := &proto.RequestBasic{A: int64(a), B: int64(b)}
	if res, err := client.Multiply(context.Background(), req); err == nil {
		response = res.Result
	} else {
		err_msg = err
	}
	return response, err_msg
}
