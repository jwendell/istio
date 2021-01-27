// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/contextgraph/v1alpha1/contextgraph.proto

package contextgraph

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContextGraphService service

type ContextGraphServiceClient interface {
	// Make a lot of assertions in one RPC.
	//
	// Empty batches will be rejected (ones with zero assertions).
	//
	// The timeline for each entity / relationship can be imagined as a sparse
	// array mapping timestamps to assertions.  When AssertBatch is called, the
	// assertions in the batch are written to this timeline, possibly overwriting
	// existing assertions for the same entities / relationships at the given
	// timestamps.
	//
	// If a given entity / relationship is asserted twice at the same timestamp,
	// whether in the same batch or not, it is undefined which one is stored in
	// the timeline.
	AssertBatch(ctx context.Context, in *AssertBatchRequest, opts ...grpc.CallOption) (*AssertBatchResponse, error)
}

type contextGraphServiceClient struct {
	cc *grpc.ClientConn
}

func NewContextGraphServiceClient(cc *grpc.ClientConn) ContextGraphServiceClient {
	return &contextGraphServiceClient{cc}
}

func (c *contextGraphServiceClient) AssertBatch(ctx context.Context, in *AssertBatchRequest, opts ...grpc.CallOption) (*AssertBatchResponse, error) {
	out := new(AssertBatchResponse)
	err := grpc.Invoke(ctx, "/google.cloud.contextgraph.v1alpha1.ContextGraphService/AssertBatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContextGraphService service

type ContextGraphServiceServer interface {
	// Make a lot of assertions in one RPC.
	//
	// Empty batches will be rejected (ones with zero assertions).
	//
	// The timeline for each entity / relationship can be imagined as a sparse
	// array mapping timestamps to assertions.  When AssertBatch is called, the
	// assertions in the batch are written to this timeline, possibly overwriting
	// existing assertions for the same entities / relationships at the given
	// timestamps.
	//
	// If a given entity / relationship is asserted twice at the same timestamp,
	// whether in the same batch or not, it is undefined which one is stored in
	// the timeline.
	AssertBatch(context.Context, *AssertBatchRequest) (*AssertBatchResponse, error)
}

func RegisterContextGraphServiceServer(s *grpc.Server, srv ContextGraphServiceServer) {
	s.RegisterService(&_ContextGraphService_serviceDesc, srv)
}

func _ContextGraphService_AssertBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssertBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextGraphServiceServer).AssertBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.contextgraph.v1alpha1.ContextGraphService/AssertBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextGraphServiceServer).AssertBatch(ctx, req.(*AssertBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContextGraphService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.contextgraph.v1alpha1.ContextGraphService",
	HandlerType: (*ContextGraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssertBatch",
			Handler:    _ContextGraphService_AssertBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/contextgraph/v1alpha1/contextgraph.proto",
}

func init() {
	proto.RegisterFile("google/cloud/contextgraph/v1alpha1/contextgraph.proto", fileDescriptor_contextgraph_904354ee905bfb2b)
}

var fileDescriptor_contextgraph_904354ee905bfb2b = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0x31, 0x4e, 0xf3, 0x40,
	0x10, 0x85, 0x95, 0xbf, 0xf8, 0x8b, 0xd0, 0x05, 0xd1, 0x58, 0x14, 0x28, 0x25, 0x85, 0xad, 0x80,
	0x48, 0x20, 0xa1, 0x21, 0x14, 0x74, 0x08, 0x41, 0x47, 0xb3, 0x9a, 0xd8, 0xa3, 0xf5, 0x4a, 0xce,
	0x8e, 0xd9, 0x9d, 0x04, 0xd2, 0x72, 0x05, 0x2e, 0xc0, 0x31, 0xb8, 0x07, 0x57, 0xe0, 0x20, 0xd8,
	0xde, 0x0d, 0x49, 0x40, 0x82, 0x4d, 0xbb, 0xf3, 0xbe, 0x99, 0x37, 0x9e, 0xe7, 0xf6, 0x89, 0x24,
	0x92, 0x05, 0x26, 0x69, 0x41, 0xb3, 0x2c, 0x49, 0x49, 0x33, 0x3e, 0xb1, 0x34, 0x50, 0xe6, 0xc9,
	0xbc, 0x07, 0x45, 0x99, 0x43, 0x6f, 0xe3, 0x35, 0x2e, 0x0d, 0x31, 0x75, 0xba, 0x0e, 0x8b, 0x1b,
	0x2c, 0xde, 0x10, 0x2c, 0xb1, 0x68, 0xdf, 0xb7, 0x86, 0x52, 0x25, 0xa0, 0x35, 0x31, 0xb0, 0x22,
	0x6d, 0x5d, 0x87, 0xa8, 0x1f, 0x30, 0x18, 0xb2, 0x4c, 0xd4, 0x15, 0x50, 0x1a, 0x8d, 0xe7, 0x42,
	0x0c, 0x83, 0xb5, 0x68, 0x58, 0x4c, 0x80, 0x53, 0x6f, 0x38, 0x1a, 0x05, 0x60, 0x8d, 0x5e, 0x48,
	0x64, 0x81, 0x9a, 0x15, 0x2b, 0x5c, 0x7a, 0x3d, 0x0b, 0x80, 0x33, 0x2c, 0x90, 0xf1, 0x87, 0xdd,
	0xd3, 0x00, 0xb4, 0x50, 0x96, 0x57, 0xe0, 0x72, 0xe8, 0x79, 0x00, 0xf9, 0x30, 0x43, 0xb3, 0x10,
	0xbe, 0x26, 0xd6, 0x0e, 0x14, 0x0d, 0x82, 0xe9, 0x6f, 0xbb, 0x8e, 0x82, 0x41, 0x8d, 0x4a, 0xe6,
	0x13, 0x32, 0x39, 0x51, 0xe6, 0xe1, 0x61, 0x00, 0x6c, 0xe0, 0xd1, 0x39, 0x15, 0xd5, 0x91, 0xe6,
	0x2a, 0xc5, 0x2d, 0x02, 0x51, 0xb3, 0x4a, 0x8b, 0x29, 0x4e, 0xc9, 0x2c, 0x1c, 0x77, 0xf4, 0xd6,
	0x6a, 0xef, 0x5e, 0x3a, 0xf5, 0x55, 0xad, 0xbe, 0x73, 0x5d, 0x3b, 0xaf, 0xad, 0xf6, 0xce, 0x45,
	0x13, 0x84, 0x71, 0x7d, 0xd7, 0x4e, 0x3f, 0xfe, 0x3b, 0xb3, 0xf1, 0x1a, 0x70, 0x8b, 0xd5, 0x9a,
	0x96, 0xa3, 0xc1, 0xd6, 0x9c, 0x2d, 0xab, 0x98, 0x63, 0xf7, 0xe0, 0xf9, 0xfd, 0xe3, 0xe5, 0x5f,
	0xd4, 0xdd, 0xfb, 0xf2, 0x3f, 0x84, 0x95, 0x6c, 0xd8, 0x3a, 0x1c, 0xdf, 0xdc, 0x5f, 0xfb, 0xde,
	0x92, 0x0a, 0xd0, 0x32, 0x26, 0x23, 0x13, 0x89, 0xba, 0x59, 0x2c, 0x71, 0xa5, 0xea, 0xef, 0xb1,
	0xbf, 0x7d, 0x93, 0xd1, 0xfa, 0xeb, 0xe4, 0x7f, 0x83, 0x1e, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xf9, 0x6a, 0xd3, 0x08, 0xd7, 0x03, 0x00, 0x00,
}