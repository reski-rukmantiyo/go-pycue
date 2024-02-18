// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/department.proto

package department

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
	DepartmentInterface_AddDepartmentName_FullMethodName    = "/department.DepartmentInterface/AddDepartmentName"
	DepartmentInterface_AddTask_FullMethodName              = "/department.DepartmentInterface/AddTask"
	DepartmentInterface_AddTasks_FullMethodName             = "/department.DepartmentInterface/AddTasks"
	DepartmentInterface_ClearTasks_FullMethodName           = "/department.DepartmentInterface/ClearTasks"
	DepartmentInterface_ClearTaskAdjustments_FullMethodName = "/department.DepartmentInterface/ClearTaskAdjustments"
	DepartmentInterface_Delete_FullMethodName               = "/department.DepartmentInterface/Delete"
	DepartmentInterface_DisableTiManaged_FullMethodName     = "/department.DepartmentInterface/DisableTiManaged"
	DepartmentInterface_EnableTiManaged_FullMethodName      = "/department.DepartmentInterface/EnableTiManaged"
	DepartmentInterface_GetDepartmentNames_FullMethodName   = "/department.DepartmentInterface/GetDepartmentNames"
	DepartmentInterface_GetTasks_FullMethodName             = "/department.DepartmentInterface/GetTasks"
	DepartmentInterface_RemoveDepartmentName_FullMethodName = "/department.DepartmentInterface/RemoveDepartmentName"
	DepartmentInterface_ReplaceTasks_FullMethodName         = "/department.DepartmentInterface/ReplaceTasks"
	DepartmentInterface_SetMangedCores_FullMethodName       = "/department.DepartmentInterface/SetMangedCores"
)

// DepartmentInterfaceClient is the client API for DepartmentInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentInterfaceClient interface {
	// Add a new department name
	AddDepartmentName(ctx context.Context, in *DeptAddDeptNameRequest, opts ...grpc.CallOption) (*DeptAddDeptNameResponse, error)
	// Adds a task to this department and return it
	AddTask(ctx context.Context, in *DeptAddTaskRequest, opts ...grpc.CallOption) (*DeptAddTaskResponse, error)
	// Adds a map of tasks to this deparmtment and return them as a list.  Offers a quick way to add many tasks with one
	// function.
	AddTasks(ctx context.Context, in *DeptAddTasksRequest, opts ...grpc.CallOption) (*DeptAddTasksResponse, error)
	// Clears all tasks
	ClearTasks(ctx context.Context, in *DeptClearTasksRequest, opts ...grpc.CallOption) (*DeptClearTasksResponse, error)
	// Clears all manual task adjustments to managed tasks.  This won't do anything unless your using Ti Managed tasks.
	ClearTaskAdjustments(ctx context.Context, in *DeptClearTaskAdjustmentsRequest, opts ...grpc.CallOption) (*DeptClearTaskAdjustmentsResponse, error)
	// Delete the department
	Delete(ctx context.Context, in *DeptDeleteRequest, opts ...grpc.CallOption) (*DeptDeleteResponse, error)
	// Disable Track-It management.  This will also clear all tasks.
	DisableTiManaged(ctx context.Context, in *DeptDisableTiManagedRequest, opts ...grpc.CallOption) (*DeptDisableTiManagedResponse, error)
	// Enable Track-It management.  This will pull a task list from Track-It and keep it synced.
	EnableTiManaged(ctx context.Context, in *DeptEnableTiManagedRequest, opts ...grpc.CallOption) (*DeptEnableTiManagedResponse, error)
	// Return a list of department names
	GetDepartmentNames(ctx context.Context, in *DeptGetDepartmentNamesRequest, opts ...grpc.CallOption) (*DeptGetDepartmentNamesResponse, error)
	// Returns the list of tasks for this department.
	GetTasks(ctx context.Context, in *DeptGetTasksRequest, opts ...grpc.CallOption) (*DeptGetTasksResponse, error)
	// Remove the department of this name
	RemoveDepartmentName(ctx context.Context, in *DeptRemoveDepartmentNameRequest, opts ...grpc.CallOption) (*DeptRemoveDepartmentNameResponse, error)
	// Replaces a map of tasks.  If the task already exists its updated, else its inserted.
	ReplaceTasks(ctx context.Context, in *DeptReplaceTaskRequest, opts ...grpc.CallOption) (*DeptReplaceTaskResponse, error)
	// Sets the minimum number of cores for the department to manage between its tasks.
	SetMangedCores(ctx context.Context, in *DeptSetManagedCoresRequest, opts ...grpc.CallOption) (*DeptSetManagedCoresResponse, error)
}

type departmentInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentInterfaceClient(cc grpc.ClientConnInterface) DepartmentInterfaceClient {
	return &departmentInterfaceClient{cc}
}

func (c *departmentInterfaceClient) AddDepartmentName(ctx context.Context, in *DeptAddDeptNameRequest, opts ...grpc.CallOption) (*DeptAddDeptNameResponse, error) {
	out := new(DeptAddDeptNameResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_AddDepartmentName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) AddTask(ctx context.Context, in *DeptAddTaskRequest, opts ...grpc.CallOption) (*DeptAddTaskResponse, error) {
	out := new(DeptAddTaskResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_AddTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) AddTasks(ctx context.Context, in *DeptAddTasksRequest, opts ...grpc.CallOption) (*DeptAddTasksResponse, error) {
	out := new(DeptAddTasksResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_AddTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) ClearTasks(ctx context.Context, in *DeptClearTasksRequest, opts ...grpc.CallOption) (*DeptClearTasksResponse, error) {
	out := new(DeptClearTasksResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_ClearTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) ClearTaskAdjustments(ctx context.Context, in *DeptClearTaskAdjustmentsRequest, opts ...grpc.CallOption) (*DeptClearTaskAdjustmentsResponse, error) {
	out := new(DeptClearTaskAdjustmentsResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_ClearTaskAdjustments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) Delete(ctx context.Context, in *DeptDeleteRequest, opts ...grpc.CallOption) (*DeptDeleteResponse, error) {
	out := new(DeptDeleteResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) DisableTiManaged(ctx context.Context, in *DeptDisableTiManagedRequest, opts ...grpc.CallOption) (*DeptDisableTiManagedResponse, error) {
	out := new(DeptDisableTiManagedResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_DisableTiManaged_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) EnableTiManaged(ctx context.Context, in *DeptEnableTiManagedRequest, opts ...grpc.CallOption) (*DeptEnableTiManagedResponse, error) {
	out := new(DeptEnableTiManagedResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_EnableTiManaged_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) GetDepartmentNames(ctx context.Context, in *DeptGetDepartmentNamesRequest, opts ...grpc.CallOption) (*DeptGetDepartmentNamesResponse, error) {
	out := new(DeptGetDepartmentNamesResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_GetDepartmentNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) GetTasks(ctx context.Context, in *DeptGetTasksRequest, opts ...grpc.CallOption) (*DeptGetTasksResponse, error) {
	out := new(DeptGetTasksResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_GetTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) RemoveDepartmentName(ctx context.Context, in *DeptRemoveDepartmentNameRequest, opts ...grpc.CallOption) (*DeptRemoveDepartmentNameResponse, error) {
	out := new(DeptRemoveDepartmentNameResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_RemoveDepartmentName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) ReplaceTasks(ctx context.Context, in *DeptReplaceTaskRequest, opts ...grpc.CallOption) (*DeptReplaceTaskResponse, error) {
	out := new(DeptReplaceTaskResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_ReplaceTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentInterfaceClient) SetMangedCores(ctx context.Context, in *DeptSetManagedCoresRequest, opts ...grpc.CallOption) (*DeptSetManagedCoresResponse, error) {
	out := new(DeptSetManagedCoresResponse)
	err := c.cc.Invoke(ctx, DepartmentInterface_SetMangedCores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentInterfaceServer is the server API for DepartmentInterface service.
// All implementations must embed UnimplementedDepartmentInterfaceServer
// for forward compatibility
type DepartmentInterfaceServer interface {
	// Add a new department name
	AddDepartmentName(context.Context, *DeptAddDeptNameRequest) (*DeptAddDeptNameResponse, error)
	// Adds a task to this department and return it
	AddTask(context.Context, *DeptAddTaskRequest) (*DeptAddTaskResponse, error)
	// Adds a map of tasks to this deparmtment and return them as a list.  Offers a quick way to add many tasks with one
	// function.
	AddTasks(context.Context, *DeptAddTasksRequest) (*DeptAddTasksResponse, error)
	// Clears all tasks
	ClearTasks(context.Context, *DeptClearTasksRequest) (*DeptClearTasksResponse, error)
	// Clears all manual task adjustments to managed tasks.  This won't do anything unless your using Ti Managed tasks.
	ClearTaskAdjustments(context.Context, *DeptClearTaskAdjustmentsRequest) (*DeptClearTaskAdjustmentsResponse, error)
	// Delete the department
	Delete(context.Context, *DeptDeleteRequest) (*DeptDeleteResponse, error)
	// Disable Track-It management.  This will also clear all tasks.
	DisableTiManaged(context.Context, *DeptDisableTiManagedRequest) (*DeptDisableTiManagedResponse, error)
	// Enable Track-It management.  This will pull a task list from Track-It and keep it synced.
	EnableTiManaged(context.Context, *DeptEnableTiManagedRequest) (*DeptEnableTiManagedResponse, error)
	// Return a list of department names
	GetDepartmentNames(context.Context, *DeptGetDepartmentNamesRequest) (*DeptGetDepartmentNamesResponse, error)
	// Returns the list of tasks for this department.
	GetTasks(context.Context, *DeptGetTasksRequest) (*DeptGetTasksResponse, error)
	// Remove the department of this name
	RemoveDepartmentName(context.Context, *DeptRemoveDepartmentNameRequest) (*DeptRemoveDepartmentNameResponse, error)
	// Replaces a map of tasks.  If the task already exists its updated, else its inserted.
	ReplaceTasks(context.Context, *DeptReplaceTaskRequest) (*DeptReplaceTaskResponse, error)
	// Sets the minimum number of cores for the department to manage between its tasks.
	SetMangedCores(context.Context, *DeptSetManagedCoresRequest) (*DeptSetManagedCoresResponse, error)
	mustEmbedUnimplementedDepartmentInterfaceServer()
}

// UnimplementedDepartmentInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedDepartmentInterfaceServer struct {
}

func (UnimplementedDepartmentInterfaceServer) AddDepartmentName(context.Context, *DeptAddDeptNameRequest) (*DeptAddDeptNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDepartmentName not implemented")
}
func (UnimplementedDepartmentInterfaceServer) AddTask(context.Context, *DeptAddTaskRequest) (*DeptAddTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedDepartmentInterfaceServer) AddTasks(context.Context, *DeptAddTasksRequest) (*DeptAddTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTasks not implemented")
}
func (UnimplementedDepartmentInterfaceServer) ClearTasks(context.Context, *DeptClearTasksRequest) (*DeptClearTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearTasks not implemented")
}
func (UnimplementedDepartmentInterfaceServer) ClearTaskAdjustments(context.Context, *DeptClearTaskAdjustmentsRequest) (*DeptClearTaskAdjustmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearTaskAdjustments not implemented")
}
func (UnimplementedDepartmentInterfaceServer) Delete(context.Context, *DeptDeleteRequest) (*DeptDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDepartmentInterfaceServer) DisableTiManaged(context.Context, *DeptDisableTiManagedRequest) (*DeptDisableTiManagedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableTiManaged not implemented")
}
func (UnimplementedDepartmentInterfaceServer) EnableTiManaged(context.Context, *DeptEnableTiManagedRequest) (*DeptEnableTiManagedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableTiManaged not implemented")
}
func (UnimplementedDepartmentInterfaceServer) GetDepartmentNames(context.Context, *DeptGetDepartmentNamesRequest) (*DeptGetDepartmentNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartmentNames not implemented")
}
func (UnimplementedDepartmentInterfaceServer) GetTasks(context.Context, *DeptGetTasksRequest) (*DeptGetTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedDepartmentInterfaceServer) RemoveDepartmentName(context.Context, *DeptRemoveDepartmentNameRequest) (*DeptRemoveDepartmentNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDepartmentName not implemented")
}
func (UnimplementedDepartmentInterfaceServer) ReplaceTasks(context.Context, *DeptReplaceTaskRequest) (*DeptReplaceTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceTasks not implemented")
}
func (UnimplementedDepartmentInterfaceServer) SetMangedCores(context.Context, *DeptSetManagedCoresRequest) (*DeptSetManagedCoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMangedCores not implemented")
}
func (UnimplementedDepartmentInterfaceServer) mustEmbedUnimplementedDepartmentInterfaceServer() {}

// UnsafeDepartmentInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentInterfaceServer will
// result in compilation errors.
type UnsafeDepartmentInterfaceServer interface {
	mustEmbedUnimplementedDepartmentInterfaceServer()
}

func RegisterDepartmentInterfaceServer(s grpc.ServiceRegistrar, srv DepartmentInterfaceServer) {
	s.RegisterService(&DepartmentInterface_ServiceDesc, srv)
}

func _DepartmentInterface_AddDepartmentName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptAddDeptNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).AddDepartmentName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_AddDepartmentName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).AddDepartmentName(ctx, req.(*DeptAddDeptNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptAddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_AddTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).AddTask(ctx, req.(*DeptAddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_AddTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptAddTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).AddTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_AddTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).AddTasks(ctx, req.(*DeptAddTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_ClearTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptClearTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).ClearTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_ClearTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).ClearTasks(ctx, req.(*DeptClearTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_ClearTaskAdjustments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptClearTaskAdjustmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).ClearTaskAdjustments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_ClearTaskAdjustments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).ClearTaskAdjustments(ctx, req.(*DeptClearTaskAdjustmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).Delete(ctx, req.(*DeptDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_DisableTiManaged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptDisableTiManagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).DisableTiManaged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_DisableTiManaged_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).DisableTiManaged(ctx, req.(*DeptDisableTiManagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_EnableTiManaged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptEnableTiManagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).EnableTiManaged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_EnableTiManaged_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).EnableTiManaged(ctx, req.(*DeptEnableTiManagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_GetDepartmentNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptGetDepartmentNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).GetDepartmentNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_GetDepartmentNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).GetDepartmentNames(ctx, req.(*DeptGetDepartmentNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptGetTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_GetTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).GetTasks(ctx, req.(*DeptGetTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_RemoveDepartmentName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptRemoveDepartmentNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).RemoveDepartmentName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_RemoveDepartmentName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).RemoveDepartmentName(ctx, req.(*DeptRemoveDepartmentNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_ReplaceTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptReplaceTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).ReplaceTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_ReplaceTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).ReplaceTasks(ctx, req.(*DeptReplaceTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentInterface_SetMangedCores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptSetManagedCoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentInterfaceServer).SetMangedCores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentInterface_SetMangedCores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentInterfaceServer).SetMangedCores(ctx, req.(*DeptSetManagedCoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentInterface_ServiceDesc is the grpc.ServiceDesc for DepartmentInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "department.DepartmentInterface",
	HandlerType: (*DepartmentInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDepartmentName",
			Handler:    _DepartmentInterface_AddDepartmentName_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _DepartmentInterface_AddTask_Handler,
		},
		{
			MethodName: "AddTasks",
			Handler:    _DepartmentInterface_AddTasks_Handler,
		},
		{
			MethodName: "ClearTasks",
			Handler:    _DepartmentInterface_ClearTasks_Handler,
		},
		{
			MethodName: "ClearTaskAdjustments",
			Handler:    _DepartmentInterface_ClearTaskAdjustments_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DepartmentInterface_Delete_Handler,
		},
		{
			MethodName: "DisableTiManaged",
			Handler:    _DepartmentInterface_DisableTiManaged_Handler,
		},
		{
			MethodName: "EnableTiManaged",
			Handler:    _DepartmentInterface_EnableTiManaged_Handler,
		},
		{
			MethodName: "GetDepartmentNames",
			Handler:    _DepartmentInterface_GetDepartmentNames_Handler,
		},
		{
			MethodName: "GetTasks",
			Handler:    _DepartmentInterface_GetTasks_Handler,
		},
		{
			MethodName: "RemoveDepartmentName",
			Handler:    _DepartmentInterface_RemoveDepartmentName_Handler,
		},
		{
			MethodName: "ReplaceTasks",
			Handler:    _DepartmentInterface_ReplaceTasks_Handler,
		},
		{
			MethodName: "SetMangedCores",
			Handler:    _DepartmentInterface_SetMangedCores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/department.proto",
}
