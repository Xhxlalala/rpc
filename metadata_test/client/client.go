package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"rpc/metadata_test/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//发送metadata
	md := metadata.New(map[string]string{
		"name":     "xtar",
		"password": "123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "星哥哥"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
