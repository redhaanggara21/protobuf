package main

import (
	"context"

	pb "gitlab.com/redha/gopotobuf/proto/pb/greet"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
