// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/depend.proto

package depend

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
	DependInterface_GetDepend_FullMethodName = "/depend.DependInterface/GetDepend"
	DependInterface_Satisfy_FullMethodName   = "/depend.DependInterface/Satisfy"
	DependInterface_Unsatisfy_FullMethodName = "/depend.DependInterface/Unsatisfy"
)

// DependInterfaceClient is the client API for DependInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DependInterfaceClient interface {
	// Get a Depend by it's id
	GetDepend(ctx context.Context, in *DependGetDependRequest, opts ...grpc.CallOption) (*DependGetDependResponse, error)
	// Satisfies the dependency which sets any frames waiting on this dependency to the WAITING state.
	Satisfy(ctx context.Context, in *DependSatisfyRequest, opts ...grpc.CallOption) (*DependSatisfyResponse, error)
	// Unsatisfies the dependency making it active again and sets matching frames to DEPEND
	Unsatisfy(ctx context.Context, in *DependUnsatisfyRequest, opts ...grpc.CallOption) (*DependUnsatisfyResponse, error)
}

type dependInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewDependInterfaceClient(cc grpc.ClientConnInterface) DependInterfaceClient {
	return &dependInterfaceClient{cc}
}

func (c *dependInterfaceClient) GetDepend(ctx context.Context, in *DependGetDependRequest, opts ...grpc.CallOption) (*DependGetDependResponse, error) {
	out := new(DependGetDependResponse)
	err := c.cc.Invoke(ctx, DependInterface_GetDepend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dependInterfaceClient) Satisfy(ctx context.Context, in *DependSatisfyRequest, opts ...grpc.CallOption) (*DependSatisfyResponse, error) {
	out := new(DependSatisfyResponse)
	err := c.cc.Invoke(ctx, DependInterface_Satisfy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dependInterfaceClient) Unsatisfy(ctx context.Context, in *DependUnsatisfyRequest, opts ...grpc.CallOption) (*DependUnsatisfyResponse, error) {
	out := new(DependUnsatisfyResponse)
	err := c.cc.Invoke(ctx, DependInterface_Unsatisfy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DependInterfaceServer is the server API for DependInterface service.
// All implementations must embed UnimplementedDependInterfaceServer
// for forward compatibility
type DependInterfaceServer interface {
	// Get a Depend by it's id
	GetDepend(context.Context, *DependGetDependRequest) (*DependGetDependResponse, error)
	// Satisfies the dependency which sets any frames waiting on this dependency to the WAITING state.
	Satisfy(context.Context, *DependSatisfyRequest) (*DependSatisfyResponse, error)
	// Unsatisfies the dependency making it active again and sets matching frames to DEPEND
	Unsatisfy(context.Context, *DependUnsatisfyRequest) (*DependUnsatisfyResponse, error)
	mustEmbedUnimplementedDependInterfaceServer()
}

// UnimplementedDependInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedDependInterfaceServer struct {
}

func (UnimplementedDependInterfaceServer) GetDepend(context.Context, *DependGetDependRequest) (*DependGetDependResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepend not implemented")
}
func (UnimplementedDependInterfaceServer) Satisfy(context.Context, *DependSatisfyRequest) (*DependSatisfyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Satisfy not implemented")
}
func (UnimplementedDependInterfaceServer) Unsatisfy(context.Context, *DependUnsatisfyRequest) (*DependUnsatisfyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsatisfy not implemented")
}
func (UnimplementedDependInterfaceServer) mustEmbedUnimplementedDependInterfaceServer() {}

// UnsafeDependInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DependInterfaceServer will
// result in compilation errors.
type UnsafeDependInterfaceServer interface {
	mustEmbedUnimplementedDependInterfaceServer()
}

func RegisterDependInterfaceServer(s grpc.ServiceRegistrar, srv DependInterfaceServer) {
	s.RegisterService(&DependInterface_ServiceDesc, srv)
}

func _DependInterface_GetDepend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependGetDependRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependInterfaceServer).GetDepend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DependInterface_GetDepend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependInterfaceServer).GetDepend(ctx, req.(*DependGetDependRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DependInterface_Satisfy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependSatisfyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependInterfaceServer).Satisfy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DependInterface_Satisfy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependInterfaceServer).Satisfy(ctx, req.(*DependSatisfyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DependInterface_Unsatisfy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependUnsatisfyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependInterfaceServer).Unsatisfy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DependInterface_Unsatisfy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependInterfaceServer).Unsatisfy(ctx, req.(*DependUnsatisfyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DependInterface_ServiceDesc is the grpc.ServiceDesc for DependInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DependInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "depend.DependInterface",
	HandlerType: (*DependInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDepend",
			Handler:    _DependInterface_GetDepend_Handler,
		},
		{
			MethodName: "Satisfy",
			Handler:    _DependInterface_Satisfy_Handler,
		},
		{
			MethodName: "Unsatisfy",
			Handler:    _DependInterface_Unsatisfy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/depend.proto",
}
