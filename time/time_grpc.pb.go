// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: time/time.proto

package time

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

// GetCurrentTimeClient is the client API for GetCurrentTime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetCurrentTimeClient interface {
	GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (*GetTimeReply, error)
}

type getCurrentTimeClient struct {
	cc grpc.ClientConnInterface
}

func NewGetCurrentTimeClient(cc grpc.ClientConnInterface) GetCurrentTimeClient {
	return &getCurrentTimeClient{cc}
}

func (c *getCurrentTimeClient) GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (*GetTimeReply, error) {
	out := new(GetTimeReply)
	err := c.cc.Invoke(ctx, "/time.getCurrentTime/getTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetCurrentTimeServer is the server API for GetCurrentTime service.
// All implementations must embed UnimplementedGetCurrentTimeServer
// for forward compatibility
type GetCurrentTimeServer interface {
	GetTime(context.Context, *GetTimeRequest) (*GetTimeReply, error)
	mustEmbedUnimplementedGetCurrentTimeServer()
}

// UnimplementedGetCurrentTimeServer must be embedded to have forward compatible implementations.
type UnimplementedGetCurrentTimeServer struct {
}

func (UnimplementedGetCurrentTimeServer) GetTime(context.Context, *GetTimeRequest) (*GetTimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}
func (UnimplementedGetCurrentTimeServer) mustEmbedUnimplementedGetCurrentTimeServer() {}

// UnsafeGetCurrentTimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetCurrentTimeServer will
// result in compilation errors.
type UnsafeGetCurrentTimeServer interface {
	mustEmbedUnimplementedGetCurrentTimeServer()
}

func RegisterGetCurrentTimeServer(s grpc.ServiceRegistrar, srv GetCurrentTimeServer) {
	s.RegisterService(&GetCurrentTime_ServiceDesc, srv)
}

func _GetCurrentTime_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetCurrentTimeServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.getCurrentTime/getTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetCurrentTimeServer).GetTime(ctx, req.(*GetTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetCurrentTime_ServiceDesc is the grpc.ServiceDesc for GetCurrentTime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetCurrentTime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "time.getCurrentTime",
	HandlerType: (*GetCurrentTimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getTime",
			Handler:    _GetCurrentTime_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time/time.proto",
}
