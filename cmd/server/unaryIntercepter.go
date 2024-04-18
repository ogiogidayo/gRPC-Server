package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func myUnaryServerIntercepter1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server intercepter 1: ", info.FullMethod)
	res, err := handler(ctx, req) // 本来の処理
	log.Println("[post] my unary server intercepter 1 ", req)
	return res, err
}
