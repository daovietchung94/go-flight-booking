// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: booking.proto

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

// MyBookingClient is the client API for MyBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyBookingClient interface {
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*Reservation, error)
	GetReservationDetails(ctx context.Context, in *GetReservationDetailsRequest, opts ...grpc.CallOption) (*Reservation, error)
}

type myBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewMyBookingClient(cc grpc.ClientConnInterface) MyBookingClient {
	return &myBookingClient{cc}
}

func (c *myBookingClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*Reservation, error) {
	out := new(Reservation)
	err := c.cc.Invoke(ctx, "/proto.MyBooking/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myBookingClient) GetReservationDetails(ctx context.Context, in *GetReservationDetailsRequest, opts ...grpc.CallOption) (*Reservation, error) {
	out := new(Reservation)
	err := c.cc.Invoke(ctx, "/proto.MyBooking/GetReservationDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyBookingServer is the server API for MyBooking service.
// All implementations must embed UnimplementedMyBookingServer
// for forward compatibility
type MyBookingServer interface {
	CreateReservation(context.Context, *CreateReservationRequest) (*Reservation, error)
	GetReservationDetails(context.Context, *GetReservationDetailsRequest) (*Reservation, error)
	mustEmbedUnimplementedMyBookingServer()
}

// UnimplementedMyBookingServer must be embedded to have forward compatible implementations.
type UnimplementedMyBookingServer struct {
}

func (UnimplementedMyBookingServer) CreateReservation(context.Context, *CreateReservationRequest) (*Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedMyBookingServer) GetReservationDetails(context.Context, *GetReservationDetailsRequest) (*Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationDetails not implemented")
}
func (UnimplementedMyBookingServer) mustEmbedUnimplementedMyBookingServer() {}

// UnsafeMyBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyBookingServer will
// result in compilation errors.
type UnsafeMyBookingServer interface {
	mustEmbedUnimplementedMyBookingServer()
}

func RegisterMyBookingServer(s grpc.ServiceRegistrar, srv MyBookingServer) {
	s.RegisterService(&MyBooking_ServiceDesc, srv)
}

func _MyBooking_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyBookingServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyBooking/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyBookingServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyBooking_GetReservationDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyBookingServer).GetReservationDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MyBooking/GetReservationDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyBookingServer).GetReservationDetails(ctx, req.(*GetReservationDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyBooking_ServiceDesc is the grpc.ServiceDesc for MyBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyBooking",
	HandlerType: (*MyBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReservation",
			Handler:    _MyBooking_CreateReservation_Handler,
		},
		{
			MethodName: "GetReservationDetails",
			Handler:    _MyBooking_GetReservationDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}