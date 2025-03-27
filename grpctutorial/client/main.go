package main

import (
	"context"
	"log"
	"time"

	pb "gintutorial/helloworld" // 根据实际生成的包路径调整

	"google.golang.org/grpc"
)

func main() {
	// 连接服务端（使用 WithInsecure 仅用于示例，生产环境建议使用 TLS）
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
