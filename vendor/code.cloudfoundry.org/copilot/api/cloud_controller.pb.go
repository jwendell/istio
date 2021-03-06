// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloud_controller.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	cloud_controller.proto
	istio.proto

It has these top-level messages:
	AddRequest
	AddResponse
	HealthRequest
	HealthResponse
	RoutesRequest
	RoutesResponse
	BackendSet
	Backend
*/
package api

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

type AddRequest struct {
	ProcessGuid string `protobuf:"bytes,1,opt,name=processGuid" json:"processGuid,omitempty"`
	Hostname    string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetProcessGuid() string {
	if m != nil {
		return m.ProcessGuid
	}
	return ""
}

func (m *AddRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type AddResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "api.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "api.AddResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CloudControllerCopilot service

type CloudControllerCopilotClient interface {
	AddRoute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type cloudControllerCopilotClient struct {
	cc *grpc.ClientConn
}

func NewCloudControllerCopilotClient(cc *grpc.ClientConn) CloudControllerCopilotClient {
	return &cloudControllerCopilotClient{cc}
}

func (c *cloudControllerCopilotClient) AddRoute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/AddRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudControllerCopilot service

type CloudControllerCopilotServer interface {
	AddRoute(context.Context, *AddRequest) (*AddResponse, error)
}

func RegisterCloudControllerCopilotServer(s *grpc.Server, srv CloudControllerCopilotServer) {
	s.RegisterService(&_CloudControllerCopilot_serviceDesc, srv)
}

func _CloudControllerCopilot_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/AddRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).AddRoute(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudControllerCopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CloudControllerCopilot",
	HandlerType: (*CloudControllerCopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRoute",
			Handler:    _CloudControllerCopilot_AddRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud_controller.proto",
}

func init() { proto.RegisterFile("cloud_controller.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x8a, 0xc2, 0x30,
	0x10, 0x86, 0xb7, 0xbb, 0xb0, 0xd6, 0xe9, 0x41, 0xc9, 0xa1, 0x94, 0x9e, 0x4a, 0x2f, 0x7a, 0xaa,
	0xa0, 0x4f, 0x20, 0x3d, 0x88, 0x1e, 0xfb, 0x02, 0x52, 0x93, 0x01, 0x03, 0xb1, 0x13, 0x33, 0xc9,
	0xfb, 0x4b, 0x23, 0xb6, 0x1e, 0xff, 0x7f, 0x98, 0x6f, 0xbe, 0x81, 0x5c, 0x1a, 0x0a, 0xea, 0x2a,
	0x69, 0xf0, 0x8e, 0x8c, 0x41, 0xd7, 0x58, 0x47, 0x9e, 0xc4, 0x5f, 0x6f, 0x75, 0x7d, 0x01, 0x38,
	0x2a, 0xd5, 0xe1, 0x33, 0x20, 0x7b, 0x51, 0x41, 0x66, 0x1d, 0x49, 0x64, 0x3e, 0x05, 0xad, 0x8a,
	0xa4, 0x4a, 0xb6, 0xcb, 0xee, 0xbb, 0x12, 0x25, 0xa4, 0x77, 0x62, 0x3f, 0xf4, 0x0f, 0x2c, 0x7e,
	0xe3, 0x78, 0xca, 0xf5, 0x06, 0xb2, 0xc8, 0x62, 0x4b, 0x03, 0xa3, 0x28, 0x60, 0xc1, 0x41, 0x8e,
	0x9b, 0x11, 0x94, 0x76, 0x9f, 0xb8, 0x3f, 0x43, 0xde, 0x8e, 0x4e, 0xed, 0xa4, 0xd4, 0x92, 0xd5,
	0x86, 0xbc, 0xd8, 0x41, 0x3a, 0x22, 0x28, 0x78, 0x14, 0xab, 0xa6, 0xb7, 0xba, 0x99, 0xed, 0xca,
	0xf5, 0x5c, 0xbc, 0x4f, 0xd4, 0x3f, 0xb7, 0xff, 0xf8, 0xcb, 0xe1, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x10, 0x3d, 0x37, 0x22, 0xe5, 0x00, 0x00, 0x00,
}
