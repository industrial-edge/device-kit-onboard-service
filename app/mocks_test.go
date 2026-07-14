/*
 * Copyright © Siemens 2020 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"log"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	"onboardservice/internal/system"
	"sync"
	"time"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

type MockClientFactory struct {
}

func (m *MockClientFactory) CreateHelper() system.Configurable {
	return &system.MockHelper{}
}

type MockListener struct {
	receivedStatus *v1.OnboardingStatus
	context        context.Context
	doneChannel    chan struct{}
	mu             sync.Mutex
	sendCalled     chan struct{}
}

func (m *MockListener) Send(onboardingStatus *v1.OnboardingStatus) error {
	log.Println("MOCK RECEVIED STATUS :", onboardingStatus.IsOnboarded)
	m.mu.Lock()
	m.receivedStatus = onboardingStatus
	m.mu.Unlock()
	if m.sendCalled != nil {
		m.sendCalled <- struct{}{}
	}
	return nil
}

func (m *MockListener) getReceivedStatus() *v1.OnboardingStatus {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.receivedStatus
}

func (m *MockListener) SetHeader(md metadata.MD) error {
	return nil
}

func (m *MockListener) SendHeader(md metadata.MD) error {
	return nil
}

func (m *MockListener) SetTrailer(md metadata.MD) {

}

type mockContext struct {
	doneChannel chan struct{}
}

func (m *mockContext) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (m *mockContext) Done() <-chan struct{} {
	return m.doneChannel
}

func (m *mockContext) Err() error {
	panic("implement me")
}

func (m *mockContext) Value(key interface{}) interface{} {
	panic("implement me")
}

func (m *MockListener) Context() context.Context {
	return m.context
}

func (m *MockListener) SendMsg(a interface{}) error {
	return nil
}

func (m *MockListener) RecvMsg(a interface{}) error {
	return nil
}

type mockedgeCoreRuntime struct {
	mock.Mock
	isGood bool
}

func (m *mockedgeCoreRuntime) Activate(configuration *v1.DeviceConfiguration) (bool, error) {
	args := m.Called(configuration)
	return args.Bool(0), args.Error(1)
}

func (m *mockedgeCoreRuntime) Onboarded() (bool, bool) {
	args := m.Called()
	return args.Bool(0), args.Bool(1)
}
