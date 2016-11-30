// Code generated by protoc-gen-go.
// source: piglow.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	piglow.proto

It has these top-level messages:
	LedRequest
	Ack
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type LedRequest struct {
	Num        int32  `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Brightness uint32 `protobuf:"varint,2,opt,name=brightness" json:"brightness,omitempty"`
}

func (m *LedRequest) Reset()                    { *m = LedRequest{} }
func (m *LedRequest) String() string            { return proto1.CompactTextString(m) }
func (*LedRequest) ProtoMessage()               {}
func (*LedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LedRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *LedRequest) GetBrightness() uint32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

type Ack struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto1.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ack) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto1.RegisterType((*LedRequest)(nil), "proto.LedRequest")
	proto1.RegisterType((*Ack)(nil), "proto.Ack")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PiGlow service

type PiGlowClient interface {
	// SetLED
	SetLED(ctx context.Context, in *LedRequest, opts ...grpc.CallOption) (*Ack, error)
}

type piGlowClient struct {
	cc *grpc.ClientConn
}

func NewPiGlowClient(cc *grpc.ClientConn) PiGlowClient {
	return &piGlowClient{cc}
}

func (c *piGlowClient) SetLED(ctx context.Context, in *LedRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/proto.PiGlow/SetLED", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PiGlow service

type PiGlowServer interface {
	// SetLED
	SetLED(context.Context, *LedRequest) (*Ack, error)
}

func RegisterPiGlowServer(s *grpc.Server, srv PiGlowServer) {
	s.RegisterService(&_PiGlow_serviceDesc, srv)
}

func _PiGlow_SetLED_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiGlowServer).SetLED(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PiGlow/SetLED",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiGlowServer).SetLED(ctx, req.(*LedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PiGlow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PiGlow",
	HandlerType: (*PiGlowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLED",
			Handler:    _PiGlow_SetLED_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "piglow.proto",
}

func init() { proto1.RegisterFile("piglow.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x4c, 0xcf,
	0xc9, 0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x76, 0x5c, 0x5c,
	0x3e, 0xa9, 0x29, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x79, 0xa5,
	0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x20, 0xa6, 0x90, 0x1c, 0x17, 0x57, 0x52, 0x51,
	0x66, 0x7a, 0x46, 0x49, 0x5e, 0x6a, 0x71, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x92,
	0x88, 0x92, 0x28, 0x17, 0xb3, 0x63, 0x72, 0xb6, 0x10, 0x1f, 0x17, 0x53, 0x7e, 0x36, 0x58, 0x1f,
	0x47, 0x10, 0x53, 0x7e, 0xb6, 0x91, 0x31, 0x17, 0x5b, 0x40, 0xa6, 0x7b, 0x4e, 0x7e, 0xb9, 0x90,
	0x26, 0x17, 0x5b, 0x70, 0x6a, 0x89, 0x8f, 0xab, 0x8b, 0x90, 0x20, 0xc4, 0x66, 0x3d, 0x84, 0x7d,
	0x52, 0x5c, 0x50, 0x21, 0xc7, 0xe4, 0x6c, 0x25, 0x86, 0x24, 0x36, 0x30, 0xc7, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x77, 0x17, 0x69, 0x53, 0xa9, 0x00, 0x00, 0x00,
}