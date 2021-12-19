package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 毛剑老师第三周作业：基于 errgroup 实现一个 http server 的启动和关闭，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
// errgroup：只要一个 goroutine 出错我们就不再等其他 goroutine 了，减少资源浪费。
func main() {
	g, ctx := errgroup.WithContext(context.Background())
	var (
		addr1 = "127.0.0.1:8001"
		addr2 = "127.0.0.1:8002"
	)
	// g1 端口8001启动
	g.Go(func() error {
		// 创建ServeMux对象，防止使用DefaultServeMux默认对象
		mux1 := http.NewServeMux()
		mux1.HandleFunc("/test1", Test1)
		mux1.HandleFunc("/test2", Test2)
		mux1.HandleFunc("/test3", Test3)
		return StartServer(ctx, addr1, mux1)
	})
	// g2 端口8002启动
	g.Go(func() error {
		mux2 := http.NewServeMux()
		mux2.HandleFunc("/test4", Test4)
		mux2.HandleFunc("/test5", Test5)
		mux2.HandleFunc("/test6", Test6)
		return StartServer(ctx, addr2, mux2)

	})
	// g3 接收信号退出
	g.Go(func() error {
		stopChannel := make(chan os.Signal, 1)
		// SIGHUP 终止进程 终端线路挂断
		// SIGINT 终止进程 中断进程 Control-C (SIGINT)
		// SIGTERM 终止进程 软件终止信号
		signal.Notify(stopChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM) // 用来监听收到的信号
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-stopChannel: // 接收终止信号，g3退出，cancel取消，context不再阻塞，g2、g1退出
			return errors.Errorf("get os signal: %v", sig)
		}
	})
	if err := g.Wait(); err != nil {
		log.Printf("errgroup exiting: %+v", err)
	} else {
		log.Println("all group done!")
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
	// 创建http.Server对象定制参数
	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	// context收到取消信号，g1、g2退出，context不再阻塞，g3退出
	go func() {
		<-ctx.Done()
		log.Printf("server %s stop\n", addr)
		s.Shutdown(ctx)
	}()
	log.Printf("server %s start\n", addr)
	return s.ListenAndServe()
}
