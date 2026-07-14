/*
 * Copyright © Siemens 2020 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"
	"onboardservice/internal/system"

	edgeutils "onboardservice/internal/utils"
	"os"
	"os/user"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const onboardCheckInterval = 10 * time.Second

// DeviceModelService defines the common service interface
type DeviceModelService interface {
	StartGRPC(args []string)
	StartApp()
}

// MainApp application type
type MainApp struct {
	Clients          *ClientPack
	serverInstance   *OnboardServer
	helper           system.Configurable
	capabilityParser *DefaultCapabilityParser
	close            chan bool
	wg               sync.WaitGroup
}

type listenerInfo struct {
	id       int32
	channel  chan *v1.OnboardingStatus
	listener v1.OnboardService_ListenOnboardStateServer
}

func createNewListenerInfo(thelistener v1.OnboardService_ListenOnboardStateServer) listenerInfo {
	return listenerInfo{
		id:       rand.Int31(),
		channel:  make(chan *v1.OnboardingStatus),
		listener: thelistener,
	}
}

// CreateServiceApp creates main App with given clientpack factory
func CreateServiceApp(factory ClientFactory) *MainApp {

	fsOps := &edgeutils.OsFileSystemOperations{}
	// Create capability parser instance using improved pattern
	capParser := NewCapabilityParser(fsOps)

	app := MainApp{
		Clients:          factory.CreateClients(),
		serverInstance:   &OnboardServer{},
		close:            make(chan bool),
		helper:           factory.CreateHelper(),
		capabilityParser: capParser,
	}

	server := OnboardServer{
		listeners: map[int32]listenerInfo{},
		status: &v1.OnboardingStatus{
			IsOnboarded: false,
			Message:     "",
		},
		parentApp: &app,
		Mutex:     sync.Mutex{},
	}
	app.serverInstance = &server
	return &app
}

// StopGRPC Sends stop signal
func (app MainApp) StopGRPC() {
	go func() { app.close <- true }()
}

// StartGRPC is the entry point
func (app MainApp) StartGRPC(args []string) error {
	const message string = `ERROR: Could not start monitor with bad arguments! \n  
		Sample usage:\n  ./xxxservice unix /tmp/iedk/xxx.sock \n 
		./xxxservice tcp localhost:50006`

	if len(args) != 3 {
		fmt.Println(message)
		return errors.New("parameter not supported")
	}
	typeOfConnection := args[1]
	address := args[2]
	if typeOfConnection != "unix" && typeOfConnection != "tcp" {
		fmt.Println(message)
		return errors.New("parameter not supported: " + typeOfConnection)
	}

	if typeOfConnection == "unix" {

		if err := os.RemoveAll(os.Args[2]); err != nil {
			return errors.New("socket could not removed: " + typeOfConnection)
		}
	}

	lis, err := net.Listen(typeOfConnection, address)

	if err != nil {
		log.Printf("failed to listen: %v", err)
		return errors.New("Failed to listen: " + err.Error())

	}
	if chownSocket(address, "root", "docker") != nil {
		log.Println("failed permissions")
		//	return err
	}

	log.Print("Started listening on : ", typeOfConnection, " - ", address)
	s := grpc.NewServer()
	go func(s *grpc.Server) {
		_ = <-app.close
		s.Stop()
	}(s)
	v1.RegisterOnboardServiceServer(s, app.serverInstance)
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
		return errors.New("Failed to serve: " + err.Error())
	}

	return nil
}

func chownSocket(address string, userName string, groupName string) error {
	us, err1 := user.Lookup(userName)
	group, err2 := user.LookupGroup(groupName)
	if err1 == nil && err2 == nil {
		uid, _ := strconv.Atoi(us.Uid)
		gid, _ := strconv.Atoi(group.Gid)
		err3 := os.Chmod(address, os.FileMode.Perm(0660))
		err4 := os.Chown(address, uid, gid)
		if err3 != nil || err4 != nil {
			return errors.New("File permissions failed")
		}
		log.Println(uid, " : ", gid)
		return nil
	}
	return errors.New("File permissions failed")
}

// StartApp starts things up
func (app *MainApp) StartApp() {
	app.StartAppWithContext(context.Background())
}

// StartAppWithContext starts things up with a context for cancellation
func (app *MainApp) StartAppWithContext(ctx context.Context) {
	log.Println("Initializing App - Starting onboard status monitoring")

	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		if app.checkAndUpdateOnboardStatus() {
			log.Println("Onboard status monitoring completed successfully")
			return
		}

		log.Printf("Initial onboard check: container not accessible yet, will retry in %v...", onboardCheckInterval)

		ticker := time.NewTicker(onboardCheckInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Println("Context cancelled, stopping onboard status monitoring")
				return
			case <-ticker.C:
				if app.checkAndUpdateOnboardStatus() {
					log.Println("Onboard status monitoring completed successfully")
					return
				}
			}
		}
	}()
}

// checks the onboard status and updates the server instance if accessible
// Returns true if the container is accessible, false otherwise
func (app *MainApp) checkAndUpdateOnboardStatus() bool {
	onboardStatus, isContainerAccessible := app.Clients.RestClient.Onboarded()
	if isContainerAccessible {
		log.Printf("Onboard status check successful. Onboarded: %v", onboardStatus)
		app.serverInstance.Mutex.Lock()
		app.serverInstance.isContainerAccessible = isContainerAccessible
		app.serverInstance.status.IsOnboarded = onboardStatus
		app.serverInstance.Mutex.Unlock()
		app.serverInstance.notifyListeners(&v1.OnboardingStatus{IsOnboarded: onboardStatus})
		return true
	}
	return false
}
