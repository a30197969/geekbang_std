package proto

import (
	"fmt"
	"testing"
)

// goim 协议的封包与解码器
func TestGDecode(t *testing.T) {
	data := GEncode("毛老师的课真棒")
	fmt.Println(data)
	GDecode(data)
}
