// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/internal/proto/examplepb/stream.proto

package examplepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	sub "github.com/grpc-ecosystem/grpc-gateway/examples/internal/proto/sub"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("examples/internal/proto/examplepb/stream.proto", fileDescriptor_cc5dba844cf4f624)
}

var fileDescriptor_cc5dba844cf4f624 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xcf, 0xaa, 0x13, 0x31,
	0x14, 0xc6, 0x89, 0x5c, 0xe4, 0x3a, 0xe2, 0x26, 0x5c, 0x2e, 0x38, 0x16, 0xc4, 0x41, 0xb0, 0x75,
	0x91, 0xb4, 0xba, 0xab, 0x2b, 0x47, 0x0b, 0x2e, 0x14, 0x17, 0xdd, 0x88, 0x9b, 0x92, 0xcc, 0x9c,
	0xce, 0x84, 0xce, 0x24, 0x21, 0x39, 0xd3, 0x3a, 0x5b, 0xc1, 0x27, 0xe8, 0x33, 0xb8, 0xf6, 0x61,
	0x7c, 0x05, 0x1f, 0x44, 0x3a, 0xff, 0xe8, 0xc2, 0xd2, 0xca, 0x5d, 0xe6, 0x9c, 0x7c, 0xdf, 0xf9,
	0xbe, 0x5f, 0xc0, 0xe0, 0x9b, 0x28, 0x6d, 0x01, 0x9e, 0x2b, 0x8d, 0xe0, 0xb4, 0x28, 0xb8, 0x75,
	0x06, 0x0d, 0xef, 0xe6, 0x56, 0x72, 0x8f, 0x0e, 0x44, 0xc9, 0x9a, 0x31, 0x1d, 0x67, 0xce, 0x26,
	0x2c, 0x13, 0x08, 0x3b, 0x51, 0x0f, 0x62, 0xd6, 0x8b, 0xd9, 0x20, 0x0b, 0x47, 0x99, 0x31, 0x59,
	0x01, 0x5c, 0x58, 0xc5, 0x85, 0xd6, 0x06, 0x05, 0x2a, 0xa3, 0x7d, 0xeb, 0x13, 0x3e, 0x3e, 0xda,
	0xe6, 0x88, 0x56, 0x9a, 0xb4, 0xee, 0x56, 0x4f, 0xba, 0x55, 0xf3, 0x92, 0xd5, 0x9a, 0x43, 0x69,
	0xb1, 0x5f, 0xbe, 0x39, 0x9f, 0x57, 0xac, 0xa4, 0xc2, 0x95, 0x59, 0xaf, 0x60, 0x0b, 0xae, 0xc6,
	0x5c, 0xe9, 0xac, 0x13, 0x4f, 0x4e, 0x89, 0x7d, 0x25, 0x79, 0x09, 0xde, 0x8b, 0x0c, 0xda, 0xaf,
	0xaf, 0x7e, 0x5d, 0x05, 0x8f, 0x96, 0x4d, 0xf1, 0x25, 0xb8, 0xad, 0x4a, 0x80, 0xee, 0x49, 0x10,
	0xc4, 0x55, 0xb1, 0x79, 0xe7, 0x40, 0x20, 0xd0, 0x39, 0xbb, 0x94, 0x04, 0x7b, 0x1b, 0x2b, 0xfc,
	0xbc, 0x5e, 0x0c, 0x69, 0xc2, 0x5b, 0xd6, 0x56, 0x64, 0x7d, 0x45, 0xb6, 0x38, 0x54, 0x8c, 0xf8,
	0xf7, 0xdf, 0x7f, 0xf6, 0xf7, 0x26, 0xd1, 0x73, 0xbe, 0x9d, 0xf5, 0x85, 0xfe, 0x55, 0x87, 0xcb,
	0xaa, 0xd8, 0xcc, 0xc9, 0xcb, 0x31, 0xa1, 0x3f, 0x48, 0x70, 0xf5, 0x51, 0x79, 0xa4, 0x27, 0x3c,
	0xc3, 0x3b, 0xe4, 0x8c, 0x5e, 0x34, 0x79, 0x9e, 0xd1, 0xa7, 0x67, 0xf2, 0x4c, 0x09, 0xfd, 0x49,
	0x82, 0xeb, 0x03, 0x9d, 0x45, 0x92, 0x1b, 0x3a, 0x3b, 0x77, 0xd3, 0x57, 0x92, 0x2d, 0xd1, 0x29,
	0x9d, 0x7d, 0x6a, 0xa9, 0x87, 0xff, 0x2f, 0xb9, 0x9c, 0x16, 0x24, 0xb9, 0x69, 0x68, 0x4d, 0x09,
	0xfd, 0x12, 0x5c, 0xbf, 0x37, 0x3b, 0x5d, 0x18, 0x91, 0x9e, 0x44, 0x76, 0xd3, 0xcf, 0x85, 0x55,
	0xec, 0x03, 0xa2, 0x8d, 0x4d, 0x5a, 0x47, 0xa3, 0xe6, 0xdc, 0x2d, 0xbd, 0x39, 0x3e, 0x97, 0x76,
	0x5e, 0x53, 0x12, 0x3f, 0xfc, 0xfa, 0x60, 0x40, 0x29, 0xef, 0x37, 0x96, 0xaf, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x74, 0xe1, 0x34, 0xfd, 0x5f, 0x03, 0x00, 0x00,
}

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
	Download(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_DownloadClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/grpc.gateway.examples.internal.examplepb.StreamService/BulkCreate", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/grpc.gateway.examples.internal.examplepb.StreamService/List", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/grpc.gateway.examples.internal.examplepb.StreamService/BulkEcho", opts...)
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

func (c *streamServiceClient) Download(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[3], "/grpc.gateway.examples.internal.examplepb.StreamService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_DownloadClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type streamServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *streamServiceDownloadClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
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
	Download(*empty.Empty, StreamService_DownloadServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) BulkCreate(srv StreamService_BulkCreateServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkCreate not implemented")
}
func (*UnimplementedStreamServiceServer) List(req *empty.Empty, srv StreamService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStreamServiceServer) BulkEcho(srv StreamService_BulkEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkEcho not implemented")
}
func (*UnimplementedStreamServiceServer) Download(req *empty.Empty, srv StreamService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
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

func _StreamService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).Download(m, &streamServiceDownloadServer{stream})
}

type StreamService_DownloadServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type streamServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *streamServiceDownloadServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.examplepb.StreamService",
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
		{
			StreamName:    "Download",
			Handler:       _StreamService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examples/internal/proto/examplepb/stream.proto",
}
