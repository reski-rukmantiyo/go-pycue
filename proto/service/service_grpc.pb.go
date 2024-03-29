// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/service.proto

package service

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
	ServiceInterface_CreateService_FullMethodName      = "/service.ServiceInterface/CreateService"
	ServiceInterface_Delete_FullMethodName             = "/service.ServiceInterface/Delete"
	ServiceInterface_GetDefaultServices_FullMethodName = "/service.ServiceInterface/GetDefaultServices"
	ServiceInterface_GetService_FullMethodName         = "/service.ServiceInterface/GetService"
	ServiceInterface_Update_FullMethodName             = "/service.ServiceInterface/Update"
)

// ServiceInterfaceClient is the client API for ServiceInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceInterfaceClient interface {
	// Create a new service.
	CreateService(ctx context.Context, in *ServiceCreateServiceRequest, opts ...grpc.CallOption) (*ServiceCreateServiceResponse, error)
	// Delete the provided service
	Delete(ctx context.Context, in *ServiceDeleteRequest, opts ...grpc.CallOption) (*ServiceDeleteResponse, error)
	// Return a the list of default services.
	GetDefaultServices(ctx context.Context, in *ServiceGetDefaultServicesRequest, opts ...grpc.CallOption) (*ServiceGetDefaultServicesResponse, error)
	// Return the given service using its name or unique id.
	GetService(ctx context.Context, in *ServiceGetServiceRequest, opts ...grpc.CallOption) (*ServiceGetServiceResponse, error)
	// Update the provided service
	Update(ctx context.Context, in *ServiceUpdateRequest, opts ...grpc.CallOption) (*ServiceUpdateResponse, error)
}

type serviceInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceInterfaceClient(cc grpc.ClientConnInterface) ServiceInterfaceClient {
	return &serviceInterfaceClient{cc}
}

func (c *serviceInterfaceClient) CreateService(ctx context.Context, in *ServiceCreateServiceRequest, opts ...grpc.CallOption) (*ServiceCreateServiceResponse, error) {
	out := new(ServiceCreateServiceResponse)
	err := c.cc.Invoke(ctx, ServiceInterface_CreateService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) Delete(ctx context.Context, in *ServiceDeleteRequest, opts ...grpc.CallOption) (*ServiceDeleteResponse, error) {
	out := new(ServiceDeleteResponse)
	err := c.cc.Invoke(ctx, ServiceInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) GetDefaultServices(ctx context.Context, in *ServiceGetDefaultServicesRequest, opts ...grpc.CallOption) (*ServiceGetDefaultServicesResponse, error) {
	out := new(ServiceGetDefaultServicesResponse)
	err := c.cc.Invoke(ctx, ServiceInterface_GetDefaultServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) GetService(ctx context.Context, in *ServiceGetServiceRequest, opts ...grpc.CallOption) (*ServiceGetServiceResponse, error) {
	out := new(ServiceGetServiceResponse)
	err := c.cc.Invoke(ctx, ServiceInterface_GetService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) Update(ctx context.Context, in *ServiceUpdateRequest, opts ...grpc.CallOption) (*ServiceUpdateResponse, error) {
	out := new(ServiceUpdateResponse)
	err := c.cc.Invoke(ctx, ServiceInterface_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceInterfaceServer is the server API for ServiceInterface service.
// All implementations must embed UnimplementedServiceInterfaceServer
// for forward compatibility
type ServiceInterfaceServer interface {
	// Create a new service.
	CreateService(context.Context, *ServiceCreateServiceRequest) (*ServiceCreateServiceResponse, error)
	// Delete the provided service
	Delete(context.Context, *ServiceDeleteRequest) (*ServiceDeleteResponse, error)
	// Return a the list of default services.
	GetDefaultServices(context.Context, *ServiceGetDefaultServicesRequest) (*ServiceGetDefaultServicesResponse, error)
	// Return the given service using its name or unique id.
	GetService(context.Context, *ServiceGetServiceRequest) (*ServiceGetServiceResponse, error)
	// Update the provided service
	Update(context.Context, *ServiceUpdateRequest) (*ServiceUpdateResponse, error)
	mustEmbedUnimplementedServiceInterfaceServer()
}

// UnimplementedServiceInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceInterfaceServer struct {
}

func (UnimplementedServiceInterfaceServer) CreateService(context.Context, *ServiceCreateServiceRequest) (*ServiceCreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceInterfaceServer) Delete(context.Context, *ServiceDeleteRequest) (*ServiceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServiceInterfaceServer) GetDefaultServices(context.Context, *ServiceGetDefaultServicesRequest) (*ServiceGetDefaultServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultServices not implemented")
}
func (UnimplementedServiceInterfaceServer) GetService(context.Context, *ServiceGetServiceRequest) (*ServiceGetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServiceInterfaceServer) Update(context.Context, *ServiceUpdateRequest) (*ServiceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServiceInterfaceServer) mustEmbedUnimplementedServiceInterfaceServer() {}

// UnsafeServiceInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceInterfaceServer will
// result in compilation errors.
type UnsafeServiceInterfaceServer interface {
	mustEmbedUnimplementedServiceInterfaceServer()
}

func RegisterServiceInterfaceServer(s grpc.ServiceRegistrar, srv ServiceInterfaceServer) {
	s.RegisterService(&ServiceInterface_ServiceDesc, srv)
}

func _ServiceInterface_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceCreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).CreateService(ctx, req.(*ServiceCreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).Delete(ctx, req.(*ServiceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_GetDefaultServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceGetDefaultServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).GetDefaultServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_GetDefaultServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).GetDefaultServices(ctx, req.(*ServiceGetDefaultServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceGetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).GetService(ctx, req.(*ServiceGetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).Update(ctx, req.(*ServiceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceInterface_ServiceDesc is the grpc.ServiceDesc for ServiceInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ServiceInterface",
	HandlerType: (*ServiceInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _ServiceInterface_CreateService_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServiceInterface_Delete_Handler,
		},
		{
			MethodName: "GetDefaultServices",
			Handler:    _ServiceInterface_GetDefaultServices_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ServiceInterface_GetService_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceInterface_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/service.proto",
}

const (
	ServiceOverrideInterface_Delete_FullMethodName = "/service.ServiceOverrideInterface/Delete"
	ServiceOverrideInterface_Update_FullMethodName = "/service.ServiceOverrideInterface/Update"
)

// ServiceOverrideInterfaceClient is the client API for ServiceOverrideInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceOverrideInterfaceClient interface {
	// Delete the service override
	Delete(ctx context.Context, in *ServiceOverrideDeleteRequest, opts ...grpc.CallOption) (*ServiceOverrideDeleteResponse, error)
	// Update this service override
	Update(ctx context.Context, in *ServiceOverrideUpdateRequest, opts ...grpc.CallOption) (*ServiceOverrideUpdateResponse, error)
}

type serviceOverrideInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceOverrideInterfaceClient(cc grpc.ClientConnInterface) ServiceOverrideInterfaceClient {
	return &serviceOverrideInterfaceClient{cc}
}

func (c *serviceOverrideInterfaceClient) Delete(ctx context.Context, in *ServiceOverrideDeleteRequest, opts ...grpc.CallOption) (*ServiceOverrideDeleteResponse, error) {
	out := new(ServiceOverrideDeleteResponse)
	err := c.cc.Invoke(ctx, ServiceOverrideInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceOverrideInterfaceClient) Update(ctx context.Context, in *ServiceOverrideUpdateRequest, opts ...grpc.CallOption) (*ServiceOverrideUpdateResponse, error) {
	out := new(ServiceOverrideUpdateResponse)
	err := c.cc.Invoke(ctx, ServiceOverrideInterface_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceOverrideInterfaceServer is the server API for ServiceOverrideInterface service.
// All implementations must embed UnimplementedServiceOverrideInterfaceServer
// for forward compatibility
type ServiceOverrideInterfaceServer interface {
	// Delete the service override
	Delete(context.Context, *ServiceOverrideDeleteRequest) (*ServiceOverrideDeleteResponse, error)
	// Update this service override
	Update(context.Context, *ServiceOverrideUpdateRequest) (*ServiceOverrideUpdateResponse, error)
	mustEmbedUnimplementedServiceOverrideInterfaceServer()
}

// UnimplementedServiceOverrideInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceOverrideInterfaceServer struct {
}

func (UnimplementedServiceOverrideInterfaceServer) Delete(context.Context, *ServiceOverrideDeleteRequest) (*ServiceOverrideDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServiceOverrideInterfaceServer) Update(context.Context, *ServiceOverrideUpdateRequest) (*ServiceOverrideUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServiceOverrideInterfaceServer) mustEmbedUnimplementedServiceOverrideInterfaceServer() {
}

// UnsafeServiceOverrideInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceOverrideInterfaceServer will
// result in compilation errors.
type UnsafeServiceOverrideInterfaceServer interface {
	mustEmbedUnimplementedServiceOverrideInterfaceServer()
}

func RegisterServiceOverrideInterfaceServer(s grpc.ServiceRegistrar, srv ServiceOverrideInterfaceServer) {
	s.RegisterService(&ServiceOverrideInterface_ServiceDesc, srv)
}

func _ServiceOverrideInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceOverrideDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceOverrideInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceOverrideInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceOverrideInterfaceServer).Delete(ctx, req.(*ServiceOverrideDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceOverrideInterface_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceOverrideUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceOverrideInterfaceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceOverrideInterface_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceOverrideInterfaceServer).Update(ctx, req.(*ServiceOverrideUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceOverrideInterface_ServiceDesc is the grpc.ServiceDesc for ServiceOverrideInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceOverrideInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ServiceOverrideInterface",
	HandlerType: (*ServiceOverrideInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _ServiceOverrideInterface_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceOverrideInterface_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/service.proto",
}
