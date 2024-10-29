package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"rpc/grpc_error_handle_test/proto"
	"time"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	//return nil, status.Errorf(codes.NotFound, "not found:%s", in.Name)
	time.Sleep(time.Second * 5)
	return &proto.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		return
	}
}
