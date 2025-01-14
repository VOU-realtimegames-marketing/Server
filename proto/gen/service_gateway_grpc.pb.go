// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.15.8
// source: gateway/service_gateway.proto

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
	Gateway_CreateUser_FullMethodName          = "/vou.proto.Gateway/CreateUser"
	Gateway_LoginUser_FullMethodName           = "/vou.proto.Gateway/LoginUser"
	Gateway_AuthorizeUser_FullMethodName       = "/vou.proto.Gateway/AuthorizeUser"
	Gateway_RenewAccessToken_FullMethodName    = "/vou.proto.Gateway/RenewAccessToken"
	Gateway_VerifyEmail_FullMethodName         = "/vou.proto.Gateway/VerifyEmail"
	Gateway_CreateStore_FullMethodName         = "/vou.proto.Gateway/CreateStore"
	Gateway_GetAllStoresOfOwner_FullMethodName = "/vou.proto.Gateway/GetAllStoresOfOwner"
	Gateway_UpdateStore_FullMethodName         = "/vou.proto.Gateway/UpdateStore"
	Gateway_DeleteStore_FullMethodName         = "/vou.proto.Gateway/DeleteStore"
	Gateway_CreateBranch_FullMethodName        = "/vou.proto.Gateway/CreateBranch"
	Gateway_GetBranchs_FullMethodName          = "/vou.proto.Gateway/GetBranchs"
	Gateway_DeleteBranch_FullMethodName        = "/vou.proto.Gateway/DeleteBranch"
	Gateway_CreateEvent_FullMethodName         = "/vou.proto.Gateway/CreateEvent"
	Gateway_GetAllEvents_FullMethodName        = "/vou.proto.Gateway/GetAllEvents"
	Gateway_GetAllEventsOfOwner_FullMethodName = "/vou.proto.Gateway/GetAllEventsOfOwner"
	Gateway_GetEventById_FullMethodName        = "/vou.proto.Gateway/GetEventById"
	Gateway_UpdateEventStatus_FullMethodName   = "/vou.proto.Gateway/UpdateEventStatus"
	Gateway_GetQuizzesByEventId_FullMethodName = "/vou.proto.Gateway/GetQuizzesByEventId"
	Gateway_GetMyVouchers_FullMethodName       = "/vou.proto.Gateway/GetMyVouchers"
	Gateway_GetAllOtherUsers_FullMethodName    = "/vou.proto.Gateway/GetAllOtherUsers"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	// BEGIN AUTH
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	AuthorizeUser(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	// BEGIN COUNTERPART
	CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error)
	GetAllStoresOfOwner(ctx context.Context, in *GetStoresOfOwnerRequest, opts ...grpc.CallOption) (*GetStoresOfOwnerResponse, error)
	UpdateStore(ctx context.Context, in *UpdateStoreRequest, opts ...grpc.CallOption) (*UpdateStoreResponse, error)
	DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*DeleteStoreResponse, error)
	CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error)
	GetBranchs(ctx context.Context, in *GetBranchsRequest, opts ...grpc.CallOption) (*GetBranchsResponse, error)
	DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*DeleteBranchResponse, error)
	// BEGIN EVENT-QUIZ
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	GetAllEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*GetAllEventsResponse, error)
	GetAllEventsOfOwner(ctx context.Context, in *GetEventsOfOwnerRequest, opts ...grpc.CallOption) (*GetEventsOfOwnerResponse, error)
	GetEventById(ctx context.Context, in *GetEventByIdRequest, opts ...grpc.CallOption) (*GetEventByIdResponse, error)
	UpdateEventStatus(ctx context.Context, in *UpdateEventStatusRequest, opts ...grpc.CallOption) (*UpdateEventStatusResponse, error)
	GetQuizzesByEventId(ctx context.Context, in *GetQuizzesByEventIdRequest, opts ...grpc.CallOption) (*GetQuizzesByEventIdResponse, error)
	GetMyVouchers(ctx context.Context, in *GetMyVouchersRequest, opts ...grpc.CallOption) (*GetMyVouchersResponse, error)
	// BEGIN USER
	GetAllOtherUsers(ctx context.Context, in *GetAllOtherUsersRequest, opts ...grpc.CallOption) (*GetAllOtherUsersResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Gateway_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AuthorizeUser(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, Gateway_AuthorizeUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RenewAccessTokenResponse)
	err := c.cc.Invoke(ctx, Gateway_RenewAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, Gateway_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStoreResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAllStoresOfOwner(ctx context.Context, in *GetStoresOfOwnerRequest, opts ...grpc.CallOption) (*GetStoresOfOwnerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStoresOfOwnerResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAllStoresOfOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateStore(ctx context.Context, in *UpdateStoreRequest, opts ...grpc.CallOption) (*UpdateStoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStoreResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*DeleteStoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteStoreResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBranchResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetBranchs(ctx context.Context, in *GetBranchsRequest, opts ...grpc.CallOption) (*GetBranchsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBranchsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetBranchs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*DeleteBranchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBranchResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAllEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*GetAllEventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllEventsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAllEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAllEventsOfOwner(ctx context.Context, in *GetEventsOfOwnerRequest, opts ...grpc.CallOption) (*GetEventsOfOwnerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventsOfOwnerResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAllEventsOfOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetEventById(ctx context.Context, in *GetEventByIdRequest, opts ...grpc.CallOption) (*GetEventByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventByIdResponse)
	err := c.cc.Invoke(ctx, Gateway_GetEventById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateEventStatus(ctx context.Context, in *UpdateEventStatusRequest, opts ...grpc.CallOption) (*UpdateEventStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEventStatusResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateEventStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetQuizzesByEventId(ctx context.Context, in *GetQuizzesByEventIdRequest, opts ...grpc.CallOption) (*GetQuizzesByEventIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQuizzesByEventIdResponse)
	err := c.cc.Invoke(ctx, Gateway_GetQuizzesByEventId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyVouchers(ctx context.Context, in *GetMyVouchersRequest, opts ...grpc.CallOption) (*GetMyVouchersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMyVouchersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMyVouchers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAllOtherUsers(ctx context.Context, in *GetAllOtherUsersRequest, opts ...grpc.CallOption) (*GetAllOtherUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllOtherUsersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAllOtherUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility.
type GatewayServer interface {
	// BEGIN AUTH
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	AuthorizeUser(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	// BEGIN COUNTERPART
	CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error)
	GetAllStoresOfOwner(context.Context, *GetStoresOfOwnerRequest) (*GetStoresOfOwnerResponse, error)
	UpdateStore(context.Context, *UpdateStoreRequest) (*UpdateStoreResponse, error)
	DeleteStore(context.Context, *DeleteStoreRequest) (*DeleteStoreResponse, error)
	CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error)
	GetBranchs(context.Context, *GetBranchsRequest) (*GetBranchsResponse, error)
	DeleteBranch(context.Context, *DeleteBranchRequest) (*DeleteBranchResponse, error)
	// BEGIN EVENT-QUIZ
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	GetAllEvents(context.Context, *GetAllEventsRequest) (*GetAllEventsResponse, error)
	GetAllEventsOfOwner(context.Context, *GetEventsOfOwnerRequest) (*GetEventsOfOwnerResponse, error)
	GetEventById(context.Context, *GetEventByIdRequest) (*GetEventByIdResponse, error)
	UpdateEventStatus(context.Context, *UpdateEventStatusRequest) (*UpdateEventStatusResponse, error)
	GetQuizzesByEventId(context.Context, *GetQuizzesByEventIdRequest) (*GetQuizzesByEventIdResponse, error)
	GetMyVouchers(context.Context, *GetMyVouchersRequest) (*GetMyVouchersResponse, error)
	// BEGIN USER
	GetAllOtherUsers(context.Context, *GetAllOtherUsersRequest) (*GetAllOtherUsersResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGatewayServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedGatewayServer) AuthorizeUser(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeUser not implemented")
}
func (UnimplementedGatewayServer) RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAccessToken not implemented")
}
func (UnimplementedGatewayServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedGatewayServer) CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedGatewayServer) GetAllStoresOfOwner(context.Context, *GetStoresOfOwnerRequest) (*GetStoresOfOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStoresOfOwner not implemented")
}
func (UnimplementedGatewayServer) UpdateStore(context.Context, *UpdateStoreRequest) (*UpdateStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStore not implemented")
}
func (UnimplementedGatewayServer) DeleteStore(context.Context, *DeleteStoreRequest) (*DeleteStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStore not implemented")
}
func (UnimplementedGatewayServer) CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranch not implemented")
}
func (UnimplementedGatewayServer) GetBranchs(context.Context, *GetBranchsRequest) (*GetBranchsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranchs not implemented")
}
func (UnimplementedGatewayServer) DeleteBranch(context.Context, *DeleteBranchRequest) (*DeleteBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBranch not implemented")
}
func (UnimplementedGatewayServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedGatewayServer) GetAllEvents(context.Context, *GetAllEventsRequest) (*GetAllEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvents not implemented")
}
func (UnimplementedGatewayServer) GetAllEventsOfOwner(context.Context, *GetEventsOfOwnerRequest) (*GetEventsOfOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEventsOfOwner not implemented")
}
func (UnimplementedGatewayServer) GetEventById(context.Context, *GetEventByIdRequest) (*GetEventByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventById not implemented")
}
func (UnimplementedGatewayServer) UpdateEventStatus(context.Context, *UpdateEventStatusRequest) (*UpdateEventStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEventStatus not implemented")
}
func (UnimplementedGatewayServer) GetQuizzesByEventId(context.Context, *GetQuizzesByEventIdRequest) (*GetQuizzesByEventIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuizzesByEventId not implemented")
}
func (UnimplementedGatewayServer) GetMyVouchers(context.Context, *GetMyVouchersRequest) (*GetMyVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyVouchers not implemented")
}
func (UnimplementedGatewayServer) GetAllOtherUsers(context.Context, *GetAllOtherUsersRequest) (*GetAllOtherUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOtherUsers not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}
func (UnimplementedGatewayServer) testEmbeddedByValue()                 {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	// If the following call pancis, it indicates UnimplementedGatewayServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AuthorizeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AuthorizeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AuthorizeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AuthorizeUser(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_RenewAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).RenewAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_RenewAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).RenewAccessToken(ctx, req.(*RenewAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateStore(ctx, req.(*CreateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAllStoresOfOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoresOfOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAllStoresOfOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAllStoresOfOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAllStoresOfOwner(ctx, req.(*GetStoresOfOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateStore(ctx, req.(*UpdateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteStore(ctx, req.(*DeleteStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateBranch(ctx, req.(*CreateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetBranchs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetBranchs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetBranchs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetBranchs(ctx, req.(*GetBranchsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteBranch(ctx, req.(*DeleteBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAllEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAllEvents(ctx, req.(*GetAllEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAllEventsOfOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsOfOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAllEventsOfOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAllEventsOfOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAllEventsOfOwner(ctx, req.(*GetEventsOfOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetEventById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetEventById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetEventById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetEventById(ctx, req.(*GetEventByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateEventStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateEventStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateEventStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateEventStatus(ctx, req.(*UpdateEventStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetQuizzesByEventId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuizzesByEventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetQuizzesByEventId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetQuizzesByEventId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetQuizzesByEventId(ctx, req.(*GetQuizzesByEventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyVouchersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMyVouchers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyVouchers(ctx, req.(*GetMyVouchersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAllOtherUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOtherUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAllOtherUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAllOtherUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAllOtherUsers(ctx, req.(*GetAllOtherUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vou.proto.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Gateway_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Gateway_LoginUser_Handler,
		},
		{
			MethodName: "AuthorizeUser",
			Handler:    _Gateway_AuthorizeUser_Handler,
		},
		{
			MethodName: "RenewAccessToken",
			Handler:    _Gateway_RenewAccessToken_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _Gateway_VerifyEmail_Handler,
		},
		{
			MethodName: "CreateStore",
			Handler:    _Gateway_CreateStore_Handler,
		},
		{
			MethodName: "GetAllStoresOfOwner",
			Handler:    _Gateway_GetAllStoresOfOwner_Handler,
		},
		{
			MethodName: "UpdateStore",
			Handler:    _Gateway_UpdateStore_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _Gateway_DeleteStore_Handler,
		},
		{
			MethodName: "CreateBranch",
			Handler:    _Gateway_CreateBranch_Handler,
		},
		{
			MethodName: "GetBranchs",
			Handler:    _Gateway_GetBranchs_Handler,
		},
		{
			MethodName: "DeleteBranch",
			Handler:    _Gateway_DeleteBranch_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Gateway_CreateEvent_Handler,
		},
		{
			MethodName: "GetAllEvents",
			Handler:    _Gateway_GetAllEvents_Handler,
		},
		{
			MethodName: "GetAllEventsOfOwner",
			Handler:    _Gateway_GetAllEventsOfOwner_Handler,
		},
		{
			MethodName: "GetEventById",
			Handler:    _Gateway_GetEventById_Handler,
		},
		{
			MethodName: "UpdateEventStatus",
			Handler:    _Gateway_UpdateEventStatus_Handler,
		},
		{
			MethodName: "GetQuizzesByEventId",
			Handler:    _Gateway_GetQuizzesByEventId_Handler,
		},
		{
			MethodName: "GetMyVouchers",
			Handler:    _Gateway_GetMyVouchers_Handler,
		},
		{
			MethodName: "GetAllOtherUsers",
			Handler:    _Gateway_GetAllOtherUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/service_gateway.proto",
}
