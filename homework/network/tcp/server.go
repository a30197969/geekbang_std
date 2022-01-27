package main

import (
	"bufio"
	"fmt"
	"geekbang_study/homework/network/proto"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			break
		}
		fmt.Println("收到cient端发来的数据：", msg)
		conn.Write([]byte(msg)) // 发送数据
	}

}
