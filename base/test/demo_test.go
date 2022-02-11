package test

import (
	"fmt"
	"testing"
)

// goto 用法
func TestDemo1(t *testing.T) {
	n := 30
	if n > 20 {
		t.Logf("很不懂\n")
		goto lable3 // 无条件跳转
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
lable3:
	fmt.Println("ddd")
	fmt.Println("eee")
	fmt.Println("fff")
}

// 值类型 和 引用类型
func TestDemo2(t *testing.T) {
	// 值类型
	var a = 10
	b := a
	a = 15
	fmt.Printf("%+v, %+v\n", a, b)

	// 值类型
	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1
	arr1[0] = 11
	fmt.Printf("%v, %p, %v, %p\n", arr1, &arr1, arr2, &arr2) // 指针 十六进制表示

	// 引用类型
	var slice1 = []int{1, 2, 3}
	slice2 := slice1
	slice1[0] = 111
	fmt.Printf("%v, %p, %v, %p\n", slice1, slice1, slice2, slice2)

	//每个变量都有两层含义，变量的内存和变量的地址
	var aaa int = 10
	fmt.Printf("a = %d\n", aaa)  //变量的内存 10
	fmt.Printf("a = %v\n", &aaa) //变量的地址 0xc042060080

	// 多维数组
	var arrMulti = [3][2]string{
		{"北京", "上海"},
		{"山东", "安徽"},
		{"山西", "浙江"},
	}
	fmt.Printf("%#v\n", arrMulti)
	fmt.Printf("%#v\n", arrMulti[0][0])
	for _, strings := range arrMulti {
		for _, s := range strings {
			fmt.Printf("%#v\n", s)
		}
	}
}

// 切片
// 长度：切片的长度就是它所包含的元素个数
// 容量：切片的容量就是从它的第一个元素开始数，到其底层数组元素末尾的个数
func TestDemo3(t *testing.T) {
	// 定义切片
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := make([]int, 3, 89)
	fmt.Println(s1, s2)

	a := []string{"北京", "上海", "山东", "安徽", "山西", "浙江"}
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", a, a, len(a), cap(a))
	b := a[2:]
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", b, b, len(b), cap(b))
	c := a[1:3]
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", c, c, len(c), cap(c))
}

// 切片复制
func TestDemo4(t *testing.T) {
	// 定义切片
	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = []int{5, 6, 7}
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, s2...)
	fmt.Println(s1, s2)
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", s1, s1, len(s1), cap(s1))

	res := copy(s1, s2)
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", s2, s2, len(s2), cap(s2))
	fmt.Printf("%v\n", res)

}

// 切片合并
func TestDemo5(t *testing.T) {
	// 定义切片
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("%v, 类型：%T, 长度 = %d, 容量 = %d\n", s1, s1, len(s1), cap(s1))
}
