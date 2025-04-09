package grpcbasic

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "ys3721.com/grpcbasic/calculator"
)

func doGrpcClientMain() {
	conn, err := grpc.Dial("localhost:9459", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	req := &pb.AddRequest{A: 10, B: 20}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call Add: %v", err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
