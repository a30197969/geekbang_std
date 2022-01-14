package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var myname string
var myage int

func init() {
	// 命令参数存储
	//第 1 个参数是用于存储该命令参数值的地址，具体到这里就是在前面声明的变量name的地址了，由表达式&name表示。
	//第 2 个参数是为了指定该命令参数的名称，这里是name。
	//第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone。
	//至于第 4 个函数参数，即是该命令参数的简短说明了，这在打印命令说明时会用到。
	flag.StringVar(&myname, "name", "everyone", "问候人")
	flag.IntVar(&myage, "age", 18, "年龄")
}

func main() {
	// 函数flag.Parse用于真正解析命令参数，并把它们的值赋给相应的变量。
	flag.Parse()
	Hello(os.Stdout, myname)
	fmt.Printf("我的年龄是：%d！\n", myage)
	fmt.Println(os.Args)

	t := time.Now()
	str := t.Format("2006-01-02 15:04:05")
	fmt.Println(str)
}
