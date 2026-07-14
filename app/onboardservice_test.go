/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"errors"
	"log"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	"os"
	"sync"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	mockClients *ClientPack
	ml          mockLed
	mn          mockNtp
	mpi         mockedgeCoreRuntime
	mNet        mockNetwork
	mSys        mockSystem
	app         *MainApp
	mu          sync.Mutex
}

func (suite *TestSuite) SetupSuite() {
	os.Remove("/tmp/temp.sock")
	suite.app = CreateServiceApp(&MockClientFactory{})
	suite.app.Clients = suite.mockClients
	suite.app.serverInstance.isContainerAccessible = true
}

func (suite *TestSuite) SetupTest() {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	initMocks(suite)
}

func initMocks(suite *TestSuite) {
	suite.ml = mockLed{Mock: mock.Mock{}}
	suite.mn = mockNtp{Mock: mock.Mock{}}
	suite.mpi = mockedgeCoreRuntime{Mock: mock.Mock{}}
	suite.mNet = mockNetwork{Mock: mock.Mock{}}
	suite.mSys = mockSystem{Mock: mock.Mock{}}

	suite.mockClients = &ClientPack{}
	suite.mockClients.RestClient = &suite.mpi
	suite.mockClients.NtpClient = &suite.mn
	suite.mockClients.SystemClient = &suite.mSys
	suite.mockClients.NetworkClient = &suite.mNet

	suite.app.Clients = suite.mockClients
}

func TestSuiteAll(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TearDownSuite() {
	suite.app.StopGRPC()
	log.Println("Test suite end")
}

func (suite *TestSuite) Test_ApplyConfiguration_CallOrder() {
	// Create a variable to track call order
	callOrder := []string{}
	orderMutex := sync.Mutex{}

	// Mock Network with custom function to track call order
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			orderMutex.Lock()
			callOrder = append(callOrder, "network")
			orderMutex.Unlock()
		}).
		Return(&emptypb.Empty{}, nil).Once()

	// Mock NTP with custom function to track call order
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			orderMutex.Lock()
			callOrder = append(callOrder, "ntp")
			orderMutex.Unlock()
		}).
		Return(&emptypb.Empty{}, nil).Once()

	// Execute the method
	_, err := suite.app.serverInstance.ApplyConfiguration(context.Background(), deviceConfig)
	suite.Nil(err)

	// Assert call order
	suite.Require().Len(callOrder, 2)
	suite.Equal("network", callOrder[0], "Network should be configured first")
	suite.Equal("ntp", callOrder[1], "NTP should be configured second")
}

func (suite *TestSuite) Test_ApplyAllConfigurations() {
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	_, err := suite.app.serverInstance.ApplyConfiguration(context.Background(), deviceConfig)
	suite.Nil(err, "apply test failed", err)
}

func (suite *TestSuite) Test_ApplyAllConfigurationsWithNetworkFail() {
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("virtual error")).Once()
	_, err := suite.app.serverInstance.ApplyConfiguration(context.Background(), deviceConfig)
	suite.NotNil(err, "apply test failed", err)
}

func (suite *TestSuite) Test_ApplyAllConfigurationsWithNtpFail() {
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).
		Return(&emptypb.Empty{}, nil).Once()

	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{},
		errors.New("virtual error")).Once()

	_, err := suite.app.serverInstance.ApplyConfiguration(context.Background(), deviceConfig)
	suite.NotNil(err, "apply test failed", err)
}

func (suite *TestSuite) TestOnboardServer_OnboardWithUSB() {
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	suite.mpi.On("Activate", mock.Anything).Return(true, nil).Once()
	suite.mpi.On("Onboarded").Return(false, true).Once()
	_, err := suite.app.serverInstance.OnboardWithUSB(context.Background(), deviceConfig)
	suite.Nil(err, "apply test failed", err)
}

func (suite *TestSuite) TestOnboardServer_OnboardWithUSB_FailActivate() {
	suite.mpi.On("Onboarded").Return(false, true).Once()
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	suite.mpi.On("Activate", mock.Anything).
		Return(false, errors.New("virtual error")).Once()
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	_, err2 := suite.app.serverInstance.OnboardWithUSB(context.Background(), deviceConfig)
	suite.NotNil(err2, "apply test failed", err2)
}

func (suite *TestSuite) TestOnboardServer_OnboardWithUSB_AlreadyOnboardedCase() {
	suite.mpi.On("Onboarded").Return(true, true).Once()
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()

	_, err2 := suite.app.serverInstance.OnboardWithUSB(context.Background(), deviceConfig)
	suite.NotNil(err2, "apply test failed", err2)
}

func (suite *TestSuite) TestOnboardServer_SetOnboardingStatus() {
	_, err := suite.app.serverInstance.SetOnboardingStatus(context.Background(), &v1.OnboardingStatus{IsOnboarded: true})
	suite.Nil(err, "apply test failed", err)
}

func (suite *TestSuite) TestOnboardServer_SetOnboardingStatusWithFalse() {
	_, err := suite.app.serverInstance.SetOnboardingStatus(context.Background(), &v1.OnboardingStatus{IsOnboarded: false})
	suite.Nil(err, "apply test failed", err)
}

func (suite *TestSuite) TestOnboardServer_SetOnboardingStatusWithNil() {
	_, err := suite.app.serverInstance.SetOnboardingStatus(context.Background(), nil)
	suite.NotNil(err, "apply test failed", err)
}

func (suite *TestSuite) TestOnboardServer_CreateHelper() {
	factory := ClientFactoryImpl{}
	suite.NotNil(factory.CreateHelper())
}

func (suite *TestSuite) TestOnboardServer_StartApp() {
	// Create a context with cancel to control the goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//first call should return error
	suite.mpi.On("Onboarded").Return(false, false).Once()
	//second call no error
	suite.mpi.On("Onboarded").Return(false, true).Once()

	// Pass context to StartApp
	suite.app.StartAppWithContext(ctx)

	// Wait for the goroutine to complete its checks
	time.Sleep(100 * time.Millisecond)

	// Cancel will stop the goroutine
	cancel()

	// Give time for cleanup
	time.Sleep(100 * time.Millisecond)
}

func (suite *TestSuite) TestOnboardServer_StartApp_TickerSuccess() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	suite.mpi.On("Onboarded").Return(false, false).Once()
	// Second call (first ticker) should return false
	suite.mpi.On("Onboarded").Return(false, false).Once()
	// Third call (second ticker) should return true and container accessible
	suite.mpi.On("Onboarded").Return(true, true).Once()

	suite.app.StartAppWithContext(ctx)

	// Wait for 2 ticker intervals plus buffer
	time.Sleep(2*onboardCheckInterval + 5*time.Second)

	suite.app.serverInstance.Mutex.Lock()
	isAccessible := suite.app.serverInstance.isContainerAccessible
	isOnboarded := suite.app.serverInstance.status.IsOnboarded
	suite.app.serverInstance.Mutex.Unlock()

	suite.True(isAccessible)
	suite.True(isOnboarded)

	// Cancel to cleanup
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func (suite *TestSuite) TestOnboardServer_checkAndUpdateOnboardStatus_Success() {
	// Mock successful onboard check
	suite.mpi.On("Onboarded").Return(true, true).Once()

	result := suite.app.checkAndUpdateOnboardStatus()

	suite.True(result, "Expected checkAndUpdateOnboardStatus to return true")
	suite.True(suite.app.serverInstance.isContainerAccessible, "Expected isContainerAccessible to be true")
	suite.True(suite.app.serverInstance.status.IsOnboarded, "Expected IsOnboarded status to be true")
}

func (suite *TestSuite) TestOnboardServer_checkAndUpdateOnboardStatus_ContainerNotAccessible() {
	// Mock container not accessible
	suite.mpi.On("Onboarded").Return(false, false).Once()

	result := suite.app.checkAndUpdateOnboardStatus()

	suite.False(result, "Expected checkAndUpdateOnboardStatus to return false")
}

func (suite *TestSuite) TestOnboardServer_checkAndUpdateOnboardStatus_NotOnboarded() {
	// Mock container accessible but not onboarded
	suite.mpi.On("Onboarded").Return(false, true).Once()

	result := suite.app.checkAndUpdateOnboardStatus()

	suite.True(result, "Expected checkAndUpdateOnboardStatus to return true when container is accessible")
	suite.True(suite.app.serverInstance.isContainerAccessible, "Expected isContainerAccessible to be true")
	suite.False(suite.app.serverInstance.status.IsOnboarded, "Expected IsOnboarded status to be false")
}

func (suite *TestSuite) TestOnboardServer_Subscribe() {
	suite.mpi.On("Onboarded").Return(false, true).Maybe()

	channel := make(chan struct{})
	listener := &MockListener{
		receivedStatus: &v1.OnboardingStatus{
			IsOnboarded: false,
			Message:     "",
		},
		context:     &mockContext{doneChannel: channel},
		doneChannel: channel,
	}

	go func(testSuite *TestSuite) {
		testSuite.app.serverInstance.ListenOnboardState(&emptypb.Empty{}, listener)
	}(suite)

	time.Sleep(200 * time.Millisecond)

	suite.app.serverInstance.SetOnboardingStatus(context.TODO(), &v1.OnboardingStatus{IsOnboarded: true})
	time.Sleep(200 * time.Millisecond)

	channel <- struct{}{}
	time.Sleep(200 * time.Millisecond)
}

func (suite *TestSuite) Test_VerifyArgsForStartGRPC_WithDummySocketForUnix() {
	//Create App to use
	tArgs := []string{"onboardservice", "unix", "/tmpa/temp.sock"}

	tErr := suite.app.StartGRPC(tArgs)

	//Connection failure is expected with dummy sock
	suite.NotNil(tErr, "Did not get expected result. got: %q", tErr)
}

func (suite *TestSuite) Test_VerifyArgsForStartGRPC() {
	done := make(chan bool)
	//Create App to use
	go func() {
		defer close(done)
		tArgs := []string{"onboardservice", "unix", "/tmp/temp.sock"}
		tErr := suite.app.StartGRPC(tArgs)

		//Connection failure is expected with dummy sock
		suite.Nil(tErr, "Did not get expected result. got: %q", tErr)
	}()

	time.Sleep(100 * time.Millisecond)
	suite.app.StopGRPC()
	<-done // Wait for goroutine to finish
}

func (suite *TestSuite) Test_VerifyArgsForStartGRPC_TCP() {
	done := make(chan bool)
	//Create App to use
	go func() {
		defer close(done)
		tArgs := []string{"onboardservice", "tcp", ":4444"}
		tErr := suite.app.StartGRPC(tArgs)
		suite.Nil(tErr, "Did not get expected result. got: %q", tErr)
	}()

	time.Sleep(100 * time.Millisecond)
	suite.app.StopGRPC()
	<-done // Wait for goroutine to finish
}

func (suite *TestSuite) Test_VerifyArgsForStartGRPCInvalidArgs() {
	//Create App to use
	tArgs := []string{"onboardservice", "unix", "hey", "/tmp/temp.sock"}
	tErr := suite.app.StartGRPC(tArgs)
	//Connection failure is expected with dummy sock
	suite.NotNil(tErr, "Did not get expected result. got: %q", tErr)
}

func (suite *TestSuite) TestOnboardServer_ClientFactory() {
	cf := ClientFactoryImpl{}
	var clients = cf.CreateClients()
	suite.NotNil(clients, "Client factory did not run")
}

var int1 = v1.Interface{
	GatewayInterface: false,
	MacAddress:       "12:13:14:15:15",
	DHCP:             "enabled",
	Static:           nil,
	DNSConfig:        nil,
}

var deviceConfig = &v1.DeviceConfiguration{
	Device: &v1.Device{
		Network:   &v1.NetworkSettings{Interfaces: []*v1.Interface{&int1}},
		NtpServer: []string{"some.server"},
	},
	Onboarding: &v1.Onboarding{PotalUrl: "edgeCoreRuntime.rest.cli"},
	Agents:     nil,
	Proxies: []*v1.ProxySettings{&v1.ProxySettings{
		ProxyId:            "",
		ProxyType:          "",
		Host:               "",
		Protocol:           "",
		AuthenticationType: "",
		User:               "",
		Password:           "",
		NoProxy:            "",
		CustomPorts:        nil,
	}},
}

func (m *MockClientFactory) CreateClients() *ClientPack {
	return &ClientPack{
		NetworkClient: &mockNetwork{},
		NtpClient:     &mockNtp{},
		RestClient:    &mockedgeCoreRuntime{},
	}
}
