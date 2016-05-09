// Code generated by protoc-gen-go.
// source: examples/examplepb/a_bit_of_everything.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import gengo_grpc_gateway_examples_sub "github.com/gengo/grpc-gateway/examples/sub"
import sub2 "github.com/gengo/grpc-gateway/examples/sub2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NumericEnum int32

const (
	NumericEnum_ZERO NumericEnum = 0
	NumericEnum_ONE  NumericEnum = 1
)

var NumericEnum_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
}
var NumericEnum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
}

func (x NumericEnum) String() string {
	return proto.EnumName(NumericEnum_name, int32(x))
}
func (NumericEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ABitOfEverything_Nested_DeepEnum int32

const (
	ABitOfEverything_Nested_FALSE ABitOfEverything_Nested_DeepEnum = 0
	ABitOfEverything_Nested_TRUE  ABitOfEverything_Nested_DeepEnum = 1
)

var ABitOfEverything_Nested_DeepEnum_name = map[int32]string{
	0: "FALSE",
	1: "TRUE",
}
var ABitOfEverything_Nested_DeepEnum_value = map[string]int32{
	"FALSE": 0,
	"TRUE":  1,
}

func (x ABitOfEverything_Nested_DeepEnum) String() string {
	return proto.EnumName(ABitOfEverything_Nested_DeepEnum_name, int32(x))
}
func (ABitOfEverything_Nested_DeepEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0, 0}
}

type ABitOfEverything struct {
	Uuid         string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Nested       []*ABitOfEverything_Nested `protobuf:"bytes,2,rep,name=nested" json:"nested,omitempty"`
	FloatValue   float32                    `protobuf:"fixed32,3,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue  float64                    `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	Int64Value   int64                      `protobuf:"varint,5,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Uint64Value  uint64                     `protobuf:"varint,6,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	Int32Value   int32                      `protobuf:"varint,7,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Fixed64Value uint64                     `protobuf:"fixed64,8,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
	Fixed32Value uint32                     `protobuf:"fixed32,9,opt,name=fixed32_value,json=fixed32Value" json:"fixed32_value,omitempty"`
	BoolValue    bool                       `protobuf:"varint,10,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	StringValue  string                     `protobuf:"bytes,11,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	// TODO(yugui) add bytes_value
	Uint32Value         uint32      `protobuf:"varint,13,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	EnumValue           NumericEnum `protobuf:"varint,14,opt,name=enum_value,json=enumValue,enum=gengo.grpc.gateway.examples.examplepb.NumericEnum" json:"enum_value,omitempty"`
	Sfixed32Value       int32       `protobuf:"fixed32,15,opt,name=sfixed32_value,json=sfixed32Value" json:"sfixed32_value,omitempty"`
	Sfixed64Value       int64       `protobuf:"fixed64,16,opt,name=sfixed64_value,json=sfixed64Value" json:"sfixed64_value,omitempty"`
	Sint32Value         int32       `protobuf:"zigzag32,17,opt,name=sint32_value,json=sint32Value" json:"sint32_value,omitempty"`
	Sint64Value         int64       `protobuf:"zigzag64,18,opt,name=sint64_value,json=sint64Value" json:"sint64_value,omitempty"`
	RepeatedStringValue []string    `protobuf:"bytes,19,rep,name=repeated_string_value,json=repeatedStringValue" json:"repeated_string_value,omitempty"`
	// Types that are valid to be assigned to OneofValue:
	//	*ABitOfEverything_OneofEmpty
	//	*ABitOfEverything_OneofString
	OneofValue        isABitOfEverything_OneofValue       `protobuf_oneof:"oneof_value"`
	MapValue          map[string]NumericEnum              `protobuf:"bytes,22,rep,name=map_value,json=mapValue" json:"map_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=gengo.grpc.gateway.examples.examplepb.NumericEnum"`
	MappedStringValue map[string]string                   `protobuf:"bytes,23,rep,name=mapped_string_value,json=mappedStringValue" json:"mapped_string_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MappedNestedValue map[string]*ABitOfEverything_Nested `protobuf:"bytes,24,rep,name=mapped_nested_value,json=mappedNestedValue" json:"mapped_nested_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ABitOfEverything) Reset()                    { *m = ABitOfEverything{} }
func (m *ABitOfEverything) String() string            { return proto.CompactTextString(m) }
func (*ABitOfEverything) ProtoMessage()               {}
func (*ABitOfEverything) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isABitOfEverything_OneofValue interface {
	isABitOfEverything_OneofValue()
}

type ABitOfEverything_OneofEmpty struct {
	OneofEmpty *EmptyMessage `protobuf:"bytes,20,opt,name=oneof_empty,json=oneofEmpty,oneof"`
}
type ABitOfEverything_OneofString struct {
	OneofString string `protobuf:"bytes,21,opt,name=oneof_string,json=oneofString,oneof"`
}

func (*ABitOfEverything_OneofEmpty) isABitOfEverything_OneofValue()  {}
func (*ABitOfEverything_OneofString) isABitOfEverything_OneofValue() {}

func (m *ABitOfEverything) GetOneofValue() isABitOfEverything_OneofValue {
	if m != nil {
		return m.OneofValue
	}
	return nil
}

func (m *ABitOfEverything) GetNested() []*ABitOfEverything_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *ABitOfEverything) GetOneofEmpty() *EmptyMessage {
	if x, ok := m.GetOneofValue().(*ABitOfEverything_OneofEmpty); ok {
		return x.OneofEmpty
	}
	return nil
}

func (m *ABitOfEverything) GetOneofString() string {
	if x, ok := m.GetOneofValue().(*ABitOfEverything_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (m *ABitOfEverything) GetMapValue() map[string]NumericEnum {
	if m != nil {
		return m.MapValue
	}
	return nil
}

func (m *ABitOfEverything) GetMappedStringValue() map[string]string {
	if m != nil {
		return m.MappedStringValue
	}
	return nil
}

func (m *ABitOfEverything) GetMappedNestedValue() map[string]*ABitOfEverything_Nested {
	if m != nil {
		return m.MappedNestedValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ABitOfEverything) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ABitOfEverything_OneofMarshaler, _ABitOfEverything_OneofUnmarshaler, _ABitOfEverything_OneofSizer, []interface{}{
		(*ABitOfEverything_OneofEmpty)(nil),
		(*ABitOfEverything_OneofString)(nil),
	}
}

func _ABitOfEverything_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ABitOfEverything)
	// oneof_value
	switch x := m.OneofValue.(type) {
	case *ABitOfEverything_OneofEmpty:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneofEmpty); err != nil {
			return err
		}
	case *ABitOfEverything_OneofString:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OneofString)
	case nil:
	default:
		return fmt.Errorf("ABitOfEverything.OneofValue has unexpected type %T", x)
	}
	return nil
}

func _ABitOfEverything_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ABitOfEverything)
	switch tag {
	case 20: // oneof_value.oneof_empty
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EmptyMessage)
		err := b.DecodeMessage(msg)
		m.OneofValue = &ABitOfEverything_OneofEmpty{msg}
		return true, err
	case 21: // oneof_value.oneof_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OneofValue = &ABitOfEverything_OneofString{x}
		return true, err
	default:
		return false, nil
	}
}

func _ABitOfEverything_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ABitOfEverything)
	// oneof_value
	switch x := m.OneofValue.(type) {
	case *ABitOfEverything_OneofEmpty:
		s := proto.Size(x.OneofEmpty)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ABitOfEverything_OneofString:
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OneofString)))
		n += len(x.OneofString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ABitOfEverything_Nested struct {
	Name   string                           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Amount uint32                           `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Ok     ABitOfEverything_Nested_DeepEnum `protobuf:"varint,3,opt,name=ok,enum=gengo.grpc.gateway.examples.examplepb.ABitOfEverything_Nested_DeepEnum" json:"ok,omitempty"`
}

func (m *ABitOfEverything_Nested) Reset()                    { *m = ABitOfEverything_Nested{} }
func (m *ABitOfEverything_Nested) String() string            { return proto.CompactTextString(m) }
func (*ABitOfEverything_Nested) ProtoMessage()               {}
func (*ABitOfEverything_Nested) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*ABitOfEverything)(nil), "gengo.grpc.gateway.examples.examplepb.ABitOfEverything")
	proto.RegisterType((*ABitOfEverything_Nested)(nil), "gengo.grpc.gateway.examples.examplepb.ABitOfEverything.Nested")
	proto.RegisterType((*EmptyMessage)(nil), "gengo.grpc.gateway.examples.examplepb.EmptyMessage")
	proto.RegisterEnum("gengo.grpc.gateway.examples.examplepb.NumericEnum", NumericEnum_name, NumericEnum_value)
	proto.RegisterEnum("gengo.grpc.gateway.examples.examplepb.ABitOfEverything_Nested_DeepEnum", ABitOfEverything_Nested_DeepEnum_name, ABitOfEverything_Nested_DeepEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for ABitOfEverythingService service

type ABitOfEverythingServiceClient interface {
	Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkCreateClient, error)
	Lookup(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error)
	List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (ABitOfEverythingService_ListClient, error)
	Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*EmptyMessage, error)
	Delete(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	Echo(ctx context.Context, in *gengo_grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkEchoClient, error)
}

type aBitOfEverythingServiceClient struct {
	cc *grpc.ClientConn
}

func NewABitOfEverythingServiceClient(cc *grpc.ClientConn) ABitOfEverythingServiceClient {
	return &aBitOfEverythingServiceClient{cc}
}

func (c *aBitOfEverythingServiceClient) Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkCreateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ABitOfEverythingService_serviceDesc.Streams[0], c.cc, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &aBitOfEverythingServiceBulkCreateClient{stream}
	return x, nil
}

type ABitOfEverythingService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*EmptyMessage, error)
	grpc.ClientStream
}

type aBitOfEverythingServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *aBitOfEverythingServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkCreateClient) CloseAndRecv() (*EmptyMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aBitOfEverythingServiceClient) Lookup(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (ABitOfEverythingService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ABitOfEverythingService_serviceDesc.Streams[1], c.cc, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &aBitOfEverythingServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ABitOfEverythingService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type aBitOfEverythingServiceListClient struct {
	grpc.ClientStream
}

func (x *aBitOfEverythingServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aBitOfEverythingServiceClient) Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Delete(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Echo(ctx context.Context, in *gengo_grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	out := new(gengo_grpc_gateway_examples_sub.StringMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkEchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ABitOfEverythingService_serviceDesc.Streams[2], c.cc, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &aBitOfEverythingServiceBulkEchoClient{stream}
	return x, nil
}

type ABitOfEverythingService_BulkEchoClient interface {
	Send(*gengo_grpc_gateway_examples_sub.StringMessage) error
	Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	grpc.ClientStream
}

type aBitOfEverythingServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *aBitOfEverythingServiceBulkEchoClient) Send(m *gengo_grpc_gateway_examples_sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkEchoClient) Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	m := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ABitOfEverythingService service

type ABitOfEverythingServiceServer interface {
	Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	BulkCreate(ABitOfEverythingService_BulkCreateServer) error
	Lookup(context.Context, *sub2.IdMessage) (*ABitOfEverything, error)
	List(*EmptyMessage, ABitOfEverythingService_ListServer) error
	Update(context.Context, *ABitOfEverything) (*EmptyMessage, error)
	Delete(context.Context, *sub2.IdMessage) (*EmptyMessage, error)
	Echo(context.Context, *gengo_grpc_gateway_examples_sub.StringMessage) (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	BulkEcho(ABitOfEverythingService_BulkEchoServer) error
}

func RegisterABitOfEverythingServiceServer(s *grpc.Server, srv ABitOfEverythingServiceServer) {
	s.RegisterService(&_ABitOfEverythingService_serviceDesc, srv)
}

func _ABitOfEverythingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Create(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_CreateBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).CreateBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).CreateBody(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ABitOfEverythingServiceServer).BulkCreate(&aBitOfEverythingServiceBulkCreateServer{stream})
}

type ABitOfEverythingService_BulkCreateServer interface {
	SendAndClose(*EmptyMessage) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type aBitOfEverythingServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *aBitOfEverythingServiceBulkCreateServer) SendAndClose(m *EmptyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ABitOfEverythingService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub2.IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Lookup(ctx, req.(*sub2.IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ABitOfEverythingServiceServer).List(m, &aBitOfEverythingServiceListServer{stream})
}

type ABitOfEverythingService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type aBitOfEverythingServiceListServer struct {
	grpc.ServerStream
}

func (x *aBitOfEverythingServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _ABitOfEverythingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Update(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub2.IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Delete(ctx, req.(*sub2.IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Echo(ctx, req.(*gengo_grpc_gateway_examples_sub.StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ABitOfEverythingServiceServer).BulkEcho(&aBitOfEverythingServiceBulkEchoServer{stream})
}

type ABitOfEverythingService_BulkEchoServer interface {
	Send(*gengo_grpc_gateway_examples_sub.StringMessage) error
	Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	grpc.ServerStream
}

type aBitOfEverythingServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *aBitOfEverythingServiceBulkEchoServer) Send(m *gengo_grpc_gateway_examples_sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkEchoServer) Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	m := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ABitOfEverythingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService",
	HandlerType: (*ABitOfEverythingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ABitOfEverythingService_Create_Handler,
		},
		{
			MethodName: "CreateBody",
			Handler:    _ABitOfEverythingService_CreateBody_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _ABitOfEverythingService_Lookup_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ABitOfEverythingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ABitOfEverythingService_Delete_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _ABitOfEverythingService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _ABitOfEverythingService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _ABitOfEverythingService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _ABitOfEverythingService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

var fileDescriptor1 = []byte{
	// 1110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x57, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xef, 0xd9, 0x8e, 0x6b, 0xef, 0xc5, 0x8e, 0xb3, 0x69, 0x53, 0x63, 0x40, 0x49, 0x8f, 0x16,
	0x82, 0xa9, 0xee, 0xda, 0x0b, 0x02, 0x14, 0x09, 0xa4, 0x9a, 0x18, 0x8a, 0x94, 0x26, 0xe2, 0xd2,
	0x06, 0x29, 0x12, 0xb2, 0xce, 0xf6, 0xc6, 0x3d, 0xc5, 0x77, 0x7b, 0xba, 0x3f, 0xa6, 0x96, 0x95,
	0x3e, 0x54, 0x3c, 0x20, 0xf1, 0x80, 0x04, 0x1f, 0x80, 0x67, 0x90, 0xe0, 0x8b, 0xf4, 0x05, 0x89,
	0xaf, 0xc0, 0x07, 0x61, 0x6e, 0xf7, 0xee, 0xba, 0xe7, 0x44, 0xd8, 0x38, 0xa8, 0xbc, 0xdd, 0xce,
	0xfe, 0xf6, 0x37, 0xbf, 0x99, 0xd9, 0x99, 0xb5, 0xd1, 0x1d, 0xf2, 0xd4, 0xb4, 0xdd, 0x21, 0xf1,
	0xb5, 0xf8, 0xc3, 0xed, 0x6a, 0x66, 0xa7, 0x6b, 0x05, 0x1d, 0x7a, 0xd2, 0x21, 0x23, 0xe2, 0x8d,
	0x83, 0x27, 0x96, 0x33, 0x50, 0x5d, 0x8f, 0x06, 0x14, 0xdf, 0x1e, 0x10, 0x67, 0x40, 0xd5, 0x81,
	0xe7, 0xf6, 0xd4, 0x81, 0x19, 0x90, 0x6f, 0xcc, 0xb1, 0x9a, 0x10, 0xa8, 0x29, 0x41, 0xe3, 0x8d,
	0x01, 0xa5, 0x83, 0x21, 0xd1, 0x4c, 0xd7, 0xd2, 0x4c, 0xc7, 0xa1, 0x81, 0x19, 0x58, 0xd4, 0xf1,
	0x39, 0x49, 0xa3, 0x91, 0xba, 0xf4, 0xc3, 0xae, 0x66, 0x13, 0xdf, 0x37, 0x07, 0x24, 0xde, 0x7b,
	0x5d, 0xdc, 0xd3, 0xb3, 0x9b, 0xca, 0x1f, 0x15, 0x54, 0xbb, 0xdf, 0xb2, 0x82, 0x83, 0x93, 0x76,
	0x2a, 0x0c, 0x63, 0x54, 0x08, 0x43, 0xab, 0x5f, 0x97, 0x36, 0xa5, 0xad, 0xb2, 0xc1, 0xbe, 0xf1,
	0x11, 0x2a, 0x3a, 0xc4, 0x0f, 0x48, 0xbf, 0x9e, 0xdb, 0xcc, 0x6f, 0xc9, 0xfa, 0x27, 0xea, 0x5c,
	0xba, 0xd5, 0x69, 0x72, 0x75, 0x9f, 0xb1, 0x18, 0x31, 0x1b, 0xde, 0x40, 0xf2, 0xc9, 0x90, 0x9a,
	0x41, 0x67, 0x64, 0x0e, 0x43, 0x52, 0xcf, 0x83, 0xcb, 0x9c, 0x81, 0x98, 0xe9, 0x28, 0xb2, 0xe0,
	0x9b, 0x68, 0xb9, 0x4f, 0xc3, 0xee, 0x90, 0xc4, 0x88, 0x02, 0x20, 0x24, 0x43, 0xe6, 0x36, 0x0e,
	0x01, 0x0e, 0xcb, 0x09, 0x3e, 0x78, 0x3f, 0x46, 0x2c, 0x01, 0x22, 0x6f, 0x20, 0x66, 0x4a, 0x39,
	0x42, 0x11, 0x51, 0x04, 0x44, 0xc1, 0x90, 0x43, 0x01, 0xc2, 0x39, 0xb6, 0xf5, 0x18, 0x71, 0x15,
	0x10, 0x4b, 0x8c, 0x63, 0x5b, 0xe7, 0x80, 0xb7, 0x50, 0xe5, 0xc4, 0x7a, 0x4a, 0xfa, 0x29, 0x49,
	0x09, 0x20, 0x45, 0x63, 0x39, 0x36, 0x66, 0x41, 0x29, 0x4f, 0x19, 0x40, 0x57, 0x63, 0x50, 0xc2,
	0xf4, 0x26, 0x42, 0x5d, 0x4a, 0x87, 0x31, 0x02, 0x01, 0xa2, 0x64, 0x94, 0x23, 0x4b, 0x2a, 0xd6,
	0x0f, 0x3c, 0x48, 0x55, 0x0c, 0x90, 0x59, 0x15, 0x64, 0x6e, 0xcb, 0xc4, 0x93, 0x7a, 0xa9, 0x00,
	0xa4, 0xc2, 0xe3, 0x49, 0x9c, 0x7c, 0x89, 0x10, 0x71, 0x42, 0x3b, 0x06, 0x54, 0x01, 0x50, 0xd5,
	0xf5, 0x39, 0x6b, 0xb6, 0x1f, 0xda, 0xc4, 0xb3, 0x7a, 0x6d, 0x38, 0x6f, 0x94, 0x23, 0x16, 0x4e,
	0x79, 0x1b, 0x55, 0xfd, 0x6c, 0x74, 0x2b, 0x40, 0xbb, 0x62, 0x54, 0xfc, 0x4c, 0x78, 0x29, 0x2c,
	0xcd, 0x54, 0x0d, 0x60, 0xb5, 0x04, 0x26, 0xd4, 0xc4, 0x17, 0x63, 0x58, 0x05, 0xd0, 0x2a, 0x84,
	0x29, 0xc4, 0x10, 0x43, 0x52, 0x1e, 0x0c, 0x10, 0xcc, 0x21, 0x09, 0x8b, 0x8e, 0xae, 0x7b, 0xc4,
	0x25, 0x10, 0x4b, 0xbf, 0x93, 0xc9, 0xda, 0x1a, 0xdc, 0xd2, 0xb2, 0xb1, 0x96, 0x6c, 0x1e, 0x0a,
	0xd9, 0x3b, 0x42, 0x32, 0x75, 0x48, 0xd4, 0x8b, 0xb6, 0x1b, 0x8c, 0xeb, 0xd7, 0x80, 0x55, 0xd6,
	0xb7, 0xe7, 0xcc, 0x4d, 0x3b, 0x3a, 0xf3, 0x90, 0xf7, 0xd0, 0x83, 0x2b, 0x06, 0x62, 0x4c, 0xcc,
	0x08, 0xc5, 0x5f, 0xe6, 0xbc, 0x5c, 0x48, 0xfd, 0x7a, 0x54, 0x38, 0xc0, 0x70, 0x6f, 0x5c, 0x01,
	0xee, 0xa2, 0xb2, 0x6d, 0xba, 0xb1, 0xc8, 0x75, 0xd6, 0x4a, 0xed, 0x45, 0x5b, 0xe9, 0xa1, 0xe9,
	0xb2, 0x88, 0xda, 0x4e, 0xe0, 0x8d, 0x8d, 0x92, 0x1d, 0x2f, 0xf1, 0x33, 0xb4, 0x06, 0xdf, 0xee,
	0x74, 0x4a, 0x6e, 0x30, 0x6f, 0xfb, 0x97, 0xf0, 0xe6, 0x66, 0x12, 0xc9, 0xdd, 0xae, 0xda, 0xd3,
	0x76, 0xc1, 0x3f, 0x6f, 0xf2, 0xd8, 0x7f, 0xfd, 0xbf, 0xf0, 0xcf, 0xc7, 0xc7, 0x79, 0xff, 0x82,
	0xbd, 0xf1, 0xbb, 0x84, 0x8a, 0x7c, 0x1d, 0x8d, 0x32, 0xc7, 0xb4, 0x49, 0x32, 0xca, 0xa2, 0x6f,
	0xbc, 0x8e, 0x8a, 0xa6, 0x4d, 0x43, 0x27, 0x80, 0x51, 0x16, 0xf5, 0x4d, 0xbc, 0xc2, 0x5f, 0xa1,
	0x1c, 0x3d, 0x65, 0x13, 0xa8, 0xaa, 0x7f, 0x7e, 0xb9, 0xf1, 0xa6, 0xee, 0x12, 0xe2, 0xb2, 0xfe,
	0x01, 0x4a, 0x65, 0x03, 0x95, 0x92, 0x35, 0x2e, 0xa3, 0xa5, 0xcf, 0xee, 0xef, 0x1d, 0xb6, 0x6b,
	0x57, 0x70, 0x09, 0x15, 0x1e, 0x19, 0x8f, 0xdb, 0x35, 0xa9, 0x41, 0x51, 0x25, 0x53, 0x4b, 0x5c,
	0x43, 0xf9, 0x53, 0x32, 0x8e, 0x55, 0x47, 0x9f, 0xf8, 0x01, 0x5a, 0xe2, 0x59, 0xcc, 0x2d, 0xdc,
	0xca, 0x9c, 0x60, 0x27, 0xf7, 0x91, 0xd4, 0xd8, 0x45, 0xeb, 0x17, 0x97, 0xf3, 0x02, 0xcf, 0xd7,
	0x44, 0xcf, 0x65, 0x91, 0xe5, 0x5b, 0x29, 0xa1, 0x99, 0xae, 0xca, 0x05, 0x34, 0x8f, 0x44, 0x9a,
	0xcb, 0xbf, 0x1f, 0x2f, 0x65, 0xb4, 0x2a, 0x49, 0x3f, 0x33, 0x93, 0x52, 0x45, 0xcb, 0x62, 0x93,
	0x36, 0x37, 0x91, 0x2c, 0x64, 0x21, 0xca, 0xfa, 0x71, 0xdb, 0x38, 0x80, 0xfc, 0x5f, 0x45, 0xf9,
	0x83, 0x7d, 0x48, 0xbf, 0xfe, 0xa2, 0x82, 0x6e, 0x4c, 0xfb, 0x39, 0x24, 0xde, 0xc8, 0xea, 0x11,
	0xfc, 0x43, 0x1e, 0x15, 0x3f, 0xf5, 0xa2, 0x19, 0x82, 0x3f, 0x5c, 0x50, 0x72, 0x63, 0xd1, 0x83,
	0xca, 0x8f, 0xb9, 0xe7, 0x7f, 0xfe, 0xf5, 0x53, 0xee, 0xfb, 0x9c, 0xf2, 0x5d, 0x4e, 0x1b, 0xdd,
	0x4b, 0x7e, 0x4f, 0x5c, 0xf4, 0x6b, 0x42, 0x9b, 0x08, 0xcf, 0xe8, 0x99, 0x36, 0x11, 0xdf, 0x4c,
	0x58, 0x0a, 0x63, 0xf4, 0x4c, 0xf3, 0x89, 0x6b, 0x7a, 0x66, 0x40, 0x3d, 0x6d, 0x12, 0x66, 0x36,
	0x26, 0xc2, 0x40, 0x86, 0x55, 0x66, 0x8a, 0x27, 0x6b, 0x61, 0xff, 0xe5, 0x2b, 0x06, 0x0b, 0x71,
	0xd4, 0x7c, 0x0c, 0x0b, 0xd7, 0x23, 0x80, 0xd7, 0x9a, 0x67, 0xdc, 0x89, 0x70, 0xcc, 0x9f, 0xe6,
	0xf1, 0xa7, 0x1d, 0xf9, 0x53, 0x07, 0x44, 0x91, 0xf8, 0x57, 0x09, 0x21, 0x5e, 0x91, 0x16, 0xed,
	0x8f, 0xff, 0x87, 0xaa, 0x34, 0x59, 0x51, 0x6e, 0x29, 0x1b, 0x33, 0x4a, 0xb2, 0x23, 0x35, 0xf1,
	0x6f, 0x20, 0xb6, 0x15, 0x0e, 0x4f, 0x2f, 0x7b, 0x85, 0x16, 0x79, 0x9e, 0x14, 0x8d, 0x09, 0x7d,
	0x57, 0xb9, 0x35, 0xeb, 0xee, 0x74, 0x41, 0x21, 0xa8, 0xdd, 0x92, 0xf0, 0x73, 0x18, 0x9d, 0x7b,
	0x94, 0x9e, 0x86, 0x2e, 0x5e, 0x51, 0xa3, 0xdf, 0x8b, 0xea, 0x17, 0xfd, 0x98, 0x6e, 0xf1, 0x84,
	0xa9, 0x4c, 0xc7, 0x16, 0x7e, 0x7b, 0xe6, 0x1d, 0x8e, 0x7e, 0x6a, 0x9e, 0xe1, 0x9f, 0x25, 0x54,
	0xd8, 0xb3, 0xfc, 0x00, 0x2f, 0x12, 0xf5, 0xe2, 0x32, 0xdf, 0x61, 0x32, 0x6f, 0xe2, 0x59, 0x75,
	0xbd, 0x2b, 0xe1, 0x5f, 0x20, 0x4d, 0x8f, 0xdd, 0xfe, 0xab, 0x2f, 0xe9, 0x3d, 0xa6, 0xf1, 0xbd,
	0xc6, 0x9c, 0xa9, 0x8c, 0xae, 0xe0, 0x33, 0x54, 0xdc, 0x25, 0x43, 0x02, 0x52, 0xcf, 0x55, 0x74,
	0x21, 0x09, 0x71, 0x35, 0x9b, 0xf3, 0x56, 0xf3, 0x05, 0x54, 0xb3, 0xdd, 0x7b, 0x42, 0xb1, 0xfa,
	0x8f, 0xde, 0x40, 0x9a, 0xca, 0x5f, 0xa3, 0x44, 0xdd, 0xbf, 0xc4, 0x2b, 0x3d, 0x26, 0xec, 0x6b,
	0x7c, 0x67, 0x96, 0x30, 0x02, 0x6a, 0xb4, 0x09, 0x1f, 0x27, 0xc7, 0xaf, 0x29, 0x35, 0x6d, 0xa4,
	0xa7, 0xf8, 0x68, 0x6f, 0x87, 0x3f, 0x2e, 0xc7, 0x18, 0x9f, 0xdb, 0x8a, 0xee, 0x66, 0x29, 0x6a,
	0xe8, 0x57, 0x12, 0xd1, 0xdc, 0x0d, 0xcc, 0x54, 0x47, 0x0d, 0x7c, 0x57, 0x6a, 0xc9, 0xc7, 0xe5,
	0xb4, 0x72, 0xdd, 0x22, 0xfb, 0x9b, 0xb7, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0xeb,
	0x23, 0xde, 0x94, 0x0e, 0x00, 0x00,
}
