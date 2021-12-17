package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type config struct {
	a []int
}

func (c *config) T() {

}

func TestGoroutine(t *testing.T) {
	var v atomic.Value
	v.Store(&config{})
	go func() {
		i := 0
		for {
			i++
			cfg := &config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 2; j++ {
				cfg := v.Load().(*config)
				cfg.T()
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

func TestStore(t *testing.T) {
	var v atomic.Value
	v.Store("111")
	v.Store("222")
	v.Store("333")
	t.Log(v.Load().(string))
	t.Log(v.Load().(string))
	t.Log(v.Load().(string))
}
