package main

import (
	"log"
	"net"
	"strconv"

	"github.com/didrocks/grpc-piglow/piglowservice"
	"google.golang.org/grpc"
)

//go:generate protoc -I proto/ proto/piglow.proto --go_out=plugins=grpc:proto
func main() {
	var l net.Listener
	var err error

	port := 3146
	for l == nil {
		l, err = net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			log.Printf("failed to listenn: %v\n", err)
			port++
		}
	}
	log.Printf("listening on port %d\n", port)

	s := grpc.NewServer()

	if err := piglowservice.RegisterPiGlowService(s); err != nil {
		log.Fatalf("couldn't create piglow proxy: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
