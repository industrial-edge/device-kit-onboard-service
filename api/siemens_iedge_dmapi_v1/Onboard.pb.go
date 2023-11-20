//
// Copyright (c) 2023 Siemens AG
// Licensed under the MIT license
// See LICENSE file in the top-level directory

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: Onboard.proto

package siemens_iedge_dmapi_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Onboarding section in IEM Onboarding JSON file.
type Onboarding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalUserName         string `protobuf:"bytes,1,opt,name=localUserName,proto3" json:"localUserName,omitempty"`
	LocalPassword         string `protobuf:"bytes,2,opt,name=localPassword,proto3" json:"localPassword,omitempty"`
	DeviceName            string `protobuf:"bytes,3,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceId              string `protobuf:"bytes,4,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	PlatformActualName    string `protobuf:"bytes,5,opt,name=platformActualName,proto3" json:"platformActualName,omitempty"`
	SoftwarePlatformName  string `protobuf:"bytes,6,opt,name=softwarePlatformName,proto3" json:"softwarePlatformName,omitempty"`
	PotalUrl              string `protobuf:"bytes,7,opt,name=potalUrl,proto3" json:"potalUrl,omitempty"`
	NodeId                string `protobuf:"bytes,8,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	UserId                string `protobuf:"bytes,9,opt,name=userId,proto3" json:"userId,omitempty"`
	DeviceType            string `protobuf:"bytes,10,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	SwPlatformId          string `protobuf:"bytes,11,opt,name=swPlatformId,proto3" json:"swPlatformId,omitempty"`
	PlatformId            string `protobuf:"bytes,12,opt,name=platformId,proto3" json:"platformId,omitempty"`
	IsActivationConfirmed bool   `protobuf:"varint,13,opt,name=isActivationConfirmed,proto3" json:"isActivationConfirmed,omitempty"`
	DeviceRole            string `protobuf:"bytes,14,opt,name=deviceRole,proto3" json:"deviceRole,omitempty"`
}

func (x *Onboarding) Reset() {
	*x = Onboarding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Onboarding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Onboarding) ProtoMessage() {}

func (x *Onboarding) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Onboarding.ProtoReflect.Descriptor instead.
func (*Onboarding) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{0}
}

func (x *Onboarding) GetLocalUserName() string {
	if x != nil {
		return x.LocalUserName
	}
	return ""
}

func (x *Onboarding) GetLocalPassword() string {
	if x != nil {
		return x.LocalPassword
	}
	return ""
}

func (x *Onboarding) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Onboarding) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Onboarding) GetPlatformActualName() string {
	if x != nil {
		return x.PlatformActualName
	}
	return ""
}

func (x *Onboarding) GetSoftwarePlatformName() string {
	if x != nil {
		return x.SoftwarePlatformName
	}
	return ""
}

func (x *Onboarding) GetPotalUrl() string {
	if x != nil {
		return x.PotalUrl
	}
	return ""
}

func (x *Onboarding) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Onboarding) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Onboarding) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Onboarding) GetSwPlatformId() string {
	if x != nil {
		return x.SwPlatformId
	}
	return ""
}

func (x *Onboarding) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *Onboarding) GetIsActivationConfirmed() bool {
	if x != nil {
		return x.IsActivationConfirmed
	}
	return false
}

func (x *Onboarding) GetDeviceRole() string {
	if x != nil {
		return x.DeviceRole
	}
	return ""
}

// Security section in IEM Onboarding JSON file.
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network             *NetworkSettings `protobuf:"bytes,1,opt,name=Network,proto3" json:"Network,omitempty"`
	NtpServer           []string         `protobuf:"bytes,2,rep,name=ntpServer,proto3" json:"ntpServer,omitempty"`
	DockerIP            string           `protobuf:"bytes,3,opt,name=dockerIP,proto3" json:"dockerIP,omitempty"`
	CustomConfiguration *anypb.Any       `protobuf:"bytes,4,opt,name=customConfiguration,proto3" json:"customConfiguration,omitempty"` //device builder specific custom configuraiton.
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetNetwork() *NetworkSettings {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *Device) GetNtpServer() []string {
	if x != nil {
		return x.NtpServer
	}
	return nil
}

func (x *Device) GetDockerIP() string {
	if x != nil {
		return x.DockerIP
	}
	return ""
}

func (x *Device) GetCustomConfiguration() *anypb.Any {
	if x != nil {
		return x.CustomConfiguration
	}
	return nil
}

// Security section in IEM Onboarding JSON file.
type Security struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseUrl                 string   `protobuf:"bytes,1,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	ClientCredentialProfile []string `protobuf:"bytes,2,rep,name=clientCredentialProfile,proto3" json:"clientCredentialProfile,omitempty"`
	CaChain                 string   `protobuf:"bytes,3,opt,name=ca_chain,json=caChain,proto3" json:"ca_chain,omitempty"`
}

func (x *Security) Reset() {
	*x = Security{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Security) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Security) ProtoMessage() {}

func (x *Security) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Security.ProtoReflect.Descriptor instead.
func (*Security) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{2}
}

func (x *Security) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *Security) GetClientCredentialProfile() []string {
	if x != nil {
		return x.ClientCredentialProfile
	}
	return nil
}

func (x *Security) GetCaChain() string {
	if x != nil {
		return x.CaChain
	}
	return ""
}

// Agent section in IEM Onboarding JSON file.
type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Proxy    []string  `protobuf:"bytes,2,rep,name=proxy,proto3" json:"proxy,omitempty"`
	AgentId  string    `protobuf:"bytes,3,opt,name=agentId,proto3" json:"agentId,omitempty"`
	Security *Security `protobuf:"bytes,4,opt,name=security,proto3" json:"security,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{3}
}

func (x *Agent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Agent) GetProxy() []string {
	if x != nil {
		return x.Proxy
	}
	return nil
}

func (x *Agent) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Agent) GetSecurity() *Security {
	if x != nil {
		return x.Security
	}
	return nil
}

type ProxySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyId            string  `protobuf:"bytes,1,opt,name=proxyId,proto3" json:"proxyId,omitempty"`
	ProxyType          string  `protobuf:"bytes,2,opt,name=proxyType,proto3" json:"proxyType,omitempty"`
	Host               string  `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Protocol           string  `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	AuthenticationType string  `protobuf:"bytes,5,opt,name=authenticationType,proto3" json:"authenticationType,omitempty"`
	User               string  `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Password           string  `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	NoProxy            string  `protobuf:"bytes,8,opt,name=noProxy,proto3" json:"noProxy,omitempty"`
	CustomPorts        []int32 `protobuf:"varint,9,rep,packed,name=customPorts,proto3" json:"customPorts,omitempty"`
}

func (x *ProxySettings) Reset() {
	*x = ProxySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxySettings) ProtoMessage() {}

func (x *ProxySettings) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxySettings.ProtoReflect.Descriptor instead.
func (*ProxySettings) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{4}
}

func (x *ProxySettings) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *ProxySettings) GetProxyType() string {
	if x != nil {
		return x.ProxyType
	}
	return ""
}

func (x *ProxySettings) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ProxySettings) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ProxySettings) GetAuthenticationType() string {
	if x != nil {
		return x.AuthenticationType
	}
	return ""
}

func (x *ProxySettings) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ProxySettings) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ProxySettings) GetNoProxy() string {
	if x != nil {
		return x.NoProxy
	}
	return ""
}

func (x *ProxySettings) GetCustomPorts() []int32 {
	if x != nil {
		return x.CustomPorts
	}
	return nil
}

// Corresponding proto message for JSON file that had downloaded from IEM. Onboarding JSON file needs to be converted to this type first.
// Compatible with IEM Onboarding JSON schema, since all field names are same.
// For more information please refer to IEM Onboarding JSON scheme.
type DeviceConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device     *Device          `protobuf:"bytes,1,opt,name=Device,proto3" json:"Device,omitempty"`         // Device seciton in IEM Onboarding JSON file.
	Onboarding *Onboarding      `protobuf:"bytes,2,opt,name=onboarding,proto3" json:"onboarding,omitempty"` // Onboarding section in IEM Onboarding JSON file.
	Agents     []*Agent         `protobuf:"bytes,3,rep,name=agents,proto3" json:"agents,omitempty"`         // Agent section in IEM Onboarding JSON file.
	Proxies    []*ProxySettings `protobuf:"bytes,4,rep,name=proxies,proto3" json:"proxies,omitempty"`       //Proxy section in IEM Onboarding JSON file.
}

func (x *DeviceConfiguration) Reset() {
	*x = DeviceConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceConfiguration) ProtoMessage() {}

func (x *DeviceConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceConfiguration.ProtoReflect.Descriptor instead.
func (*DeviceConfiguration) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceConfiguration) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *DeviceConfiguration) GetOnboarding() *Onboarding {
	if x != nil {
		return x.Onboarding
	}
	return nil
}

func (x *DeviceConfiguration) GetAgents() []*Agent {
	if x != nil {
		return x.Agents
	}
	return nil
}

func (x *DeviceConfiguration) GetProxies() []*ProxySettings {
	if x != nil {
		return x.Proxies
	}
	return nil
}

// Type indicates the Onboarding Status.
type OnboardingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOnboarded bool   `protobuf:"varint,1,opt,name=isOnboarded,proto3" json:"isOnboarded,omitempty"` // true if onboarding is successful,otherwise false.
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`          // Additonal information for onboarding status.
}

func (x *OnboardingStatus) Reset() {
	*x = OnboardingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Onboard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnboardingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnboardingStatus) ProtoMessage() {}

func (x *OnboardingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Onboard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnboardingStatus.ProtoReflect.Descriptor instead.
func (*OnboardingStatus) Descriptor() ([]byte, []int) {
	return file_Onboard_proto_rawDescGZIP(), []int{6}
}

func (x *OnboardingStatus) GetIsOnboarded() bool {
	if x != nil {
		return x.IsOnboarded
	}
	return false
}

func (x *OnboardingStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_Onboard_proto protoreflect.FileDescriptor

var file_Onboard_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x03, 0x0a, 0x0a, 0x4f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x12, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a,
	0x14, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x6f, 0x66,
	0x74, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x77, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x77, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x15, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x50, 0x12, 0x46, 0x0a, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x13, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x79, 0x0a, 0x08, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x61, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x69, 0x65,
	0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0x93,
	0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x22, 0xa9, 0x02, 0x0a, 0x13, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0a,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65,
	0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65,
	0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73,
	0x22, 0x4e, 0x0a, 0x10, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x95, 0x03, 0x0a, 0x0e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x57, 0x69,
	0x74, 0x68, 0x55, 0x53, 0x42, 0x12, 0x33, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e,
	0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x61, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65,
	0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5f, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x2e, 0x73,
	0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x30, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x69,
	0x65, 0x6d, 0x65, 0x6e, 0x73, 0x5f, 0x69, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x64, 0x6d, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Onboard_proto_rawDescOnce sync.Once
	file_Onboard_proto_rawDescData = file_Onboard_proto_rawDesc
)

func file_Onboard_proto_rawDescGZIP() []byte {
	file_Onboard_proto_rawDescOnce.Do(func() {
		file_Onboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_Onboard_proto_rawDescData)
	})
	return file_Onboard_proto_rawDescData
}

var file_Onboard_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Onboard_proto_goTypes = []interface{}{
	(*Onboarding)(nil),          // 0: siemens.iedge.dmapi.onboard.v1.Onboarding
	(*Device)(nil),              // 1: siemens.iedge.dmapi.onboard.v1.Device
	(*Security)(nil),            // 2: siemens.iedge.dmapi.onboard.v1.Security
	(*Agent)(nil),               // 3: siemens.iedge.dmapi.onboard.v1.Agent
	(*ProxySettings)(nil),       // 4: siemens.iedge.dmapi.onboard.v1.ProxySettings
	(*DeviceConfiguration)(nil), // 5: siemens.iedge.dmapi.onboard.v1.DeviceConfiguration
	(*OnboardingStatus)(nil),    // 6: siemens.iedge.dmapi.onboard.v1.OnboardingStatus
	(*NetworkSettings)(nil),     // 7: siemens.iedge.dmapi.network.v1.NetworkSettings
	(*anypb.Any)(nil),           // 8: google.protobuf.Any
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_Onboard_proto_depIdxs = []int32{
	7,  // 0: siemens.iedge.dmapi.onboard.v1.Device.Network:type_name -> siemens.iedge.dmapi.network.v1.NetworkSettings
	8,  // 1: siemens.iedge.dmapi.onboard.v1.Device.customConfiguration:type_name -> google.protobuf.Any
	2,  // 2: siemens.iedge.dmapi.onboard.v1.Agent.security:type_name -> siemens.iedge.dmapi.onboard.v1.Security
	1,  // 3: siemens.iedge.dmapi.onboard.v1.DeviceConfiguration.Device:type_name -> siemens.iedge.dmapi.onboard.v1.Device
	0,  // 4: siemens.iedge.dmapi.onboard.v1.DeviceConfiguration.onboarding:type_name -> siemens.iedge.dmapi.onboard.v1.Onboarding
	3,  // 5: siemens.iedge.dmapi.onboard.v1.DeviceConfiguration.agents:type_name -> siemens.iedge.dmapi.onboard.v1.Agent
	4,  // 6: siemens.iedge.dmapi.onboard.v1.DeviceConfiguration.proxies:type_name -> siemens.iedge.dmapi.onboard.v1.ProxySettings
	5,  // 7: siemens.iedge.dmapi.onboard.v1.OnboardService.OnboardWithUSB:input_type -> siemens.iedge.dmapi.onboard.v1.DeviceConfiguration
	5,  // 8: siemens.iedge.dmapi.onboard.v1.OnboardService.ApplyConfiguration:input_type -> siemens.iedge.dmapi.onboard.v1.DeviceConfiguration
	6,  // 9: siemens.iedge.dmapi.onboard.v1.OnboardService.SetOnboardingStatus:input_type -> siemens.iedge.dmapi.onboard.v1.OnboardingStatus
	9,  // 10: siemens.iedge.dmapi.onboard.v1.OnboardService.ListenOnboardState:input_type -> google.protobuf.Empty
	9,  // 11: siemens.iedge.dmapi.onboard.v1.OnboardService.OnboardWithUSB:output_type -> google.protobuf.Empty
	9,  // 12: siemens.iedge.dmapi.onboard.v1.OnboardService.ApplyConfiguration:output_type -> google.protobuf.Empty
	9,  // 13: siemens.iedge.dmapi.onboard.v1.OnboardService.SetOnboardingStatus:output_type -> google.protobuf.Empty
	6,  // 14: siemens.iedge.dmapi.onboard.v1.OnboardService.ListenOnboardState:output_type -> siemens.iedge.dmapi.onboard.v1.OnboardingStatus
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_Onboard_proto_init() }
func file_Onboard_proto_init() {
	if File_Onboard_proto != nil {
		return
	}
	file_Network_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Onboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Onboarding); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Onboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Onboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Security); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Onboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Onboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxySettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Onboard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Onboard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnboardingStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Onboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Onboard_proto_goTypes,
		DependencyIndexes: file_Onboard_proto_depIdxs,
		MessageInfos:      file_Onboard_proto_msgTypes,
	}.Build()
	File_Onboard_proto = out.File
	file_Onboard_proto_rawDesc = nil
	file_Onboard_proto_goTypes = nil
	file_Onboard_proto_depIdxs = nil
}
