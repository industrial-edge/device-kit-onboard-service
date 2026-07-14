/*
 * Copyright © Siemens 2020 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"fmt"
	"log"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	"strings"

	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	// GSM connection name constants
	//
	// The GSM connection name is constructed using the device name and the "ied-gsm-primary" suffix.
	// This same connection name is used in the edgecorelite repository for GSM connection creation during factory reset.
	//
	// CRITICAL: If this name must be changed, follow these steps in order:
	// 1. Update the corresponding cleanup code in the edgecorelite repository FIRST
	// 2. Ensure IERT hard reset tests pass with the new name
	// 3. Coordinate release deployment to avoid breaking factory reset functionality
	//
	gsmConnectionName       = "ied-gsm-primary"
	gsmConnectionNameFormat = "ied-gsm-primary_%s"
)

//Internals for OnboardServer

func (o *OnboardServer) configureNetwork(ctx context.Context, configuration *v1.DeviceConfiguration) error {
	if configuration.Device.Network != nil && configuration.Device.Network.Interfaces != nil &&
		len(configuration.Device.Network.Interfaces) > 0 {
		log.Println("Applying  Network settings ...")
		_, err2 := o.parentApp.Clients.NetworkClient.ApplySettings(ctx, configuration.Device.Network)
		if err2 != nil {
			log.Println("Failed Network settings")
			return status.New(codes.Internal, "Applying Network settings failed: "+err2.Error()).Err()
		}
		log.Println("Applied Network settings")
	}
	return nil
}

func (o *OnboardServer) configureNTP(ctx context.Context, configuration *v1.DeviceConfiguration) error {
	if len(configuration.Device.NtpServer) > 0 {
		ntp := v1.Ntp{NtpServer: configuration.Device.NtpServer}
		log.Println("Applying  Ntp settings ...")
		_, err := o.parentApp.Clients.NtpClient.SetNtpServer(ctx, &ntp)
		if err != nil {
			log.Println("Failed Ntp settings")
			return status.New(codes.Internal, "Ntpserver set failed: "+err.Error()).Err()
		}
		log.Println("Applied Ntp settings")
	}
	return nil
}

func (o *OnboardServer) configureCustomSettings(ctx context.Context, customConfiguration *anypb.Any) error {
	if customConfiguration != nil {
		log.Println("ApplyCustomSettings  rpc method will be called...")
		_, err2 := o.parentApp.Clients.SystemClient.ApplyCustomSettings(ctx, customConfiguration)
		if err2 != nil {
			log.Println("Failed apply custom settings")
			return status.New(codes.Internal, "Applying custom settings failed: "+err2.Error()).Err()
		}
		log.Println("Applied custom settings")
	}

	return nil
}

// write2Device writes given content to given path.
func (o *OnboardServer) write2Device(content string, filePath string) error {
	// create file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
	return nil
}

// hasGSMConfiguration checks if the device configuration contains valid GSM configuration
// Returns the GSM configuration if found, otherwise returns nil
func (o *OnboardServer) hasGSMConfiguration(configuration *v1.DeviceConfiguration) (gsmConfig *v1.Interface_GsmConf) {
	// Check if configuration structure is valid
	if configuration == nil || configuration.Device == nil || configuration.Device.Network == nil {
		log.Println("Device configuration or network settings is nil")
		return nil
	}

	// Check if interfaces exist
	if len(configuration.Device.Network.Interfaces) == 0 {
		log.Println("No network interfaces found in device configuration")
		return nil
	}

	for _, iface := range configuration.Device.Network.Interfaces {
		if iface.GsmConfiguration != nil {
			// Validate GSM configuration has required fields
			if o.isValidGSMConfiguration(iface.GsmConfiguration) {
				log.Printf("Found valid GSM configuration on interface: %s", iface.InterfaceName)
				return iface.GsmConfiguration
			}
		}
	}
	return nil
}

// isValidGSMConfiguration validates that GSM configuration has required non-empty fields
func (o *OnboardServer) isValidGSMConfiguration(gsmConfig *v1.Interface_GsmConf) bool {

	// Check required fields: APN must not be empty
	if len(gsmConfig.Apn) == 0 {
		log.Println("GSM configuration validation failed: APN is empty")
		return false
	}
	return true
}

// checkGSMCapability validates that the device supports GSM capabilities
// Should be called only after confirming GSM configuration exists
func (o *OnboardServer) checkGSMCapability() error {
	log.Println("GSM interface detected, checking device capabilities...")
	err := o.parentApp.capabilityParser.HasGSMCapability()
	if err != nil {
		log.Printf("GSM capability check failed: %v", err)
		return err
	}

	return nil
}

// getGSMDeviceName calls NetworkClient.GetAllInterfaces to find GSM interfaces and returns device name
func (o *OnboardServer) getGSMDeviceName(ctx context.Context) (string, error) {
	// Call NetworkClient to get all interfaces
	allInterfaces, err := o.parentApp.Clients.NetworkClient.GetAllInterfaces(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Failed to get all interfaces: %v", err)
		return "", err
	}

	// Look for GSM interface type
	for _, iface := range allInterfaces.Interfaces {

		if iface.InterfaceType != nil && *iface.InterfaceType == v1.Interface_GSM {
			if len(iface.InterfaceName) == 0 {
				log.Printf("GSM interface found but interface name is empty: %v", iface)
				return "", fmt.Errorf("GSM interface found but interface name is empty")
			}
			log.Printf("Found GSM interface: %s", iface.InterfaceName)
			return iface.InterfaceName, nil
		}
	}

	// No GSM interface found
	return "", fmt.Errorf("no GSM interface found")
}

// createGSMConnectionSettings creates connection settings from validated GSM configuration
// The gsmConf parameter should already be validated using isValidGSMConfiguration
func (o *OnboardServer) createGSMConnectionSettings(gsmConf *v1.Interface_GsmConf, deviceName string) *v1.ConnectionSettings {

	// Create connection name: ied-gsm-primary_deviceName
	connectionName := fmt.Sprintf(gsmConnectionNameFormat, deviceName)

	log.Printf("Creating GSM connection settings for APN: %s, Connection Name: %s", gsmConf.Apn, connectionName)
	connectionSettings := &v1.ConnectionSettings{
		Name:           connectionName,
		ConnectionType: v1.ConnectionSettings_GSM,
		ConnectionConfig: &v1.ConnectionSettings_Gsm{
			Gsm: gsmConf,
		},
	}
	return connectionSettings
}

// handleGSMConfiguration encapsulates all GSM-related logic in a self-contained method
// This method checks for GSM configuration, validates capability, and creates GSM connections
func (o *OnboardServer) handleGSMConfiguration(ctx context.Context, configuration *v1.DeviceConfiguration) error {
	// Check if GSM configuration is present and valid
	gsmConfig := o.hasGSMConfiguration(configuration)
	if gsmConfig == nil {
		// No GSM configuration found, continue with normal flow
		return nil
	}

	// Verify GSM capability is supported on the device
	// Check gsm.connectivity enabled and that network service has both create and remove for having a rollback ability
	if err := o.checkGSMCapability(); err != nil {
		log.Println("GSM capability check failed: ", err)
		return err
	}

	// GSM capability check succeeded, now get GSM device name and create GSM connections
	log.Println("GSM capability check succeeded...")

	if err := o.verifyGatewayInterface(configuration); err != nil {
		log.Println("Gateway interface verification failed: " + err.Error())
		return err
	}

	deviceName, err := o.getGSMDeviceName(ctx)
	if err != nil {
		log.Printf("Failed to get GSM device name: %v", err)
		return err
	}

	connectionSettings := o.createGSMConnectionSettings(gsmConfig, deviceName)
	if connectionSettings == nil {
		return status.New(codes.Internal, "Failed to create GSM connection settings").Err()
	}

	// Call CreateConnection for GSM interfaces
	// for onboarding retry we need a clean starting point,  BUT in future if we allow the GSM connections to be editable 
   //  we should introduce a proper rollback for the CreateConnection, similar to the ApplySettings()
	if _, err := o.parentApp.Clients.NetworkClient.CreateConnection(ctx, connectionSettings); err != nil {
		log.Printf("Network service CreateConnection failed: %v", err)	

		return err
	}

	log.Println("GSM connections created successfully")
	return nil
}

// verifyGatewayInterface checks that exactly one gateway interface is defined in the device configuration
func (o *OnboardServer) verifyGatewayInterface(configuration *v1.DeviceConfiguration) error {

	var gatewayInterfaceCount int
	var gatewayInterface *v1.Interface
	for _, iface := range configuration.Device.Network.Interfaces {
		if iface.GatewayInterface {
			gatewayInterfaceCount++
			gatewayInterface = iface
		}
	}
	if gatewayInterfaceCount == 0 {
		log.Println("No gateway interface found in device configuration")
		return status.New(codes.InvalidArgument, "no gateway interface found in device configuration").Err()
	}
	if gatewayInterfaceCount > 1 {
		log.Println("Multiple gateway interfaces found in device configuration")
		return status.New(codes.InvalidArgument, "multiple gateway interfaces are not supported").Err()
	}
	 // Normalize DHCP value to lowercase for consistent comparison
	 // as this same comparison logic will be used in network service for deciding whether to apply static IP validation or not

	gatewayInterface.DHCP=strings.ToLower(gatewayInterface.DHCP)    
	if gatewayInterface.DHCP == enabledStatus {
		log.Println("Gateway interface has DHCP enabled, skipping static IP validation")
		return nil
	}
	
	if gatewayInterface.Static != nil {	
		if gatewayInterface.Static.Gateway == "" {
			log.Println("Gateway interface has static IP configuration but gateway is empty")
			return status.New(codes.InvalidArgument, "gateway interface has static IP configuration but gateway is empty").Err()
		}
	}
	return nil
}
