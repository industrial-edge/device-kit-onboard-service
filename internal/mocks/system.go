/*
 * Copyright © Siemens 2025 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package mocks

import (
	"context"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockSystem struct {
	mock.Mock
}

func (m MockSystem) HardReset(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (m MockSystem) GetCustomSettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*anypb.Any, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*anypb.Any), args.Error(1)
}

func (m MockSystem) ApplyCustomSettings(ctx context.Context, in *anypb.Any, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m MockSystem) GetLogFile(ctx context.Context, in *v1.LogRequest, opts ...grpc.CallOption) (*v1.LogResponse, error) {
	panic("implement me")
}

func (m MockSystem) RestartDevice(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m MockSystem) ShutdownDevice(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m MockSystem) FactoryReset(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m MockSystem) GetModelNumber(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.ModelNumber, error) {
	args := m.Called(in)
	return &v1.ModelNumber{
		ModelNumber: "ipc227e",
	}, args.Error(1)
}

func (m MockSystem) GetFirmwareInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.FirmwareInfo, error) {
	args := m.Called(in)
	return &v1.FirmwareInfo{Version: "v1.0.0"}, args.Error(1)
}

func (m MockSystem) GetResourceStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Stats, error) {
	args := m.Called(in)
	return &v1.Stats{}, args.Error(1)
}

func (m MockSystem) GetLimits(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Limits, error) {
	args := m.Called(in)
	return &v1.Limits{}, args.Error(1)
}

func (m MockSystem) UpdateHostname(ctx context.Context, in *v1.Hostname, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m MockSystem) GetHostname(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Hostname, error) {
	args := m.Called(in)
	return &v1.Hostname{}, args.Error(1)
}

func (m MockSystem) GetDeviceInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.DeviceInfo, error) {
	args := m.Called(in)
	return &v1.DeviceInfo{}, args.Error(1)
}
