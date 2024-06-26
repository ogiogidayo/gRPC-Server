package main

import (
	"context"
	"errors"
	"fmt"
	hellopb "gRPC_Server/pkg/grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
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

	s := grpc.NewServer(
		grpc.UnaryInterceptor(myUnaryServerIntercepter1),
	)

	hellopb.RegisterGreetingServiceServer(s, NewServer())

	reflection.Register(s)

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

func (s *gRPCServer) Hello(_ context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	//return &hellopb.HelloResponse{Message: fmt.Sprintf("Hello, %s", req.GetName())}, nil
	stat := status.New(codes.Unknown, "unknown error occurred")
	stat, _ = stat.WithDetails(&errdetails.DebugInfo{
		Detail: "detail reason of err",
	})
	err := stat.Err()
	return nil, err
}

func (s *gRPCServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func (s *gRPCServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			message := fmt.Sprintf("Hello %v!", nameList)
			return stream.SendAndClose(&hellopb.HelloResponse{
				Message: message,
			})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

func (s *gRPCServer) HelloBiStreams(stream hellopb.GreetingService_HelloBiStreamsServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello, %v", req.GetName())
		if err := stream.Send(&hellopb.HelloResponse{
			Message: message,
		}); err != nil {
			return err
		}
	}
}
