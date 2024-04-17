package main

import (
	"context"
	"fmt"
	hellopb "gRPC_Server/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func NewServer() *gRPCServer {
	return &gRPCServer{}
}

type gRPCServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	hellopb.RegisterGreetingServiceServer(s, NewServer())

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Stopping gRPC server...")
	s.GracefulStop()
}

func (s *gRPCServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{Message: fmt.Sprintf("Hello, %s", req.GetName())}, nil
}
