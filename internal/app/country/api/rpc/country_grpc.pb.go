// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/app/country/api/rpc/country.proto

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

// CountryServiceClient is the client API for CountryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountryServiceClient interface {
	GetCountryByPhoneCode(ctx context.Context, in *GetCountryByPhoneCodeCmd, opts ...grpc.CallOption) (*Country, error)
}

type countryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryServiceClient(cc grpc.ClientConnInterface) CountryServiceClient {
	return &countryServiceClient{cc}
}

func (c *countryServiceClient) GetCountryByPhoneCode(ctx context.Context, in *GetCountryByPhoneCodeCmd, opts ...grpc.CallOption) (*Country, error) {
	out := new(Country)
	err := c.cc.Invoke(ctx, "/country.CountryService/GetCountryByPhoneCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServiceServer is the server API for CountryService service.
// All implementations must embed UnimplementedCountryServiceServer
// for forward compatibility
type CountryServiceServer interface {
	GetCountryByPhoneCode(context.Context, *GetCountryByPhoneCodeCmd) (*Country, error)
	mustEmbedUnimplementedCountryServiceServer()
}

// UnimplementedCountryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCountryServiceServer struct {
}

func (UnimplementedCountryServiceServer) GetCountryByPhoneCode(context.Context, *GetCountryByPhoneCodeCmd) (*Country, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountryByPhoneCode not implemented")
}
func (UnimplementedCountryServiceServer) mustEmbedUnimplementedCountryServiceServer() {}

// UnsafeCountryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountryServiceServer will
// result in compilation errors.
type UnsafeCountryServiceServer interface {
	mustEmbedUnimplementedCountryServiceServer()
}

func RegisterCountryServiceServer(s grpc.ServiceRegistrar, srv CountryServiceServer) {
	s.RegisterService(&CountryService_ServiceDesc, srv)
}

func _CountryService_GetCountryByPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryByPhoneCodeCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).GetCountryByPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/country.CountryService/GetCountryByPhoneCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).GetCountryByPhoneCode(ctx, req.(*GetCountryByPhoneCodeCmd))
	}
	return interceptor(ctx, in, info, handler)
}

// CountryService_ServiceDesc is the grpc.ServiceDesc for CountryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "country.CountryService",
	HandlerType: (*CountryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCountryByPhoneCode",
			Handler:    _CountryService_GetCountryByPhoneCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/country/api/rpc/country.proto",
}
