package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerHandler struct {

}

func (s *ServerHandler) GetName(id int, item *Item) error {
	log.Printf("receive GetName call, id:%d", id)
	item.Id = id
	item.Name = "Oscar"
	return nil
}

func (s *ServerHandler) SaveName(item Item, r RpcResponse) error  {
	log.Printf("receive SaveName call, item: %v", item)
	r.Ok = true
	r.Id = item.Id
	r.Message = "存储成功！"
	return nil
}

func main() {
	// 初始化RPC服务端
	server := rpc.NewServer()

	// 监听端口 8080
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("监听端口失败：%v", err)
	}
	defer listener.Close()

	log.Println("Start listener on port localhost:8083")

	// 初始化服务器
	serviceHandler := &ServerHandler{}
	// 注册处理器
	err = server.Register(serviceHandler)
	if err != nil {
		log.Fatalf("注册服务处理器失败:%v", err)
	}

	// 等待并处理客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("接收客户端连接请求失败:%v", err)
		}
		// 自定义RPC 处理器：新建一个 jsonrpc 编码器给 RPC 服务端处理器 指定编解码器
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
