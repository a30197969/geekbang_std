package main

import "fmt"

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- ServeDebug(stop)
	}()
	go func() {
		done <- ServeApp(stop)
	}()
	var stopped bool
	for i := 0; i < cap(done); i++ {
		fmt.Println("1111")            // 这里会被执行
		if err := <-done; err != nil { // 这里会被阻塞，直到有错误产生
			fmt.Printf("error: %v\n", err)
		}
		fmt.Println("2222") // 没有错误这里不会输出
		if !stopped {
			fmt.Printf("%v\n", stopped)
			stopped = true
			close(stop)
		}
	}
}
