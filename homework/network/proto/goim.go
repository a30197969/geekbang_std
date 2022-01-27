package proto

import (
	"encoding/binary"
	"fmt"
)

// Package Length，包长度 4bytes
// Header Length，头长度	2bytes
// Protocol Version，协议版本 2bytes
// Operation，操作码 4bytes
// Sequence 请求序号 ID 4bytes
// Body，包内容

// GDecode goim 协议的解码器
func GDecode(data []byte) string {
	if len(data) < 16 {
		return ""
	}
	packetLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("packetLen:%v\n", packetLen)
	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("version:%v\n", version)

	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("operation:%v\n", operation)

	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("sequence:%v\n", sequence)

	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
	return body
}

// GEncode goim 协议的封包
func GEncode(body string) []byte {
	headerLen := 16
	packetLen := len(body) + headerLen
	ret := make([]byte, packetLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(packetLen))

	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))

	version := 5
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))

	operation := 6
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation))

	sequence := 7
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[16:], byteBody)
	return ret
}
