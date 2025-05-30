/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

//
// Copyright (c) 2023 Siemens AG
// Licensed under the MIT license
// See LICENSE file in the top-level directory

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: Network.proto

package siemens_iedge_dmapi_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Contains MAC address, used for retrieving specified Network Interface settings.
type NetworkInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac string `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *NetworkInterfaceRequest) Reset() {
	*x = NetworkInterfaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceRequest) ProtoMessage() {}

func (x *NetworkInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceRequest.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkInterfaceRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

// Returns an Network Interface with
type NetworkInterfaceRequestWithLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *NetworkInterfaceRequestWithLabel) Reset() {
	*x = NetworkInterfaceRequestWithLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterfaceRequestWithLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceRequestWithLabel) ProtoMessage() {}

func (x *NetworkInterfaceRequestWithLabel) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceRequestWithLabel.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceRequestWithLabel) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkInterfaceRequestWithLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Interface type holds settings for a Network Interface.
type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayInterface bool                  `protobuf:"varint,1,opt,name=GatewayInterface,proto3" json:"GatewayInterface,omitempty"` // if true, route metric will be set to 1. Otherwise route metric is -1. Similarly, when the interface is requested,return value will be true if route metric is 1.
	MacAddress       string                `protobuf:"bytes,2,opt,name=MacAddress,proto3" json:"MacAddress,omitempty"`              // "20:87:56:b5:ed:e0"
	DHCP             string                `protobuf:"bytes,3,opt,name=DHCP,proto3" json:"DHCP,omitempty"`                          // values can be 'enabled' or 'disabled'. for compatiblity reasons it is not boolean.
	Static           *Interface_StaticConf `protobuf:"bytes,4,opt,name=Static,proto3" json:"Static,omitempty"`                      // Static field is StaticConf type instance.
	DNSConfig        *Interface_Dns        `protobuf:"bytes,5,opt,name=DNSConfig,proto3" json:"DNSConfig,omitempty"`                // DNSConfig is dns type instance.
	L2Conf           *Interface_L2         `protobuf:"bytes,6,opt,name=L2Conf,proto3" json:"L2Conf,omitempty"`
	InterfaceName    string                `protobuf:"bytes,7,opt,name=InterfaceName,proto3" json:"InterfaceName,omitempty"` // ens2p
	Label            string                `protobuf:"bytes,8,opt,name=Label,proto3" json:"Label,omitempty"`                 // x1
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{2}
}

func (x *Interface) GetGatewayInterface() bool {
	if x != nil {
		return x.GatewayInterface
	}
	return false
}

func (x *Interface) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *Interface) GetDHCP() string {
	if x != nil {
		return x.DHCP
	}
	return ""
}

func (x *Interface) GetStatic() *Interface_StaticConf {
	if x != nil {
		return x.Static
	}
	return nil
}

func (x *Interface) GetDNSConfig() *Interface_Dns {
	if x != nil {
		return x.DNSConfig
	}
	return nil
}

func (x *Interface) GetL2Conf() *Interface_L2 {
	if x != nil {
		return x.L2Conf
	}
	return nil
}

func (x *Interface) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *Interface) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Contains multiple network interface settings. It can be used to apply or get the settings.
type NetworkSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interfaces []*Interface      `protobuf:"bytes,1,rep,name=Interfaces,proto3" json:"Interfaces,omitempty"`                                                                                     // Network settings contains an array of Interfaces.Applying new settings or receiving current settings is supported for multiple ethernet typed network interfaces supported.
	LabelMap   map[string]string `protobuf:"bytes,2,rep,name=LabelMap,proto3" json:"LabelMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // LabelMap contains port label and corresponding interface-name. e.g key : x1 value: enp2s0
}

func (x *NetworkSettings) Reset() {
	*x = NetworkSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkSettings) ProtoMessage() {}

func (x *NetworkSettings) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkSettings.ProtoReflect.Descriptor instead.
func (*NetworkSettings) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkSettings) GetInterfaces() []*Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

func (x *NetworkSettings) GetLabelMap() map[string]string {
	if x != nil {
		return x.LabelMap
	}
	return nil
}

// StaticConf type holds IP Netmask and Gateway information
type Interface_StaticConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPv4    string `protobuf:"bytes,1,opt,name=IPv4,proto3" json:"IPv4,omitempty"`       // e.g: 192.168.0.2
	NetMask string `protobuf:"bytes,2,opt,name=NetMask,proto3" json:"NetMask,omitempty"` // e.g: 255.255.255.0
	Gateway string `protobuf:"bytes,3,opt,name=Gateway,proto3" json:"Gateway,omitempty"` // e.g: 192.168.0.1
}

func (x *Interface_StaticConf) Reset() {
	*x = Interface_StaticConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface_StaticConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface_StaticConf) ProtoMessage() {}

func (x *Interface_StaticConf) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface_StaticConf.ProtoReflect.Descriptor instead.
func (*Interface_StaticConf) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Interface_StaticConf) GetIPv4() string {
	if x != nil {
		return x.IPv4
	}
	return ""
}

func (x *Interface_StaticConf) GetNetMask() string {
	if x != nil {
		return x.NetMask
	}
	return ""
}

func (x *Interface_StaticConf) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

// Type that contains Primary and Secondary DNS.
type Interface_Dns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryDNS   string `protobuf:"bytes,1,opt,name=PrimaryDNS,proto3" json:"PrimaryDNS,omitempty"`     // e.g:  "1.1.1.2"
	SecondaryDNS string `protobuf:"bytes,2,opt,name=SecondaryDNS,proto3" json:"SecondaryDNS,omitempty"` // e.g: "1.1.1.1"
}

func (x *Interface_Dns) Reset() {
	*x = Interface_Dns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface_Dns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface_Dns) ProtoMessage() {}

func (x *Interface_Dns) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface_Dns.ProtoReflect.Descriptor instead.
func (*Interface_Dns) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Interface_Dns) GetPrimaryDNS() string {
	if x != nil {
		return x.PrimaryDNS
	}
	return ""
}

func (x *Interface_Dns) GetSecondaryDNS() string {
	if x != nil {
		return x.SecondaryDNS
	}
	return ""
}

type Interface_L2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartingAddressIPv4 string            `protobuf:"bytes,1,opt,name=StartingAddressIPv4,proto3" json:"StartingAddressIPv4,omitempty"`                                                                                       // e.g: 192.168.0.2
	NetMask             string            `protobuf:"bytes,2,opt,name=NetMask,proto3" json:"NetMask,omitempty"`                                                                                                               // e.g: 255.255.255.0
	Range               string            `protobuf:"bytes,3,opt,name=Range,proto3" json:"Range,omitempty"`                                                                                                                   // e.g: 16
	Gateway             string            `protobuf:"bytes,4,opt,name=Gateway,proto3" json:"Gateway,omitempty"`                                                                                                               //e.g: 192.168.2.1
	AuxiliaryAddresses  map[string]string `protobuf:"bytes,5,rep,name=AuxiliaryAddresses,proto3" json:"AuxiliaryAddresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //Preserved addresses for other devices. These addresses won't be assigned to containers. e.g  my_plc,  192.168.0.5
}

func (x *Interface_L2) Reset() {
	*x = Interface_L2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface_L2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface_L2) ProtoMessage() {}

func (x *Interface_L2) ProtoReflect() protoreflect.Message {
	mi := &file_Network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface_L2.ProtoReflect.Descriptor instead.
func (*Interface_L2) Descriptor() ([]byte, []int) {
	return file_Network_proto_rawDescGZIP(), []int{2, 2}
}

func (x *Interface_L2) GetStartingAddressIPv4() string {
	if x != nil {
		return x.StartingAddressIPv4
	}
	return ""
}

func (x *Interface_L2) GetNetMask() string {
	if x != nil {
		return x.NetMask
	}
	return ""
}

func (x *Interface_L2) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

func (x *Interface_L2) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Interface_L2) GetAuxiliaryAddresses() map[string]string {
	if x != nil {
		return x.AuxiliaryAddresses
	}
	return nil
}

var File_Network_proto protoreflect.FileDescriptor

var file_Network_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x17,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x22, 0x38, 0x0a, 0x20, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0xe9, 0x06, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x48, 0x43, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x48, 0x43,
	0x50, 0x12, 0x4c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12,
	0x4b, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x6e,
	0x73, 0x52, 0x09, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x06,
	0x4c, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73,
	0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x32, 0x52, 0x06, 0x4c, 0x32, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x54,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04,
	0x49, 0x50, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x50, 0x76, 0x34,
	0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x1a, 0x49, 0x0a, 0x03, 0x44, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x4e, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x4e, 0x53, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x44, 0x4e, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x44, 0x4e, 0x53, 0x1a,
	0xbd, 0x02, 0x0a, 0x02, 0x4c, 0x32, 0x12, 0x30, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x34, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x50, 0x76, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x4d,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x74, 0x0a, 0x12, 0x41, 0x75, 0x78, 0x69, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44,
	0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x32, 0x2e, 0x41, 0x75, 0x78,
	0x69, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x41, 0x75, 0x78, 0x69, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x45, 0x0a, 0x17, 0x41, 0x75, 0x78, 0x69,
	0x6c, 0x69, 0x61, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xf4, 0x01, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e,
	0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x59,
	0x0a, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x1a, 0x3b, 0x0a, 0x0d, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xc9, 0x03, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e,
	0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x79, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x63, 0x12, 0x37, 0x2e,
	0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73,
	0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x84, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x40, 0x2e, 0x73, 0x69,
	0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x29, 0x2e,
	0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2f, 0x2e, 0x73, 0x69, 0x65, 0x6d,
	0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x5f,
	0x69, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Network_proto_rawDescOnce sync.Once
	file_Network_proto_rawDescData = file_Network_proto_rawDesc
)

func file_Network_proto_rawDescGZIP() []byte {
	file_Network_proto_rawDescOnce.Do(func() {
		file_Network_proto_rawDescData = protoimpl.X.CompressGZIP(file_Network_proto_rawDescData)
	})
	return file_Network_proto_rawDescData
}

var file_Network_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Network_proto_goTypes = []interface{}{
	(*NetworkInterfaceRequest)(nil),          // 0: siemens.iedge.dmapi.network.v1.NetworkInterfaceRequest
	(*NetworkInterfaceRequestWithLabel)(nil), // 1: siemens.iedge.dmapi.network.v1.NetworkInterfaceRequestWithLabel
	(*Interface)(nil),                        // 2: siemens.iedge.dmapi.network.v1.Interface
	(*NetworkSettings)(nil),                  // 3: siemens.iedge.dmapi.network.v1.NetworkSettings
	(*Interface_StaticConf)(nil),             // 4: siemens.iedge.dmapi.network.v1.Interface.StaticConf
	(*Interface_Dns)(nil),                    // 5: siemens.iedge.dmapi.network.v1.Interface.Dns
	(*Interface_L2)(nil),                     // 6: siemens.iedge.dmapi.network.v1.Interface.L2
	nil,                                      // 7: siemens.iedge.dmapi.network.v1.Interface.L2.AuxiliaryAddressesEntry
	nil,                                      // 8: siemens.iedge.dmapi.network.v1.NetworkSettings.LabelMapEntry
	(*emptypb.Empty)(nil),                    // 9: google.protobuf.Empty
}
var file_Network_proto_depIdxs = []int32{
	4,  // 0: siemens.iedge.dmapi.network.v1.Interface.Static:type_name -> siemens.iedge.dmapi.network.v1.Interface.StaticConf
	5,  // 1: siemens.iedge.dmapi.network.v1.Interface.DNSConfig:type_name -> siemens.iedge.dmapi.network.v1.Interface.Dns
	6,  // 2: siemens.iedge.dmapi.network.v1.Interface.L2Conf:type_name -> siemens.iedge.dmapi.network.v1.Interface.L2
	2,  // 3: siemens.iedge.dmapi.network.v1.NetworkSettings.Interfaces:type_name -> siemens.iedge.dmapi.network.v1.Interface
	8,  // 4: siemens.iedge.dmapi.network.v1.NetworkSettings.LabelMap:type_name -> siemens.iedge.dmapi.network.v1.NetworkSettings.LabelMapEntry
	7,  // 5: siemens.iedge.dmapi.network.v1.Interface.L2.AuxiliaryAddresses:type_name -> siemens.iedge.dmapi.network.v1.Interface.L2.AuxiliaryAddressesEntry
	9,  // 6: siemens.iedge.dmapi.network.v1.NetworkService.GetAllInterfaces:input_type -> google.protobuf.Empty
	0,  // 7: siemens.iedge.dmapi.network.v1.NetworkService.GetInterfaceWithMac:input_type -> siemens.iedge.dmapi.network.v1.NetworkInterfaceRequest
	1,  // 8: siemens.iedge.dmapi.network.v1.NetworkService.GetInterfaceWithLabel:input_type -> siemens.iedge.dmapi.network.v1.NetworkInterfaceRequestWithLabel
	3,  // 9: siemens.iedge.dmapi.network.v1.NetworkService.ApplySettings:input_type -> siemens.iedge.dmapi.network.v1.NetworkSettings
	3,  // 10: siemens.iedge.dmapi.network.v1.NetworkService.GetAllInterfaces:output_type -> siemens.iedge.dmapi.network.v1.NetworkSettings
	2,  // 11: siemens.iedge.dmapi.network.v1.NetworkService.GetInterfaceWithMac:output_type -> siemens.iedge.dmapi.network.v1.Interface
	2,  // 12: siemens.iedge.dmapi.network.v1.NetworkService.GetInterfaceWithLabel:output_type -> siemens.iedge.dmapi.network.v1.Interface
	9,  // 13: siemens.iedge.dmapi.network.v1.NetworkService.ApplySettings:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_Network_proto_init() }
func file_Network_proto_init() {
	if File_Network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterfaceRequest); i {
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
		file_Network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterfaceRequestWithLabel); i {
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
		file_Network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
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
		file_Network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkSettings); i {
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
		file_Network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface_StaticConf); i {
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
		file_Network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface_Dns); i {
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
		file_Network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface_L2); i {
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
			RawDescriptor: file_Network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Network_proto_goTypes,
		DependencyIndexes: file_Network_proto_depIdxs,
		MessageInfos:      file_Network_proto_msgTypes,
	}.Build()
	File_Network_proto = out.File
	file_Network_proto_rawDesc = nil
	file_Network_proto_goTypes = nil
	file_Network_proto_depIdxs = nil
}
