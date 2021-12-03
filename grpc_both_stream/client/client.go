package main

import (
	"context"
	"geekbang_study/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

const (
	Address string = ":8000"
)

var streamClient proto.StreamClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v\n", err)
	}
	defer conn.Close()
	streamClient = proto.NewStreamClient(conn)
	conversations()
}

// conversations 调用服务端的Conversations方法
func conversations() {
	// 调用服务端的Conversations方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v\n", err)
	}
	for i := 0; i < 10; i++ {
		err = stream.Send(&proto.StreamRequest{
			Question: "stream client rpc " + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v\n", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	// 最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v\n", err)
	}
}
