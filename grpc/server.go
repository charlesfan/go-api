package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/charlesfan/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const Protocol string = "tcp"

func Run(port string) {
	fmt.Printf("[gRPC test] gRPC start service with %s on %s\n", Protocol, port)
	lis, err := net.Listen(Protocol, port)
	if err != nil {
		log.Fatalf("gRPC binding port failed: %v", err)
	}
	// Create gRPC Server and regist Service
	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &login{})

	reflection.Register(s)
	// Start
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC run error: %v", err)
	}
}
