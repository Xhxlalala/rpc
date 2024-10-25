package client_proxy

import (
	"net/rpc"

	"rpc/helloworld/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

//在go语言中没有类、对象就意味着没有初始化方法

func NewHelloServiceClient(protocol, addr string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, addr)
	if err != nil {
		panic("连接失败")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
