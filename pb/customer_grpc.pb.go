// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: customer.proto

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

// MyCustomerClient is the client API for MyCustomer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyCustomerClient interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	GetCustomerDetails(ctx context.Context, in *GetCustomerDetailsRequest, opts ...grpc.CallOption) (*Customer, error)
	ChangeCustomerPassword(ctx context.Context, in *ChangeCustomerPasswordRequest, opts ...grpc.CallOption) (*Empty, error)
}

type myCustomerClient struct {
	cc grpc.ClientConnInterface
}

func NewMyCustomerClient(cc grpc.ClientConnInterface) MyCustomerClient {
	return &myCustomerClient{cc}
}

func (c *myCustomerClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/proto.MyCustomer/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myCustomerClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/proto.MyCustomer/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myCustomerClient) GetCustomerDetails(ctx context.Context, in *GetCustomerDetailsRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/proto.MyCustomer/GetCustomerDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myCustomerClient) ChangeCustomerPassword(ctx context.Context, in *ChangeCustomerPasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.MyCustomer/ChangeCustomerPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyCustomerServer is the server API for MyCustomer service.
// All implementations must embed UnimplementedMyCustomerServer
// for forward compatibility
type MyCustomerServer interface {
	CreateCustomer(context.Context, *CreateCustomerRequest) (*Customer, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*Customer, error)
	GetCustomerDetails(context.Context, *GetCustomerDetailsRequest) (*Customer, error)
	ChangeCustomerPassword(context.Context, *ChangeCustomerPasswordRequest) (*Empty, error)
	mustEmbedUnimplementedMyCustomerServer()
}

// UnimplementedMyCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedMyCustomerServer struct {
}

func (UnimplementedMyCustomerServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedMyCustomerServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedMyCustomerServer) GetCustomerDetails(context.Context, *GetCustomerDetailsRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerDetails not implemented")
}
func (UnimplementedMyCustomerServer) ChangeCustomerPassword(context.Context, *ChangeCustomerPasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCustomerPassword not implemented")
}
func (UnimplementedMyCustomerServer) mustEmbedUnimplementedMyCustomerServer() {}

// UnsafeMyCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyCustomerServer will
// result in compilation errors.
type UnsafeMyCustomerServer interface {
	mustEmbedUnimplementedMyCustomerServer()
}

func RegisterMyCustomerServer(s grpc.ServiceRegistrar, srv MyCustomerServer) {
	s.RegisterService(&MyCustomer_ServiceDesc, srv)
}

func _MyCustomer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyCustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyCustomer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyCustomerServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyCustomer_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyCustomerServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyCustomer/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyCustomerServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyCustomer_GetCustomerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyCustomerServer).GetCustomerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyCustomer/GetCustomerDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyCustomerServer).GetCustomerDetails(ctx, req.(*GetCustomerDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyCustomer_ChangeCustomerPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCustomerPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyCustomerServer).ChangeCustomerPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyCustomer/ChangeCustomerPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyCustomerServer).ChangeCustomerPassword(ctx, req.(*ChangeCustomerPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyCustomer_ServiceDesc is the grpc.ServiceDesc for MyCustomer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyCustomer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyCustomer",
	HandlerType: (*MyCustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _MyCustomer_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _MyCustomer_UpdateCustomer_Handler,
		},
		{
			MethodName: "GetCustomerDetails",
			Handler:    _MyCustomer_GetCustomerDetails_Handler,
		},
		{
			MethodName: "ChangeCustomerPassword",
			Handler:    _MyCustomer_ChangeCustomerPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
