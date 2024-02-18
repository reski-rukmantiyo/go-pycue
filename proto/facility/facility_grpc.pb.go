// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/facility.proto

package facility

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
	FacilityInterface_Create_FullMethodName         = "/facility.FacilityInterface/Create"
	FacilityInterface_Delete_FullMethodName         = "/facility.FacilityInterface/Delete"
	FacilityInterface_Get_FullMethodName            = "/facility.FacilityInterface/Get"
	FacilityInterface_Rename_FullMethodName         = "/facility.FacilityInterface/Rename"
	FacilityInterface_GetAllocations_FullMethodName = "/facility.FacilityInterface/GetAllocations"
)

// FacilityInterfaceClient is the client API for FacilityInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FacilityInterfaceClient interface {
	// Create a facility.
	Create(ctx context.Context, in *FacilityCreateRequest, opts ...grpc.CallOption) (*FacilityCreateResponse, error)
	// A facility is never really deleted, just make inactive.
	Delete(ctx context.Context, in *FacilityDeleteRequest, opts ...grpc.CallOption) (*FacilityDeleteResponse, error)
	// Look up a facility by name.
	Get(ctx context.Context, in *FacilityGetRequest, opts ...grpc.CallOption) (*FacilityGetResponse, error)
	// Rename the facility to a given name.
	Rename(ctx context.Context, in *FacilityRenameRequest, opts ...grpc.CallOption) (*FacilityRenameResponse, error)
	// Get Allocations associated with this facility.
	GetAllocations(ctx context.Context, in *FacilityGetAllocationsRequest, opts ...grpc.CallOption) (*FacilityGetAllocationsResponse, error)
}

type facilityInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewFacilityInterfaceClient(cc grpc.ClientConnInterface) FacilityInterfaceClient {
	return &facilityInterfaceClient{cc}
}

func (c *facilityInterfaceClient) Create(ctx context.Context, in *FacilityCreateRequest, opts ...grpc.CallOption) (*FacilityCreateResponse, error) {
	out := new(FacilityCreateResponse)
	err := c.cc.Invoke(ctx, FacilityInterface_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facilityInterfaceClient) Delete(ctx context.Context, in *FacilityDeleteRequest, opts ...grpc.CallOption) (*FacilityDeleteResponse, error) {
	out := new(FacilityDeleteResponse)
	err := c.cc.Invoke(ctx, FacilityInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facilityInterfaceClient) Get(ctx context.Context, in *FacilityGetRequest, opts ...grpc.CallOption) (*FacilityGetResponse, error) {
	out := new(FacilityGetResponse)
	err := c.cc.Invoke(ctx, FacilityInterface_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facilityInterfaceClient) Rename(ctx context.Context, in *FacilityRenameRequest, opts ...grpc.CallOption) (*FacilityRenameResponse, error) {
	out := new(FacilityRenameResponse)
	err := c.cc.Invoke(ctx, FacilityInterface_Rename_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facilityInterfaceClient) GetAllocations(ctx context.Context, in *FacilityGetAllocationsRequest, opts ...grpc.CallOption) (*FacilityGetAllocationsResponse, error) {
	out := new(FacilityGetAllocationsResponse)
	err := c.cc.Invoke(ctx, FacilityInterface_GetAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FacilityInterfaceServer is the server API for FacilityInterface service.
// All implementations must embed UnimplementedFacilityInterfaceServer
// for forward compatibility
type FacilityInterfaceServer interface {
	// Create a facility.
	Create(context.Context, *FacilityCreateRequest) (*FacilityCreateResponse, error)
	// A facility is never really deleted, just make inactive.
	Delete(context.Context, *FacilityDeleteRequest) (*FacilityDeleteResponse, error)
	// Look up a facility by name.
	Get(context.Context, *FacilityGetRequest) (*FacilityGetResponse, error)
	// Rename the facility to a given name.
	Rename(context.Context, *FacilityRenameRequest) (*FacilityRenameResponse, error)
	// Get Allocations associated with this facility.
	GetAllocations(context.Context, *FacilityGetAllocationsRequest) (*FacilityGetAllocationsResponse, error)
	mustEmbedUnimplementedFacilityInterfaceServer()
}

// UnimplementedFacilityInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedFacilityInterfaceServer struct {
}

func (UnimplementedFacilityInterfaceServer) Create(context.Context, *FacilityCreateRequest) (*FacilityCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFacilityInterfaceServer) Delete(context.Context, *FacilityDeleteRequest) (*FacilityDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFacilityInterfaceServer) Get(context.Context, *FacilityGetRequest) (*FacilityGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFacilityInterfaceServer) Rename(context.Context, *FacilityRenameRequest) (*FacilityRenameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedFacilityInterfaceServer) GetAllocations(context.Context, *FacilityGetAllocationsRequest) (*FacilityGetAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllocations not implemented")
}
func (UnimplementedFacilityInterfaceServer) mustEmbedUnimplementedFacilityInterfaceServer() {}

// UnsafeFacilityInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FacilityInterfaceServer will
// result in compilation errors.
type UnsafeFacilityInterfaceServer interface {
	mustEmbedUnimplementedFacilityInterfaceServer()
}

func RegisterFacilityInterfaceServer(s grpc.ServiceRegistrar, srv FacilityInterfaceServer) {
	s.RegisterService(&FacilityInterface_ServiceDesc, srv)
}

func _FacilityInterface_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilityCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityInterfaceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityInterface_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityInterfaceServer).Create(ctx, req.(*FacilityCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacilityInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilityDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityInterfaceServer).Delete(ctx, req.(*FacilityDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacilityInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilityGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityInterfaceServer).Get(ctx, req.(*FacilityGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacilityInterface_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilityRenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityInterfaceServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityInterface_Rename_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityInterfaceServer).Rename(ctx, req.(*FacilityRenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacilityInterface_GetAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilityGetAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityInterfaceServer).GetAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityInterface_GetAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityInterfaceServer).GetAllocations(ctx, req.(*FacilityGetAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FacilityInterface_ServiceDesc is the grpc.ServiceDesc for FacilityInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FacilityInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "facility.FacilityInterface",
	HandlerType: (*FacilityInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FacilityInterface_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FacilityInterface_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FacilityInterface_Get_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _FacilityInterface_Rename_Handler,
		},
		{
			MethodName: "GetAllocations",
			Handler:    _FacilityInterface_GetAllocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/facility.proto",
}

const (
	AllocationInterface_Create_FullMethodName           = "/facility.AllocationInterface/Create"
	AllocationInterface_Delete_FullMethodName           = "/facility.AllocationInterface/Delete"
	AllocationInterface_Find_FullMethodName             = "/facility.AllocationInterface/Find"
	AllocationInterface_FindHosts_FullMethodName        = "/facility.AllocationInterface/FindHosts"
	AllocationInterface_Get_FullMethodName              = "/facility.AllocationInterface/Get"
	AllocationInterface_GetAll_FullMethodName           = "/facility.AllocationInterface/GetAll"
	AllocationInterface_GetHosts_FullMethodName         = "/facility.AllocationInterface/GetHosts"
	AllocationInterface_GetSubscriptions_FullMethodName = "/facility.AllocationInterface/GetSubscriptions"
	AllocationInterface_ReparentHosts_FullMethodName    = "/facility.AllocationInterface/ReparentHosts"
	AllocationInterface_SetBillable_FullMethodName      = "/facility.AllocationInterface/SetBillable"
	AllocationInterface_SetName_FullMethodName          = "/facility.AllocationInterface/SetName"
	AllocationInterface_SetTag_FullMethodName           = "/facility.AllocationInterface/SetTag"
	AllocationInterface_GetDefault_FullMethodName       = "/facility.AllocationInterface/GetDefault"
	AllocationInterface_SetDefault_FullMethodName       = "/facility.AllocationInterface/SetDefault"
)

// AllocationInterfaceClient is the client API for AllocationInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AllocationInterfaceClient interface {
	Create(ctx context.Context, in *AllocCreateRequest, opts ...grpc.CallOption) (*AllocCreateResponse, error)
	// Delete this allocation
	Delete(ctx context.Context, in *AllocDeleteRequest, opts ...grpc.CallOption) (*AllocDeleteResponse, error)
	// Find an allocation by name
	Find(ctx context.Context, in *AllocFindRequest, opts ...grpc.CallOption) (*AllocFindResponse, error)
	// Use HostSearchCriteria to find a list of hosts
	FindHosts(ctx context.Context, in *AllocFindHostsRequest, opts ...grpc.CallOption) (*AllocFindHostsResponse, error)
	// Look up an allocation by id
	Get(ctx context.Context, in *AllocGetRequest, opts ...grpc.CallOption) (*AllocGetResponse, error)
	// Return a list of all allocations
	GetAll(ctx context.Context, in *AllocGetAllRequest, opts ...grpc.CallOption) (*AllocGetAllResponse, error)
	// Returns the list of hosts in this allocation
	GetHosts(ctx context.Context, in *AllocGetHostsRequest, opts ...grpc.CallOption) (*AllocGetHostsResponse, error)
	// Returns all subscriptions for this allocation
	GetSubscriptions(ctx context.Context, in *AllocGetSubscriptionsRequest, opts ...grpc.CallOption) (*AllocGetSubscriptionsResponse, error)
	// Assigns a list of hosts to this allocation.
	ReparentHosts(ctx context.Context, in *AllocReparentHostsRequest, opts ...grpc.CallOption) (*AllocReparentHostsResponse, error)
	// Set an allocation billable or not billable.
	SetBillable(ctx context.Context, in *AllocSetBillableRequest, opts ...grpc.CallOption) (*AllocSetBillableResponse, error)
	// Set the allocation name
	SetName(ctx context.Context, in *AllocSetNameRequest, opts ...grpc.CallOption) (*AllocSetNameResponse, error)
	// Set the allocation tag.  Setting this will re-tag all the hosts in this allocation.
	SetTag(ctx context.Context, in *AllocSetTagRequest, opts ...grpc.CallOption) (*AllocSetTagResponse, error)
	// Return the default allocation.
	GetDefault(ctx context.Context, in *AllocGetDefaultRequest, opts ...grpc.CallOption) (*AllocGetDefaultResponse, error)
	// Set the default allocation.
	SetDefault(ctx context.Context, in *AllocSetDefaultRequest, opts ...grpc.CallOption) (*AllocSetDefaultResponse, error)
}

type allocationInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewAllocationInterfaceClient(cc grpc.ClientConnInterface) AllocationInterfaceClient {
	return &allocationInterfaceClient{cc}
}

func (c *allocationInterfaceClient) Create(ctx context.Context, in *AllocCreateRequest, opts ...grpc.CallOption) (*AllocCreateResponse, error) {
	out := new(AllocCreateResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) Delete(ctx context.Context, in *AllocDeleteRequest, opts ...grpc.CallOption) (*AllocDeleteResponse, error) {
	out := new(AllocDeleteResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) Find(ctx context.Context, in *AllocFindRequest, opts ...grpc.CallOption) (*AllocFindResponse, error) {
	out := new(AllocFindResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) FindHosts(ctx context.Context, in *AllocFindHostsRequest, opts ...grpc.CallOption) (*AllocFindHostsResponse, error) {
	out := new(AllocFindHostsResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_FindHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) Get(ctx context.Context, in *AllocGetRequest, opts ...grpc.CallOption) (*AllocGetResponse, error) {
	out := new(AllocGetResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) GetAll(ctx context.Context, in *AllocGetAllRequest, opts ...grpc.CallOption) (*AllocGetAllResponse, error) {
	out := new(AllocGetAllResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) GetHosts(ctx context.Context, in *AllocGetHostsRequest, opts ...grpc.CallOption) (*AllocGetHostsResponse, error) {
	out := new(AllocGetHostsResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_GetHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) GetSubscriptions(ctx context.Context, in *AllocGetSubscriptionsRequest, opts ...grpc.CallOption) (*AllocGetSubscriptionsResponse, error) {
	out := new(AllocGetSubscriptionsResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_GetSubscriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) ReparentHosts(ctx context.Context, in *AllocReparentHostsRequest, opts ...grpc.CallOption) (*AllocReparentHostsResponse, error) {
	out := new(AllocReparentHostsResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_ReparentHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) SetBillable(ctx context.Context, in *AllocSetBillableRequest, opts ...grpc.CallOption) (*AllocSetBillableResponse, error) {
	out := new(AllocSetBillableResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_SetBillable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) SetName(ctx context.Context, in *AllocSetNameRequest, opts ...grpc.CallOption) (*AllocSetNameResponse, error) {
	out := new(AllocSetNameResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_SetName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) SetTag(ctx context.Context, in *AllocSetTagRequest, opts ...grpc.CallOption) (*AllocSetTagResponse, error) {
	out := new(AllocSetTagResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_SetTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) GetDefault(ctx context.Context, in *AllocGetDefaultRequest, opts ...grpc.CallOption) (*AllocGetDefaultResponse, error) {
	out := new(AllocGetDefaultResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_GetDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationInterfaceClient) SetDefault(ctx context.Context, in *AllocSetDefaultRequest, opts ...grpc.CallOption) (*AllocSetDefaultResponse, error) {
	out := new(AllocSetDefaultResponse)
	err := c.cc.Invoke(ctx, AllocationInterface_SetDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllocationInterfaceServer is the server API for AllocationInterface service.
// All implementations must embed UnimplementedAllocationInterfaceServer
// for forward compatibility
type AllocationInterfaceServer interface {
	Create(context.Context, *AllocCreateRequest) (*AllocCreateResponse, error)
	// Delete this allocation
	Delete(context.Context, *AllocDeleteRequest) (*AllocDeleteResponse, error)
	// Find an allocation by name
	Find(context.Context, *AllocFindRequest) (*AllocFindResponse, error)
	// Use HostSearchCriteria to find a list of hosts
	FindHosts(context.Context, *AllocFindHostsRequest) (*AllocFindHostsResponse, error)
	// Look up an allocation by id
	Get(context.Context, *AllocGetRequest) (*AllocGetResponse, error)
	// Return a list of all allocations
	GetAll(context.Context, *AllocGetAllRequest) (*AllocGetAllResponse, error)
	// Returns the list of hosts in this allocation
	GetHosts(context.Context, *AllocGetHostsRequest) (*AllocGetHostsResponse, error)
	// Returns all subscriptions for this allocation
	GetSubscriptions(context.Context, *AllocGetSubscriptionsRequest) (*AllocGetSubscriptionsResponse, error)
	// Assigns a list of hosts to this allocation.
	ReparentHosts(context.Context, *AllocReparentHostsRequest) (*AllocReparentHostsResponse, error)
	// Set an allocation billable or not billable.
	SetBillable(context.Context, *AllocSetBillableRequest) (*AllocSetBillableResponse, error)
	// Set the allocation name
	SetName(context.Context, *AllocSetNameRequest) (*AllocSetNameResponse, error)
	// Set the allocation tag.  Setting this will re-tag all the hosts in this allocation.
	SetTag(context.Context, *AllocSetTagRequest) (*AllocSetTagResponse, error)
	// Return the default allocation.
	GetDefault(context.Context, *AllocGetDefaultRequest) (*AllocGetDefaultResponse, error)
	// Set the default allocation.
	SetDefault(context.Context, *AllocSetDefaultRequest) (*AllocSetDefaultResponse, error)
	mustEmbedUnimplementedAllocationInterfaceServer()
}

// UnimplementedAllocationInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedAllocationInterfaceServer struct {
}

func (UnimplementedAllocationInterfaceServer) Create(context.Context, *AllocCreateRequest) (*AllocCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAllocationInterfaceServer) Delete(context.Context, *AllocDeleteRequest) (*AllocDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAllocationInterfaceServer) Find(context.Context, *AllocFindRequest) (*AllocFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedAllocationInterfaceServer) FindHosts(context.Context, *AllocFindHostsRequest) (*AllocFindHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHosts not implemented")
}
func (UnimplementedAllocationInterfaceServer) Get(context.Context, *AllocGetRequest) (*AllocGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAllocationInterfaceServer) GetAll(context.Context, *AllocGetAllRequest) (*AllocGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAllocationInterfaceServer) GetHosts(context.Context, *AllocGetHostsRequest) (*AllocGetHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHosts not implemented")
}
func (UnimplementedAllocationInterfaceServer) GetSubscriptions(context.Context, *AllocGetSubscriptionsRequest) (*AllocGetSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (UnimplementedAllocationInterfaceServer) ReparentHosts(context.Context, *AllocReparentHostsRequest) (*AllocReparentHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReparentHosts not implemented")
}
func (UnimplementedAllocationInterfaceServer) SetBillable(context.Context, *AllocSetBillableRequest) (*AllocSetBillableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBillable not implemented")
}
func (UnimplementedAllocationInterfaceServer) SetName(context.Context, *AllocSetNameRequest) (*AllocSetNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetName not implemented")
}
func (UnimplementedAllocationInterfaceServer) SetTag(context.Context, *AllocSetTagRequest) (*AllocSetTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTag not implemented")
}
func (UnimplementedAllocationInterfaceServer) GetDefault(context.Context, *AllocGetDefaultRequest) (*AllocGetDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefault not implemented")
}
func (UnimplementedAllocationInterfaceServer) SetDefault(context.Context, *AllocSetDefaultRequest) (*AllocSetDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefault not implemented")
}
func (UnimplementedAllocationInterfaceServer) mustEmbedUnimplementedAllocationInterfaceServer() {}

// UnsafeAllocationInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AllocationInterfaceServer will
// result in compilation errors.
type UnsafeAllocationInterfaceServer interface {
	mustEmbedUnimplementedAllocationInterfaceServer()
}

func RegisterAllocationInterfaceServer(s grpc.ServiceRegistrar, srv AllocationInterfaceServer) {
	s.RegisterService(&AllocationInterface_ServiceDesc, srv)
}

func _AllocationInterface_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).Create(ctx, req.(*AllocCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).Delete(ctx, req.(*AllocDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).Find(ctx, req.(*AllocFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_FindHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocFindHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).FindHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_FindHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).FindHosts(ctx, req.(*AllocFindHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).Get(ctx, req.(*AllocGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).GetAll(ctx, req.(*AllocGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_GetHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocGetHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).GetHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_GetHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).GetHosts(ctx, req.(*AllocGetHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_GetSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocGetSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).GetSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_GetSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).GetSubscriptions(ctx, req.(*AllocGetSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_ReparentHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocReparentHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).ReparentHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_ReparentHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).ReparentHosts(ctx, req.(*AllocReparentHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_SetBillable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocSetBillableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).SetBillable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_SetBillable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).SetBillable(ctx, req.(*AllocSetBillableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_SetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocSetNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).SetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_SetName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).SetName(ctx, req.(*AllocSetNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_SetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocSetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).SetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_SetTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).SetTag(ctx, req.(*AllocSetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_GetDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocGetDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).GetDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_GetDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).GetDefault(ctx, req.(*AllocGetDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationInterface_SetDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocSetDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationInterfaceServer).SetDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationInterface_SetDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationInterfaceServer).SetDefault(ctx, req.(*AllocSetDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AllocationInterface_ServiceDesc is the grpc.ServiceDesc for AllocationInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AllocationInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "facility.AllocationInterface",
	HandlerType: (*AllocationInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AllocationInterface_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AllocationInterface_Delete_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _AllocationInterface_Find_Handler,
		},
		{
			MethodName: "FindHosts",
			Handler:    _AllocationInterface_FindHosts_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AllocationInterface_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AllocationInterface_GetAll_Handler,
		},
		{
			MethodName: "GetHosts",
			Handler:    _AllocationInterface_GetHosts_Handler,
		},
		{
			MethodName: "GetSubscriptions",
			Handler:    _AllocationInterface_GetSubscriptions_Handler,
		},
		{
			MethodName: "ReparentHosts",
			Handler:    _AllocationInterface_ReparentHosts_Handler,
		},
		{
			MethodName: "SetBillable",
			Handler:    _AllocationInterface_SetBillable_Handler,
		},
		{
			MethodName: "SetName",
			Handler:    _AllocationInterface_SetName_Handler,
		},
		{
			MethodName: "SetTag",
			Handler:    _AllocationInterface_SetTag_Handler,
		},
		{
			MethodName: "GetDefault",
			Handler:    _AllocationInterface_GetDefault_Handler,
		},
		{
			MethodName: "SetDefault",
			Handler:    _AllocationInterface_SetDefault_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/facility.proto",
}