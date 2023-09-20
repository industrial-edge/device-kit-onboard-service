// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: Onboard.proto

package siemens_iedge_dmapi_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OnboardService_OnboardWithUSB_FullMethodName      = "/siemens.iedge.dmapi.onboard.v1.OnboardService/OnboardWithUSB"
	OnboardService_ApplyConfiguration_FullMethodName  = "/siemens.iedge.dmapi.onboard.v1.OnboardService/ApplyConfiguration"
	OnboardService_SetOnboardingStatus_FullMethodName = "/siemens.iedge.dmapi.onboard.v1.OnboardService/SetOnboardingStatus"
	OnboardService_ListenOnboardState_FullMethodName  = "/siemens.iedge.dmapi.onboard.v1.OnboardService/ListenOnboardState"
)

// OnboardServiceClient is the client API for OnboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnboardServiceClient interface {
	// Starts onboarding sequence,applies  all settings via ApplyConfiguration() and finally calls  EdgeCoreRuntime REST API to onboard the device.
	OnboardWithUSB(ctx context.Context, in *DeviceConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Applies all settings. Does not call EdgeCoreRuntime REST API.Applicable for UI Onboarding.
	ApplyConfiguration(ctx context.Context, in *DeviceConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Sets state for device onboard status.In Onboarding from UI, EdgeCoreRuntime calls this method to set the onboarding status.
	SetOnboardingStatus(ctx context.Context, in *OnboardingStatus, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Server stream for onboarding status changes.
	ListenOnboardState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OnboardService_ListenOnboardStateClient, error)
}

type onboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnboardServiceClient(cc grpc.ClientConnInterface) OnboardServiceClient {
	return &onboardServiceClient{cc}
}

func (c *onboardServiceClient) OnboardWithUSB(ctx context.Context, in *DeviceConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OnboardService_OnboardWithUSB_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardServiceClient) ApplyConfiguration(ctx context.Context, in *DeviceConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OnboardService_ApplyConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardServiceClient) SetOnboardingStatus(ctx context.Context, in *OnboardingStatus, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OnboardService_SetOnboardingStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardServiceClient) ListenOnboardState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OnboardService_ListenOnboardStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &OnboardService_ServiceDesc.Streams[0], OnboardService_ListenOnboardState_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &onboardServiceListenOnboardStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OnboardService_ListenOnboardStateClient interface {
	Recv() (*OnboardingStatus, error)
	grpc.ClientStream
}

type onboardServiceListenOnboardStateClient struct {
	grpc.ClientStream
}

func (x *onboardServiceListenOnboardStateClient) Recv() (*OnboardingStatus, error) {
	m := new(OnboardingStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OnboardServiceServer is the server API for OnboardService service.
// All implementations must embed UnimplementedOnboardServiceServer
// for forward compatibility
type OnboardServiceServer interface {
	// Starts onboarding sequence,applies  all settings via ApplyConfiguration() and finally calls  EdgeCoreRuntime REST API to onboard the device.
	OnboardWithUSB(context.Context, *DeviceConfiguration) (*emptypb.Empty, error)
	// Applies all settings. Does not call EdgeCoreRuntime REST API.Applicable for UI Onboarding.
	ApplyConfiguration(context.Context, *DeviceConfiguration) (*emptypb.Empty, error)
	// Sets state for device onboard status.In Onboarding from UI, EdgeCoreRuntime calls this method to set the onboarding status.
	SetOnboardingStatus(context.Context, *OnboardingStatus) (*emptypb.Empty, error)
	// Server stream for onboarding status changes.
	ListenOnboardState(*emptypb.Empty, OnboardService_ListenOnboardStateServer) error
	mustEmbedUnimplementedOnboardServiceServer()
}

// UnimplementedOnboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOnboardServiceServer struct {
}

func (UnimplementedOnboardServiceServer) OnboardWithUSB(context.Context, *DeviceConfiguration) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnboardWithUSB not implemented")
}
func (UnimplementedOnboardServiceServer) ApplyConfiguration(context.Context, *DeviceConfiguration) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyConfiguration not implemented")
}
func (UnimplementedOnboardServiceServer) SetOnboardingStatus(context.Context, *OnboardingStatus) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOnboardingStatus not implemented")
}
func (UnimplementedOnboardServiceServer) ListenOnboardState(*emptypb.Empty, OnboardService_ListenOnboardStateServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenOnboardState not implemented")
}
func (UnimplementedOnboardServiceServer) mustEmbedUnimplementedOnboardServiceServer() {}

// UnsafeOnboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnboardServiceServer will
// result in compilation errors.
type UnsafeOnboardServiceServer interface {
	mustEmbedUnimplementedOnboardServiceServer()
}

func RegisterOnboardServiceServer(s grpc.ServiceRegistrar, srv OnboardServiceServer) {
	s.RegisterService(&OnboardService_ServiceDesc, srv)
}

func _OnboardService_OnboardWithUSB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardServiceServer).OnboardWithUSB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnboardService_OnboardWithUSB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardServiceServer).OnboardWithUSB(ctx, req.(*DeviceConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardService_ApplyConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardServiceServer).ApplyConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnboardService_ApplyConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardServiceServer).ApplyConfiguration(ctx, req.(*DeviceConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardService_SetOnboardingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnboardingStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardServiceServer).SetOnboardingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnboardService_SetOnboardingStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardServiceServer).SetOnboardingStatus(ctx, req.(*OnboardingStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardService_ListenOnboardState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OnboardServiceServer).ListenOnboardState(m, &onboardServiceListenOnboardStateServer{stream})
}

type OnboardService_ListenOnboardStateServer interface {
	Send(*OnboardingStatus) error
	grpc.ServerStream
}

type onboardServiceListenOnboardStateServer struct {
	grpc.ServerStream
}

func (x *onboardServiceListenOnboardStateServer) Send(m *OnboardingStatus) error {
	return x.ServerStream.SendMsg(m)
}

// OnboardService_ServiceDesc is the grpc.ServiceDesc for OnboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "siemens.iedge.dmapi.onboard.v1.OnboardService",
	HandlerType: (*OnboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnboardWithUSB",
			Handler:    _OnboardService_OnboardWithUSB_Handler,
		},
		{
			MethodName: "ApplyConfiguration",
			Handler:    _OnboardService_ApplyConfiguration_Handler,
		},
		{
			MethodName: "SetOnboardingStatus",
			Handler:    _OnboardService_SetOnboardingStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenOnboardState",
			Handler:       _OnboardService_ListenOnboardState_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Onboard.proto",
}
