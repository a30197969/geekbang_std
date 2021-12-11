package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	// 命令参数存储
	flag.StringVar(&name, "name", "everyone", "问候人")
}
func main() {
	// 函数flag.Parse用于真正解析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("你好，%s！\n", name)
}
