// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/subscription.proto

package subscription

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

const (
	SubscriptionInterface_Delete_FullMethodName   = "/subscription.SubscriptionInterface/Delete"
	SubscriptionInterface_Find_FullMethodName     = "/subscription.SubscriptionInterface/Find"
	SubscriptionInterface_Get_FullMethodName      = "/subscription.SubscriptionInterface/Get"
	SubscriptionInterface_SetBurst_FullMethodName = "/subscription.SubscriptionInterface/SetBurst"
	SubscriptionInterface_SetSize_FullMethodName  = "/subscription.SubscriptionInterface/SetSize"
)

// SubscriptionInterfaceClient is the client API for SubscriptionInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionInterfaceClient interface {
	// Delete a subscription
	Delete(ctx context.Context, in *SubscriptionDeleteRequest, opts ...grpc.CallOption) (*SubscriptionDeleteResponse, error)
	// Find a subscription by name
	Find(ctx context.Context, in *SubscriptionFindRequest, opts ...grpc.CallOption) (*SubscriptionFindResponse, error)
	// Locate a subscription by id
	Get(ctx context.Context, in *SubscriptionGetRequest, opts ...grpc.CallOption) (*SubscriptionGetResponse, error)
	// Set the burst size of a subscription
	SetBurst(ctx context.Context, in *SubscriptionSetBurstRequest, opts ...grpc.CallOption) (*SubscriptionSetBurstResponse, error)
	// Set the size of a subscription
	SetSize(ctx context.Context, in *SubscriptionSetSizeRequest, opts ...grpc.CallOption) (*SubscriptionSetSizeResponse, error)
}

type subscriptionInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionInterfaceClient(cc grpc.ClientConnInterface) SubscriptionInterfaceClient {
	return &subscriptionInterfaceClient{cc}
}

func (c *subscriptionInterfaceClient) Delete(ctx context.Context, in *SubscriptionDeleteRequest, opts ...grpc.CallOption) (*SubscriptionDeleteResponse, error) {
	out := new(SubscriptionDeleteResponse)
	err := c.cc.Invoke(ctx, SubscriptionInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionInterfaceClient) Find(ctx context.Context, in *SubscriptionFindRequest, opts ...grpc.CallOption) (*SubscriptionFindResponse, error) {
	out := new(SubscriptionFindResponse)
	err := c.cc.Invoke(ctx, SubscriptionInterface_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionInterfaceClient) Get(ctx context.Context, in *SubscriptionGetRequest, opts ...grpc.CallOption) (*SubscriptionGetResponse, error) {
	out := new(SubscriptionGetResponse)
	err := c.cc.Invoke(ctx, SubscriptionInterface_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionInterfaceClient) SetBurst(ctx context.Context, in *SubscriptionSetBurstRequest, opts ...grpc.CallOption) (*SubscriptionSetBurstResponse, error) {
	out := new(SubscriptionSetBurstResponse)
	err := c.cc.Invoke(ctx, SubscriptionInterface_SetBurst_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionInterfaceClient) SetSize(ctx context.Context, in *SubscriptionSetSizeRequest, opts ...grpc.CallOption) (*SubscriptionSetSizeResponse, error) {
	out := new(SubscriptionSetSizeResponse)
	err := c.cc.Invoke(ctx, SubscriptionInterface_SetSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionInterfaceServer is the server API for SubscriptionInterface service.
// All implementations must embed UnimplementedSubscriptionInterfaceServer
// for forward compatibility
type SubscriptionInterfaceServer interface {
	// Delete a subscription
	Delete(context.Context, *SubscriptionDeleteRequest) (*SubscriptionDeleteResponse, error)
	// Find a subscription by name
	Find(context.Context, *SubscriptionFindRequest) (*SubscriptionFindResponse, error)
	// Locate a subscription by id
	Get(context.Context, *SubscriptionGetRequest) (*SubscriptionGetResponse, error)
	// Set the burst size of a subscription
	SetBurst(context.Context, *SubscriptionSetBurstRequest) (*SubscriptionSetBurstResponse, error)
	// Set the size of a subscription
	SetSize(context.Context, *SubscriptionSetSizeRequest) (*SubscriptionSetSizeResponse, error)
	mustEmbedUnimplementedSubscriptionInterfaceServer()
}

// UnimplementedSubscriptionInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionInterfaceServer struct {
}

func (UnimplementedSubscriptionInterfaceServer) Delete(context.Context, *SubscriptionDeleteRequest) (*SubscriptionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSubscriptionInterfaceServer) Find(context.Context, *SubscriptionFindRequest) (*SubscriptionFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedSubscriptionInterfaceServer) Get(context.Context, *SubscriptionGetRequest) (*SubscriptionGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSubscriptionInterfaceServer) SetBurst(context.Context, *SubscriptionSetBurstRequest) (*SubscriptionSetBurstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBurst not implemented")
}
func (UnimplementedSubscriptionInterfaceServer) SetSize(context.Context, *SubscriptionSetSizeRequest) (*SubscriptionSetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSize not implemented")
}
func (UnimplementedSubscriptionInterfaceServer) mustEmbedUnimplementedSubscriptionInterfaceServer() {}

// UnsafeSubscriptionInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionInterfaceServer will
// result in compilation errors.
type UnsafeSubscriptionInterfaceServer interface {
	mustEmbedUnimplementedSubscriptionInterfaceServer()
}

func RegisterSubscriptionInterfaceServer(s grpc.ServiceRegistrar, srv SubscriptionInterfaceServer) {
	s.RegisterService(&SubscriptionInterface_ServiceDesc, srv)
}

func _SubscriptionInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionInterfaceServer).Delete(ctx, req.(*SubscriptionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionInterface_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionInterfaceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionInterface_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionInterfaceServer).Find(ctx, req.(*SubscriptionFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionInterfaceServer).Get(ctx, req.(*SubscriptionGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionInterface_SetBurst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionSetBurstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionInterfaceServer).SetBurst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionInterface_SetBurst_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionInterfaceServer).SetBurst(ctx, req.(*SubscriptionSetBurstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionInterface_SetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionSetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionInterfaceServer).SetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionInterface_SetSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionInterfaceServer).SetSize(ctx, req.(*SubscriptionSetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionInterface_ServiceDesc is the grpc.ServiceDesc for SubscriptionInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.SubscriptionInterface",
	HandlerType: (*SubscriptionInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _SubscriptionInterface_Delete_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _SubscriptionInterface_Find_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SubscriptionInterface_Get_Handler,
		},
		{
			MethodName: "SetBurst",
			Handler:    _SubscriptionInterface_SetBurst_Handler,
		},
		{
			MethodName: "SetSize",
			Handler:    _SubscriptionInterface_SetSize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/subscription.proto",
}
