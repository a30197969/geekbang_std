package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // 生成随机数种子
	fmt.Println(rand.Intn(11))       // 前闭后开区间 0-10取
	fmt.Println(rand.Int31n(99999))  // 前闭后开区间 0-取
	fmt.Println(rand.Float32())      // 前闭后开区间 0-1取
	fmt.Println(rand.Float64())      // 前闭后开区间 0-1取

}
