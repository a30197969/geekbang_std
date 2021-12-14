package channel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch1 := make(chan int, 2)
	// 发送方
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender：sending element %v\n", i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}
	fmt.Println("End.")
}

func TestChannel2(t *testing.T) {
	intChan := getIntChan()
	fmt.Printf("%#v\n", intChan)
	// 带有range子句的for语句
	for elem := range intChan {
		fmt.Printf("%+v\n", elem)
		fmt.Printf("%#v\n", elem)
	}
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func TestSelectChannel(t *testing.T) {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	fmt.Printf("%#v\n", intChannels)
	index := rand.Intn(3)
	fmt.Printf("%v\n", index)
	intChannels[index] <- index
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}
func TestSelectChannelBreak(t *testing.T) {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})

	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
