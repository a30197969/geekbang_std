package main

import (
	"fmt"
	"geekbang_study/homework/network/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("client conn failed, err:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "how are you?"
		date, err := proto.Encode(msg)
		_, err = conn.Write(date) // 发送数据
		if err != nil {
			return
		}
		//buf := [512]byte{}
		buf := make([]byte, 512, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("recv failed, err:", err)
		}
		fmt.Println(string(buf[:n]))
	}
}
