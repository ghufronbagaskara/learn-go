// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: interceptor.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	Token(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TokenResponse, error)
	Protected(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProtectedResponse, error)
	ProtectedStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (SecretService_ProtectedStreamClient, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) Token(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/metdata.SecretService/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) Protected(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProtectedResponse, error) {
	out := new(ProtectedResponse)
	err := c.cc.Invoke(ctx, "/metdata.SecretService/Protected", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ProtectedStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (SecretService_ProtectedStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecretService_ServiceDesc.Streams[0], "/metdata.SecretService/ProtectedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretServiceProtectedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecretService_ProtectedStreamClient interface {
	Recv() (*ProtectedResponse, error)
	grpc.ClientStream
}

type secretServiceProtectedStreamClient struct {
	grpc.ClientStream
}

func (x *secretServiceProtectedStreamClient) Recv() (*ProtectedResponse, error) {
	m := new(ProtectedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	Token(context.Context, *emptypb.Empty) (*TokenResponse, error)
	Protected(context.Context, *emptypb.Empty) (*ProtectedResponse, error)
	ProtectedStream(*emptypb.Empty, SecretService_ProtectedStreamServer) error
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) Token(context.Context, *emptypb.Empty) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedSecretServiceServer) Protected(context.Context, *emptypb.Empty) (*ProtectedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Protected not implemented")
}
func (UnimplementedSecretServiceServer) ProtectedStream(*emptypb.Empty, SecretService_ProtectedStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProtectedStream not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metdata.SecretService/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).Token(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_Protected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).Protected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metdata.SecretService/Protected",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).Protected(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ProtectedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecretServiceServer).ProtectedStream(m, &secretServiceProtectedStreamServer{stream})
}

type SecretService_ProtectedStreamServer interface {
	Send(*ProtectedResponse) error
	grpc.ServerStream
}

type secretServiceProtectedStreamServer struct {
	grpc.ServerStream
}

func (x *secretServiceProtectedStreamServer) Send(m *ProtectedResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metdata.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _SecretService_Token_Handler,
		},
		{
			MethodName: "Protected",
			Handler:    _SecretService_Protected_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProtectedStream",
			Handler:       _SecretService_ProtectedStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "interceptor.proto",
}
