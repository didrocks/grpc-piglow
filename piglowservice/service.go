package piglowservice

import (
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/didrocks/grpc-piglow/proto"

	piglow "github.com/wjessop/go-piglow"
)

// service receiving piglow rpc requests
type service struct {
	p *piglow.Piglow
}

// RegisterPiGlowService create a piglow object, register it with the matching protocol and grpc server
func RegisterPiGlowService(gserver *grpc.Server) error {

	// Create a new Piglow
	p, err := piglow.NewPiglow()
	if err != nil {
		return fmt.Errorf("Couldn't create a Piglow: %v", err)
	}

	pb.RegisterPiGlowServer(gserver, &service{p})
	return nil
}
