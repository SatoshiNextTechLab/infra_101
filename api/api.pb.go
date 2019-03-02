// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Note
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

type Note struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Note) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Note)(nil), "api.Note")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	ChatStream(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatStreamClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) ChatStream(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[0], c.cc, "/api.Chat/ChatStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatChatStreamClient{stream}
	return x, nil
}

type Chat_ChatStreamClient interface {
	Send(*Note) error
	Recv() (*Note, error)
	grpc.ClientStream
}

type chatChatStreamClient struct {
	grpc.ClientStream
}

func (x *chatChatStreamClient) Send(m *Note) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatChatStreamClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Chat service

type ChatServer interface {
	ChatStream(Chat_ChatStreamServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_ChatStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).ChatStream(&chatChatStreamServer{stream})
}

type Chat_ChatStreamServer interface {
	Send(*Note) error
	Recv() (*Note, error)
	grpc.ServerStream
}

type chatChatStreamServer struct {
	grpc.ServerStream
}

func (x *chatChatStreamServer) Send(m *Note) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatChatStreamServer) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatStream",
			Handler:       _Chat_ChatStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52,
	0xe0, 0x62, 0xf1, 0xcb, 0x2f, 0x49, 0x15, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x8d, 0xf4, 0xb8, 0x58, 0x9c, 0x33,
	0x12, 0x4b, 0x84, 0xd4, 0xb8, 0xb8, 0x40, 0x74, 0x70, 0x49, 0x51, 0x6a, 0x62, 0xae, 0x10, 0xa7,
	0x1e, 0xc8, 0x20, 0x90, 0x56, 0x29, 0x04, 0x53, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x6c, 0xba,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x79, 0xc3, 0xe7, 0x2b, 0x6e, 0x00, 0x00, 0x00,
}
