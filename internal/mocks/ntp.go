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

type MockNtp struct {
	mock.Mock
}

func (m *MockNtp) SetNtpServer(ctx context.Context, in *v1.Ntp, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(in)
	return &emptypb.Empty{}, args.Error(1)

}

func (m *MockNtp) GetNtpServer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Ntp, error) {
	args := m.Called(in)
	return &v1.Ntp{}, args.Error(1)

}

func (m *MockNtp) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Status, error) {
	args := m.Called(in)
	return &v1.Status{}, args.Error(1)

}
