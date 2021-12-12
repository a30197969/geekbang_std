package channel

import (
	"fmt"
	"testing"
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
