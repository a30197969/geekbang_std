package test

import (
	"fmt"
	"testing"
)

type A interface{} // 空接口，可以表示任何类型

func TestName(t *testing.T) {
	var a A
	var str = "你好"
	a = str // 让字符串实现A接口   相当于  var _ 接口类型 = 结构体
	fmt.Printf("值：%v 类型：%T\n", a, a)

	var num = 20
	a = num
	fmt.Printf("值：%v 类型：%T\n", a, a)

	var flag = true
	a = flag
	fmt.Printf("值：%v 类型：%T\n", a, a)

}

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Printf("%v 手机启动\n", p.Name)
}

func (p Phone) stop() {
	fmt.Printf("%v 手机关机\n", p.Name)
}

func TestName2(t *testing.T) {
	// 结构体值接收者：实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
	// 结构体指针接收者：只有实例化后的结构体指针类型可以赋值给接口变量
	p := Phone{
		Name: "小米",
	}
	var u Usber = p // 让手机实现USB接口
	u.start()

	p2 := &Phone{
		Name: "苹果",
	}
	var u2 Usber = p2
	u2.start()

}
