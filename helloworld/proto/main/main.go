package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"rpc/helloworld/proto"
)

type Hello struct {
	Name string `json:"name"`
}

func main() {
	req := pb.HelloRequest{
		Name: "bobby",
		Age:  18,
		Courses: []string{
			"math",
			"english",
		},
	}
	//jsonStruct := Hello{Name: "bobby"}
	//jsonRsp, _ := json.Marshal(jsonStruct)
	//fmt.Println(string(jsonRsp))
	rsp, _ := proto.Marshal(&req)
	newReq := pb.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(string(rsp))
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
