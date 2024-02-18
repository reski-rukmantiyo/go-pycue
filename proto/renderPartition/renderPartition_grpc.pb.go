// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/renderPartition.proto

package renderPartition

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
	RenderPartitionInterface_Delete_FullMethodName          = "/renderPartition.RenderPartitionInterface/Delete"
	RenderPartitionInterface_SetMaxResources_FullMethodName = "/renderPartition.RenderPartitionInterface/SetMaxResources"
)

// RenderPartitionInterfaceClient is the client API for RenderPartitionInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RenderPartitionInterfaceClient interface {
	// Deletes the host local setup.  Any proc that is running from this should be killed or else the counts
	// could be off.
	Delete(ctx context.Context, in *RenderPartDeleteRequest, opts ...grpc.CallOption) (*RenderPartDeleteResponse, error)
	// Reset the maximum amount of cores and memory for this render partition.
	SetMaxResources(ctx context.Context, in *RenderPartSetMaxResourcesRequest, opts ...grpc.CallOption) (*RenderPartSetMaxResourcesResponse, error)
}

type renderPartitionInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewRenderPartitionInterfaceClient(cc grpc.ClientConnInterface) RenderPartitionInterfaceClient {
	return &renderPartitionInterfaceClient{cc}
}

func (c *renderPartitionInterfaceClient) Delete(ctx context.Context, in *RenderPartDeleteRequest, opts ...grpc.CallOption) (*RenderPartDeleteResponse, error) {
	out := new(RenderPartDeleteResponse)
	err := c.cc.Invoke(ctx, RenderPartitionInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *renderPartitionInterfaceClient) SetMaxResources(ctx context.Context, in *RenderPartSetMaxResourcesRequest, opts ...grpc.CallOption) (*RenderPartSetMaxResourcesResponse, error) {
	out := new(RenderPartSetMaxResourcesResponse)
	err := c.cc.Invoke(ctx, RenderPartitionInterface_SetMaxResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RenderPartitionInterfaceServer is the server API for RenderPartitionInterface service.
// All implementations must embed UnimplementedRenderPartitionInterfaceServer
// for forward compatibility
type RenderPartitionInterfaceServer interface {
	// Deletes the host local setup.  Any proc that is running from this should be killed or else the counts
	// could be off.
	Delete(context.Context, *RenderPartDeleteRequest) (*RenderPartDeleteResponse, error)
	// Reset the maximum amount of cores and memory for this render partition.
	SetMaxResources(context.Context, *RenderPartSetMaxResourcesRequest) (*RenderPartSetMaxResourcesResponse, error)
	mustEmbedUnimplementedRenderPartitionInterfaceServer()
}

// UnimplementedRenderPartitionInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedRenderPartitionInterfaceServer struct {
}

func (UnimplementedRenderPartitionInterfaceServer) Delete(context.Context, *RenderPartDeleteRequest) (*RenderPartDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRenderPartitionInterfaceServer) SetMaxResources(context.Context, *RenderPartSetMaxResourcesRequest) (*RenderPartSetMaxResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMaxResources not implemented")
}
func (UnimplementedRenderPartitionInterfaceServer) mustEmbedUnimplementedRenderPartitionInterfaceServer() {
}

// UnsafeRenderPartitionInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RenderPartitionInterfaceServer will
// result in compilation errors.
type UnsafeRenderPartitionInterfaceServer interface {
	mustEmbedUnimplementedRenderPartitionInterfaceServer()
}

func RegisterRenderPartitionInterfaceServer(s grpc.ServiceRegistrar, srv RenderPartitionInterfaceServer) {
	s.RegisterService(&RenderPartitionInterface_ServiceDesc, srv)
}

func _RenderPartitionInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderPartDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenderPartitionInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RenderPartitionInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenderPartitionInterfaceServer).Delete(ctx, req.(*RenderPartDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RenderPartitionInterface_SetMaxResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderPartSetMaxResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenderPartitionInterfaceServer).SetMaxResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RenderPartitionInterface_SetMaxResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenderPartitionInterfaceServer).SetMaxResources(ctx, req.(*RenderPartSetMaxResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RenderPartitionInterface_ServiceDesc is the grpc.ServiceDesc for RenderPartitionInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RenderPartitionInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "renderPartition.RenderPartitionInterface",
	HandlerType: (*RenderPartitionInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _RenderPartitionInterface_Delete_Handler,
		},
		{
			MethodName: "SetMaxResources",
			Handler:    _RenderPartitionInterface_SetMaxResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/renderPartition.proto",
}