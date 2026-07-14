/*
 * Copyright © Siemens 2020 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	edgeutils "onboardservice/internal/mocks"

	"os"
	"sync"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	mockClients *ClientPack
	mn          edgeutils.MockNtp
	mpi         mockedgeCoreRuntime
	mNet        edgeutils.MockNetwork
	mSys        edgeutils.MockSystem
	app         *MainApp
	mu          sync.Mutex
}

func (suite *TestSuite) SetupSuite() {
	os.Remove("/tmp/temp.sock")
	suite.app = CreateServiceApp(&MockClientFactory{})
	capabilityParser := &edgeutils.MockFileSystem{}
	NewCapabilityParser(capabilityParser)
	suite.app.Clients = suite.mockClients
	suite.app.serverInstance.isContainerAccessible = true
}

func (suite *TestSuite) SetupTest() {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	initMocks(suite)
}

func initMocks(suite *TestSuite) {	
	suite.mn = edgeutils.MockNtp{Mock: mock.Mock{}}
	suite.mpi = mockedgeCoreRuntime{Mock: mock.Mock{}}
	suite.mNet = edgeutils.MockNetwork{Mock: mock.Mock{}}
	suite.mSys = edgeutils.MockSystem{Mock: mock.Mock{}}

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
	suite.mNet.On("CreateConnection", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Maybe()
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
	suite.mNet.On("CreateConnection", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Maybe()
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

func (suite *TestSuite) TestOnboardServer_OnboardWithUSB_WithGSMConfig() {
	// Create device configuration with GSM interface
	gsmInterface := &v1.Interface{
		InterfaceName:    "gsm0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:18",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: &v1.Interface_GsmConf{
			Apn:      "internet.provider.com",
			Username: "gsmuser",
			Password: "gsmpass",
			Pin:      "1234",
		},
	}

	regularInterface := &v1.Interface{
		InterfaceName:    "eth0",
		GatewayInterface: true,
		MacAddress:       "12:13:14:15:17",
		DHCP:             "enabled",
		Static:           nil,
		DNSConfig:        nil,
	}

	deviceConfigWithGSM := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{regularInterface, gsmInterface},
			},
			NtpServer: []string{"ntp.server.com"},
		},
		Onboarding: &v1.Onboarding{PotalUrl: "edgeCoreRuntime.rest.cli"},
		Agents:     nil,
		Proxies:    []*v1.ProxySettings{},
	}

	// Mock capability parser to succeed GSM capability check
	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Mock GetAllInterfaces to return GSM interface
	interfaceType := v1.Interface_GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		&v1.NetworkSettings{
			Interfaces: []*v1.Interface{
				{
					InterfaceName: "gsm0",
					InterfaceType: &interfaceType,
				},
			},
		}, nil).Once()

	// Mock network settings application
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()

	// Mock GSM connection creation
	suite.mNet.On("CreateConnection", mock.Anything, mock.MatchedBy(func(settings *v1.ConnectionSettings) bool {
		return settings.Name == "ied-gsm-primary_gsm0" &&
			settings.ConnectionType == v1.ConnectionSettings_GSM &&
			settings.GetGsm().Apn == "internet.provider.com" &&
			settings.GetGsm().Username == "gsmuser" &&
			settings.GetGsm().Password == "gsmpass" &&
			settings.GetGsm().Pin == "1234"
	})).Return(&emptypb.Empty{}, nil).Once()

	// Mock NTP configuration
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()

	// Mock activation process
	suite.mpi.On("Onboarded").Return(false, true).Once()
	suite.mpi.On("Activate", mock.Anything).Return(true, nil).Once()

	// Call OnboardWithUSB
	_, err = suite.app.serverInstance.OnboardWithUSB(context.Background(), deviceConfigWithGSM)

	// Should succeed with GSM configuration
	suite.Nil(err, "OnboardWithUSB should succeed with GSM configuration")

	// Verify all mock expectations
	mockFileSystem.AssertExpectations(suite.T())
	suite.mNet.AssertExpectations(suite.T())
	suite.mn.AssertExpectations(suite.T())
	suite.mpi.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestOnboardWithUSB_WithGSMConfiguration_CreateConnectionFails_ReturnsError() {
	// Create device configuration with GSM interface
	gsmInterface := &v1.Interface{
		InterfaceName:    "gsm0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:18",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: &v1.Interface_GsmConf{
			Apn:      "internet.provider.com",
			Username: "gsmuser",
			Password: "gsmpass",
			Pin:      "1234",
		},
	}

	regularInterface := &v1.Interface{
		InterfaceName:    "eth0",
		GatewayInterface: true,
		MacAddress:       "12:13:14:15:17",
		DHCP:             "enabled",
		Static:           nil,
		DNSConfig:        nil,
	}

	deviceConfigWithGSM := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{regularInterface, gsmInterface},
			},
			NtpServer: []string{"ntp.server.com"},
		},
		Onboarding: &v1.Onboarding{PotalUrl: "edgeCoreRuntime.rest.cli"},
		Agents:     nil,
		Proxies:    []*v1.ProxySettings{},
	}

	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)
	// Mock GetAllInterfaces to return GSM interface
	interfaceType := v1.Interface_GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		&v1.NetworkSettings{
			Interfaces: []*v1.Interface{
				{
					InterfaceName: "gsm0",
					InterfaceType: &interfaceType,
				},
			},
		}, nil).Once()

	// Mock network settings application
	suite.mNet.On("ApplySettings", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()

	// Mock GSM connection creation
	suite.mNet.On("CreateConnection", mock.Anything, mock.MatchedBy(func(settings *v1.ConnectionSettings) bool {
		return settings.Name == "ied-gsm-primary_gsm0" &&
			settings.ConnectionType == v1.ConnectionSettings_GSM &&
			settings.GetGsm().Apn == "internet.provider.com" &&
			settings.GetGsm().Username == "gsmuser" &&
			settings.GetGsm().Password == "gsmpass" &&
			settings.GetGsm().Pin == "1234"
	})).Return(nil, errors.New("failed to create GSM connection")).Once()

	// Mock NTP configuration
	suite.mn.On("SetNtpServer", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()

	// Mock activation process
	suite.mpi.On("Onboarded").Return(false, true).Once()
	suite.mpi.On("Activate", mock.Anything).Return(true, nil).Once()

	// Call OnboardWithUSB
	_, err = suite.app.serverInstance.OnboardWithUSB(context.Background(), deviceConfigWithGSM)

	// Should fail with GSM configuration
	suite.NotNil(err, "OnboardWithUSB should fail when GSM connection creation fails")

}

func (suite *TestSuite) TestCreateGSMConnectionSettings_ShouldNeverReturnNil() {
	gsmConf := &v1.Interface_GsmConf{
		Apn: "internet",
		Pin: "",
	}

	deviceName := "wwan0"

	connSettings := suite.app.serverInstance.createGSMConnectionSettings(gsmConf, deviceName)

	assert.NotNil(suite.T(), connSettings, "createGSMConnectionSettings should never return nil")
	assert.Equal(suite.T(), v1.ConnectionSettings_GSM, connSettings.ConnectionType)
	assert.Equal(suite.T(), "internet", connSettings.GetGsm().Apn)
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
	GatewayInterface: true,
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
		NetworkClient: &edgeutils.MockNetwork{},
		NtpClient:     &edgeutils.MockNtp{},
		RestClient:    &mockedgeCoreRuntime{},
	}
}

func (suite *TestSuite) TestHandleGSMConfiguration_WithoutGSMInterface() {
	// Create device configuration with GSM interface
	gsmInterface := &v1.Interface{
		InterfaceName:    "wwan0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:16",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{gsmInterface},
			},
		},
	}

	gsmConfig := suite.app.serverInstance.handleGSMConfiguration(context.Background(), deviceConfig)
	assert.Nil(suite.T(), gsmConfig, "GSM configuration should not be nil")

}

// Tests for hasGSMConfiguration method
func (suite *TestSuite) TestHasGSMConfiguration_WithGSMInterface() {
	// Create device configuration with GSM interface
	gsmInterface := &v1.Interface{
		InterfaceName:    "wwan0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:16",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: &v1.Interface_GsmConf{
			Apn: "internet",
			Pin: "1234",
		},
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{gsmInterface},
			},
		},
	}

	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)
	assert.NotNil(suite.T(), gsmConfig, "GSM configuration should not be nil")

}

func (suite *TestSuite) TestHasGSMConfiguration_WithEmptyAPN_ShouldFail() {
	// Create device configuration with GSM interface but empty APN
	gsmInterface := &v1.Interface{
		InterfaceName:    "wwan0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:16",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: &v1.Interface_GsmConf{
			Apn: "", // Empty APN should cause validation failure
			Pin: "1234",
		},
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{gsmInterface},
			},
		},
	}

	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)

	// Expect nil because APN is empty
	assert.Nil(suite.T(), gsmConfig, "GSM configuration should be nil when APN is empty")
}

func (suite *TestSuite) TestHasGSMConfiguration_NilGSMInterface() {
	// Create device configuration without GSM interface
	regularInterface := &v1.Interface{
		InterfaceName:    "eth0",
		GatewayInterface: true,
		MacAddress:       "12:13:14:15:17",
		DHCP:             "enabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: nil, // No GSM configuration
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{regularInterface},
			},
		},
	}

	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)
	assert.Nil(suite.T(), gsmConfig, "GSM configuration is nil")
}

func (suite *TestSuite) TestHasGSMConfiguration_WithEmptyPin_ShouldPass() {
	// Create device configuration with GSM interface and empty PIN
	gsmInterface := &v1.Interface{
		InterfaceName:    "wwan0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:16",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: &v1.Interface_GsmConf{
			Apn: "internet", // APN is valid
			Pin: "",         // Empty PIN should be allowed
		},
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{gsmInterface},
			},
		},
	}

	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)

	// Expect GSM configuration to be returned since APN is present
	assert.NotNil(suite.T(), gsmConfig, "GSM configuration should not be nil when APN is provided")
	assert.Equal(suite.T(), "internet", gsmConfig.Apn)
	assert.Equal(suite.T(), "", gsmConfig.Pin)
}

func (suite *TestSuite) TestisValidGSMConfiguration_WithEmptyAPN_ShouldFail() {
	// Create GSM configuration with empty APN

	GsmConfiguration := &v1.Interface_GsmConf{
		Apn: "", // Empty APN should cause validation failure
		Pin: "1234",
	}

	isValidgsmConfig := suite.app.serverInstance.isValidGSMConfiguration(GsmConfiguration)

	// Expect false because APN is empty
	assert.False(suite.T(), isValidgsmConfig, "GSM configuration should be invalid when APN is empty")
}

func (suite *TestSuite) TestisValidGSMConfiguration_WithEmptyPin_ShouldPass() {
	// Create GSM configuration with empty APN

	GsmConfiguration := &v1.Interface_GsmConf{
		Apn: "internet", // APN is valid
		Pin: "",         // Empty PIN should be allowed
	}

	isValidgsmConfig := suite.app.serverInstance.isValidGSMConfiguration(GsmConfiguration)

	// Expect true because APN is valid and PIN is empty
	assert.True(suite.T(), isValidgsmConfig, "GSM configuration should be valid when APN is provided and PIN is empty")
}

func (suite *TestSuite) TestisValidGSMConfiguration_WithPin_ShouldPass() {
	// Create GSM configuration with empty APN

	GsmConfiguration := &v1.Interface_GsmConf{
		Apn: "internet", // APN is valid
		Pin: "1234",     // PIN should be allowed
	}

	isValidgsmConfig := suite.app.serverInstance.isValidGSMConfiguration(GsmConfiguration)

	// Expect false because APN is empty
	assert.True(suite.T(), isValidgsmConfig, "GSM configuration should be invalid when APN is empty")
}

func (suite *TestSuite) TestHasGSMConfiguration_EmptyNetworkSettings() {
	// Create device configuration with empty network settings
	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{},
			},
		},
	}

	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)
	assert.Nil(suite.T(), gsmConfig, "GSM configuration is nil")
}

func (suite *TestSuite) TestHasGSMConfiguration_NilNetworkSettings() {
	// Create device configuration with nil network settings
	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: nil,
		},
	}

	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)
	assert.Nil(suite.T(), gsmConfig, "GSM configuration is nil")
}

func (suite *TestSuite) TestHasGSMConfiguration_MixedInterfaces() {
	// Create device configuration with both regular and GSM interfaces
	regularInterface := &v1.Interface{
		InterfaceName:    "eth0",
		GatewayInterface: true,
		MacAddress:       "12:13:14:15:17",
		DHCP:             "enabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: nil,
	}

	gsmInterface := &v1.Interface{
		InterfaceName:    "wwan0",
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:18",
		DHCP:             "disabled",
		Static:           nil,
		DNSConfig:        nil,
		GsmConfiguration: &v1.Interface_GsmConf{
			Apn: "internet",
			Pin: "1234",
		},
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{regularInterface, gsmInterface},
			},
		},
	}
	gsmConfig := suite.app.serverInstance.hasGSMConfiguration(deviceConfig)
	assert.NotNil(suite.T(), gsmConfig, "GSM configuration should not be nil")
}

// Tests for checkGSMCapability method
func (suite *TestSuite) TestCheckGSMCapability_Success() {
	// Setup mock file provider
	mockFileSystem := &edgeutils.MockFileSystem{}
	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Test checkGSMCapability
	err = suite.app.serverInstance.checkGSMCapability()
	suite.Nil(err, "checkGSMCapability should succeed with valid capabilities")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestCheckGSMCapability_MissingGSMCapability() {
	// Setup mock file provider
	mockFileSystem := &edgeutils.MockFileSystem{}
	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities without GSM capability
	testCapabilities := []Capability{
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Test checkGSMCapability
	err = suite.app.serverInstance.checkGSMCapability()
	suite.NotNil(err, "checkGSMCapability should fail when GSM capability is missing")
	suite.Contains(err.Error(), "GSM connectivity not supported")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestCheckConnection_Settings() {
	gsmConfiguration := &v1.Interface_GsmConf{
		Apn: "internet",
		Pin: "1234",
	}
	// Test createGSMConnectionSettings with device name
	deviceName := "wwan0"
	connectionSetting := suite.app.serverInstance.createGSMConnectionSettings(gsmConfiguration, deviceName)
	suite.NotNil(connectionSetting, "Connection settings should not be nil")
	expectedConnectionName := fmt.Sprintf("%s_%s", gsmConnectionName, deviceName)
	suite.Equal(expectedConnectionName, connectionSetting.Name, "Connection name should include device name")
}

func (suite *TestSuite) TestCheckGSMCapability_FileReadError() {
	// Setup mock file provider
	mockFileSystem := &edgeutils.MockFileSystem{}
	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup mock expectations to return file read error
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return([]byte{}, errors.New("permission denied"))

	// Test checkGSMCapability
	err := suite.app.serverInstance.checkGSMCapability()
	suite.NotNil(err, "checkGSMCapability should fail when file read fails")
	suite.Contains(err.Error(), "Error parsing capabilities file")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestConfigureCustomSettings_Success() {
	// Create test custom configuration
	testConfig := &anypb.Any{
		TypeUrl: "type.googleapis.com/example.CustomConfig",
		Value:   []byte("test custom settings data"),
	}

	// Mock successful ApplyCustomSettings
	suite.mSys.On("ApplyCustomSettings", mock.Anything, testConfig).Return(&emptypb.Empty{}, nil).Once()

	// Test configureCustomSettings
	err := suite.app.serverInstance.configureCustomSettings(context.Background(), testConfig)
	suite.Nil(err, "configureCustomSettings should succeed")

	// Verify mock expectations
	suite.mSys.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestConfigureCustomSettings_Failure() {
	// Create test custom configuration
	testConfig := &anypb.Any{
		TypeUrl: "type.googleapis.com/example.CustomConfig",
		Value:   []byte("test custom settings data"),
	}

	// Mock failed ApplyCustomSettings
	suite.mSys.On("ApplyCustomSettings", mock.Anything, testConfig).
		Return(nil, errors.New("custom settings apply failed")).Once()

	// Test configureCustomSettings
	err := suite.app.serverInstance.configureCustomSettings(context.Background(), testConfig)
	suite.NotNil(err, "configureCustomSettings should fail when ApplyCustomSettings fails")
	suite.Contains(err.Error(), "Applying custom settings failed")

	// Verify mock expectations
	suite.mSys.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestConfigureCustomSettings_NilConfiguration() {
	// Test configureCustomSettings with nil configuration
	err := suite.app.serverInstance.configureCustomSettings(context.Background(), nil)
	suite.Nil(err, "configureCustomSettings should succeed with nil configuration")

	// No mock expectations should be called for nil configuration
}

func (suite *TestSuite) TestGetGSMDeviceName_Success() {
	// Create mock response with GSM interface
	gsmconnectionType := v1.Interface_GSM
	ethernetConnectionType := v1.Interface_ETHERNET
	mockNetworkSettings := &v1.NetworkSettings{
		Interfaces: []*v1.Interface{
			{
				InterfaceName: "eth0",
				InterfaceType: &ethernetConnectionType,
			},
			{
				InterfaceName: "wwan0",
				InterfaceType: &gsmconnectionType,
			},
		},
	}

	// Mock successful GetAllInterfaces
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).
		Return(mockNetworkSettings, nil).Once()

	// Test getGSMDeviceName
	deviceName, err := suite.app.serverInstance.getGSMDeviceName(context.Background())
	suite.Nil(err, "getGSMDeviceName should succeed")
	suite.Equal("wwan0", deviceName, "Should return GSM interface name")

	// Verify mock expectations
	suite.mNet.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestGetGSMDeviceName_NoGSMInterface() {
	// Create mock response without GSM interface

	ethernetConnectionType := v1.Interface_ETHERNET
	mockNetworkSettings := &v1.NetworkSettings{
		Interfaces: []*v1.Interface{
			{
				InterfaceName: "eth0",
				InterfaceType: &ethernetConnectionType,
			},
			{
				InterfaceName: "wlan0",
				InterfaceType: &ethernetConnectionType,
			},
		},
	}

	// Mock successful GetAllInterfaces but no GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).
		Return(mockNetworkSettings, nil).Once()

	// Test getGSMDeviceName
	deviceName, err := suite.app.serverInstance.getGSMDeviceName(context.Background())
	suite.NotNil(err, "getGSMDeviceName should fail when no GSM interface found")
	suite.Equal("", deviceName, "Should return empty string")
	suite.Contains(err.Error(), "no GSM interface found")

	// Verify mock expectations
	suite.mNet.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestGetGSMDeviceName_NetworkClientError() {
	// Mock GetAllInterfaces to fail
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).
		Return(nil, errors.New("network client error")).Once()

	// Test getGSMDeviceName
	deviceName, err := suite.app.serverInstance.getGSMDeviceName(context.Background())
	suite.NotNil(err, "getGSMDeviceName should fail when NetworkClient fails")
	suite.Equal("", deviceName, "Should return empty string")
	suite.Contains(err.Error(), "network client error")

	// Verify mock expectations
	suite.mNet.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestEmptyGSMDeviceName_InGSMInterface() {
	// Create mock response without GSM interface

	ethernetConnectionType := v1.Interface_ETHERNET
	gsmConnectionType := v1.Interface_GSM
	mockNetworkSettings := &v1.NetworkSettings{
		Interfaces: []*v1.Interface{
			{
				InterfaceName: "eth0",
				InterfaceType: &ethernetConnectionType,
			},
			{
				InterfaceName: "",
				InterfaceType: &gsmConnectionType,
			},
		},
	}

	// Mock successful GetAllInterfaces but no GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).
		Return(mockNetworkSettings, nil).Once()

	// Test getGSMDeviceName
	deviceName, err := suite.app.serverInstance.getGSMDeviceName(context.Background())
    suite.NotNil(err, "GSM interface found but interface name is empty")
	suite.Equal("", deviceName, "Should return empty string")
	suite.Contains(err.Error(), "GSM interface found but interface name is empty")

	// Verify mock expectations
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_NoGSMConfig tests the scenario where no GSM configuration is provided
func (suite *TestSuite) Test_handleGSMConfiguration_WithoutGSMConfig() {
	// Create configuration without GSM
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "eth0",
						GsmConfiguration: nil, // No GSM config
					},
				},
			},
		},
	}

	// Call method under test
	err := suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should succeed with no error since no GSM configuration is present
	suite.Nil(err)

	// No network client calls should be made
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_InvalidGSMConfig tests the scenario with invalid GSM configuration
func (suite *TestSuite) Test_handleGSMConfiguration_InvalidGSMConfigPresent() {
	// Create configuration with invalid GSM (empty APN)
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn: "", // Invalid: empty APN
						},
					},
				},
			},
		},
	}

	// Call method under test
	err := suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should succeed with no error since GSM configuration is invalid (empty APN)
	// handleGSMConfiguration returns early when hasGSMConfiguration returns nil
	suite.Nil(err)

	// No network client calls should be made
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_GSMCapabilityCheckFails tests GSM capability check failure
func (suite *TestSuite) Test_handleGSMConfiguration_ErrorInGSMCapabilityCheck() {
	// Create valid GSM configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn: "internet.provider.com",
						},
					},
				},
			},
		},
	}

	// Mock capability parser to fail GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "disabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Call method under test
	err = suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should fail because GSM capability is not found in capabilities file
	suite.NotNil(err)
	suite.Contains(err.Error(), "GSM connectivity not supported on this device")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_GetGSMDeviceNameFails tests the scenario when getting GSM device name fails
func (suite *TestSuite) Test_handleGSMConfiguration_ErrorInGetGSMDeviceName() {
	// Create valid GSM configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn: "internet.provider.com",
						},
						GatewayInterface: true,
					},
				},
			},
		},
	}

	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)
	// Mock GetAllInterfaces to fail
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		nil, errors.New("network service unavailable"))

	// Call method under test
	err = suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should fail because GetAllInterfaces failed
	suite.NotNil(err)
	suite.Contains(err.Error(), "network service unavailable")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_NoGSMInterfaceFound tests when no GSM interface is found in device
func (suite *TestSuite) Test_handleGSMConfiguration_NoGSMInterfaceFoundError() {
	// Create valid GSM configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn: "internet.provider.com",
						},
						GatewayInterface: true,
					},
				},
			},
		},
	}

	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Mock GetAllInterfaces to return interfaces without GSM type
	interfaceType := v1.Interface_ETHERNET
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		&v1.NetworkSettings{
			Interfaces: []*v1.Interface{
				{
					InterfaceName: "eth0",
					InterfaceType: &interfaceType,
				},
			},
		}, nil)

	// Call method under test
	err = suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should fail because no GSM interface is found
	suite.NotNil(err)
	suite.Contains(err.Error(), "no GSM interface found")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_EmptyGSMDeviceName tests when GSM interface has empty name
func (suite *TestSuite) Test_handleGSMConfiguration_EmptyGSMDeviceNameError() {
	// Create valid GSM configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn: "internet.provider.com",
						},
						GatewayInterface: true,
					},
				},
			},
		},
	}

	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Mock GetAllInterfaces to return GSM interface with empty name
	interfaceType := v1.Interface_GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		&v1.NetworkSettings{
			Interfaces: []*v1.Interface{
				{
					InterfaceName: "", // Empty name
					InterfaceType: &interfaceType,
				},
			},
		}, nil)

	// Call method under test
	err = suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should fail because GSM interface has empty name
	suite.NotNil(err)
	suite.Contains(err.Error(), "GSM interface found but interface name is empty")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
	// suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_CreateConnectionFails tests when CreateConnection fails
func (suite *TestSuite) Test_handleGSMConfiguration_CreateConnectionFails() {
	// Create valid GSM configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn: "internet.provider.com",
						},
						GatewayInterface: true,
					},
				},
			},
		},
	}

	// Mock capability parser to succeed GSM capability check
	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Mock GetAllInterfaces to return valid GSM interface
	interfaceType := v1.Interface_GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		&v1.NetworkSettings{
			Interfaces: []*v1.Interface{
				{
					InterfaceName: "gsm0",
					InterfaceType: &interfaceType,
				},
			},
		}, nil)

	// Mock CreateConnection to fail
	suite.mNet.On("CreateConnection", mock.Anything, mock.MatchedBy(func(settings *v1.ConnectionSettings) bool {
		return settings.Name == "ied-gsm-primary_gsm0" &&
			settings.ConnectionType == v1.ConnectionSettings_GSM &&
			settings.GetGsm().Apn == "internet.provider.com"
	})).Return(nil, errors.New("connection creation failed"))

	// Call method under test
	err = suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should fail because CreateConnection failed
	suite.NotNil(err)
	suite.Contains(err.Error(), "connection creation failed")

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_Success tests successful GSM configuration
func (suite *TestSuite) Test_handleGSMConfiguration_Success() {
	// Create valid GSM configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName: "gsm0",
						GsmConfiguration: &v1.Interface_GsmConf{
							Apn:      "internet.provider.com",
							Username: "testuser",
							Password: "testpass",
						},
						GatewayInterface: true,
					},
				},
			},
		},
	}

	// Mock capability parser to succeed GSM capability check
	mockFileSystem := &edgeutils.MockFileSystem{}

	suite.app.capabilityParser = &DefaultCapabilityParser{fileUtils: mockFileSystem}

	// Setup test capabilities
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Name:         "GSM Connectivity",
			Description:  "GSM connectivity support",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Name:         "Network Service",
			Description:  "Network service operations",
			Scope:        "protected",
			Property:     "optional",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	// Marshal test data
	expectedData, err := json.Marshal(testCapabilities)
	suite.Nil(err, "Failed to marshal test data")

	// Setup mock expectations
	mockFileSystem.On("Stat", capabilitiesFilePath).Return(nil)
	mockFileSystem.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Mock GetAllInterfaces to return valid GSM interface
	interfaceType := v1.Interface_GSM
	suite.mNet.On("GetAllInterfaces", mock.Anything, mock.Anything).Return(
		&v1.NetworkSettings{
			Interfaces: []*v1.Interface{
				{
					InterfaceName: "gsm0",
					InterfaceType: &interfaceType,
				},
			},
		}, nil)

	// Mock CreateConnection to succeed
	suite.mNet.On("CreateConnection", mock.Anything, mock.MatchedBy(func(settings *v1.ConnectionSettings) bool {
		return settings.Name == "ied-gsm-primary_gsm0" &&
			settings.ConnectionType == v1.ConnectionSettings_GSM &&
			settings.GetGsm().Apn == "internet.provider.com" &&
			settings.GetGsm().Username == "testuser" &&
			settings.GetGsm().Password == "testpass"
	})).Return(&emptypb.Empty{}, nil)

	// Call method under test
	err = suite.app.serverInstance.handleGSMConfiguration(context.Background(), config)

	// Should succeed
	suite.Nil(err)

	// Verify mock expectations
	mockFileSystem.AssertExpectations(suite.T())
	suite.mNet.AssertExpectations(suite.T())
}

// Test_handleGSMConfiguration_NilConfiguration tests with nil configuration
func (suite *TestSuite) Test_handleGSMConfiguration_NilConfiguration() {
	// Call method with nil configuration
	err := suite.app.serverInstance.handleGSMConfiguration(context.Background(), nil)

	// Should succeed with no error since no GSM configuration is present
	suite.Nil(err)

	// No network client calls should be made
	suite.mNet.AssertExpectations(suite.T())
}

func (suite *TestSuite) TestVerifyGatewayInterface_NoInterfaces() {
	// Create device configuration with no interfaces
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.NotNil(err, "Expected error when no interfaces are present")
}

func (suite *TestSuite) TestVerifyGatewayInterface_WithDHCP() {
	// Create device configuration with DHCP enabled and nil static configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						DHCP:  "enabled",
						Static: nil,
						GatewayInterface: true,
					},
					
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.Nil(err, "Expected no error when DHCP interfaces is present")
}

func (suite *TestSuite) TestVerifyGatewayInterface_WithDHCPAndEmptyStatic() {
	// Create device configuration with DHCP enabled and empty static configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						DHCP:  "enabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "",
							Gateway: "",
							NetMask: "",
						},
						GatewayInterface: true,
					},
					
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.Nil(err, "Expected no error when DHCP interfaces and empty static configuration are present")
}

func (suite *TestSuite) TestVerifyGatewayInterface_WithDifferentDHCPAndEmptyStatic() {
	// Create device configuration with DHCP disabled and empty static configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						DHCP:  "EnAbLed",
						Static: &v1.Interface_StaticConf{
							IPv4:    "",
							Gateway: "",
							NetMask: "",
						},
						GatewayInterface: true,
					},
					
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.Nil(err, "Expected no error when DHCP string is different case and empty static configuration are present")
}

func (suite *TestSuite) TestVerifyGatewayInterface_WithDHCPDisabledAndEmptyStatic() {
	// Create device configuration with DHCP disabled and empty static configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						DHCP:  "disabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "",
							Gateway: "",
							NetMask: "",
						},
						GatewayInterface: true,
					},
					
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.NotNil(err, "Expected error when DHCP is disabled and static configuration is empty")
}

func (suite *TestSuite) TestVerifyGatewayInterface_WithDHCPDisabledAndNonEmptyStatic() {
	// Create device configuration with DHCP disabled and non-empty static configuration
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						DHCP:  "disabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "192.168.1.100",
							Gateway: "192.168.1.1",
							NetMask: "255.255.255.0",
						},
						GatewayInterface: true,
					},
					
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.Nil(err, "Expected no error when DHCP is disabled and static configuration is non-empty")
}



func (suite *TestSuite) TestVerifyGatewayInterface_NoGatewayInterface() {
	// Create device configuration with interfaces but none marked as gateway
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "eth0",
						GatewayInterface: false,
					},
					{
						InterfaceName:    "wlan0",
						GatewayInterface: false,
					},
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.NotNil(err, "Expected error when no gateway interface is present")
	suite.Contains(err.Error(), "no gateway interface found in device configuration")
}

func (suite *TestSuite) TestVerifyGatewayInterface_MultipleGatewayInterfaces() {
	// Create device configuration with multiple gateway interfaces
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "eth0",
						GatewayInterface: true,
					},
					{
						InterfaceName:    "wlan0",
						GatewayInterface: true,
					},
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.NotNil(err, "Expected error when multiple gateway interfaces are present")
	suite.Contains(err.Error(), "multiple gateway interfaces are not supported")
}

func (suite *TestSuite) TestVerifyGatewayInterface_InValidConfigurationWithStatic() {
	// Create device configuration with invalid configuration with static
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "eth0",
						GatewayInterface: false,
						DHCP:             "enabled",
					},
					{
						InterfaceName:    "wlan0",
						GatewayInterface: true,
						DHCP:             "disabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "192.168.1.100",
							Gateway: "",
							NetMask: "255.255.255.0",
						},
					},
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.NotNil(err, "Expected no error when exactly one gateway interface is present")
	suite.Contains(err.Error(), "gateway interface has static IP configuration but gateway is empty")
}

func (suite *TestSuite) TestVerifyGatewayInterface_ValidConfigurationWithDHCP() {
	// Create device configuration with valid configuration with DHCP
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "eth0",
						GatewayInterface: true,
						DHCP:             "enabled",
					},
					{
						InterfaceName:    "wlan0",
						GatewayInterface: false,
						DHCP:             "disabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "192.168.1.100",
							Gateway: "192.168.1.1",
							NetMask: "255.255.255.0",
						},
					},
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.Nil(err, "Expected no error when exactly one gateway interface is present")
}

func (suite *TestSuite) TestVerifyGatewayInterface_ValidConfigurationWithStatic() {
	// Create device configuration with valid configuration with static IP
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "eth0",
						GatewayInterface: false,
						DHCP:             "enabled",
					},
					{
						InterfaceName:    "wlan0",
						GatewayInterface: true,
						DHCP:             "disabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "192.168.1.100",
							Gateway: "192.168.1.1",
							NetMask: "255.255.255.0",
						},
					},
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.Nil(err, "Expected no error when exactly one gateway interface is present")
}

func (suite *TestSuite) TestVerifyGatewayInterface_InValidConfigurationWithStaticOrder() {
	// Create device configuration with invalid configuration with static where gateway interface is not in order
	config := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network: &v1.NetworkSettings{
				Interfaces: []*v1.Interface{
					{
						InterfaceName:    "wlan0",
						GatewayInterface: true,
						DHCP:             "disabled",
						Static: &v1.Interface_StaticConf{
							IPv4:    "192.168.1.100",
							Gateway: "",
							NetMask: "255.255.255.0",
						},
					},
					{
						InterfaceName:    "eth0",
						GatewayInterface: false,
						DHCP:             "enabled",
					},
				},
			},
		},
	}

	err := suite.app.serverInstance.verifyGatewayInterface(config)
	suite.NotNil(err, "Expected error when exactly one gateway interface is present")
	suite.Contains(err.Error(), "gateway interface has static IP configuration but gateway is empty")
}
