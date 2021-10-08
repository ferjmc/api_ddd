// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sessionService

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

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	GetSessionByID(ctx context.Context, in *GetSessionByIDRequest, opts ...grpc.CallOption) (*GetSessionByIDResponse, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
	CreateCsrfToken(ctx context.Context, in *CreateCsrfTokenRequest, opts ...grpc.CallOption) (*CreateCsrfTokenResponse, error)
	CheckCsrfToken(ctx context.Context, in *CheckCsrfTokenRequest, opts ...grpc.CallOption) (*CheckCsrfTokenResponse, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetSessionByID(ctx context.Context, in *GetSessionByIDRequest, opts ...grpc.CallOption) (*GetSessionByIDResponse, error) {
	out := new(GetSessionByIDResponse)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/GetSessionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateCsrfToken(ctx context.Context, in *CreateCsrfTokenRequest, opts ...grpc.CallOption) (*CreateCsrfTokenResponse, error) {
	out := new(CreateCsrfTokenResponse)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/CreateCsrfToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CheckCsrfToken(ctx context.Context, in *CheckCsrfTokenRequest, opts ...grpc.CallOption) (*CheckCsrfTokenResponse, error) {
	out := new(CheckCsrfTokenResponse)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/CheckCsrfToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations must embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	GetSessionByID(context.Context, *GetSessionByIDRequest) (*GetSessionByIDResponse, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	CreateCsrfToken(context.Context, *CreateCsrfTokenRequest) (*CreateCsrfTokenResponse, error)
	CheckCsrfToken(context.Context, *CheckCsrfTokenRequest) (*CheckCsrfTokenResponse, error)
	mustEmbedUnimplementedAuthorizationServiceServer()
}

// UnimplementedAuthorizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetSessionByID(context.Context, *GetSessionByIDRequest) (*GetSessionByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionByID not implemented")
}
func (UnimplementedAuthorizationServiceServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedAuthorizationServiceServer) CreateCsrfToken(context.Context, *CreateCsrfTokenRequest) (*CreateCsrfTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCsrfToken not implemented")
}
func (UnimplementedAuthorizationServiceServer) CheckCsrfToken(context.Context, *CheckCsrfTokenRequest) (*CheckCsrfTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCsrfToken not implemented")
}
func (UnimplementedAuthorizationServiceServer) mustEmbedUnimplementedAuthorizationServiceServer() {}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetSessionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetSessionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/GetSessionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetSessionByID(ctx, req.(*GetSessionByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateCsrfToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCsrfTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateCsrfToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/CreateCsrfToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateCsrfToken(ctx, req.(*CreateCsrfTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CheckCsrfToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCsrfTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CheckCsrfToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/CheckCsrfToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CheckCsrfToken(ctx, req.(*CheckCsrfTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sessionService.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _AuthorizationService_CreateSession_Handler,
		},
		{
			MethodName: "GetSessionByID",
			Handler:    _AuthorizationService_GetSessionByID_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AuthorizationService_DeleteSession_Handler,
		},
		{
			MethodName: "CreateCsrfToken",
			Handler:    _AuthorizationService_CreateCsrfToken_Handler,
		},
		{
			MethodName: "CheckCsrfToken",
			Handler:    _AuthorizationService_CheckCsrfToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
