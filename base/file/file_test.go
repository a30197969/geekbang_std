package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// file 方式
func TestName(t *testing.T) {
	fileObj, err := os.Open("./hello.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	fmt.Println(fileObj)
	tempSlice := make([]byte, 128, 128) // 一次读取4个字节
	var contentSlice []byte
	for {
		n, err := fileObj.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		contentSlice = append(contentSlice, tempSlice[:n]...)
	}
	fmt.Println(string(contentSlice))
}

// bufio 方式
func TestName2(t *testing.T) {
	fileObj, err := os.Open("./hello.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	fmt.Println(fileObj)
	readerObj := bufio.NewReader(fileObj)
	var content string
	for {
		tempStr, err := readerObj.ReadString('\n')
		if err == io.EOF {
			content += tempStr
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		content += tempStr
	}
	fmt.Println(content)
}

// ioutil 方式，适用于小文件
func TestName3(t *testing.T) {
	contentSlice, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	fmt.Println(string(contentSlice))
}
