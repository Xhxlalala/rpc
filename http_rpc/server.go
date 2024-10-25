package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值实现的
	*reply = "hello, " + request
	return nil
}

func main() {
	// 1.注册服务
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}
		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	_ = http.ListenAndServe(":1234", nil)

}
