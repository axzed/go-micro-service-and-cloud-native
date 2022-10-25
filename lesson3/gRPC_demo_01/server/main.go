package main

import (
	"context"
	"fmt"
	"net"
	"server/pb"

	"google.golang.org/grpc"
)

// grpc server
type server struct {
	// 不需要实现全部proto中的方法
	pb.UnimplementedGreeterServer
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":5050")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	s := grpc.NewServer() //创建gRPC服务
	// 注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve, err:%v\n", err)
		return
	}
}
