/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"log"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

//  write2Device writes given content to given path.
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

