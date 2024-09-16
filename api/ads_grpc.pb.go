// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: api/ads.proto

package api

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
	AdsService_CreateAd_FullMethodName       = "/api.AdsService/CreateAd"
	AdsService_GetAd_FullMethodName          = "/api.AdsService/GetAd"
	AdsService_CheckAdIsValid_FullMethodName = "/api.AdsService/CheckAdIsValid"
)

// AdsServiceClient is the client API for AdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdsServiceClient interface {
	CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*Ad, error)
	GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*Ad, error)
	CheckAdIsValid(ctx context.Context, in *CheckAdIsValidRequest, opts ...grpc.CallOption) (*AdValidity, error)
}

type adsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdsServiceClient(cc grpc.ClientConnInterface) AdsServiceClient {
	return &adsServiceClient{cc}
}

func (c *adsServiceClient) CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*Ad, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ad)
	err := c.cc.Invoke(ctx, AdsService_CreateAd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adsServiceClient) GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*Ad, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ad)
	err := c.cc.Invoke(ctx, AdsService_GetAd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adsServiceClient) CheckAdIsValid(ctx context.Context, in *CheckAdIsValidRequest, opts ...grpc.CallOption) (*AdValidity, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdValidity)
	err := c.cc.Invoke(ctx, AdsService_CheckAdIsValid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdsServiceServer is the server API for AdsService service.
// All implementations must embed UnimplementedAdsServiceServer
// for forward compatibility.
type AdsServiceServer interface {
	CreateAd(context.Context, *CreateAdRequest) (*Ad, error)
	GetAd(context.Context, *GetAdRequest) (*Ad, error)
	CheckAdIsValid(context.Context, *CheckAdIsValidRequest) (*AdValidity, error)
	mustEmbedUnimplementedAdsServiceServer()
}

// UnimplementedAdsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdsServiceServer struct{}

func (UnimplementedAdsServiceServer) CreateAd(context.Context, *CreateAdRequest) (*Ad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAd not implemented")
}
func (UnimplementedAdsServiceServer) GetAd(context.Context, *GetAdRequest) (*Ad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAd not implemented")
}
func (UnimplementedAdsServiceServer) CheckAdIsValid(context.Context, *CheckAdIsValidRequest) (*AdValidity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAdIsValid not implemented")
}
func (UnimplementedAdsServiceServer) mustEmbedUnimplementedAdsServiceServer() {}
func (UnimplementedAdsServiceServer) testEmbeddedByValue()                    {}

// UnsafeAdsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdsServiceServer will
// result in compilation errors.
type UnsafeAdsServiceServer interface {
	mustEmbedUnimplementedAdsServiceServer()
}

func RegisterAdsServiceServer(s grpc.ServiceRegistrar, srv AdsServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdsService_ServiceDesc, srv)
}

func _AdsService_CreateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).CreateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdsService_CreateAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).CreateAd(ctx, req.(*CreateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdsService_GetAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).GetAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdsService_GetAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).GetAd(ctx, req.(*GetAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdsService_CheckAdIsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAdIsValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).CheckAdIsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdsService_CheckAdIsValid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).CheckAdIsValid(ctx, req.(*CheckAdIsValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdsService_ServiceDesc is the grpc.ServiceDesc for AdsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AdsService",
	HandlerType: (*AdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAd",
			Handler:    _AdsService_CreateAd_Handler,
		},
		{
			MethodName: "GetAd",
			Handler:    _AdsService_GetAd_Handler,
		},
		{
			MethodName: "CheckAdIsValid",
			Handler:    _AdsService_CheckAdIsValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ads.proto",
}

const (
	AggregatorService_SendClick_FullMethodName       = "/api.AggregatorService/SendClick"
	AggregatorService_GetClickCounter_FullMethodName = "/api.AggregatorService/GetClickCounter"
)

// AggregatorServiceClient is the client API for AggregatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorServiceClient interface {
	SendClick(ctx context.Context, in *SendClickRequest, opts ...grpc.CallOption) (*Click, error)
	GetClickCounter(ctx context.Context, in *GetClicksCounterRequest, opts ...grpc.CallOption) (*ClickCounter, error)
}

type aggregatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorServiceClient(cc grpc.ClientConnInterface) AggregatorServiceClient {
	return &aggregatorServiceClient{cc}
}

func (c *aggregatorServiceClient) SendClick(ctx context.Context, in *SendClickRequest, opts ...grpc.CallOption) (*Click, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Click)
	err := c.cc.Invoke(ctx, AggregatorService_SendClick_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorServiceClient) GetClickCounter(ctx context.Context, in *GetClicksCounterRequest, opts ...grpc.CallOption) (*ClickCounter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClickCounter)
	err := c.cc.Invoke(ctx, AggregatorService_GetClickCounter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServiceServer is the server API for AggregatorService service.
// All implementations must embed UnimplementedAggregatorServiceServer
// for forward compatibility.
type AggregatorServiceServer interface {
	SendClick(context.Context, *SendClickRequest) (*Click, error)
	GetClickCounter(context.Context, *GetClicksCounterRequest) (*ClickCounter, error)
	mustEmbedUnimplementedAggregatorServiceServer()
}

// UnimplementedAggregatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAggregatorServiceServer struct{}

func (UnimplementedAggregatorServiceServer) SendClick(context.Context, *SendClickRequest) (*Click, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClick not implemented")
}
func (UnimplementedAggregatorServiceServer) GetClickCounter(context.Context, *GetClicksCounterRequest) (*ClickCounter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClickCounter not implemented")
}
func (UnimplementedAggregatorServiceServer) mustEmbedUnimplementedAggregatorServiceServer() {}
func (UnimplementedAggregatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeAggregatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServiceServer will
// result in compilation errors.
type UnsafeAggregatorServiceServer interface {
	mustEmbedUnimplementedAggregatorServiceServer()
}

func RegisterAggregatorServiceServer(s grpc.ServiceRegistrar, srv AggregatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedAggregatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AggregatorService_ServiceDesc, srv)
}

func _AggregatorService_SendClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendClickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).SendClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_SendClick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).SendClick(ctx, req.(*SendClickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregatorService_GetClickCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClicksCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetClickCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetClickCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetClickCounter(ctx, req.(*GetClicksCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregatorService_ServiceDesc is the grpc.ServiceDesc for AggregatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AggregatorService",
	HandlerType: (*AggregatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClick",
			Handler:    _AggregatorService_SendClick_Handler,
		},
		{
			MethodName: "GetClickCounter",
			Handler:    _AggregatorService_GetClickCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ads.proto",
}
