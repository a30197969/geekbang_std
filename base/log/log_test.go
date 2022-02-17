package log

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	dirpath, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败", err.Error())
		return
	}
	fileObj, err := os.OpenFile(dirpath+"/111.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败", err.Error())
		return
	}

	log.SetOutput(fileObj) // 设置输出位置
	//log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)         // 配置日志参数
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate) // 配置日志参数
	log.SetPrefix("[前缀信息]")                                     // 设置日志前缀信息
	log.Println("这是一条很普通的日志。")                                  // 写入日志
}
