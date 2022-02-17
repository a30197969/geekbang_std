package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// 复制文件 音视频文件也可以复制
// 1、 ioutil.ReadAll比ioutil.ReadFile耗时长
// 2、 ioutil.ReadAll比ioutil.ReadFile占用的内存空间大
// 3、 ioutil.ReadFile只能读取有具体地址的文件，而ioutil.ReadAll能读取上传文件的内容
func TestName111(t *testing.T) {
	dirPath, err := os.Getwd() // 获取项目根路径
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	contentSlice, err := ioutil.ReadFile(dirPath + "/hello.txt")
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	fmt.Println(string(contentSlice))
	fmt.Println("文件内容长度", len(contentSlice))
	fmt.Println("文件开辟的内存", cap(contentSlice))
	err = ioutil.WriteFile(dirPath+"/copy.txt", contentSlice, 0666)
	if err != nil {
		fmt.Println("复制文件失败")
		return
	}
}

// file 方式复制
func TestName222(t *testing.T) {
	fileObj, err := os.Open("./hello.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	fileObj2, err := os.OpenFile("./copy2.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	defer fileObj2.Close()
	if err != nil {
		fmt.Println("新建文件失败")
		return
	}
	tempSlice := make([]byte, 128, 128)
	for {
		n, err := fileObj.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		n2, err := fileObj2.Write(tempSlice[:n])
		if err != nil {
			fmt.Println("写入文件失败")
			return
		}
		fmt.Println(n2)
	}
}

// bufio 方式复制
func TestName333(t *testing.T) {
	dirPath, err := os.Getwd() // 获取项目根路径
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	fileObj, err := os.Open(dirPath + "/hello.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	fileObj3, err := os.OpenFile(dirPath+"/copy3.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer fileObj3.Close()
	if err != nil {
		fmt.Println("新建文件失败")
		return
	}
	tempSlice := make([]byte, 128, 128)
	readerObj := bufio.NewReader(fileObj)
	writerObj := bufio.NewWriter(fileObj3)
	for {
		n, err := readerObj.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		writerObj.Write(tempSlice[:n])
	}
	writerObj.Flush()
}

// 创建目录
func TestName444(t *testing.T) {
	dirPath, err := os.Getwd() // 获取项目根路径
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	os.Mkdir(dirPath+"/abc", 0777)
	err = os.MkdirAll(dirPath+"/a/b/c", 0777)
	if err != nil {
		fmt.Println("创建目录失败")
		return
	}
}

// 删除目录
func TestName555(t *testing.T) {
	dirPath, err := os.Getwd() // 获取项目根路径
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	//err = os.Remove(dirPath + "/abc")
	//if err != nil {
	//	fmt.Println("删除目录失败", err.Error())
	//	return
	//}
	err = os.RemoveAll(dirPath + "/a")
	if err != nil {
		fmt.Println("删除目录失败", err.Error())
		return
	}
}

// 重命名
func TestName666(t *testing.T) {
	dirPath, err := os.Getwd() // 获取项目根路径
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	err = os.Rename(dirPath+"/copy3.txt", dirPath+"/copy333.txt")
	if err != nil {
		fmt.Println("重命名失败", err.Error())
		return
	}
}
