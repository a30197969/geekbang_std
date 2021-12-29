package five

import (
	"fmt"
	"sync"
)

type SlideWindowStatistics struct {
	intervalInMs int          // 统计周期，滑动窗口的总体长度，单位毫秒
	count        int          // 划分个数
	windowSize   int          // 窗口长度，统计周期的精度 = 统计周期 / 划分个数
	incrSlice    []int        // 窗口计数器
	qps          int          // 允许的最大QPS
	mux          sync.RWMutex // 读写锁处理并发计数统计
}

func NewSlideWindowStatistics(intervalInMs int, count int, qps int) *SlideWindowStatistics {
	windowSize := intervalInMs / count
	incrSlice := make([]int, count, count)
	return &SlideWindowStatistics{
		intervalInMs: intervalInMs,
		count:        count,
		windowSize:   windowSize, // 窗口大小为100ms
		incrSlice:    incrSlice,
		qps:          qps,
	}
}

// 只有高效的统计单位时间内的请求数量，才能根据请求数量是否超出限制，来判断下一次请求是否被允许
// 测试结果如下：
// 当前总请求数为80，进来的请求数为19
// 当前总请求数为89，进来的请求数为5
// 当前总请求数为92，进来的请求数为0
// --------------------------------限流了，当前总请求数为92，进来的请求数为9
// 当前总请求数为85，进来的请求数为9
// 当前总请求数为85，进来的请求数为12
// --------------------------------限流了，当前总请求数为94，进来的请求数为17
// 当前总请求数为97，进来的请求数为3
// --------------------------------限流了，当前总请求数为97，进来的请求数为6
// 当前总请求数为88，进来的请求数为8
// 当前总请求数为92，进来的请求数为3
// 当前总请求数为88，进来的请求数为5
// 当前总请求数为80，进来的请求数为3
// 当前总请求数为80，进来的请求数为6

// SlidingWindow 移动窗口统计
func (s *SlideWindowStatistics) SlidingWindow(msTime int64, requestCount int) {
	index := s.LocationIndex(msTime)
	//fmt.Println(msTime, requestCount, index)
	// 统计周期的请求总数
	total := s.GetTotalRequest()
	if total+requestCount > s.qps {
		fmt.Printf("--------------------------------限流了，当前总请求数为%v，进来的请求数为%v\n", total, requestCount)
		return
	}
	fmt.Printf("当前总请求数为%v，进来的请求数为%v\n", total, requestCount)
	s.mux.Lock()
	defer s.mux.Unlock()
	s.incrSlice[index] = requestCount
}

// LocationIndex 获取当前时间戳在窗口中的位置
func (s *SlideWindowStatistics) LocationIndex(msTime int64) int64 {
	return msTime % int64(s.count)
}

// GetTotalRequest 获取统计周期内的总请求数
func (s *SlideWindowStatistics) GetTotalRequest() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	var total int
	// fmt.Println(s.incrSlice)
	for _, val := range s.incrSlice {
		total += val
	}
	return total
}
