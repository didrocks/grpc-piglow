package main

import (
	"log"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/didrocks/grpc-piglow/remote"
	piglow "github.com/wjessop/go-piglow"
)

// service
type service struct {
	p *piglow.Piglow
}

func (s *service) SetLED(ctx context.Context, in *pb.LedRequest) (*pb.Ack, error) {
	s.p.SetLED((int8)(in.Num), (uint8)(in.Brightness))
	err := s.p.Apply()
	if err != nil { // Apply the changes
		log.Println("Couldn't apply changes: ", err)
	}
	return &pb.Ack{Ok: true}, err
}

//go:generate protoc -I remote/ remote/piglow.proto --go_out=plugins=grpc:remote
func main() {
	var p *piglow.Piglow
	var err error

	// Create a new Piglow
	p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	lis, err := net.Listen("tcp", ":9875")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPiGlowServer(s, &service{p})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
