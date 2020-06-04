// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package schema

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloAgainRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname              string   `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloAgainRequest) Reset()         { *m = HelloAgainRequest{} }
func (m *HelloAgainRequest) String() string { return proto.CompactTextString(m) }
func (*HelloAgainRequest) ProtoMessage()    {}
func (*HelloAgainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *HelloAgainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloAgainRequest.Unmarshal(m, b)
}
func (m *HelloAgainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloAgainRequest.Marshal(b, m, deterministic)
}
func (m *HelloAgainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloAgainRequest.Merge(m, src)
}
func (m *HelloAgainRequest) XXX_Size() int {
	return xxx_messageInfo_HelloAgainRequest.Size(m)
}
func (m *HelloAgainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloAgainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloAgainRequest proto.InternalMessageInfo

func (m *HelloAgainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloAgainRequest) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

type HelloAgainReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloAgainReply) Reset()         { *m = HelloAgainReply{} }
func (m *HelloAgainReply) String() string { return proto.CompactTextString(m) }
func (*HelloAgainReply) ProtoMessage()    {}
func (*HelloAgainReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *HelloAgainReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloAgainReply.Unmarshal(m, b)
}
func (m *HelloAgainReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloAgainReply.Marshal(b, m, deterministic)
}
func (m *HelloAgainReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloAgainReply.Merge(m, src)
}
func (m *HelloAgainReply) XXX_Size() int {
	return xxx_messageInfo_HelloAgainReply.Size(m)
}
func (m *HelloAgainReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloAgainReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloAgainReply proto.InternalMessageInfo

func (m *HelloAgainReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloAgainReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*HelloAgainRequest)(nil), "helloworld.HelloAgainRequest")
	proto.RegisterType((*HelloAgainReply)(nil), "helloworld.HelloAgainReply")
}

func init() {
	proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2)
}

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xad, 0xa8, 0xab, 0x83, 0xa2, 0x0e, 0x22, 0x65, 0x45, 0x90, 0x1c, 0xfc, 0x73, 0x68,
	0x0b, 0x7a, 0x17, 0xea, 0x45, 0x2f, 0x5e, 0xd6, 0x9b, 0xb7, 0x74, 0x1d, 0xba, 0x8b, 0x69, 0x52,
	0x27, 0x29, 0x4b, 0xbf, 0x89, 0x1f, 0x57, 0x9a, 0x6d, 0x31, 0xa0, 0xdb, 0x5b, 0xde, 0xbc, 0x97,
	0x5f, 0xf2, 0x12, 0x38, 0x59, 0x90, 0x52, 0x66, 0x65, 0x58, 0x7d, 0xa4, 0x35, 0x1b, 0x67, 0x10,
	0x7e, 0x27, 0x42, 0xc0, 0xe1, 0x4b, 0xa7, 0x66, 0xf4, 0xd5, 0x90, 0x75, 0x88, 0xb0, 0xa3, 0x65,
	0x45, 0x71, 0x74, 0x15, 0xdd, 0x1e, 0xcc, 0xfc, 0x5a, 0x5c, 0x03, 0xf4, 0x99, 0x5a, 0xb5, 0x18,
	0xc3, 0xa4, 0x22, 0x6b, 0x65, 0x39, 0x84, 0x06, 0x29, 0x72, 0x38, 0xf5, 0xb9, 0xbc, 0x94, 0x4b,
	0x3d, 0x02, 0xec, 0x10, 0xb6, 0x61, 0x3f, 0xde, 0x5e, 0x23, 0x7a, 0x29, 0x72, 0x38, 0x0e, 0x11,
	0xa3, 0xe7, 0xe1, 0x19, 0xec, 0x12, 0xb3, 0xe1, 0x1e, 0xb2, 0x16, 0xf7, 0xdf, 0x11, 0x4c, 0x9e,
	0x99, 0xc8, 0x11, 0xe3, 0x23, 0xec, 0xbf, 0xc9, 0xd6, 0x13, 0x31, 0x4e, 0x83, 0x87, 0x08, 0x3b,
	0x4f, 0xcf, 0xff, 0x71, 0x6a, 0xd5, 0x8a, 0x2d, 0x7c, 0x85, 0xa3, 0x61, 0xbf, 0xbf, 0x11, 0x5e,
	0xfe, 0x89, 0x86, 0x65, 0xa7, 0x17, 0x9b, 0x6c, 0x8f, 0x7b, 0xba, 0x7b, 0xbf, 0x29, 0x96, 0xae,
	0x68, 0xe6, 0x9f, 0xe4, 0x52, 0xc3, 0x65, 0xa6, 0x4d, 0xd2, 0xd5, 0x4e, 0x4a, 0x59, 0x51, 0xa6,
	0x75, 0xb2, 0xb2, 0x99, 0x9d, 0x2f, 0xa8, 0x92, 0xc5, 0x9e, 0xff, 0xaa, 0x87, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x30, 0x1e, 0x5f, 0x59, 0xbe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloAgainRequest, opts ...grpc.CallOption) (*HelloAgainReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloAgainRequest, opts ...grpc.CallOption) (*HelloAgainReply, error) {
	out := new(HelloAgainReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloAgainRequest) (*HelloAgainReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) SayHelloAgain(ctx context.Context, req *HelloAgainRequest) (*HelloAgainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloAgainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloAgainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
