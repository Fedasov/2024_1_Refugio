//go:generate mockgen -source=./folder_grpc.pb.go -destination=../mock/folder_grpc_mock.go -package=mock proto FolderServiceClient
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: folder.proto

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

// FolderServiceClient is the client API for FolderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FolderServiceClient interface {
	CreateFolder(ctx context.Context, in *Folder, opts ...grpc.CallOption) (*FolderWithID, error)
	GetAllFolders(ctx context.Context, in *GetAllFoldersData, opts ...grpc.CallOption) (*Folders, error)
	UpdateFolder(ctx context.Context, in *Folder, opts ...grpc.CallOption) (*FolderStatus, error)
	DeleteFolder(ctx context.Context, in *DeleteFolderData, opts ...grpc.CallOption) (*FolderStatus, error)
	AddEmailInFolder(ctx context.Context, in *FolderEmail, opts ...grpc.CallOption) (*FolderEmailStatus, error)
	DeleteEmailInFolder(ctx context.Context, in *FolderEmail, opts ...grpc.CallOption) (*FolderEmailStatus, error)
	GetAllEmailsInFolder(ctx context.Context, in *GetAllEmailsInFolderData, opts ...grpc.CallOption) (*ObjectsEmail, error)
	CheckFolderProfile(ctx context.Context, in *FolderProfile, opts ...grpc.CallOption) (*FolderEmailStatus, error)
	CheckEmailProfile(ctx context.Context, in *EmailProfile, opts ...grpc.CallOption) (*FolderEmailStatus, error)
}

type folderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFolderServiceClient(cc grpc.ClientConnInterface) FolderServiceClient {
	return &folderServiceClient{cc}
}

func (c *folderServiceClient) CreateFolder(ctx context.Context, in *Folder, opts ...grpc.CallOption) (*FolderWithID, error) {
	out := new(FolderWithID)
	err := c.cc.Invoke(ctx, "/proto.FolderService/CreateFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) GetAllFolders(ctx context.Context, in *GetAllFoldersData, opts ...grpc.CallOption) (*Folders, error) {
	out := new(Folders)
	err := c.cc.Invoke(ctx, "/proto.FolderService/GetAllFolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) UpdateFolder(ctx context.Context, in *Folder, opts ...grpc.CallOption) (*FolderStatus, error) {
	out := new(FolderStatus)
	err := c.cc.Invoke(ctx, "/proto.FolderService/UpdateFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) DeleteFolder(ctx context.Context, in *DeleteFolderData, opts ...grpc.CallOption) (*FolderStatus, error) {
	out := new(FolderStatus)
	err := c.cc.Invoke(ctx, "/proto.FolderService/DeleteFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) AddEmailInFolder(ctx context.Context, in *FolderEmail, opts ...grpc.CallOption) (*FolderEmailStatus, error) {
	out := new(FolderEmailStatus)
	err := c.cc.Invoke(ctx, "/proto.FolderService/AddEmailInFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) DeleteEmailInFolder(ctx context.Context, in *FolderEmail, opts ...grpc.CallOption) (*FolderEmailStatus, error) {
	out := new(FolderEmailStatus)
	err := c.cc.Invoke(ctx, "/proto.FolderService/DeleteEmailInFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) GetAllEmailsInFolder(ctx context.Context, in *GetAllEmailsInFolderData, opts ...grpc.CallOption) (*ObjectsEmail, error) {
	out := new(ObjectsEmail)
	err := c.cc.Invoke(ctx, "/proto.FolderService/GetAllEmailsInFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) CheckFolderProfile(ctx context.Context, in *FolderProfile, opts ...grpc.CallOption) (*FolderEmailStatus, error) {
	out := new(FolderEmailStatus)
	err := c.cc.Invoke(ctx, "/proto.FolderService/CheckFolderProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) CheckEmailProfile(ctx context.Context, in *EmailProfile, opts ...grpc.CallOption) (*FolderEmailStatus, error) {
	out := new(FolderEmailStatus)
	err := c.cc.Invoke(ctx, "/proto.FolderService/CheckEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FolderServiceServer is the server API for FolderService service.
// All implementations must embed UnimplementedFolderServiceServer
// for forward compatibility
type FolderServiceServer interface {
	CreateFolder(context.Context, *Folder) (*FolderWithID, error)
	GetAllFolders(context.Context, *GetAllFoldersData) (*Folders, error)
	UpdateFolder(context.Context, *Folder) (*FolderStatus, error)
	DeleteFolder(context.Context, *DeleteFolderData) (*FolderStatus, error)
	AddEmailInFolder(context.Context, *FolderEmail) (*FolderEmailStatus, error)
	DeleteEmailInFolder(context.Context, *FolderEmail) (*FolderEmailStatus, error)
	GetAllEmailsInFolder(context.Context, *GetAllEmailsInFolderData) (*ObjectsEmail, error)
	CheckFolderProfile(context.Context, *FolderProfile) (*FolderEmailStatus, error)
	CheckEmailProfile(context.Context, *EmailProfile) (*FolderEmailStatus, error)
	mustEmbedUnimplementedFolderServiceServer()
}

// UnimplementedFolderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFolderServiceServer struct {
}

func (UnimplementedFolderServiceServer) CreateFolder(context.Context, *Folder) (*FolderWithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFolder not implemented")
}
func (UnimplementedFolderServiceServer) GetAllFolders(context.Context, *GetAllFoldersData) (*Folders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFolders not implemented")
}
func (UnimplementedFolderServiceServer) UpdateFolder(context.Context, *Folder) (*FolderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFolder not implemented")
}
func (UnimplementedFolderServiceServer) DeleteFolder(context.Context, *DeleteFolderData) (*FolderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}
func (UnimplementedFolderServiceServer) AddEmailInFolder(context.Context, *FolderEmail) (*FolderEmailStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmailInFolder not implemented")
}
func (UnimplementedFolderServiceServer) DeleteEmailInFolder(context.Context, *FolderEmail) (*FolderEmailStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmailInFolder not implemented")
}
func (UnimplementedFolderServiceServer) GetAllEmailsInFolder(context.Context, *GetAllEmailsInFolderData) (*ObjectsEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEmailsInFolder not implemented")
}
func (UnimplementedFolderServiceServer) CheckFolderProfile(context.Context, *FolderProfile) (*FolderEmailStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckFolderProfile not implemented")
}
func (UnimplementedFolderServiceServer) CheckEmailProfile(context.Context, *EmailProfile) (*FolderEmailStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmailProfile not implemented")
}
func (UnimplementedFolderServiceServer) mustEmbedUnimplementedFolderServiceServer() {}

// UnsafeFolderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FolderServiceServer will
// result in compilation errors.
type UnsafeFolderServiceServer interface {
	mustEmbedUnimplementedFolderServiceServer()
}

func RegisterFolderServiceServer(s grpc.ServiceRegistrar, srv FolderServiceServer) {
	s.RegisterService(&FolderService_ServiceDesc, srv)
}

func _FolderService_CreateFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Folder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).CreateFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/CreateFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).CreateFolder(ctx, req.(*Folder))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_GetAllFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFoldersData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).GetAllFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/GetAllFolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).GetAllFolders(ctx, req.(*GetAllFoldersData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_UpdateFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Folder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).UpdateFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/UpdateFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).UpdateFolder(ctx, req.(*Folder))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_DeleteFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFolderData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).DeleteFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/DeleteFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).DeleteFolder(ctx, req.(*DeleteFolderData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_AddEmailInFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).AddEmailInFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/AddEmailInFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).AddEmailInFolder(ctx, req.(*FolderEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_DeleteEmailInFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).DeleteEmailInFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/DeleteEmailInFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).DeleteEmailInFolder(ctx, req.(*FolderEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_GetAllEmailsInFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEmailsInFolderData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).GetAllEmailsInFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/GetAllEmailsInFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).GetAllEmailsInFolder(ctx, req.(*GetAllEmailsInFolderData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_CheckFolderProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).CheckFolderProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/CheckFolderProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).CheckFolderProfile(ctx, req.(*FolderProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_CheckEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).CheckEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FolderService/CheckEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).CheckEmailProfile(ctx, req.(*EmailProfile))
	}
	return interceptor(ctx, in, info, handler)
}

// FolderService_ServiceDesc is the grpc.ServiceDesc for FolderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FolderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FolderService",
	HandlerType: (*FolderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFolder",
			Handler:    _FolderService_CreateFolder_Handler,
		},
		{
			MethodName: "GetAllFolders",
			Handler:    _FolderService_GetAllFolders_Handler,
		},
		{
			MethodName: "UpdateFolder",
			Handler:    _FolderService_UpdateFolder_Handler,
		},
		{
			MethodName: "DeleteFolder",
			Handler:    _FolderService_DeleteFolder_Handler,
		},
		{
			MethodName: "AddEmailInFolder",
			Handler:    _FolderService_AddEmailInFolder_Handler,
		},
		{
			MethodName: "DeleteEmailInFolder",
			Handler:    _FolderService_DeleteEmailInFolder_Handler,
		},
		{
			MethodName: "GetAllEmailsInFolder",
			Handler:    _FolderService_GetAllEmailsInFolder_Handler,
		},
		{
			MethodName: "CheckFolderProfile",
			Handler:    _FolderService_CheckFolderProfile_Handler,
		},
		{
			MethodName: "CheckEmailProfile",
			Handler:    _FolderService_CheckEmailProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "folder.proto",
}
