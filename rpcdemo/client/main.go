package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"time"
)

type ResponseData struct {
	Data int `json:"result"`
}

// RPC远程过程调用
func add(a, b int) int {
	req := HttpRequest.NewRequest()
	// Set the timeout to 5 seconds
	req.SetTimeout(5 * time.Second)
	res, err := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))

	if err != nil {
		fmt.Println("Error making request:", err)
		return 0
	}
	body, err := res.Body()
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}
	fmt.Println(string(body))
	rspData := ResponseData{}
	err = json.Unmarshal(body, &rspData)
	if err != nil {
		return 0
	}
	return rspData.Data
}

func main() {
	fmt.Println(add(1, 2))
}
