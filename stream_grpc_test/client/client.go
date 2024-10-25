package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"

	"rpc/stream_grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "星星星"})
	for {
		msg, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(msg.Data)
	}

	// 客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{Data: fmt.Sprintf("客户端流模式 %v", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "小星星"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
