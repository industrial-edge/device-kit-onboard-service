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
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockNetwork struct {
	mock.Mock
}

func (m *MockNetwork) GetAllInterfaces(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.NetworkSettings, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.NetworkSettings), args.Error(1)
}

func (m *MockNetwork) GetInterfaceWithMac(ctx context.Context, in *v1.NetworkInterfaceRequest, opts ...grpc.CallOption) (*v1.Interface, error) {
	panic("implement me")
}

func (m *MockNetwork) GetInterfaceWithLabel(ctx context.Context, in *v1.NetworkInterfaceRequestWithLabel, opts ...grpc.CallOption) (*v1.Interface, error) {
	panic("implement me")
}

func (m *MockNetwork) ApplySettings(ctx context.Context, in *v1.NetworkSettings, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m *MockNetwork) CreateConnection(ctx context.Context, in *v1.ConnectionSettings, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m *MockNetwork) RemoveConnection(ctx context.Context, in *v1.ConnectionSettings, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return &emptypb.Empty{}, args.Error(1)
}
