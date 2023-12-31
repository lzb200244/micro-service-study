// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package orders

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

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersServiceClient interface {
	Create(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	CreateRevert(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) Create(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) CreateRevert(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/CreateRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations must embed UnimplementedOrdersServiceServer
// for forward compatibility
type OrdersServiceServer interface {
	Create(context.Context, *OrdersRequest) (*OrdersResponse, error)
	CreateRevert(context.Context, *OrdersRequest) (*OrdersResponse, error)
	mustEmbedUnimplementedOrdersServiceServer()
}

// UnimplementedOrdersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (UnimplementedOrdersServiceServer) Create(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrdersServiceServer) CreateRevert(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRevert not implemented")
}
func (UnimplementedOrdersServiceServer) mustEmbedUnimplementedOrdersServiceServer() {}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Create(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_CreateRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).CreateRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/CreateRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).CreateRevert(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrdersService_Create_Handler,
		},
		{
			MethodName: "CreateRevert",
			Handler:    _OrdersService_CreateRevert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders.proto",
}
