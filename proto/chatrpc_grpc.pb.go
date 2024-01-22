// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/chatrpc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Chat_SendMessage_FullMethodName    = "/chat.Chat/SendMessage"
	Chat_RecieveMessage_FullMethodName = "/chat.Chat/RecieveMessage"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error)
	RecieveMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_RecieveMessageClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], Chat_SendMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatSendMessageClient{stream}
	return x, nil
}

type Chat_SendMessageClient interface {
	Send(*MessageSend) error
	Recv() (*MessageFlag, error)
	grpc.ClientStream
}

type chatSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatSendMessageClient) Send(m *MessageSend) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatSendMessageClient) Recv() (*MessageFlag, error) {
	m := new(MessageFlag)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) RecieveMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_RecieveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[1], Chat_RecieveMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRecieveMessageClient{stream}
	return x, nil
}

type Chat_RecieveMessageClient interface {
	Send(*MessageFlag) error
	Recv() (*MessageRecieve, error)
	grpc.ClientStream
}

type chatRecieveMessageClient struct {
	grpc.ClientStream
}

func (x *chatRecieveMessageClient) Send(m *MessageFlag) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRecieveMessageClient) Recv() (*MessageRecieve, error) {
	m := new(MessageRecieve)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	SendMessage(Chat_SendMessageServer) error
	RecieveMessage(Chat_RecieveMessageServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) SendMessage(Chat_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) RecieveMessage(Chat_RecieveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method RecieveMessage not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).SendMessage(&chatSendMessageServer{stream})
}

type Chat_SendMessageServer interface {
	Send(*MessageFlag) error
	Recv() (*MessageSend, error)
	grpc.ServerStream
}

type chatSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatSendMessageServer) Send(m *MessageFlag) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatSendMessageServer) Recv() (*MessageSend, error) {
	m := new(MessageSend)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_RecieveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).RecieveMessage(&chatRecieveMessageServer{stream})
}

type Chat_RecieveMessageServer interface {
	Send(*MessageRecieve) error
	Recv() (*MessageFlag, error)
	grpc.ServerStream
}

type chatRecieveMessageServer struct {
	grpc.ServerStream
}

func (x *chatRecieveMessageServer) Send(m *MessageRecieve) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRecieveMessageServer) Recv() (*MessageFlag, error) {
	m := new(MessageFlag)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _Chat_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RecieveMessage",
			Handler:       _Chat_RecieveMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chatrpc.proto",
}
