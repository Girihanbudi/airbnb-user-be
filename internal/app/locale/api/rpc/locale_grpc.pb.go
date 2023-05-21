// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/app/locale/api/rpc/locale.proto

package rpc

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

// LocaleServiceClient is the client API for LocaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocaleServiceClient interface {
	GetLocales(ctx context.Context, in *GetLocalesCmd, opts ...grpc.CallOption) (*Locales, error)
}

type localeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocaleServiceClient(cc grpc.ClientConnInterface) LocaleServiceClient {
	return &localeServiceClient{cc}
}

func (c *localeServiceClient) GetLocales(ctx context.Context, in *GetLocalesCmd, opts ...grpc.CallOption) (*Locales, error) {
	out := new(Locales)
	err := c.cc.Invoke(ctx, "/locale.LocaleService/GetLocales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocaleServiceServer is the server API for LocaleService service.
// All implementations must embed UnimplementedLocaleServiceServer
// for forward compatibility
type LocaleServiceServer interface {
	GetLocales(context.Context, *GetLocalesCmd) (*Locales, error)
	mustEmbedUnimplementedLocaleServiceServer()
}

// UnimplementedLocaleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocaleServiceServer struct {
}

func (UnimplementedLocaleServiceServer) GetLocales(context.Context, *GetLocalesCmd) (*Locales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocales not implemented")
}
func (UnimplementedLocaleServiceServer) mustEmbedUnimplementedLocaleServiceServer() {}

// UnsafeLocaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocaleServiceServer will
// result in compilation errors.
type UnsafeLocaleServiceServer interface {
	mustEmbedUnimplementedLocaleServiceServer()
}

func RegisterLocaleServiceServer(s grpc.ServiceRegistrar, srv LocaleServiceServer) {
	s.RegisterService(&LocaleService_ServiceDesc, srv)
}

func _LocaleService_GetLocales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalesCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocaleServiceServer).GetLocales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/locale.LocaleService/GetLocales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocaleServiceServer).GetLocales(ctx, req.(*GetLocalesCmd))
	}
	return interceptor(ctx, in, info, handler)
}

// LocaleService_ServiceDesc is the grpc.ServiceDesc for LocaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "locale.LocaleService",
	HandlerType: (*LocaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocales",
			Handler:    _LocaleService_GetLocales_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/locale/api/rpc/locale.proto",
}
