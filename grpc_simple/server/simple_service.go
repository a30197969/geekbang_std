package main

import (
	"context"
	"geekbang_study/proto"
)

// SimpleService 定义我们的服务
type SimpleService struct {
}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, request *proto.SimpleRequest) (*proto.SimpleResponse, error) {
	res := proto.SimpleResponse{
		Code:  200,
		Value: "hello " + request.Data,
	}
	return &res, nil
}
