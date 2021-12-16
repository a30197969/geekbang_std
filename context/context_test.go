package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextWithValue(t *testing.T) {
	// 生成顶层 Context
	ctx := context.Background()
	//ctx := context.TODO()
	// WithValue 基于 parent Context 生成一个新的 Context，常常用来传递上下文，链式查找
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0002")
	ctx = context.WithValue(ctx, "key3", "0003")
	ctx = context.WithValue(ctx, "key4", "0004")
	fmt.Println(ctx.Value("key1"))
}
func TestContextWithCancel(t *testing.T) {
	// WithCancel 方法返回 parent 的副本，用于主动取消长时间的任务
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx)
	cancel()
}
func TestContextWithTimeout(t *testing.T) {
	// 创建一个子节点的context，超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	fmt.Println(ctx)
	cancel()
}
func TestContextWithDeadline(t *testing.T) {
	// 创建一个子节点的context，5秒后自动取消
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	fmt.Println(ctx)
	cancel()
}

// parent Context 被 cancel 后，子 context 是否也立刻被 cancel 了
func TestParent(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	child := context.WithValue(ctx, "test1", "abcd")
	go func() {
		for {
			select {
			case <-child.Done():
				return
			default:
				val := child.Value("test1")
				fmt.Println(val)
				time.Sleep(time.Second)

			}
		}
	}()
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 2)
}



