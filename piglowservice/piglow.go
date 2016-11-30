package piglowservice

import (
	"log"

	pb "github.com/didrocks/grpc-piglow/proto"
	context "golang.org/x/net/context"
)

func (s *service) SetLED(ctx context.Context, in *pb.LedRequest) (*pb.Ack, error) {
	s.p.SetLED((int8)(in.Num), (uint8)(in.Brightness))
	err := s.p.Apply()
	if err != nil { // Apply the changes
		log.Println("Couldn't apply changes: ", err)
	}
	return &pb.Ack{Ok: true}, err
}
