package five

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewSlideWindowStatistics(t *testing.T) {
	intervalInMs := 1000
	count := 10
	qps := 100
	rand.Seed(time.Now().UnixNano())
	s := NewSlideWindowStatistics(intervalInMs, count, qps)

	//for i := 0; i < 100; i++ {
	//	requestCount := rand.Intn(30)
	//	// 每100毫秒随机分配几个请求
	//	s.SlidingWindow(time.Now().UnixMilli(), requestCount)
	//	time.Sleep(time.Millisecond * 34)
	//}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				requestCount := rand.Intn(20)
				// 每100毫秒随机分配几个请求
				s.SlidingWindow(time.Now().UnixMilli(), requestCount)
				time.Sleep(time.Millisecond * 34)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
