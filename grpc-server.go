package main

import (
	"log"
	"net"

	pb "github.com/grpc-up-and-running/samples/ch02/productinfo/go/proto"
)

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
