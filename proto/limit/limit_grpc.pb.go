// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/limit.proto

package limit

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
	LimitInterface_Create_FullMethodName      = "/limit.LimitInterface/Create"
	LimitInterface_Delete_FullMethodName      = "/limit.LimitInterface/Delete"
	LimitInterface_Find_FullMethodName        = "/limit.LimitInterface/Find"
	LimitInterface_Get_FullMethodName         = "/limit.LimitInterface/Get"
	LimitInterface_GetAll_FullMethodName      = "/limit.LimitInterface/GetAll"
	LimitInterface_Rename_FullMethodName      = "/limit.LimitInterface/Rename"
	LimitInterface_SetMaxValue_FullMethodName = "/limit.LimitInterface/SetMaxValue"
)

// LimitInterfaceClient is the client API for LimitInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LimitInterfaceClient interface {
	// Create a new limit with a name and max value
	Create(ctx context.Context, in *LimitCreateRequest, opts ...grpc.CallOption) (*LimitCreateResponse, error)
	// Delete the specified limit.
	Delete(ctx context.Context, in *LimitDeleteRequest, opts ...grpc.CallOption) (*LimitDeleteResponse, error)
	// Find a limit by name
	Find(ctx context.Context, in *LimitFindRequest, opts ...grpc.CallOption) (*LimitFindResponse, error)
	// Get a limit by id
	Get(ctx context.Context, in *LimitGetRequest, opts ...grpc.CallOption) (*LimitGetResponse, error)
	// Get all limits
	GetAll(ctx context.Context, in *LimitGetAllRequest, opts ...grpc.CallOption) (*LimitGetAllResponse, error)
	// Rename an existing limit
	Rename(ctx context.Context, in *LimitRenameRequest, opts ...grpc.CallOption) (*LimitRenameResponse, error)
	SetMaxValue(ctx context.Context, in *LimitSetMaxValueRequest, opts ...grpc.CallOption) (*LimitSetMaxValueResponse, error)
}

type limitInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewLimitInterfaceClient(cc grpc.ClientConnInterface) LimitInterfaceClient {
	return &limitInterfaceClient{cc}
}

func (c *limitInterfaceClient) Create(ctx context.Context, in *LimitCreateRequest, opts ...grpc.CallOption) (*LimitCreateResponse, error) {
	out := new(LimitCreateResponse)
	err := c.cc.Invoke(ctx, LimitInterface_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitInterfaceClient) Delete(ctx context.Context, in *LimitDeleteRequest, opts ...grpc.CallOption) (*LimitDeleteResponse, error) {
	out := new(LimitDeleteResponse)
	err := c.cc.Invoke(ctx, LimitInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitInterfaceClient) Find(ctx context.Context, in *LimitFindRequest, opts ...grpc.CallOption) (*LimitFindResponse, error) {
	out := new(LimitFindResponse)
	err := c.cc.Invoke(ctx, LimitInterface_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitInterfaceClient) Get(ctx context.Context, in *LimitGetRequest, opts ...grpc.CallOption) (*LimitGetResponse, error) {
	out := new(LimitGetResponse)
	err := c.cc.Invoke(ctx, LimitInterface_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitInterfaceClient) GetAll(ctx context.Context, in *LimitGetAllRequest, opts ...grpc.CallOption) (*LimitGetAllResponse, error) {
	out := new(LimitGetAllResponse)
	err := c.cc.Invoke(ctx, LimitInterface_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitInterfaceClient) Rename(ctx context.Context, in *LimitRenameRequest, opts ...grpc.CallOption) (*LimitRenameResponse, error) {
	out := new(LimitRenameResponse)
	err := c.cc.Invoke(ctx, LimitInterface_Rename_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitInterfaceClient) SetMaxValue(ctx context.Context, in *LimitSetMaxValueRequest, opts ...grpc.CallOption) (*LimitSetMaxValueResponse, error) {
	out := new(LimitSetMaxValueResponse)
	err := c.cc.Invoke(ctx, LimitInterface_SetMaxValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LimitInterfaceServer is the server API for LimitInterface service.
// All implementations must embed UnimplementedLimitInterfaceServer
// for forward compatibility
type LimitInterfaceServer interface {
	// Create a new limit with a name and max value
	Create(context.Context, *LimitCreateRequest) (*LimitCreateResponse, error)
	// Delete the specified limit.
	Delete(context.Context, *LimitDeleteRequest) (*LimitDeleteResponse, error)
	// Find a limit by name
	Find(context.Context, *LimitFindRequest) (*LimitFindResponse, error)
	// Get a limit by id
	Get(context.Context, *LimitGetRequest) (*LimitGetResponse, error)
	// Get all limits
	GetAll(context.Context, *LimitGetAllRequest) (*LimitGetAllResponse, error)
	// Rename an existing limit
	Rename(context.Context, *LimitRenameRequest) (*LimitRenameResponse, error)
	SetMaxValue(context.Context, *LimitSetMaxValueRequest) (*LimitSetMaxValueResponse, error)
	mustEmbedUnimplementedLimitInterfaceServer()
}

// UnimplementedLimitInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedLimitInterfaceServer struct {
}

func (UnimplementedLimitInterfaceServer) Create(context.Context, *LimitCreateRequest) (*LimitCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLimitInterfaceServer) Delete(context.Context, *LimitDeleteRequest) (*LimitDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLimitInterfaceServer) Find(context.Context, *LimitFindRequest) (*LimitFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedLimitInterfaceServer) Get(context.Context, *LimitGetRequest) (*LimitGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLimitInterfaceServer) GetAll(context.Context, *LimitGetAllRequest) (*LimitGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedLimitInterfaceServer) Rename(context.Context, *LimitRenameRequest) (*LimitRenameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedLimitInterfaceServer) SetMaxValue(context.Context, *LimitSetMaxValueRequest) (*LimitSetMaxValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMaxValue not implemented")
}
func (UnimplementedLimitInterfaceServer) mustEmbedUnimplementedLimitInterfaceServer() {}

// UnsafeLimitInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LimitInterfaceServer will
// result in compilation errors.
type UnsafeLimitInterfaceServer interface {
	mustEmbedUnimplementedLimitInterfaceServer()
}

func RegisterLimitInterfaceServer(s grpc.ServiceRegistrar, srv LimitInterfaceServer) {
	s.RegisterService(&LimitInterface_ServiceDesc, srv)
}

func _LimitInterface_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).Create(ctx, req.(*LimitCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).Delete(ctx, req.(*LimitDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitInterface_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).Find(ctx, req.(*LimitFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).Get(ctx, req.(*LimitGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitInterface_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).GetAll(ctx, req.(*LimitGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitInterface_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitRenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_Rename_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).Rename(ctx, req.(*LimitRenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitInterface_SetMaxValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitSetMaxValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitInterfaceServer).SetMaxValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitInterface_SetMaxValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitInterfaceServer).SetMaxValue(ctx, req.(*LimitSetMaxValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LimitInterface_ServiceDesc is the grpc.ServiceDesc for LimitInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LimitInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "limit.LimitInterface",
	HandlerType: (*LimitInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LimitInterface_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LimitInterface_Delete_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _LimitInterface_Find_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LimitInterface_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _LimitInterface_GetAll_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _LimitInterface_Rename_Handler,
		},
		{
			MethodName: "SetMaxValue",
			Handler:    _LimitInterface_SetMaxValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/limit.proto",
}
