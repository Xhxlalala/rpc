package main

import (
	"fmt"
	"rpc/helloworld/client_proxy"
)

func main() {
	//1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	//conn, err := net.Dial("tcp", "localhost:1234")

	//2.调用服务
	var reply string
	//client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err := client.Hello("小星星", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
