package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var done = false // 互斥锁需要保护的条件变量
// sync.Cond 经常用在多个 goroutine 等待，一个 goroutine 通知（事件发生）的场景。如果是一个通知，一个等待，使用互斥锁或 channel 就能搞定了。
// 一个协程在异步地接收数据，剩下的多个协程必须等待这个协程接收完数据，才能读取到正确的数据
func TestCond(t *testing.T) {
	cond := sync.NewCond(&sync.RWMutex{})
	go read("read1", cond)
	go read("read2", cond)
	go read("read3", cond)
	go read("read4", cond)
	write("writer", cond)
	time.Sleep(time.Second * 3)
}

// 调用 Wait() 等待通知，直到 done 为 true
func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	fmt.Println(name, "start read")
	c.L.Unlock()
}

// 接收数据，接收完成后，将 done 置为 true，调用 Broadcast() 通知所有等待的协程
func write(name string, c *sync.Cond) {
	fmt.Println(name, "start write")
	// 暂停了 1s，一方面是模拟耗时，另一方面是确保前面的 3 个 read 协程都执行到 Wait()，处于等待状态。TestCond 函数最后暂停了 3s，确保所有操作执行完毕
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	fmt.Println(name, "wake all")
	c.Broadcast()
}
