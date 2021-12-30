package five

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// positive feedback：用户总是积极重试，访问一个不可达的服务

func TestRetry(t *testing.T) {
	cb := NewClientBackoff()
	time := cb.Backoff(8)
	fmt.Printf("%+v\n", time)
}

type ClientBackoff struct {
	BaseDelay         time.Duration // 第一次失败重试前后需等待多久
	Multiplier        float64       // 在失败的重试后乘以的倍数
	Jitter            float64       // 随机抖动因子
	MaxDelay          time.Duration // backoff上限
	MinConnectTimeout time.Duration // 最短重试间隔
}

func NewClientBackoff() *ClientBackoff {
	return &ClientBackoff{
		BaseDelay:         1.0 * time.Second,
		Multiplier:        1.6,
		Jitter:            0.2,
		MaxDelay:          10 * time.Second,
		MinConnectTimeout: 3 * time.Second,
	}
}

// 当我们连接到一个失败的后端时，通常希望不要立即重试(以避免泛滥的网络或服务器的请求)，而是做某种形式的指数backoff
func (cb *ClientBackoff) Backoff(retries int) time.Duration {
	if retries == 0 {
		return cb.BaseDelay
	}
	backoff, max := float64(cb.BaseDelay), float64(cb.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= cb.Multiplier
		retries--
	}
	if backoff > max {
		backoff = max
	}
	backoff *= 1 + cb.Jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
