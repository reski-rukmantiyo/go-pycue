// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: OpenCue/proto/cue.proto

package cue

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
	CueInterface_GetSystemStats_FullMethodName = "/cue.CueInterface/GetSystemStats"
)

// CueInterfaceClient is the client API for CueInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CueInterfaceClient interface {
	// returns the current server statistics
	GetSystemStats(ctx context.Context, in *CueGetSystemStatsRequest, opts ...grpc.CallOption) (*CueGetSystemStatsResponse, error)
}

type cueInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewCueInterfaceClient(cc grpc.ClientConnInterface) CueInterfaceClient {
	return &cueInterfaceClient{cc}
}

func (c *cueInterfaceClient) GetSystemStats(ctx context.Context, in *CueGetSystemStatsRequest, opts ...grpc.CallOption) (*CueGetSystemStatsResponse, error) {
	out := new(CueGetSystemStatsResponse)
	err := c.cc.Invoke(ctx, CueInterface_GetSystemStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CueInterfaceServer is the server API for CueInterface service.
// All implementations must embed UnimplementedCueInterfaceServer
// for forward compatibility
type CueInterfaceServer interface {
	// returns the current server statistics
	GetSystemStats(context.Context, *CueGetSystemStatsRequest) (*CueGetSystemStatsResponse, error)
	mustEmbedUnimplementedCueInterfaceServer()
}

// UnimplementedCueInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedCueInterfaceServer struct {
}

func (UnimplementedCueInterfaceServer) GetSystemStats(context.Context, *CueGetSystemStatsRequest) (*CueGetSystemStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemStats not implemented")
}
func (UnimplementedCueInterfaceServer) mustEmbedUnimplementedCueInterfaceServer() {}

// UnsafeCueInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CueInterfaceServer will
// result in compilation errors.
type UnsafeCueInterfaceServer interface {
	mustEmbedUnimplementedCueInterfaceServer()
}

func RegisterCueInterfaceServer(s grpc.ServiceRegistrar, srv CueInterfaceServer) {
	s.RegisterService(&CueInterface_ServiceDesc, srv)
}

func _CueInterface_GetSystemStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CueGetSystemStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CueInterfaceServer).GetSystemStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CueInterface_GetSystemStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CueInterfaceServer).GetSystemStats(ctx, req.(*CueGetSystemStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CueInterface_ServiceDesc is the grpc.ServiceDesc for CueInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CueInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cue.CueInterface",
	HandlerType: (*CueInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystemStats",
			Handler:    _CueInterface_GetSystemStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OpenCue/proto/cue.proto",
}
