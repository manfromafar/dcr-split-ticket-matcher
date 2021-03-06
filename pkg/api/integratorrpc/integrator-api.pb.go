// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integrator-api.proto

/*
Package integratorrpc is a generated protocol buffer package.

It is generated from these files:
	integrator-api.proto

It has these top-level messages:
	ValidateVoteAddressRequest
	ValidateVoteAddressResponse
	ValidatePoolSubsidyAddressRequest
	ValidatePoolSubsidyAddressResponse
*/
package integratorrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ValidateVoteAddressRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *ValidateVoteAddressRequest) Reset()                    { *m = ValidateVoteAddressRequest{} }
func (m *ValidateVoteAddressRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateVoteAddressRequest) ProtoMessage()               {}
func (*ValidateVoteAddressRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ValidateVoteAddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ValidateVoteAddressResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *ValidateVoteAddressResponse) Reset()                    { *m = ValidateVoteAddressResponse{} }
func (m *ValidateVoteAddressResponse) String() string            { return proto.CompactTextString(m) }
func (*ValidateVoteAddressResponse) ProtoMessage()               {}
func (*ValidateVoteAddressResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidateVoteAddressResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ValidatePoolSubsidyAddressRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *ValidatePoolSubsidyAddressRequest) Reset()         { *m = ValidatePoolSubsidyAddressRequest{} }
func (m *ValidatePoolSubsidyAddressRequest) String() string { return proto.CompactTextString(m) }
func (*ValidatePoolSubsidyAddressRequest) ProtoMessage()    {}
func (*ValidatePoolSubsidyAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

func (m *ValidatePoolSubsidyAddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ValidatePoolSubsidyAddressResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *ValidatePoolSubsidyAddressResponse) Reset()         { *m = ValidatePoolSubsidyAddressResponse{} }
func (m *ValidatePoolSubsidyAddressResponse) String() string { return proto.CompactTextString(m) }
func (*ValidatePoolSubsidyAddressResponse) ProtoMessage()    {}
func (*ValidatePoolSubsidyAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *ValidatePoolSubsidyAddressResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidateVoteAddressRequest)(nil), "integratorrpc.ValidateVoteAddressRequest")
	proto.RegisterType((*ValidateVoteAddressResponse)(nil), "integratorrpc.ValidateVoteAddressResponse")
	proto.RegisterType((*ValidatePoolSubsidyAddressRequest)(nil), "integratorrpc.ValidatePoolSubsidyAddressRequest")
	proto.RegisterType((*ValidatePoolSubsidyAddressResponse)(nil), "integratorrpc.ValidatePoolSubsidyAddressResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VotePoolIntegratorService service

type VotePoolIntegratorServiceClient interface {
	ValidateVoteAddress(ctx context.Context, in *ValidateVoteAddressRequest, opts ...grpc.CallOption) (*ValidateVoteAddressResponse, error)
	ValidatePoolSubsidyAddress(ctx context.Context, in *ValidatePoolSubsidyAddressRequest, opts ...grpc.CallOption) (*ValidatePoolSubsidyAddressResponse, error)
}

type votePoolIntegratorServiceClient struct {
	cc *grpc.ClientConn
}

func NewVotePoolIntegratorServiceClient(cc *grpc.ClientConn) VotePoolIntegratorServiceClient {
	return &votePoolIntegratorServiceClient{cc}
}

func (c *votePoolIntegratorServiceClient) ValidateVoteAddress(ctx context.Context, in *ValidateVoteAddressRequest, opts ...grpc.CallOption) (*ValidateVoteAddressResponse, error) {
	out := new(ValidateVoteAddressResponse)
	err := grpc.Invoke(ctx, "/integratorrpc.VotePoolIntegratorService/ValidateVoteAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votePoolIntegratorServiceClient) ValidatePoolSubsidyAddress(ctx context.Context, in *ValidatePoolSubsidyAddressRequest, opts ...grpc.CallOption) (*ValidatePoolSubsidyAddressResponse, error) {
	out := new(ValidatePoolSubsidyAddressResponse)
	err := grpc.Invoke(ctx, "/integratorrpc.VotePoolIntegratorService/ValidatePoolSubsidyAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VotePoolIntegratorService service

type VotePoolIntegratorServiceServer interface {
	ValidateVoteAddress(context.Context, *ValidateVoteAddressRequest) (*ValidateVoteAddressResponse, error)
	ValidatePoolSubsidyAddress(context.Context, *ValidatePoolSubsidyAddressRequest) (*ValidatePoolSubsidyAddressResponse, error)
}

func RegisterVotePoolIntegratorServiceServer(s *grpc.Server, srv VotePoolIntegratorServiceServer) {
	s.RegisterService(&_VotePoolIntegratorService_serviceDesc, srv)
}

func _VotePoolIntegratorService_ValidateVoteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateVoteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotePoolIntegratorServiceServer).ValidateVoteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integratorrpc.VotePoolIntegratorService/ValidateVoteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotePoolIntegratorServiceServer).ValidateVoteAddress(ctx, req.(*ValidateVoteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VotePoolIntegratorService_ValidatePoolSubsidyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePoolSubsidyAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotePoolIntegratorServiceServer).ValidatePoolSubsidyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integratorrpc.VotePoolIntegratorService/ValidatePoolSubsidyAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotePoolIntegratorServiceServer).ValidatePoolSubsidyAddress(ctx, req.(*ValidatePoolSubsidyAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VotePoolIntegratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "integratorrpc.VotePoolIntegratorService",
	HandlerType: (*VotePoolIntegratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateVoteAddress",
			Handler:    _VotePoolIntegratorService_ValidateVoteAddress_Handler,
		},
		{
			MethodName: "ValidatePoolSubsidyAddress",
			Handler:    _VotePoolIntegratorService_ValidatePoolSubsidyAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrator-api.proto",
}

func init() { proto.RegisterFile("integrator-api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xcc, 0x2b, 0x49,
	0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x4d, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x45, 0x88, 0x16, 0x15, 0x24, 0x2b, 0x99, 0x71, 0x49, 0x85, 0x25, 0xe6, 0x64, 0xa6,
	0x24, 0x96, 0xa4, 0x86, 0xe5, 0x97, 0xa4, 0x3a, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x42, 0x44, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x63, 0x2e, 0x69, 0xac, 0xfa, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0xa0, 0xda, 0x20, 0x1c, 0x25, 0x5b,
	0x2e, 0x45, 0x98, 0xa6, 0x80, 0xfc, 0xfc, 0x9c, 0xe0, 0xd2, 0xa4, 0xe2, 0xcc, 0x94, 0x4a, 0xa2,
	0xed, 0xb4, 0xe2, 0x52, 0xc2, 0xa7, 0x1d, 0x9f, 0xd5, 0x46, 0xbd, 0x4c, 0x5c, 0x92, 0x20, 0x87,
	0x82, 0x34, 0x7a, 0xc2, 0x43, 0x20, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x87, 0x4b,
	0x18, 0x8b, 0x6f, 0x84, 0x34, 0xf5, 0x50, 0x02, 0x4b, 0x0f, 0x77, 0x48, 0x49, 0x69, 0x11, 0xa3,
	0x14, 0xea, 0xc2, 0x46, 0x46, 0x44, 0xa0, 0x63, 0x7a, 0x44, 0xc8, 0x00, 0x87, 0x51, 0x38, 0x83,
	0x4c, 0xca, 0x90, 0x04, 0x1d, 0x10, 0x37, 0x24, 0xb1, 0x81, 0x53, 0x83, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x21, 0x69, 0x45, 0x25, 0x02, 0x00, 0x00,
}
