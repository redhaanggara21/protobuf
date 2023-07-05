package main

import (
	"context"
	"log"
	"net"

	"gitlab.com/redha/gopotobuf/proto/pb/invoicer"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrer := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrer, service)
	err = serverRegistrer.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
