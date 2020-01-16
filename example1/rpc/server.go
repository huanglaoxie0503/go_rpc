package main

import (
	"errors"
	"fmt"
	"go_rpc/example1/rpc/utils"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct {

}

func (m *MathService) Multiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args *utils.Args, reply *int) error  {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	// 启动RPC
	math := new(MathService)
	err := rpc.Register(math)
	if err != nil {
		fmt.Println(err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动服务监听失败：", err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动HTTP服务失败：", err)
	}
}
