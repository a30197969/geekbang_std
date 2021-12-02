package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "imooc/routers" // 只执行routers的init方法
)

func main() {
	beego.Run()
}
