package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rpc/grpc_token_auth_test/proto"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "123456",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	//interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	fmt.Println("接收到了一个新的请求")
	//	md := metadata.New(map[string]string{
	//		"appid":  "101010",
	//		"appkey": "123456",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Println("请求耗时：", time.Since(start))
	//	return err
	//}

	grpc.WithPerRPCCredentials(customCredential{})
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreaterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "星哥哥"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
