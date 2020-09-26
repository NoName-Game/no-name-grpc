// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/safeplanet_crafter.proto

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

// CrafterStart
type CrafterStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID     uint32           `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Resources    map[uint32]int32 `protobuf:"bytes,2,rep,name=Resources,proto3" json:"Resources,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Price        int32            `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
	ItemType     string           `protobuf:"bytes,4,opt,name=ItemType,proto3" json:"ItemType,omitempty"`
	ItemCategory string           `protobuf:"bytes,5,opt,name=ItemCategory,proto3" json:"ItemCategory,omitempty"`
}

func (x *CrafterStartRequest) Reset() {
	*x = CrafterStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_safeplanet_crafter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrafterStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrafterStartRequest) ProtoMessage() {}

func (x *CrafterStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_safeplanet_crafter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrafterStartRequest.ProtoReflect.Descriptor instead.
func (*CrafterStartRequest) Descriptor() ([]byte, []int) {
	return file_proto_safeplanet_crafter_proto_rawDescGZIP(), []int{0}
}

func (x *CrafterStartRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *CrafterStartRequest) GetResources() map[uint32]int32 {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *CrafterStartRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CrafterStartRequest) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *CrafterStartRequest) GetItemCategory() string {
	if x != nil {
		return x.ItemCategory
	}
	return ""
}

type CrafterStartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CraftingEndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CraftingEndTime,proto3" json:"CraftingEndTime,omitempty"`
}

func (x *CrafterStartResponse) Reset() {
	*x = CrafterStartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_safeplanet_crafter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrafterStartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrafterStartResponse) ProtoMessage() {}

func (x *CrafterStartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_safeplanet_crafter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrafterStartResponse.ProtoReflect.Descriptor instead.
func (*CrafterStartResponse) Descriptor() ([]byte, []int) {
	return file_proto_safeplanet_crafter_proto_rawDescGZIP(), []int{1}
}

func (x *CrafterStartResponse) GetCraftingEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.CraftingEndTime
	}
	return nil
}

// CrafterEnd
type CrafterEndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *CrafterEndRequest) Reset() {
	*x = CrafterEndRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_safeplanet_crafter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrafterEndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrafterEndRequest) ProtoMessage() {}

func (x *CrafterEndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_safeplanet_crafter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrafterEndRequest.ProtoReflect.Descriptor instead.
func (*CrafterEndRequest) Descriptor() ([]byte, []int) {
	return file_proto_safeplanet_crafter_proto_rawDescGZIP(), []int{2}
}

func (x *CrafterEndRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type CrafterEndResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapon *Weapon `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	Armor  *Armor  `protobuf:"bytes,2,opt,name=Armor,proto3" json:"Armor,omitempty"`
}

func (x *CrafterEndResponse) Reset() {
	*x = CrafterEndResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_safeplanet_crafter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrafterEndResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrafterEndResponse) ProtoMessage() {}

func (x *CrafterEndResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_safeplanet_crafter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrafterEndResponse.ProtoReflect.Descriptor instead.
func (*CrafterEndResponse) Descriptor() ([]byte, []int) {
	return file_proto_safeplanet_crafter_proto_rawDescGZIP(), []int{3}
}

func (x *CrafterEndResponse) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

func (x *CrafterEndResponse) GetArmor() *Armor {
	if x != nil {
		return x.Armor
	}
	return nil
}

// CrafterCheck
type CrafterCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *CrafterCheckRequest) Reset() {
	*x = CrafterCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_safeplanet_crafter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrafterCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrafterCheckRequest) ProtoMessage() {}

func (x *CrafterCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_safeplanet_crafter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrafterCheckRequest.ProtoReflect.Descriptor instead.
func (*CrafterCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_safeplanet_crafter_proto_rawDescGZIP(), []int{4}
}

func (x *CrafterCheckRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type CrafterCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CraftInProgress bool                 `protobuf:"varint,1,opt,name=CraftInProgress,proto3" json:"CraftInProgress,omitempty"`
	FinishCrafting  bool                 `protobuf:"varint,2,opt,name=FinishCrafting,proto3" json:"FinishCrafting,omitempty"`
	CraftingEndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CraftingEndTime,proto3" json:"CraftingEndTime,omitempty"`
}

func (x *CrafterCheckResponse) Reset() {
	*x = CrafterCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_safeplanet_crafter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrafterCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrafterCheckResponse) ProtoMessage() {}

func (x *CrafterCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_safeplanet_crafter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrafterCheckResponse.ProtoReflect.Descriptor instead.
func (*CrafterCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_safeplanet_crafter_proto_rawDescGZIP(), []int{5}
}

func (x *CrafterCheckResponse) GetCraftInProgress() bool {
	if x != nil {
		return x.CraftInProgress
	}
	return false
}

func (x *CrafterCheckResponse) GetFinishCrafting() bool {
	if x != nil {
		return x.FinishCrafting
	}
	return false
}

func (x *CrafterCheckResponse) GetCraftingEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.CraftingEndTime
	}
	return nil
}

var File_proto_safeplanet_crafter_proto protoreflect.FileDescriptor

var file_proto_safeplanet_crafter_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x5f, 0x63, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x73, 0x61, 0x66, 0x65, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x5f, 0x63, 0x72, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x6d,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a,
	0x13, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x54, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x61, 0x66, 0x65, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x5f, 0x63, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x3c, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x14, 0x43, 0x72,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e,
	0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x72, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x60, 0x0a, 0x12, 0x43, 0x72, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52,
	0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x52, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x22, 0x31, 0x0a, 0x13, 0x43,
	0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0xae,
	0x01, 0x0a, 0x14, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x72, 0x61, 0x66, 0x74,
	0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x43, 0x72, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x43, 0x72, 0x61, 0x66, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x72, 0x61,
	0x66, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_safeplanet_crafter_proto_rawDescOnce sync.Once
	file_proto_safeplanet_crafter_proto_rawDescData = file_proto_safeplanet_crafter_proto_rawDesc
)

func file_proto_safeplanet_crafter_proto_rawDescGZIP() []byte {
	file_proto_safeplanet_crafter_proto_rawDescOnce.Do(func() {
		file_proto_safeplanet_crafter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_safeplanet_crafter_proto_rawDescData)
	})
	return file_proto_safeplanet_crafter_proto_rawDescData
}

var file_proto_safeplanet_crafter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_safeplanet_crafter_proto_goTypes = []interface{}{
	(*CrafterStartRequest)(nil),  // 0: safeplanet_crafter.CrafterStartRequest
	(*CrafterStartResponse)(nil), // 1: safeplanet_crafter.CrafterStartResponse
	(*CrafterEndRequest)(nil),    // 2: safeplanet_crafter.CrafterEndRequest
	(*CrafterEndResponse)(nil),   // 3: safeplanet_crafter.CrafterEndResponse
	(*CrafterCheckRequest)(nil),  // 4: safeplanet_crafter.CrafterCheckRequest
	(*CrafterCheckResponse)(nil), // 5: safeplanet_crafter.CrafterCheckResponse
	nil,                          // 6: safeplanet_crafter.CrafterStartRequest.ResourcesEntry
	(*timestamp.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*Weapon)(nil),               // 8: weapon.Weapon
	(*Armor)(nil),                // 9: armor.Armor
}
var file_proto_safeplanet_crafter_proto_depIdxs = []int32{
	6, // 0: safeplanet_crafter.CrafterStartRequest.Resources:type_name -> safeplanet_crafter.CrafterStartRequest.ResourcesEntry
	7, // 1: safeplanet_crafter.CrafterStartResponse.CraftingEndTime:type_name -> google.protobuf.Timestamp
	8, // 2: safeplanet_crafter.CrafterEndResponse.Weapon:type_name -> weapon.Weapon
	9, // 3: safeplanet_crafter.CrafterEndResponse.Armor:type_name -> armor.Armor
	7, // 4: safeplanet_crafter.CrafterCheckResponse.CraftingEndTime:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_safeplanet_crafter_proto_init() }
func file_proto_safeplanet_crafter_proto_init() {
	if File_proto_safeplanet_crafter_proto != nil {
		return
	}
	file_proto_armor_proto_init()
	file_proto_weapon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_safeplanet_crafter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrafterStartRequest); i {
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
		file_proto_safeplanet_crafter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrafterStartResponse); i {
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
		file_proto_safeplanet_crafter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrafterEndRequest); i {
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
		file_proto_safeplanet_crafter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrafterEndResponse); i {
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
		file_proto_safeplanet_crafter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrafterCheckRequest); i {
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
		file_proto_safeplanet_crafter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrafterCheckResponse); i {
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
			RawDescriptor: file_proto_safeplanet_crafter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_safeplanet_crafter_proto_goTypes,
		DependencyIndexes: file_proto_safeplanet_crafter_proto_depIdxs,
		MessageInfos:      file_proto_safeplanet_crafter_proto_msgTypes,
	}.Build()
	File_proto_safeplanet_crafter_proto = out.File
	file_proto_safeplanet_crafter_proto_rawDesc = nil
	file_proto_safeplanet_crafter_proto_goTypes = nil
	file_proto_safeplanet_crafter_proto_depIdxs = nil
}
