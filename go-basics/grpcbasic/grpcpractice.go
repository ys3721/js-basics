package grpcbasic

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "ys3721.com/grpcbasic/calculator"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := (*req).A + req.B
	fmt.Printf("Received: %d + %d = %d\n", req.A, (*req).B, result)
	return &pb.AddResponse{Result: result}, nil
}

func doGrpcBasicMain() {
	listener, err := net.Listen("tcp", ":9459")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &CalculatorServer{})

	fmt.Println("gRpc server running on port 9459.....")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
