// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: sso/v1/sso.proto

package sso

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_RegisterV1_FullMethodName = "/sso.v1.AuthService/RegisterV1"
	AuthService_LoginV1_FullMethodName    = "/sso.v1.AuthService/LoginV1"
	AuthService_IsRootV1_FullMethodName   = "/sso.v1.AuthService/IsRootV1"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	RegisterV1(ctx context.Context, in *RegisterRequestV1, opts ...grpc.CallOption) (*RegisterResponseV1, error)
	LoginV1(ctx context.Context, in *LoginRequestV1, opts ...grpc.CallOption) (*LoginResponseV1, error)
	IsRootV1(ctx context.Context, in *IsRootRequestV1, opts ...grpc.CallOption) (*IsRootResponseV1, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RegisterV1(ctx context.Context, in *RegisterRequestV1, opts ...grpc.CallOption) (*RegisterResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponseV1)
	err := c.cc.Invoke(ctx, AuthService_RegisterV1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LoginV1(ctx context.Context, in *LoginRequestV1, opts ...grpc.CallOption) (*LoginResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponseV1)
	err := c.cc.Invoke(ctx, AuthService_LoginV1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsRootV1(ctx context.Context, in *IsRootRequestV1, opts ...grpc.CallOption) (*IsRootResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsRootResponseV1)
	err := c.cc.Invoke(ctx, AuthService_IsRootV1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	RegisterV1(context.Context, *RegisterRequestV1) (*RegisterResponseV1, error)
	LoginV1(context.Context, *LoginRequestV1) (*LoginResponseV1, error)
	IsRootV1(context.Context, *IsRootRequestV1) (*IsRootResponseV1, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) RegisterV1(context.Context, *RegisterRequestV1) (*RegisterResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterV1 not implemented")
}
func (UnimplementedAuthServiceServer) LoginV1(context.Context, *LoginRequestV1) (*LoginResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginV1 not implemented")
}
func (UnimplementedAuthServiceServer) IsRootV1(context.Context, *IsRootRequestV1) (*IsRootResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsRootV1 not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_RegisterV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RegisterV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterV1(ctx, req.(*RegisterRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LoginV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoginV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_LoginV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoginV1(ctx, req.(*LoginRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsRootV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRootRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsRootV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IsRootV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsRootV1(ctx, req.(*IsRootRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sso.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterV1",
			Handler:    _AuthService_RegisterV1_Handler,
		},
		{
			MethodName: "LoginV1",
			Handler:    _AuthService_LoginV1_Handler,
		},
		{
			MethodName: "IsRootV1",
			Handler:    _AuthService_IsRootV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/v1/sso.proto",
}
