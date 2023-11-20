//
// Copyright (c) 2023 Siemens AG
// Licensed under the MIT license
// See LICENSE file in the top-level directory

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: Ntp.proto

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

// Type contains an array of ntp server addresses.
type Ntp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NtpServer []string `protobuf:"bytes,1,rep,name=ntpServer,proto3" json:"ntpServer,omitempty"` // array of multiple ntp server address.
}

func (x *Ntp) Reset() {
	*x = Ntp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ntp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ntp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ntp) ProtoMessage() {}

func (x *Ntp) ProtoReflect() protoreflect.Message {
	mi := &file_Ntp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ntp.ProtoReflect.Descriptor instead.
func (*Ntp) Descriptor() ([]byte, []int) {
	return file_Ntp_proto_rawDescGZIP(), []int{0}
}

func (x *Ntp) GetNtpServer() []string {
	if x != nil {
		return x.NtpServer
	}
	return nil
}

// Peer Details from ntpq -p output
type PeerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteServer string  `protobuf:"bytes,1,opt,name=remoteServer,proto3" json:"remoteServer,omitempty"` // NTP server address
	ReferenceID  string  `protobuf:"bytes,2,opt,name=referenceID,proto3" json:"referenceID,omitempty"`   // Reference id for the NTP server
	Stratum      string  `protobuf:"bytes,3,opt,name=stratum,proto3" json:"stratum,omitempty"`           // Stratum for the NTP Server
	Type         string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`                 // Type of server (local, unicast, multicast, or broadcast)
	Poll         int32   `protobuf:"varint,5,opt,name=poll,proto3" json:"poll,omitempty"`                // How frequently to query server (in seconds)
	When         int32   `protobuf:"varint,6,opt,name=when,proto3" json:"when,omitempty"`                // How many seconds passed after the last poll.
	Reach        string  `protobuf:"bytes,7,opt,name=reach,proto3" json:"reach,omitempty"`               // octal bitmask of success or failure of last 8 queries (left-shifted). eg:375
	Delay        float32 `protobuf:"fixed32,8,opt,name=delay,proto3" json:"delay,omitempty"`             // network round trip time (in milliseconds)
	Offset       float32 `protobuf:"fixed32,9,opt,name=offset,proto3" json:"offset,omitempty"`           // difference between local clock and remote clock (in milliseconds)
	Jitter       float32 `protobuf:"fixed32,10,opt,name=jitter,proto3" json:"jitter,omitempty"`          // Difference of successive time values from server (in milliseconds)
}

func (x *PeerDetails) Reset() {
	*x = PeerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ntp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerDetails) ProtoMessage() {}

func (x *PeerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_Ntp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerDetails.ProtoReflect.Descriptor instead.
func (*PeerDetails) Descriptor() ([]byte, []int) {
	return file_Ntp_proto_rawDescGZIP(), []int{1}
}

func (x *PeerDetails) GetRemoteServer() string {
	if x != nil {
		return x.RemoteServer
	}
	return ""
}

func (x *PeerDetails) GetReferenceID() string {
	if x != nil {
		return x.ReferenceID
	}
	return ""
}

func (x *PeerDetails) GetStratum() string {
	if x != nil {
		return x.Stratum
	}
	return ""
}

func (x *PeerDetails) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PeerDetails) GetPoll() int32 {
	if x != nil {
		return x.Poll
	}
	return 0
}

func (x *PeerDetails) GetWhen() int32 {
	if x != nil {
		return x.When
	}
	return 0
}

func (x *PeerDetails) GetReach() string {
	if x != nil {
		return x.Reach
	}
	return ""
}

func (x *PeerDetails) GetDelay() float32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *PeerDetails) GetOffset() float32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PeerDetails) GetJitter() float32 {
	if x != nil {
		return x.Jitter
	}
	return 0
}

// Type for ntp current sync status
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNtpServiceRunning   bool           `protobuf:"varint,1,opt,name=isNtpServiceRunning,proto3" json:"isNtpServiceRunning,omitempty"`    // indicates that ntp service is running or not
	IsSynced              bool           `protobuf:"varint,2,opt,name=isSynced,proto3" json:"isSynced,omitempty"`                          // indicates NTP server synced or not
	LastConfigurationTime string         `protobuf:"bytes,3,opt,name=lastConfigurationTime,proto3" json:"lastConfigurationTime,omitempty"` // time of the last performed iedk ntp configuration.
	LastSyncTime          string         `protobuf:"bytes,4,opt,name=lastSyncTime,proto3" json:"lastSyncTime,omitempty"`                   // time of the last ntp sync operation.
	PeerDetails           []*PeerDetails `protobuf:"bytes,5,rep,name=peerDetails,proto3" json:"peerDetails,omitempty"`                     // NTPQ peer information array. Only exist after ntp configuration done.
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ntp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_Ntp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_Ntp_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetIsNtpServiceRunning() bool {
	if x != nil {
		return x.IsNtpServiceRunning
	}
	return false
}

func (x *Status) GetIsSynced() bool {
	if x != nil {
		return x.IsSynced
	}
	return false
}

func (x *Status) GetLastConfigurationTime() string {
	if x != nil {
		return x.LastConfigurationTime
	}
	return ""
}

func (x *Status) GetLastSyncTime() string {
	if x != nil {
		return x.LastSyncTime
	}
	return ""
}

func (x *Status) GetPeerDetails() []*PeerDetails {
	if x != nil {
		return x.PeerDetails
	}
	return nil
}

var File_Ntp_proto protoreflect.FileDescriptor

var file_Ntp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x4e, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x69, 0x65,
	0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x03, 0x4e, 0x74, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x85, 0x02, 0x0a, 0x0b, 0x50, 0x65,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x61, 0x74, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x72, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x6c,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x22, 0xfb, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x13,
	0x69, 0x73, 0x4e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x73, 0x4e, 0x74, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x6c, 0x61,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x69, 0x65, 0x6d,
	0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32,
	0xe7, 0x01, 0x0a, 0x0a, 0x4e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64,
	0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x74, 0x70, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x74,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1f, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x69, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x74, 0x70,
	0x12, 0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x73, 0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x2e,
	0x69, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x64, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x74, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73,
	0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x5f, 0x69, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x64, 0x6d, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Ntp_proto_rawDescOnce sync.Once
	file_Ntp_proto_rawDescData = file_Ntp_proto_rawDesc
)

func file_Ntp_proto_rawDescGZIP() []byte {
	file_Ntp_proto_rawDescOnce.Do(func() {
		file_Ntp_proto_rawDescData = protoimpl.X.CompressGZIP(file_Ntp_proto_rawDescData)
	})
	return file_Ntp_proto_rawDescData
}

var file_Ntp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Ntp_proto_goTypes = []interface{}{
	(*Ntp)(nil),           // 0: siemens.iedge.dmapi.ntp.v1.Ntp
	(*PeerDetails)(nil),   // 1: siemens.iedge.dmapi.ntp.v1.PeerDetails
	(*Status)(nil),        // 2: siemens.iedge.dmapi.ntp.v1.Status
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_Ntp_proto_depIdxs = []int32{
	1, // 0: siemens.iedge.dmapi.ntp.v1.Status.peerDetails:type_name -> siemens.iedge.dmapi.ntp.v1.PeerDetails
	0, // 1: siemens.iedge.dmapi.ntp.v1.NtpService.SetNtpServer:input_type -> siemens.iedge.dmapi.ntp.v1.Ntp
	3, // 2: siemens.iedge.dmapi.ntp.v1.NtpService.GetNtpServer:input_type -> google.protobuf.Empty
	3, // 3: siemens.iedge.dmapi.ntp.v1.NtpService.GetStatus:input_type -> google.protobuf.Empty
	3, // 4: siemens.iedge.dmapi.ntp.v1.NtpService.SetNtpServer:output_type -> google.protobuf.Empty
	0, // 5: siemens.iedge.dmapi.ntp.v1.NtpService.GetNtpServer:output_type -> siemens.iedge.dmapi.ntp.v1.Ntp
	2, // 6: siemens.iedge.dmapi.ntp.v1.NtpService.GetStatus:output_type -> siemens.iedge.dmapi.ntp.v1.Status
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Ntp_proto_init() }
func file_Ntp_proto_init() {
	if File_Ntp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Ntp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ntp); i {
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
		file_Ntp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerDetails); i {
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
		file_Ntp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_Ntp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Ntp_proto_goTypes,
		DependencyIndexes: file_Ntp_proto_depIdxs,
		MessageInfos:      file_Ntp_proto_msgTypes,
	}.Build()
	File_Ntp_proto = out.File
	file_Ntp_proto_rawDesc = nil
	file_Ntp_proto_goTypes = nil
	file_Ntp_proto_depIdxs = nil
}
