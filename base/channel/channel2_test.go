package channel

import (
	"fmt"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	// 统计 1-120000 中有多少个素数
	// 质数又称素数，质数是指在大于1的自然数中，除了1和它本身以外不再有其他因数的自然数。
	// 从 intChan 去除数据，判断是否为素数，写入 primeChan
	num := 72
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	var wg sync.WaitGroup
	wg.Add(1)
	go putNum(intChan, num, &wg)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, &wg)
	}
	wg.Wait()
	close(primeChan)
	fmt.Println("共有素数：", len(primeChan))
	wg.Add(1)
	go printPrime(primeChan, &wg)
	wg.Wait()
}

// 存放数据
func putNum(intChan chan int, num int, wg *sync.WaitGroup) {
	for i := 2; i < num; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// 统计素数
func primeNum(intChan chan int, primeChan chan int, wg *sync.WaitGroup) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num, "是素数")
			primeChan <- num
		}
	}
	wg.Done()
}

// 打印素数
func printPrime(primeChan chan int, wg *sync.WaitGroup) {
	for num := range primeChan {
		fmt.Println(num)
	}
	wg.Done()
}
