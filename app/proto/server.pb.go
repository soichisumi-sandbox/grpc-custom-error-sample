// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/server.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/server.proto

It has these top-level messages:
	GetDataRequest
	GetDataResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GetDataRequest struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *GetDataRequest) Reset()                    { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()               {}
func (*GetDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetDataRequest) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetDataResponse struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
	Num uint32 `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (m *GetDataResponse) Reset()                    { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()               {}
func (*GetDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetDataResponse) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *GetDataResponse) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto1.RegisterType((*GetDataRequest)(nil), "proto.GetDataRequest")
	proto1.RegisterType((*GetDataResponse)(nil), "proto.GetDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Server service

type ServerClient interface {
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := grpc.Invoke(ctx, "/proto.Server/GetData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _Server_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server.proto",
}

func init() { proto1.RegisterFile("proto/server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x94,
	0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e,
	0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x91, 0x92, 0x16, 0x17, 0x9f, 0x7b, 0x6a, 0x89,
	0x4b, 0x62, 0x49, 0x62, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0x7b, 0x71,
	0x69, 0x72, 0x72, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x8c, 0xab, 0x64,
	0xca, 0xc5, 0x0f, 0x57, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc0, 0xc5, 0x5c, 0x5c,
	0x52, 0x04, 0x56, 0xc8, 0x19, 0x04, 0x62, 0x82, 0x44, 0xf2, 0x4a, 0x73, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x78, 0x83, 0x40, 0x4c, 0xa3, 0x40, 0x2e, 0xb6, 0x60, 0xb0, 0xbb, 0x84, 0xdc, 0xb9, 0xd8,
	0xa1, 0x06, 0x08, 0x89, 0x42, 0xec, 0xd7, 0x43, 0xb5, 0x5c, 0x4a, 0x0c, 0x5d, 0x18, 0x62, 0x8f,
	0x12, 0x6f, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xd8, 0x85, 0x58, 0xf5, 0x53, 0x12, 0x4b, 0x12, 0x93,
	0xd8, 0xc0, 0xaa, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x5a, 0x04, 0xc4, 0xf7, 0x00,
	0x00, 0x00,
}
