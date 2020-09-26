// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/ship_repair.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StartShipRepairRequest_RapairTypeEnum int32

const (
	StartShipRepairRequest_PARTIAL StartShipRepairRequest_RapairTypeEnum = 0
	StartShipRepairRequest_FULL    StartShipRepairRequest_RapairTypeEnum = 1
)

// Enum value maps for StartShipRepairRequest_RapairTypeEnum.
var (
	StartShipRepairRequest_RapairTypeEnum_name = map[int32]string{
		0: "PARTIAL",
		1: "FULL",
	}
	StartShipRepairRequest_RapairTypeEnum_value = map[string]int32{
		"PARTIAL": 0,
		"FULL":    1,
	}
)

func (x StartShipRepairRequest_RapairTypeEnum) Enum() *StartShipRepairRequest_RapairTypeEnum {
	p := new(StartShipRepairRequest_RapairTypeEnum)
	*p = x
	return p
}

func (x StartShipRepairRequest_RapairTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartShipRepairRequest_RapairTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ship_repair_proto_enumTypes[0].Descriptor()
}

func (StartShipRepairRequest_RapairTypeEnum) Type() protoreflect.EnumType {
	return &file_proto_ship_repair_proto_enumTypes[0]
}

func (x StartShipRepairRequest_RapairTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartShipRepairRequest_RapairTypeEnum.Descriptor instead.
func (StartShipRepairRequest_RapairTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{2, 0}
}

// GetShipRepairInfo
type GetShipRepairInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipID uint32 `protobuf:"varint,1,opt,name=ShipID,proto3" json:"ShipID,omitempty"`
}

func (x *GetShipRepairInfoRequest) Reset() {
	*x = GetShipRepairInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipRepairInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipRepairInfoRequest) ProtoMessage() {}

func (x *GetShipRepairInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipRepairInfoRequest.ProtoReflect.Descriptor instead.
func (*GetShipRepairInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{0}
}

func (x *GetShipRepairInfoRequest) GetShipID() uint32 {
	if x != nil {
		return x.ShipID
	}
	return 0
}

type GetShipRepairInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedRepairs bool                                      `protobuf:"varint,1,opt,name=NeedRepairs,proto3" json:"NeedRepairs,omitempty"`
	Partial     *GetShipRepairInfoResponse_ShipRepairInfo `protobuf:"bytes,2,opt,name=Partial,proto3" json:"Partial,omitempty"`
	Full        *GetShipRepairInfoResponse_ShipRepairInfo `protobuf:"bytes,3,opt,name=Full,proto3" json:"Full,omitempty"`
}

func (x *GetShipRepairInfoResponse) Reset() {
	*x = GetShipRepairInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipRepairInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipRepairInfoResponse) ProtoMessage() {}

func (x *GetShipRepairInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipRepairInfoResponse.ProtoReflect.Descriptor instead.
func (*GetShipRepairInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{1}
}

func (x *GetShipRepairInfoResponse) GetNeedRepairs() bool {
	if x != nil {
		return x.NeedRepairs
	}
	return false
}

func (x *GetShipRepairInfoResponse) GetPartial() *GetShipRepairInfoResponse_ShipRepairInfo {
	if x != nil {
		return x.Partial
	}
	return nil
}

func (x *GetShipRepairInfoResponse) GetFull() *GetShipRepairInfoResponse_ShipRepairInfo {
	if x != nil {
		return x.Full
	}
	return nil
}

// StartShipRepair
type StartShipRepairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   uint32                                `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	RapairType StartShipRepairRequest_RapairTypeEnum `protobuf:"varint,2,opt,name=RapairType,proto3,enum=ship_repair.StartShipRepairRequest_RapairTypeEnum" json:"RapairType,omitempty"`
}

func (x *StartShipRepairRequest) Reset() {
	*x = StartShipRepairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartShipRepairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShipRepairRequest) ProtoMessage() {}

func (x *StartShipRepairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShipRepairRequest.ProtoReflect.Descriptor instead.
func (*StartShipRepairRequest) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{2}
}

func (x *StartShipRepairRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *StartShipRepairRequest) GetRapairType() StartShipRepairRequest_RapairTypeEnum {
	if x != nil {
		return x.RapairType
	}
	return StartShipRepairRequest_PARTIAL
}

type StartShipRepairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepairingEndTime *timestamp.Timestamp                       `protobuf:"bytes,1,opt,name=RepairingEndTime,proto3" json:"RepairingEndTime,omitempty"`
	StartShipRepair  []*StartShipRepairResponse_RepairResources `protobuf:"bytes,2,rep,name=StartShipRepair,proto3" json:"StartShipRepair,omitempty"`
}

func (x *StartShipRepairResponse) Reset() {
	*x = StartShipRepairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartShipRepairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShipRepairResponse) ProtoMessage() {}

func (x *StartShipRepairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShipRepairResponse.ProtoReflect.Descriptor instead.
func (*StartShipRepairResponse) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{3}
}

func (x *StartShipRepairResponse) GetRepairingEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.RepairingEndTime
	}
	return nil
}

func (x *StartShipRepairResponse) GetStartShipRepair() []*StartShipRepairResponse_RepairResources {
	if x != nil {
		return x.StartShipRepair
	}
	return nil
}

// CheckShipRepair
type CheckShipRepairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *CheckShipRepairRequest) Reset() {
	*x = CheckShipRepairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckShipRepairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckShipRepairRequest) ProtoMessage() {}

func (x *CheckShipRepairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckShipRepairRequest.ProtoReflect.Descriptor instead.
func (*CheckShipRepairRequest) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{4}
}

func (x *CheckShipRepairRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type CheckShipRepairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepairInProgress bool                 `protobuf:"varint,1,opt,name=RepairInProgress,proto3" json:"RepairInProgress,omitempty"`
	FinishRepairing  bool                 `protobuf:"varint,2,opt,name=FinishRepairing,proto3" json:"FinishRepairing,omitempty"`
	RepairingEndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=RepairingEndTime,proto3" json:"RepairingEndTime,omitempty"`
}

func (x *CheckShipRepairResponse) Reset() {
	*x = CheckShipRepairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckShipRepairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckShipRepairResponse) ProtoMessage() {}

func (x *CheckShipRepairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckShipRepairResponse.ProtoReflect.Descriptor instead.
func (*CheckShipRepairResponse) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{5}
}

func (x *CheckShipRepairResponse) GetRepairInProgress() bool {
	if x != nil {
		return x.RepairInProgress
	}
	return false
}

func (x *CheckShipRepairResponse) GetFinishRepairing() bool {
	if x != nil {
		return x.FinishRepairing
	}
	return false
}

func (x *CheckShipRepairResponse) GetRepairingEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.RepairingEndTime
	}
	return nil
}

// EndShipRepair
type EndShipRepairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *EndShipRepairRequest) Reset() {
	*x = EndShipRepairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndShipRepairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndShipRepairRequest) ProtoMessage() {}

func (x *EndShipRepairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndShipRepairRequest.ProtoReflect.Descriptor instead.
func (*EndShipRepairRequest) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{6}
}

func (x *EndShipRepairRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type EndShipRepairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndShipRepairResponse) Reset() {
	*x = EndShipRepairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndShipRepairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndShipRepairResponse) ProtoMessage() {}

func (x *EndShipRepairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndShipRepairResponse.ProtoReflect.Descriptor instead.
func (*EndShipRepairResponse) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{7}
}

type GetShipRepairInfoResponse_ShipRepairInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepairTime        int32  `protobuf:"varint,1,opt,name=RepairTime,proto3" json:"RepairTime,omitempty"`
	QuantityResources int32  `protobuf:"varint,2,opt,name=QuantityResources,proto3" json:"QuantityResources,omitempty"`
	Integrity         uint32 `protobuf:"varint,3,opt,name=Integrity,proto3" json:"Integrity,omitempty"`
}

func (x *GetShipRepairInfoResponse_ShipRepairInfo) Reset() {
	*x = GetShipRepairInfoResponse_ShipRepairInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipRepairInfoResponse_ShipRepairInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipRepairInfoResponse_ShipRepairInfo) ProtoMessage() {}

func (x *GetShipRepairInfoResponse_ShipRepairInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipRepairInfoResponse_ShipRepairInfo.ProtoReflect.Descriptor instead.
func (*GetShipRepairInfoResponse_ShipRepairInfo) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetShipRepairInfoResponse_ShipRepairInfo) GetRepairTime() int32 {
	if x != nil {
		return x.RepairTime
	}
	return 0
}

func (x *GetShipRepairInfoResponse_ShipRepairInfo) GetQuantityResources() int32 {
	if x != nil {
		return x.QuantityResources
	}
	return 0
}

func (x *GetShipRepairInfoResponse_ShipRepairInfo) GetIntegrity() uint32 {
	if x != nil {
		return x.Integrity
	}
	return 0
}

type StartShipRepairResponse_RepairResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceID uint32 `protobuf:"varint,1,opt,name=ResourceID,proto3" json:"ResourceID,omitempty"`
	Quantity   int32  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *StartShipRepairResponse_RepairResources) Reset() {
	*x = StartShipRepairResponse_RepairResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_repair_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartShipRepairResponse_RepairResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShipRepairResponse_RepairResources) ProtoMessage() {}

func (x *StartShipRepairResponse_RepairResources) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_repair_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShipRepairResponse_RepairResources.ProtoReflect.Descriptor instead.
func (*StartShipRepairResponse_RepairResources) Descriptor() ([]byte, []int) {
	return file_proto_ship_repair_proto_rawDescGZIP(), []int{3, 0}
}

func (x *StartShipRepairResponse_RepairResources) GetResourceID() uint32 {
	if x != nil {
		return x.ResourceID
	}
	return 0
}

func (x *StartShipRepairResponse_RepairResources) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_proto_ship_repair_proto protoreflect.FileDescriptor

var file_proto_ship_repair_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x5f,
	0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x68, 0x69, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x68, 0x69, 0x70, 0x49, 0x44, 0x22, 0xd7, 0x02, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4e, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x73, 0x12, 0x4f, 0x0a, 0x07, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x49, 0x0a, 0x04,
	0x46, 0x75, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70,
	0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x46, 0x75, 0x6c, 0x6c, 0x1a, 0x7c, 0x0a, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x52,
	0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52,
	0x65, 0x70, 0x61, 0x69, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x69, 0x74, 0x79, 0x22, 0xb1, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x52, 0x0a, 0x0a,
	0x52, 0x61, 0x70, 0x61, 0x69, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x61, 0x70, 0x61, 0x69, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0a, 0x52, 0x61, 0x70, 0x61, 0x69, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x27, 0x0a, 0x0e, 0x52, 0x61, 0x70, 0x61, 0x69, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x22, 0x90, 0x02, 0x0a, 0x17, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x52, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5e, 0x0a,
	0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x72, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x0f, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x1a, 0x4d, 0x0a,
	0x0f, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x34, 0x0a, 0x16,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x22, 0xb7, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x69, 0x70,
	0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72,
	0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x61, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x52, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x14,
	0x45, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x17, 0x0a, 0x15, 0x45, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x61, 0x69,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ship_repair_proto_rawDescOnce sync.Once
	file_proto_ship_repair_proto_rawDescData = file_proto_ship_repair_proto_rawDesc
)

func file_proto_ship_repair_proto_rawDescGZIP() []byte {
	file_proto_ship_repair_proto_rawDescOnce.Do(func() {
		file_proto_ship_repair_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ship_repair_proto_rawDescData)
	})
	return file_proto_ship_repair_proto_rawDescData
}

var file_proto_ship_repair_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ship_repair_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_ship_repair_proto_goTypes = []interface{}{
	(StartShipRepairRequest_RapairTypeEnum)(0),       // 0: ship_repair.StartShipRepairRequest.RapairTypeEnum
	(*GetShipRepairInfoRequest)(nil),                 // 1: ship_repair.GetShipRepairInfoRequest
	(*GetShipRepairInfoResponse)(nil),                // 2: ship_repair.GetShipRepairInfoResponse
	(*StartShipRepairRequest)(nil),                   // 3: ship_repair.StartShipRepairRequest
	(*StartShipRepairResponse)(nil),                  // 4: ship_repair.StartShipRepairResponse
	(*CheckShipRepairRequest)(nil),                   // 5: ship_repair.CheckShipRepairRequest
	(*CheckShipRepairResponse)(nil),                  // 6: ship_repair.CheckShipRepairResponse
	(*EndShipRepairRequest)(nil),                     // 7: ship_repair.EndShipRepairRequest
	(*EndShipRepairResponse)(nil),                    // 8: ship_repair.EndShipRepairResponse
	(*GetShipRepairInfoResponse_ShipRepairInfo)(nil), // 9: ship_repair.GetShipRepairInfoResponse.ShipRepairInfo
	(*StartShipRepairResponse_RepairResources)(nil),  // 10: ship_repair.StartShipRepairResponse.RepairResources
	(*timestamp.Timestamp)(nil),                      // 11: google.protobuf.Timestamp
}
var file_proto_ship_repair_proto_depIdxs = []int32{
	9,  // 0: ship_repair.GetShipRepairInfoResponse.Partial:type_name -> ship_repair.GetShipRepairInfoResponse.ShipRepairInfo
	9,  // 1: ship_repair.GetShipRepairInfoResponse.Full:type_name -> ship_repair.GetShipRepairInfoResponse.ShipRepairInfo
	0,  // 2: ship_repair.StartShipRepairRequest.RapairType:type_name -> ship_repair.StartShipRepairRequest.RapairTypeEnum
	11, // 3: ship_repair.StartShipRepairResponse.RepairingEndTime:type_name -> google.protobuf.Timestamp
	10, // 4: ship_repair.StartShipRepairResponse.StartShipRepair:type_name -> ship_repair.StartShipRepairResponse.RepairResources
	11, // 5: ship_repair.CheckShipRepairResponse.RepairingEndTime:type_name -> google.protobuf.Timestamp
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_ship_repair_proto_init() }
func file_proto_ship_repair_proto_init() {
	if File_proto_ship_repair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ship_repair_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShipRepairInfoRequest); i {
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
		file_proto_ship_repair_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShipRepairInfoResponse); i {
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
		file_proto_ship_repair_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartShipRepairRequest); i {
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
		file_proto_ship_repair_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartShipRepairResponse); i {
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
		file_proto_ship_repair_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckShipRepairRequest); i {
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
		file_proto_ship_repair_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckShipRepairResponse); i {
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
		file_proto_ship_repair_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndShipRepairRequest); i {
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
		file_proto_ship_repair_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndShipRepairResponse); i {
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
		file_proto_ship_repair_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShipRepairInfoResponse_ShipRepairInfo); i {
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
		file_proto_ship_repair_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartShipRepairResponse_RepairResources); i {
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
			RawDescriptor: file_proto_ship_repair_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ship_repair_proto_goTypes,
		DependencyIndexes: file_proto_ship_repair_proto_depIdxs,
		EnumInfos:         file_proto_ship_repair_proto_enumTypes,
		MessageInfos:      file_proto_ship_repair_proto_msgTypes,
	}.Build()
	File_proto_ship_repair_proto = out.File
	file_proto_ship_repair_proto_rawDesc = nil
	file_proto_ship_repair_proto_goTypes = nil
	file_proto_ship_repair_proto_depIdxs = nil
}
