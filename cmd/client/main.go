package client

import (
	"bufio"
	"fmt"
	hellopb "gRPC_Server/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 標準入力から文字列を受け取る
	scanner = bufio.NewScanner(os.Stdin)

	// gROCサーバーとのコネクションを確率
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("connection failed.")
		return
	}
	defer conn.Close()

	// gRPC Clientを生成
	client = hellopb.NewGreetingServiceClient(conn)
	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()

		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}
