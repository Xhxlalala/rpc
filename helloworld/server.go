package main

import (
	"net"
	"net/rpc"

	"rpc/helloworld/handler"
	"rpc/helloworld/server_proxy"
)

func main() {
	// 1.注册服务
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	// 2.监听端口
	listener, _ := net.Listen("tcp", ":1234")
	// 3.启动服务
	// 使用协程异步处理
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
		//将序列化协议改为使用json
		//go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
