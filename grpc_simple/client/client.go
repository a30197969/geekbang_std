package main

import (
	"geekbang_study/proto"
	"google.golang.org/grpc"
	"log"
)

import (
	"context"
)

const (
	Address string = ":8080"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v\n", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := proto.NewSimpleClient(conn)
	// 创建发送结构体
	req := proto.SimpleRequest{
		Data: "grpc",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v\n", err)
	}
	// 打印返回值
	log.Println(res)
}
