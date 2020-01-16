package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	// 30秒超时时间
	conn, err := net.DialTimeout("tcp", "localhost:8083", 30 * time.Second)
	if err != nil {
		log.Fatalf("客户端发起连接失败：%v", err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item Item
	err = client.Call("ServerHandler.GetName", 1, &item)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ServerHandler.GetName 返回结果：%v\n", item)

	var resp RpcResponse
	item = Item{
		Id:   2,
		Name: "Oscar",
	}
	err = client.Call("ServerHandler.SaveName", item, &resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ServerHandler.SaveName 返回结果：%v\n", resp)
}
