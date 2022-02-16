package file

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	fileObj, err := os.Open("./hello.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	fmt.Println(fileObj)
	tempSlice := make([]byte, 4, 4) // 一次读取4个字节
	var contentSlice []byte
	for {
		_, err := fileObj.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		contentSlice = append(contentSlice, tempSlice...)
	}
	fmt.Println(string(contentSlice))
}
