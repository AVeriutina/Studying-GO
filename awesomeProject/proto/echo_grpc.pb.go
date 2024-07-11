// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/echo.proto

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

// AccountManagerClient is the client API for AccountManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountManagerClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error)
	ChangeNameAccount(ctx context.Context, in *ChangeNameAccountRequest, opts ...grpc.CallOption) (*ChangeNameAccountReply, error)
	ChangeAmountAccount(ctx context.Context, in *ChangeAmountAccountRequest, opts ...grpc.CallOption) (*ChangeAmountAccountReply, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error)
}

type accountManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountManagerClient(cc grpc.ClientConnInterface) AccountManagerClient {
	return &accountManagerClient{cc}
}

func (c *accountManagerClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error) {
	out := new(CreateAccountReply)
	err := c.cc.Invoke(ctx, "/proto.AccountManager/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountManagerClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error) {
	out := new(GetAccountReply)
	err := c.cc.Invoke(ctx, "/proto.AccountManager/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountManagerClient) ChangeNameAccount(ctx context.Context, in *ChangeNameAccountRequest, opts ...grpc.CallOption) (*ChangeNameAccountReply, error) {
	out := new(ChangeNameAccountReply)
	err := c.cc.Invoke(ctx, "/proto.AccountManager/ChangeNameAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountManagerClient) ChangeAmountAccount(ctx context.Context, in *ChangeAmountAccountRequest, opts ...grpc.CallOption) (*ChangeAmountAccountReply, error) {
	out := new(ChangeAmountAccountReply)
	err := c.cc.Invoke(ctx, "/proto.AccountManager/ChangeAmountAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountManagerClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error) {
	out := new(DeleteAccountReply)
	err := c.cc.Invoke(ctx, "/proto.AccountManager/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountManagerServer is the server API for AccountManager service.
// All implementations must embed UnimplementedAccountManagerServer
// for forward compatibility
type AccountManagerServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error)
	ChangeNameAccount(context.Context, *ChangeNameAccountRequest) (*ChangeNameAccountReply, error)
	ChangeAmountAccount(context.Context, *ChangeAmountAccountRequest) (*ChangeAmountAccountReply, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error)
	mustEmbedUnimplementedAccountManagerServer()
}

// UnimplementedAccountManagerServer must be embedded to have forward compatible implementations.
type UnimplementedAccountManagerServer struct {
}

func (UnimplementedAccountManagerServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountManagerServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountManagerServer) ChangeNameAccount(context.Context, *ChangeNameAccountRequest) (*ChangeNameAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNameAccount not implemented")
}
func (UnimplementedAccountManagerServer) ChangeAmountAccount(context.Context, *ChangeAmountAccountRequest) (*ChangeAmountAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAmountAccount not implemented")
}
func (UnimplementedAccountManagerServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountManagerServer) mustEmbedUnimplementedAccountManagerServer() {}

// UnsafeAccountManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountManagerServer will
// result in compilation errors.
type UnsafeAccountManagerServer interface {
	mustEmbedUnimplementedAccountManagerServer()
}

func RegisterAccountManagerServer(s grpc.ServiceRegistrar, srv AccountManagerServer) {
	s.RegisterService(&AccountManager_ServiceDesc, srv)
}

func _AccountManager_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountManager/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountManager_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountManager/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountManager_ChangeNameAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNameAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).ChangeNameAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountManager/ChangeNameAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).ChangeNameAccount(ctx, req.(*ChangeNameAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountManager_ChangeAmountAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAmountAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).ChangeAmountAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountManager/ChangeAmountAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).ChangeAmountAccount(ctx, req.(*ChangeAmountAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountManager_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountManager/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountManager_ServiceDesc is the grpc.ServiceDesc for AccountManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AccountManager",
	HandlerType: (*AccountManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountManager_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountManager_GetAccount_Handler,
		},
		{
			MethodName: "ChangeNameAccount",
			Handler:    _AccountManager_ChangeNameAccount_Handler,
		},
		{
			MethodName: "ChangeAmountAccount",
			Handler:    _AccountManager_ChangeAmountAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountManager_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/echo.proto",
}
