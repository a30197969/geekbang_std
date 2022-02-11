package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name   string
	Age    int
	Sex    string
	Height int
}

func (p *Person) PrintInfo() {
	fmt.Printf("姓名：%v，年龄：%d\n", p.Name, p.Age)
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func TestStruct1(t *testing.T) {
	p := &Person{
		Name:   "小李",
		Age:    18,
		Sex:    "男",
		Height: 178,
	}
	p.PrintInfo()
	p.SetInfo("小彭", 65)
	p.PrintInfo()

	p2 := &Person{
		Name:   "王五",
		Age:    22,
		Sex:    "男",
		Height: 187,
	}
	p2.PrintInfo()
	p2.SetInfo("嘎嘎", 65)
	p2.PrintInfo()
}

type Person2 struct {
	Name    string
	Age     int
	Hobby   []string
	Subject map[string]string
}

func TestStruct2(t *testing.T) {
	hobby := []string{"美术", "体育", "日语"}
	subject := make(map[string]string, 3)
	subject["美术"] = "凑活"
	subject["体育"] = "差"
	subject["日语"] = "优秀"
	p := &Person2{
		Name:    "阿拉法",
		Age:     20,
		Hobby:   hobby,
		Subject: subject,
	}
	fmt.Printf("%#v\n", p)
}

type User struct {
	ID       int
	Name     string
	Password string
	Address  *Address
}

type Address struct {
	Userid   int
	Province string
	City     string
}

func TestStruct3(t *testing.T) {
	u := &User{
		ID:       1,
		Name:     "小嘎嘎",
		Password: "123456",
	}
	u.Address = &Address{
		Userid:   1,
		Province: "北京",
		City:     "北京",
	}
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.Address)
}

