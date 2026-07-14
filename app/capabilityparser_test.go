/*
 * Copyright © Siemens 2025 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	edgeutils "onboardservice/internal/mocks"
)

func Test_parseCapabilitiesFile_Success(t *testing.T) {
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
				APICapability: `{"CreateConnection":true}`,
			},
		},
	}

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", capabilitiesFilePath).Return(nil)
	fProvider.On("ReadFile", capabilitiesFilePath).Return(expectedData, nil)

	// Test using the provider
	capabilities, err := defaultCapabilityParser.parseCapabilitiesFile()
	assert.Nil(t, err, "Expected no error for valid capabilities file")
	assert.NotNil(t, capabilities, "Capabilities should not be nil")
	assert.Len(t, capabilities, 2, "Should have 2 capabilities")
	assert.Equal(t, "host.gsmconnectivity", capabilities[0].CapabilityID)
	assert.Equal(t, "host.networkservice", capabilities[1].CapabilityID)
	assert.Equal(t, "enabled", capabilities[0].Details.DefaultStatus)
	assert.Equal(t, "enabled", capabilities[1].Details.DefaultStatus)

}

func Test_parseCapabilitiesFile_NonExistentFile(t *testing.T) {
	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Mock file not existing - Stat should return an error
	fProvider.On("Stat", tmpFile).Return(errors.New("file not found"))

	// Test using the provider
	capabilities, err := defaultCapabilityParser.parseCapabilitiesFile()

	assert.NotNil(t, err, "Expected error for non-existent file")
	assert.Equal(t, "capabilities file not found at /etc/capabilities.json - please check path and retry again", err.Error())
	assert.Nil(t, capabilities, "Capabilities should be nil")
}

func Test_parseCapabilitiesFile_EmptyFile(t *testing.T) {
	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Mock file exists but is empty
	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return([]byte(""), nil)

	// Test using the provider
	capabilities, err := defaultCapabilityParser.parseCapabilitiesFile()

	assert.NotNil(t, err, "Expected error for empty file")
	assert.Equal(t, "unable to parse capability file - unexpected end of JSON input", err.Error())
	assert.Nil(t, capabilities, "Capabilities should be nil")
}

func Test_hasGSMCapability_FileNotFound(t *testing.T) {
	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Mock file not existing
	fProvider.On("Stat", tmpFile).Return(errors.New("file not found"))

	// Test using the provider
	err := defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for missing capabilities file")
	assert.Equal(t, "rpc error: code = Internal desc = Error parsing capabilities file: capabilities file not found at /etc/capabilities.json - please check path and retry again", err.Error())
}

func Test_hasGSMCapability_GSMCapabilityNotFound(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for missing GSM capability")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = GSM connectivity not supported on this device", err.Error())
}

func Test_hasGSMCapability_NetworkCapabilityNotFound(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for missing network capability")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = Network service capability not supported on this device", err.Error())
}

func Test_hasGSMCapability_GSMCapabilityDisabled(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "disabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":false}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for CreateConnection not supported")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = GSM connectivity not supported on this device", err.Error())
}

func Test_hasGSMCapability_NetworkCapabilityDisabled(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "disabled",
				APICapability: `{"CreateConnection":false}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for Network capability disabled")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = Network service capability not supported on this device", err.Error())
}

func Test_hasGSMCapability_CaseSensitivity(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "Disabled",
				APICapability: `{"CreateConnection":false}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for DefaultStatus case sensitivity")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = Network service capability not supported on this device", err.Error())
}

func Test_hasGSMCapability_CreateConnectionNotSupported(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":false}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for CreateConnection not supported")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = CreateConnection not supported in network service", err.Error())
}

func Test_hasGSMCapability_RemoveConnectionNotSupported(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":false}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for RemoveConnection not supported")
	assert.Equal(t, "rpc error: code = FailedPrecondition desc = RemoveConnection not supported in network service", err.Error())
}

func Test_hasGSMCapability_Success(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: `{"CreateConnection":true,"RemoveConnection":true}`,
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.Nil(t, err, "Expected no error for valid GSM capability")
}

func Test_parseCapabilitiesFile_ReadFileError(t *testing.T) {
	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Mock file exists but read fails
	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return([]byte{}, errors.New("permission denied"))

	// Test using the provider
	capabilities, err := defaultCapabilityParser.parseCapabilitiesFile()
	assert.NotNil(t, err, "Expected error for read failure")
	assert.Equal(t, "unable to read capability file - permission denied", err.Error())
	assert.Nil(t, capabilities, "Capabilities should be nil")
}

func Test_getNetworkCapabilitiesObj_InvalidJSON(t *testing.T) {
	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Create capability details with invalid JSON
	details := CapabilityDetails{
		APICapability: "invalid json format",
	}

	// Test using the provider
	result, err := defaultCapabilityParser.getNetworkCapabilitiesObj(details)
	assert.NotNil(t, err, "Expected error for invalid JSON")
	assert.Contains(t, err.Error(), "unable to parse capability details")
	assert.Nil(t, result, "Result should be nil")
}

func Test_hasGSMCapability_EmptyAPICapability(t *testing.T) {
	testCapabilities := []Capability{
		{
			CapabilityID: "host.gsmconnectivity",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
			},
		},
		{
			CapabilityID: "host.networkservice",
			Details: CapabilityDetails{
				DefaultStatus: "enabled",
				APICapability: "",
			},
		},
	}

	tmpFile := filepath.Join("/etc", "capabilities.json")

	fProvider := &edgeutils.MockFileSystem{}
	defaultCapabilityParser := NewCapabilityParser(fProvider)

	// Marshal test data to bytes for the mock
	expectedData, err := json.Marshal(testCapabilities)
	assert.Nil(t, err, "Failed to marshal test data")

	fProvider.On("Stat", tmpFile).Return(nil)
	fProvider.On("ReadFile", tmpFile).Return(expectedData, nil)

	// Test using the provider
	err = defaultCapabilityParser.HasGSMCapability()
	assert.NotNil(t, err, "Expected error for invalid JSON")
	assert.Contains(t, err.Error(), "unable to parse capability details")
}
