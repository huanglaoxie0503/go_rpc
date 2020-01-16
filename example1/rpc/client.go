package main

import (
	"fmt"
	"go_rpc/example1/rpc/utils"
	"log"
	"net/rpc"
)

func main() {
	var serverAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress + ":8080")
	if err != nil {
		log.Fatal("建立与服务端连接失败:", err)
	}

	// 同步调用 RPC
	args := &utils.Args{10, 10}
	var reply int
	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)

	// 通过 Go 方法以异步方式进行 RPC 调用
	divideCall := client.Go("MathService.Divide", args, &reply, nil)
	for {
		select {
		case <- divideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, reply)
		}
	}
}
