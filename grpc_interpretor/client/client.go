package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rpc/grpc_test/proto"
	"time"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		fmt.Println("接收到了一个新的请求")
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("请求耗时：", time.Since(start))
		return err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "星哥哥"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
