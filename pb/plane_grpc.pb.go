// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: plane.proto

package pb

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

// MyPlaneClient is the client API for MyPlane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyPlaneClient interface {
	GetPlanes(ctx context.Context, in *GetPlanesRequest, opts ...grpc.CallOption) (*GetPlanesResponse, error)
	GetPlaneDetails(ctx context.Context, in *GetPlaneDetailsRequest, opts ...grpc.CallOption) (*Plane, error)
	GetPlaneStatus(ctx context.Context, in *GetPlaneStatusRequest, opts ...grpc.CallOption) (*Plane, error)
	GetPlaneByNumber(ctx context.Context, in *GetPlaneByNumberRequest, opts ...grpc.CallOption) (*Plane, error)
	CreatePlane(ctx context.Context, in *CreatePlaneRequest, opts ...grpc.CallOption) (*Plane, error)
	UpdatePlane(ctx context.Context, in *UpdatePlaneRequest, opts ...grpc.CallOption) (*Plane, error)
	UpdatePlaneStatus(ctx context.Context, in *UpdatePlaneStatusRequest, opts ...grpc.CallOption) (*Plane, error)
	DeletePlane(ctx context.Context, in *DeletePlaneRequest, opts ...grpc.CallOption) (*DeletePlaneResponse, error)
}

type myPlaneClient struct {
	cc grpc.ClientConnInterface
}

func NewMyPlaneClient(cc grpc.ClientConnInterface) MyPlaneClient {
	return &myPlaneClient{cc}
}

func (c *myPlaneClient) GetPlanes(ctx context.Context, in *GetPlanesRequest, opts ...grpc.CallOption) (*GetPlanesResponse, error) {
	out := new(GetPlanesResponse)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/GetPlanes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) GetPlaneDetails(ctx context.Context, in *GetPlaneDetailsRequest, opts ...grpc.CallOption) (*Plane, error) {
	out := new(Plane)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/GetPlaneDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) GetPlaneStatus(ctx context.Context, in *GetPlaneStatusRequest, opts ...grpc.CallOption) (*Plane, error) {
	out := new(Plane)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/GetPlaneStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) GetPlaneByNumber(ctx context.Context, in *GetPlaneByNumberRequest, opts ...grpc.CallOption) (*Plane, error) {
	out := new(Plane)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/GetPlaneByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) CreatePlane(ctx context.Context, in *CreatePlaneRequest, opts ...grpc.CallOption) (*Plane, error) {
	out := new(Plane)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/CreatePlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) UpdatePlane(ctx context.Context, in *UpdatePlaneRequest, opts ...grpc.CallOption) (*Plane, error) {
	out := new(Plane)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/UpdatePlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) UpdatePlaneStatus(ctx context.Context, in *UpdatePlaneStatusRequest, opts ...grpc.CallOption) (*Plane, error) {
	out := new(Plane)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/UpdatePlaneStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPlaneClient) DeletePlane(ctx context.Context, in *DeletePlaneRequest, opts ...grpc.CallOption) (*DeletePlaneResponse, error) {
	out := new(DeletePlaneResponse)
	err := c.cc.Invoke(ctx, "/proto.MyPlane/DeletePlane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyPlaneServer is the server API for MyPlane service.
// All implementations must embed UnimplementedMyPlaneServer
// for forward compatibility
type MyPlaneServer interface {
	GetPlanes(context.Context, *GetPlanesRequest) (*GetPlanesResponse, error)
	GetPlaneDetails(context.Context, *GetPlaneDetailsRequest) (*Plane, error)
	GetPlaneStatus(context.Context, *GetPlaneStatusRequest) (*Plane, error)
	GetPlaneByNumber(context.Context, *GetPlaneByNumberRequest) (*Plane, error)
	CreatePlane(context.Context, *CreatePlaneRequest) (*Plane, error)
	UpdatePlane(context.Context, *UpdatePlaneRequest) (*Plane, error)
	UpdatePlaneStatus(context.Context, *UpdatePlaneStatusRequest) (*Plane, error)
	DeletePlane(context.Context, *DeletePlaneRequest) (*DeletePlaneResponse, error)
	mustEmbedUnimplementedMyPlaneServer()
}

// UnimplementedMyPlaneServer must be embedded to have forward compatible implementations.
type UnimplementedMyPlaneServer struct {
}

func (UnimplementedMyPlaneServer) GetPlanes(context.Context, *GetPlanesRequest) (*GetPlanesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanes not implemented")
}
func (UnimplementedMyPlaneServer) GetPlaneDetails(context.Context, *GetPlaneDetailsRequest) (*Plane, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaneDetails not implemented")
}
func (UnimplementedMyPlaneServer) GetPlaneStatus(context.Context, *GetPlaneStatusRequest) (*Plane, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaneStatus not implemented")
}
func (UnimplementedMyPlaneServer) GetPlaneByNumber(context.Context, *GetPlaneByNumberRequest) (*Plane, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaneByNumber not implemented")
}
func (UnimplementedMyPlaneServer) CreatePlane(context.Context, *CreatePlaneRequest) (*Plane, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlane not implemented")
}
func (UnimplementedMyPlaneServer) UpdatePlane(context.Context, *UpdatePlaneRequest) (*Plane, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlane not implemented")
}
func (UnimplementedMyPlaneServer) UpdatePlaneStatus(context.Context, *UpdatePlaneStatusRequest) (*Plane, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaneStatus not implemented")
}
func (UnimplementedMyPlaneServer) DeletePlane(context.Context, *DeletePlaneRequest) (*DeletePlaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlane not implemented")
}
func (UnimplementedMyPlaneServer) mustEmbedUnimplementedMyPlaneServer() {}

// UnsafeMyPlaneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyPlaneServer will
// result in compilation errors.
type UnsafeMyPlaneServer interface {
	mustEmbedUnimplementedMyPlaneServer()
}

func RegisterMyPlaneServer(s grpc.ServiceRegistrar, srv MyPlaneServer) {
	s.RegisterService(&MyPlane_ServiceDesc, srv)
}

func _MyPlane_GetPlanes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).GetPlanes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/GetPlanes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).GetPlanes(ctx, req.(*GetPlanesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_GetPlaneDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaneDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).GetPlaneDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/GetPlaneDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).GetPlaneDetails(ctx, req.(*GetPlaneDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_GetPlaneStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaneStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).GetPlaneStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/GetPlaneStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).GetPlaneStatus(ctx, req.(*GetPlaneStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_GetPlaneByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaneByNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).GetPlaneByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/GetPlaneByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).GetPlaneByNumber(ctx, req.(*GetPlaneByNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_CreatePlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).CreatePlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/CreatePlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).CreatePlane(ctx, req.(*CreatePlaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_UpdatePlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).UpdatePlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/UpdatePlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).UpdatePlane(ctx, req.(*UpdatePlaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_UpdatePlaneStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlaneStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).UpdatePlaneStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/UpdatePlaneStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).UpdatePlaneStatus(ctx, req.(*UpdatePlaneStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPlane_DeletePlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPlaneServer).DeletePlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyPlane/DeletePlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPlaneServer).DeletePlane(ctx, req.(*DeletePlaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyPlane_ServiceDesc is the grpc.ServiceDesc for MyPlane service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyPlane_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyPlane",
	HandlerType: (*MyPlaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlanes",
			Handler:    _MyPlane_GetPlanes_Handler,
		},
		{
			MethodName: "GetPlaneDetails",
			Handler:    _MyPlane_GetPlaneDetails_Handler,
		},
		{
			MethodName: "GetPlaneStatus",
			Handler:    _MyPlane_GetPlaneStatus_Handler,
		},
		{
			MethodName: "GetPlaneByNumber",
			Handler:    _MyPlane_GetPlaneByNumber_Handler,
		},
		{
			MethodName: "CreatePlane",
			Handler:    _MyPlane_CreatePlane_Handler,
		},
		{
			MethodName: "UpdatePlane",
			Handler:    _MyPlane_UpdatePlane_Handler,
		},
		{
			MethodName: "UpdatePlaneStatus",
			Handler:    _MyPlane_UpdatePlaneStatus_Handler,
		},
		{
			MethodName: "DeletePlane",
			Handler:    _MyPlane_DeletePlane_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plane.proto",
}
