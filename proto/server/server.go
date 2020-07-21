package main

import (
	"fmt"
	"net"

	pb "ankr.com/log-agent/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) BlockHeight(ctx context.Context, in *pb.BlockHeightRequest) (*pb.BlockHeightReply, error) {
	return &pb.BlockHeightReply{Message: in.Name}, nil
}

//不更新报错： (missing SayHelloAgain method)
//func (s *server) SayHelloAgain(ctx context.Context, in *pb.BlockHeightRequest) (*pb.BlockHeightReply, error) {
//	return &pb.BlockHeightReply{Message: "Hello again " + in.GetName()}, nil
//}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterAppinfoServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

