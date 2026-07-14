/*
 * Copyright © Siemens 2025 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"encoding/json"
	"errors"
	"log"

	edgeutils "onboardservice/internal/utils"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	capabilitiesFilePath       = "/etc/capabilities.json"
	gsmCapabilityID            = "host.gsmconnectivity"
	networkServiceCapabilityID = "host.networkservice"
	enabledStatus              = "enabled"
)

// NewCapabilityParser creates and returns a new capability parser instance
// This is the recommended way to create a capability parser with dependency injection
func NewCapabilityParser(fs edgeutils.FileSystemOperations) *DefaultCapabilityParser {
	provider := &DefaultCapabilityParser{
		fileUtils: fs,
	}

	return provider
}

// DefaultCapabilityParser provides the default capabilities file path
// and implements CapabilityParser interface
type DefaultCapabilityParser struct {
	fileUtils edgeutils.FileSystemOperations
}

// CapabilityDetails represents the details section of a capability
type CapabilityDetails struct {
	DefaultStatus string `json:"defaultStatus"`
	APICapability string `json:"apiCapability,omitempty"`
}

// Capability represents a single capability entry
type Capability struct {
	CapabilityID string            `json:"capabilityId"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Scope        string            `json:"scope"`
	Property     string            `json:"property"`
	Details      CapabilityDetails `json:"details"`
}

// NetworkCapabilityObj holds the capability details for the Device-kit Network service
type NetworkCapabilityObj struct {
	IsEnabled             bool `json:"isEnabled,omitempty"`
	ApplySettings         bool `json:"ApplySettings,omitempty"`
	GetAllInterfaces      bool `json:"GetAllInterfaces,omitempty"`
	GetInterfaceWithMac   bool `json:"GetInterfaceWithMac,omitempty"`
	GetInterfaceWithLabel bool `json:"GetInterfaceWithLabel,omitempty"`
	CreateConnection      bool `json:"CreateConnection,omitempty"`
	RemoveConnection      bool `json:"RemoveConnection,omitempty"`
}

// parseCapabilitiesFile - reads capabilities.json file and parses the contents into a CapabilitiesBaseVO array
// Uses the default capability parser
func (d *DefaultCapabilityParser) parseCapabilitiesFile() (capabilities []Capability, err error) {

	isFileExist := d.fileUtils.Stat(capabilitiesFilePath) == nil
	if !isFileExist {
		return nil, errors.New("capabilities file not found at " + capabilitiesFilePath + " - please check path and retry again")
	}

	contentByte, err := d.fileUtils.ReadFile(capabilitiesFilePath)
	if err != nil {
		return nil, errors.New("unable to read capability file - " + err.Error())
	}

	err = json.Unmarshal(contentByte, &capabilities)
	if err != nil {
		log.Println("unable to parse capability file - ", err)
		return nil, errors.New("unable to parse capability file - " + err.Error())
	}

	return capabilities, nil
}

// checkGSMCapability checks if GSM is supported when GSM interfaces are present
func (d *DefaultCapabilityParser) HasGSMCapability() error {

	capabilities, err := d.parseCapabilitiesFile()
	if err != nil {
		log.Println("Error parsing capabilities file: ", err)
		return status.New(codes.Internal, "Error parsing capabilities file: "+err.Error()).Err()
	}

	gsmCapability := d.getCapability(capabilities, gsmCapabilityID)
	if gsmCapability == nil {
		return status.New(codes.FailedPrecondition, "GSM connectivity not supported on this device").Err()
	}

	networkcapability := d.getCapability(capabilities, networkServiceCapabilityID)
	if networkcapability == nil {
		return status.New(codes.FailedPrecondition, "Network service capability not supported on this device").Err()
	}

	networkCapabilityObj, err := d.getNetworkCapabilitiesObj(networkcapability.Details)
	if err != nil {
		log.Println("Error parsing network service capability details: ", err)
		return status.New(codes.Internal, "Error parsing network service capability details: "+err.Error()).Err()
	}

	if !networkCapabilityObj.CreateConnection {
		return status.New(codes.FailedPrecondition, "CreateConnection not supported in network service").Err()
	}

	if !networkCapabilityObj.RemoveConnection {
		return status.New(codes.FailedPrecondition, "RemoveConnection not supported in network service").Err()
	}

	return nil
}

// isCapabilityEnabled checks if capability is present and enabled using VO
func (d *DefaultCapabilityParser) getCapability(capabilitiesVO []Capability, capabilityID string) *Capability {
	for _, capability := range capabilitiesVO {
		if capability.CapabilityID == capabilityID && strings.ToLower(capability.Details.DefaultStatus) == enabledStatus {
			return &capability
		}
	}
	return nil
}

func (d *DefaultCapabilityParser) getNetworkCapabilitiesObj(details CapabilityDetails) (*NetworkCapabilityObj, error) {
	var networkCapabilities NetworkCapabilityObj

	err := json.Unmarshal([]byte(details.APICapability), &networkCapabilities)
	if err != nil {
		log.Println("unable to parse capability details - ", err)
		return nil, status.New(codes.Internal, "unable to parse capability details - "+err.Error()).Err()
	}

	return &networkCapabilities, nil
}
