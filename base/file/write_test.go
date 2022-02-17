package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// file 写入
func TestName11(t *testing.T) {
	fileObj, err := os.OpenFile("./hello.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	fmt.Println(fileObj)
	n, err := fileObj.WriteString("蜂鸟网123131313131311\r\n")
	n2, err := fileObj.Write([]byte("大家都挺3222好吧\r\n"))
	fmt.Println(n, n2)
}

// bufio 写入
func TestName22(t *testing.T) {
	fileObj, err := os.OpenFile("./hello.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	fmt.Println(fileObj)
	writer := bufio.NewWriter(fileObj)
	n, err := writer.WriteString("真TM棒棒的\r\n")
	writer.Flush()
	fmt.Println(n)
}

// ioutil 只可以覆盖模式写入
func TestName33(t *testing.T) {
	err := ioutil.WriteFile("./hello.txt", []byte("ni ye 不错\r\n"), 0644)
	fmt.Println(err)
}
