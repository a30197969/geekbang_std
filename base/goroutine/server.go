package main

import (
	"context"
	"net/http"
)

func Serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())

	}()
	return s.ListenAndServe()
}
func ServeDebug(stop <-chan struct{}) error {
	return Serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}
func ServeApp(stop <-chan struct{}) error {
	return Serve("127.0.0.1:8000", http.DefaultServeMux, stop)
}
