package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var mailbox uint8                        // mailbox 代表信箱。值为0则表示信箱中没有情报，而当它的值为1时则说明信箱中有情报
	var lock sync.RWMutex                    // lock 代表信箱上的锁
	sendCond := sync.NewCond(&lock)          // 为放置情报而准备的条件变量
	recvCond := sync.NewCond(lock.RLocker()) // 为获取情报而准备的条件变量
	sign := make(chan struct{}, 3)           // sign 用于传递演示完成的信号
	max := 5

	// 发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			fmt.Printf("sender [%d]: the mailbox is empty.\n", i)
			mailbox = 1
			fmt.Printf("sender [%d]: the letter has been sent.\n", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 0; j < max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			fmt.Printf("receiver [%d]: the mailbox is full.\n", j)
			mailbox = 0
			fmt.Printf("receiver [%d]: the letter has been received.\n", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)
	// 利用chan 阻塞主线程以使goroutine完全跑完
	<-sign
	<-sign
}
