// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: AuthService.proto

package __

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
	Authentification_GetPrivateKey_FullMethodName = "/AuthService.Authentification/GetPrivateKey"
	Authentification_Register_FullMethodName      = "/AuthService.Authentification/Register"
	Authentification_Login_FullMethodName         = "/AuthService.Authentification/Login"
	Authentification_UpdateTokens_FullMethodName  = "/AuthService.Authentification/UpdateTokens"
)

// AuthentificationClient is the client API for Authentification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthentificationClient interface {
	GetPrivateKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*PrivateKey, error)
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthData, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthData, error)
	UpdateTokens(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*AuthData, error)
}

type authentificationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthentificationClient(cc grpc.ClientConnInterface) AuthentificationClient {
	return &authentificationClient{cc}
}

func (c *authentificationClient) GetPrivateKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*PrivateKey, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateKey)
	err := c.cc.Invoke(ctx, Authentification_GetPrivateKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentificationClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthData)
	err := c.cc.Invoke(ctx, Authentification_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentificationClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthData)
	err := c.cc.Invoke(ctx, Authentification_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentificationClient) UpdateTokens(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*AuthData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthData)
	err := c.cc.Invoke(ctx, Authentification_UpdateTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthentificationServer is the server API for Authentification service.
// All implementations must embed UnimplementedAuthentificationServer
// for forward compatibility.
type AuthentificationServer interface {
	GetPrivateKey(context.Context, *KeyRequest) (*PrivateKey, error)
	Register(context.Context, *User) (*AuthData, error)
	Login(context.Context, *User) (*AuthData, error)
	UpdateTokens(context.Context, *AuthData) (*AuthData, error)
	mustEmbedUnimplementedAuthentificationServer()
}

// UnimplementedAuthentificationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthentificationServer struct{}

func (UnimplementedAuthentificationServer) GetPrivateKey(context.Context, *KeyRequest) (*PrivateKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedAuthentificationServer) Register(context.Context, *User) (*AuthData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthentificationServer) Login(context.Context, *User) (*AuthData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthentificationServer) UpdateTokens(context.Context, *AuthData) (*AuthData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTokens not implemented")
}
func (UnimplementedAuthentificationServer) mustEmbedUnimplementedAuthentificationServer() {}
func (UnimplementedAuthentificationServer) testEmbeddedByValue()                          {}

// UnsafeAuthentificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthentificationServer will
// result in compilation errors.
type UnsafeAuthentificationServer interface {
	mustEmbedUnimplementedAuthentificationServer()
}

func RegisterAuthentificationServer(s grpc.ServiceRegistrar, srv AuthentificationServer) {
	// If the following call pancis, it indicates UnimplementedAuthentificationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Authentification_ServiceDesc, srv)
}

func _Authentification_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentification_GetPrivateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServer).GetPrivateKey(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentification_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentification_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentification_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentification_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentification_UpdateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServer).UpdateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentification_UpdateTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServer).UpdateTokens(ctx, req.(*AuthData))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentification_ServiceDesc is the grpc.ServiceDesc for Authentification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService.Authentification",
	HandlerType: (*AuthentificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrivateKey",
			Handler:    _Authentification_GetPrivateKey_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Authentification_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authentification_Login_Handler,
		},
		{
			MethodName: "UpdateTokens",
			Handler:    _Authentification_UpdateTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AuthService.proto",
}
