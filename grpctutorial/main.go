package main

import (
	"context"
	pb "gintutorial/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
	func main() {
		r := gin.Default()

		r.GET("/ping", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		r.Run(":8080")
	}
*/
type server struct {
	pb.UnimplementedGreeterServer
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
