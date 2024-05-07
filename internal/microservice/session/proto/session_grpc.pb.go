//go:generate mockgen -source=./session_grpc.pb.go -destination=../mock/session_grpc_mock.go -package=mock proto SessionServiceClient
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: session.proto

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

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionReply, error)
	GetLoginBySession(ctx context.Context, in *GetLoginBySessionRequest, opts ...grpc.CallOption) (*GetLoginBySessionReply, error)
	GetProfileIDBySession(ctx context.Context, in *GetLoginBySessionRequest, opts ...grpc.CallOption) (*GetProfileIDBySessionReply, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionReply, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionReply, error)
	CleanupExpiredSessions(ctx context.Context, in *CleanupExpiredSessionsRequest, opts ...grpc.CallOption) (*CleanupExpiredSessionsReply, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionReply, error) {
	out := new(GetSessionReply)
	err := c.cc.Invoke(ctx, "/proto.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetLoginBySession(ctx context.Context, in *GetLoginBySessionRequest, opts ...grpc.CallOption) (*GetLoginBySessionReply, error) {
	out := new(GetLoginBySessionReply)
	err := c.cc.Invoke(ctx, "/proto.SessionService/GetLoginBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetProfileIDBySession(ctx context.Context, in *GetLoginBySessionRequest, opts ...grpc.CallOption) (*GetProfileIDBySessionReply, error) {
	out := new(GetProfileIDBySessionReply)
	err := c.cc.Invoke(ctx, "/proto.SessionService/GetProfileIDBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionReply, error) {
	out := new(CreateSessionReply)
	err := c.cc.Invoke(ctx, "/proto.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionReply, error) {
	out := new(DeleteSessionReply)
	err := c.cc.Invoke(ctx, "/proto.SessionService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CleanupExpiredSessions(ctx context.Context, in *CleanupExpiredSessionsRequest, opts ...grpc.CallOption) (*CleanupExpiredSessionsReply, error) {
	out := new(CleanupExpiredSessionsReply)
	err := c.cc.Invoke(ctx, "/proto.SessionService/CleanupExpiredSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	GetSession(context.Context, *GetSessionRequest) (*GetSessionReply, error)
	GetLoginBySession(context.Context, *GetLoginBySessionRequest) (*GetLoginBySessionReply, error)
	GetProfileIDBySession(context.Context, *GetLoginBySessionRequest) (*GetProfileIDBySessionReply, error)
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionReply, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionReply, error)
	CleanupExpiredSessions(context.Context, *CleanupExpiredSessionsRequest) (*CleanupExpiredSessionsReply, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) GetSession(context.Context, *GetSessionRequest) (*GetSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedSessionServiceServer) GetLoginBySession(context.Context, *GetLoginBySessionRequest) (*GetLoginBySessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginBySession not implemented")
}
func (UnimplementedSessionServiceServer) GetProfileIDBySession(context.Context, *GetLoginBySessionRequest) (*GetProfileIDBySessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileIDBySession not implemented")
}
func (UnimplementedSessionServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedSessionServiceServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedSessionServiceServer) CleanupExpiredSessions(context.Context, *CleanupExpiredSessionsRequest) (*CleanupExpiredSessionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanupExpiredSessions not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetLoginBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginBySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetLoginBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/GetLoginBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetLoginBySession(ctx, req.(*GetLoginBySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetProfileIDBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginBySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetProfileIDBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/GetProfileIDBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetProfileIDBySession(ctx, req.(*GetLoginBySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CleanupExpiredSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupExpiredSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CleanupExpiredSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/CleanupExpiredSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CleanupExpiredSessions(ctx, req.(*CleanupExpiredSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
		{
			MethodName: "GetLoginBySession",
			Handler:    _SessionService_GetLoginBySession_Handler,
		},
		{
			MethodName: "GetProfileIDBySession",
			Handler:    _SessionService_GetProfileIDBySession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _SessionService_DeleteSession_Handler,
		},
		{
			MethodName: "CleanupExpiredSessions",
			Handler:    _SessionService_CleanupExpiredSessions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
