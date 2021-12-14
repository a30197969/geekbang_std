package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTracker(t *testing.T) {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")
	_ = tr.Event(context.Background(), "test4")
	_ = tr.Event(context.Background(), "test5")
	time.Sleep(time.Second * 5)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	tr.Shutdown(ctx)
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

type Tracker struct {
	ch   chan string
	stop chan struct{} // 控制chan的暂停
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// 不停的从goroutine中消费数据，模拟上报
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{} // close信号后，发送一个信号给stop
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch) // 不会发送数据了
	select {    // Run函数退出后，接收stop信号
	case <-t.stop:
	case <-ctx.Done():
	}
}
