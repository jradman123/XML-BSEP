// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: message_service.proto

package message_service

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	GetAllSent(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error)
	GetAllReceived(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) GetAllSent(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error) {
	out := new(GetMultipleResponse)
	err := c.cc.Invoke(ctx, "/message_service.MessageService/getAllSent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetAllReceived(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error) {
	out := new(GetMultipleResponse)
	err := c.cc.Invoke(ctx, "/message_service.MessageService/getAllReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/message_service.MessageService/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/message_service.MessageService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	GetAllSent(context.Context, *GetRequest) (*GetMultipleResponse, error)
	GetAllReceived(context.Context, *GetRequest) (*GetMultipleResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*Empty, error)
	GetAll(context.Context, *Empty) (*GetAllResponse, error)
	MustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) GetAllSent(context.Context, *GetRequest) (*GetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSent not implemented")
}
func (UnimplementedMessageServiceServer) GetAllReceived(context.Context, *GetRequest) (*GetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReceived not implemented")
}
func (UnimplementedMessageServiceServer) SendMessage(context.Context, *SendMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessageServiceServer) GetAll(context.Context, *Empty) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMessageServiceServer) MustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	MustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_GetAllSent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetAllSent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_service.MessageService/getAllSent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetAllSent(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetAllReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetAllReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_service.MessageService/getAllReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetAllReceived(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_service.MessageService/sendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message_service.MessageService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message_service.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAllSent",
			Handler:    _MessageService_GetAllSent_Handler,
		},
		{
			MethodName: "getAllReceived",
			Handler:    _MessageService_GetAllReceived_Handler,
		},
		{
			MethodName: "sendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MessageService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_service.proto",
}