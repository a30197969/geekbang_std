package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	g, errCtx := errgroup.WithContext(ctx)
	var (
		addr1 = "127.0.0.1:8001"
		addr2 = "127.0.0.1:8002"
	)
	// 端口8001启动
	g.Go(func() error {
		mux1 := http.NewServeMux()
		mux1.HandleFunc("/test1", Test1)
		mux1.HandleFunc("/test2", Test2)
		mux1.HandleFunc("/test3", Test3)
		return StartServer(errCtx, addr1, mux1)
	})
	// 端口8002启动
	g.Go(func() error {
		mux2 := http.NewServeMux()
		mux2.HandleFunc("/test4", Test4)
		mux2.HandleFunc("/test5", Test5)
		mux2.HandleFunc("/test6", Test6)
		return StartServer(errCtx, addr2, mux2)

	})
	chanel := make(chan os.Signal, 1) //这里要用 buffer 为1的 chan
	signal.Notify(chanel)
	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): // 因为 cancel、timeout、deadline 都可能导致 Done 被 close
				return errCtx.Err()
			case <-chanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("all group done!")
	}
}

func Test1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 1111111")
}
func Test2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 2222222")
}
func Test3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 3333333")
}
func Test4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 4444444")
}
func Test5(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 5555555")
}
func Test6(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 6666666")
}
func StartServer(ctx context.Context, addr string, handler http.Handler) error {
	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-ctx.Done()
		fmt.Printf("%s stop\n", addr)
		s.Shutdown(ctx)
	}()
	fmt.Printf("%s start\n", addr)
	return s.ListenAndServe()
}
