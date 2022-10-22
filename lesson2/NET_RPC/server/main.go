package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

// ServiceA 自定义一个结构体类型
type ServiceA struct{}

// Add 为ServiceA类型增加的一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	rpc.HandleHTTP()	// 基于HTTP协议
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}