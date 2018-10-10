// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/examplepb/stream.proto

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import sub "github.com/grpc-ecosystem/grpc-gateway/examples/proto/sub"
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

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/grpc.gateway.examples.examplepb.StreamService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkCreateClient{stream}
	return x, nil
}

type StreamService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type streamServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkCreateClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/grpc.gateway.examples.examplepb.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/grpc.gateway.examples.examplepb.StreamService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkEchoClient{stream}
	return x, nil
}

type StreamService_BulkEchoClient interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ClientStream
}

type streamServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkEchoClient) Send(m *sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkEchoClient) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	BulkCreate(StreamService_BulkCreateServer) error
	List(*empty.Empty, StreamService_ListServer) error
	BulkEcho(StreamService_BulkEchoServer) error
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkCreate(&streamServiceBulkCreateServer{stream})
}

type StreamService_BulkCreateServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type streamServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkCreateServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkEcho(&streamServiceBulkEchoServer{stream})
}

type StreamService_BulkEchoServer interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ServerStream
}

type streamServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkEchoServer) Send(m *sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkEchoServer) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _StreamService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _StreamService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/proto/examplepb/stream.proto",
}

func init() {
	proto.RegisterFile("examples/proto/examplepb/stream.proto", fileDescriptor_stream_c8bbbb01c5cf053e)
}

var fileDescriptor_stream_c8bbbb01c5cf053e = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x40, 0x08, 0x8c, 0x58, 0x3c, 0x30, 0x04, 0xa4, 0x42, 0x05, 0xa2, 0x30, 0xd8,
	0x6d, 0xd9, 0xd8, 0x28, 0xea, 0x06, 0x62, 0xe8, 0xc6, 0x52, 0xd9, 0xd1, 0xd5, 0xb5, 0x9a, 0xc4,
	0x96, 0x7d, 0x29, 0x54, 0x62, 0x62, 0x64, 0xed, 0x8b, 0xf0, 0x2e, 0xbc, 0x02, 0x0f, 0x82, 0xea,
	0xa4, 0x1d, 0x10, 0x51, 0xcb, 0x78, 0xbe, 0xff, 0xf7, 0xfd, 0xdf, 0x4f, 0x2e, 0xe0, 0x55, 0x64,
	0x36, 0x05, 0xcf, 0xad, 0x33, 0x68, 0x78, 0x35, 0x5a, 0xc9, 0x3d, 0x3a, 0x10, 0x19, 0x0b, 0xcf,
	0xb4, 0xa1, 0x9c, 0x4d, 0x98, 0x12, 0x08, 0x2f, 0x62, 0xc6, 0x96, 0x1e, 0xb6, 0x52, 0xc7, 0x27,
	0xca, 0x18, 0x95, 0x02, 0x17, 0x56, 0x73, 0x91, 0xe7, 0x06, 0x05, 0x6a, 0x93, 0xfb, 0xd2, 0x1e,
	0x1f, 0x57, 0xdb, 0x30, 0xc9, 0x62, 0xc4, 0x21, 0xb3, 0x38, 0xab, 0x96, 0xdd, 0xda, 0x08, 0x62,
	0x28, 0x35, 0x0e, 0xcd, 0x68, 0x08, 0x53, 0x70, 0x33, 0x1c, 0xeb, 0x5c, 0x55, 0x9e, 0xd3, 0x5f,
	0x1e, 0x5f, 0x48, 0x9e, 0x81, 0xf7, 0x42, 0x41, 0xa9, 0xe8, 0x7e, 0x6e, 0x93, 0xc3, 0x41, 0x40,
	0x18, 0x80, 0x9b, 0xea, 0x04, 0xe8, 0x47, 0x44, 0x48, 0xaf, 0x48, 0x27, 0xf7, 0x0e, 0x04, 0x02,
	0xed, 0xb0, 0x35, 0x4c, 0xec, 0xae, 0xa7, 0xf1, 0x69, 0xd4, 0x5f, 0xdd, 0x8e, 0x8f, 0x58, 0xc9,
	0xc1, 0x96, 0x1c, 0xac, 0xbf, 0xe0, 0x68, 0xf2, 0xf7, 0xaf, 0xef, 0xf9, 0xd6, 0x55, 0xf3, 0x9c,
	0x4f, 0x3b, 0xcb, 0xf8, 0x7f, 0x85, 0xe7, 0xb2, 0x48, 0x27, 0xb7, 0xd1, 0x75, 0x2b, 0xa2, 0x6f,
	0x64, 0xe7, 0x41, 0x7b, 0xa4, 0x35, 0x5f, 0xc6, 0xff, 0x4f, 0xd7, 0xbc, 0x0c, 0x29, 0xce, 0x68,
	0x63, 0x4d, 0x8a, 0x76, 0x44, 0xe7, 0x11, 0xd9, 0x5b, 0x54, 0xd1, 0x4f, 0xc6, 0x86, 0xb6, 0x6a,
	0x4e, 0xf9, 0x42, 0xb2, 0x01, 0x3a, 0x9d, 0xab, 0xc7, 0xb2, 0xd9, 0x78, 0x63, 0xe5, 0xe6, 0x8d,
	0x40, 0x32, 0x36, 0xa1, 0x91, 0x76, 0xd4, 0x3b, 0x78, 0xde, 0x5f, 0xe1, 0xc9, 0xdd, 0x50, 0xc8,
	0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x63, 0x7a, 0xd9, 0xa1, 0x02, 0x00, 0x00,
}
