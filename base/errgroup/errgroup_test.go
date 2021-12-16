package errgroup

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
	"testing"
)

var urls = []string{
	"http://www.golang.org/",
	"http://www.baidu.com/",
	"http://www.noexist11111111.com/",
}

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp.Body.Close()
		}(url)
	}
	wg.Wait()
}
func TestErrgroup(t *testing.T) {
	eg := new(errgroup.Group)
	for _, url := range urls {
		url := url
		eg.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Printf("get [%s] success: [%d] \n", url, resp.StatusCode)
			return resp.Body.Close()
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All success!")
	}
}
