package piglowservice

import (
	"log"

	"fmt"

	pb "github.com/didrocks/grpc-piglow/proto"
	context "golang.org/x/net/context"
)

type brightnessFunc func(uint8)

// Set LED n to brightness
func (s *service) SetLED(ctx context.Context, in *pb.LedRequest) (*pb.Ack, error) {
	var err error

	n, err := ensureNumLed(in.Num)
	if err != nil {
		return nil, err
	}
	b, err := ensureBrightness(in.Brightness)
	if err != nil {
		return nil, err
	}

	s.p.SetLED(n, b)
	return s.apply()
}

// Set all LEDs to brightness
func (s *service) SetAll(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetAll)
}

// SetWhite all White LEDs to brightness
func (s *service) SetWhite(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetWhite)
}

// SetBlue all Blue LEDs to brightness
func (s *service) SetBlue(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetBlue)
}

// SetGreen all Green LEDs to brightness
func (s *service) SetGreen(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetGreen)
}

// SetYellow all WhYellowite LEDs to brightness
func (s *service) SetYellow(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetYellow)
}

// SetOrange all Orange LEDs to brightness
func (s *service) SetOrange(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetOrange)
}

// SetRed all Red LEDs to brightness
func (s *service) SetRed(ctx context.Context, in *pb.BrightnessRequest) (*pb.Ack, error) {
	return s.setBrightnessWithFunc(ctx, in, s.p.SetRed)
}

// convert and ensure num led is valid
func ensureNumLed(n int32) (int8, error) {
	if n < 0 || n > 17 {
		return 0, fmt.Errorf("invalid led number: %d", n)
	}
	return int8(n), nil
}

// convert and ensure brightness is valid
func ensureBrightness(b uint32) (uint8, error) {
	if b > 255 {
		return 0, fmt.Errorf("invalid brightness value: %d", b)
	}
	return uint8(b), nil
}

// internal apply correct changes functions
func (s *service) apply() (ack *pb.Ack, err error) {
	ack = &pb.Ack{Ok: true}
	if err = s.p.Apply(); err != nil {
		// server side logging
		log.Println("Couldn't apply changes: ", err)
	}
	return ack, nil
}

// internal helper taking any piglow functions which change only brightness
func (s *service) setBrightnessWithFunc(ctx context.Context, in *pb.BrightnessRequest, fn brightnessFunc) (*pb.Ack, error) {
	b, err := ensureBrightness(in.Brightness)
	if err != nil {
		return nil, err
	}

	fn(b)
	return s.apply()
}
