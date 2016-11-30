package main

import (
	"log"
	"net"

	"github.com/didrocks/grpc-piglow/piglowservice"

	"google.golang.org/grpc"
)

//go:generate protoc -I proto/ proto/piglow.proto --go_out=plugins=grpc:proto
func main() {

	lis, err := net.Listen("tcp", ":9875")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	if err := piglowservice.RegisterPiGlowService(s); err != nil {
		log.Fatalf("couldn't create piglow proxy: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
