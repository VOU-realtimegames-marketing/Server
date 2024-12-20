// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.15.8
// source: counterpart/service_counterpart.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CounterpartService_CreateStore_FullMethodName = "/vou.proto.CounterpartService/CreateStore"
)

// CounterpartServiceClient is the client API for CounterpartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterpartServiceClient interface {
	CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error)
}

type counterpartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterpartServiceClient(cc grpc.ClientConnInterface) CounterpartServiceClient {
	return &counterpartServiceClient{cc}
}

func (c *counterpartServiceClient) CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStoreResponse)
	err := c.cc.Invoke(ctx, CounterpartService_CreateStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterpartServiceServer is the server API for CounterpartService service.
// All implementations must embed UnimplementedCounterpartServiceServer
// for forward compatibility.
type CounterpartServiceServer interface {
	CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error)
	mustEmbedUnimplementedCounterpartServiceServer()
}

// UnimplementedCounterpartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCounterpartServiceServer struct{}

func (UnimplementedCounterpartServiceServer) CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedCounterpartServiceServer) mustEmbedUnimplementedCounterpartServiceServer() {}
func (UnimplementedCounterpartServiceServer) testEmbeddedByValue()                            {}

// UnsafeCounterpartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterpartServiceServer will
// result in compilation errors.
type UnsafeCounterpartServiceServer interface {
	mustEmbedUnimplementedCounterpartServiceServer()
}

func RegisterCounterpartServiceServer(s grpc.ServiceRegistrar, srv CounterpartServiceServer) {
	// If the following call pancis, it indicates UnimplementedCounterpartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CounterpartService_ServiceDesc, srv)
}

func _CounterpartService_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterpartServiceServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CounterpartService_CreateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterpartServiceServer).CreateStore(ctx, req.(*CreateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CounterpartService_ServiceDesc is the grpc.ServiceDesc for CounterpartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterpartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vou.proto.CounterpartService",
	HandlerType: (*CounterpartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStore",
			Handler:    _CounterpartService_CreateStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counterpart/service_counterpart.proto",
}
