/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package restclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	PORT = "3131"
)

type TestSuite struct {
	suite.Suite
	server *mockServer
}

func (suite *TestSuite) SetupSuite() {
	log.Println("suite starting")
	suite.server = &mockServer{}
	go func() { suite.server.Start() }()
	time.Sleep(1 * time.Second)
	log.Println("suite started")
}

func TestSuiteAll(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TearDownSuite() {
	log.Println("Closing Web Server")

	err := suite.server.Stop()
	if err != nil {
		log.Println("Closed web server")
	}

	log.Println("Test end")
}

func (suite *TestSuite) Test_TryActivate() {

	int1 := v1.Interface{
		GatewayInterface: false,
		MacAddress:       "12:13:14:15:15",
		DHCP:             "disabled",
		Static: &v1.Interface_StaticConf{
			IPv4:    "192.168.41.41",
			NetMask: "255.255.255.0",
			Gateway: "",
		},
		DNSConfig: nil,
	}

	deviceConfig := &v1.DeviceConfiguration{
		Device: &v1.Device{
			Network:   &v1.NetworkSettings{Interfaces: []*v1.Interface{&int1}},
			NtpServer: nil,
		},
		Onboarding: nil,
		Agents: []*v1.Agent{{
			Name:    "Smith",
			Proxy:   []string{"some proxy id"},
			AgentId: "47",
			Security: &v1.Security{
				BaseUrl:                 "secure by default",
				ClientCredentialProfile: []string{"#StayAtHome"},
				CaChain:                 "This product has been created during covid19 pandemic in April2020",
			}},
		},
		Proxies: nil,
	}

	client := NewClient(fmt.Sprintf("http://localhost:%s", PORT))
	val, err := client.Activate(deviceConfig)
	suite.True(val)
	suite.Nil(err)

}

func (suite *TestSuite) Test_CheckOnboardedWithMock() {
	client := NewClient(fmt.Sprintf("http://localhost:%s", PORT))
	isOnboarded, isContainerAccessible := client.Onboarded()
	suite.True(isOnboarded)
	suite.True(isContainerAccessible)
}

func (suite *TestSuite) Test_OnboardedWithInvalidPath() {
	client := NewClient(fmt.Sprintf("http://localhost:%s/invalid-path", PORT))
	isOnboarded, isContainerAccessible := client.Onboarded()
	suite.False(isOnboarded)
	suite.False(isContainerAccessible)
}

func (suite *TestSuite) Test_OnboardedFailingWith503StatusCode() {
	client := NewClient(fmt.Sprintf("http://localhost:%s/response-status-code/503", PORT))
	isOnboarded, isContainerAccessible := client.Onboarded()
	suite.False(isOnboarded)
	suite.False(isContainerAccessible)
}

func (suite *TestSuite) Test_OnboardedWith531StatusCode() {
	client := NewClient(fmt.Sprintf("http://localhost:%s/response-status-code/531", PORT))
	isOnboarded, isContainerAccessible := client.Onboarded()
	suite.False(isOnboarded)
	suite.True(isContainerAccessible)
}

func (suite *TestSuite) Test_OnboardedWithJSONResponse() {
	client := NewClient(fmt.Sprintf("http://localhost:%s/response-status-code/401", PORT))
	isOnboarded, isContainerAccessible := client.Onboarded()
	suite.False(isOnboarded)
	suite.False(isContainerAccessible)
}

type mockServer struct {
	server *http.Server
}

func (m *mockServer) Start() {
	s := &http.Server{
		Addr:         "127.0.0.1:" + PORT,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	m.server = s

	http.HandleFunc("/v1/activate", createHandler(200))
	http.HandleFunc("/v1", createHandler(200))
	http.HandleFunc("/response-status-code/503/v1", createHandler(503))
	http.HandleFunc("/response-status-code/531/v1", createHandler(531))
	http.HandleFunc("/response-status-code/401/v1", createHandler(401))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	s.ListenAndServe()
}

func createHandler(status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		body, err := io.ReadAll(r.Body)
		if err == nil {
			fmt.Fprintf(w, "Hello Request was, %q", string(body))
		}
	}
}

func (m *mockServer) Stop() error {
	return m.server.Close()
}
