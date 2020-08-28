// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/armor.proto

//option go_package = "bitbucket.org/no-name-game/nn-grpc/build/pb/armor";

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type Armor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              uint32         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name            string         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	RarityID        uint32         `protobuf:"varint,3,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity          *Rarity        `protobuf:"bytes,4,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ArmorCategoryID uint32         `protobuf:"varint,5,opt,name=ArmorCategoryID,proto3" json:"ArmorCategoryID,omitempty"`
	ArmorCategory   *ArmorCategory `protobuf:"bytes,6,opt,name=ArmorCategory,proto3" json:"ArmorCategory,omitempty"`
	PlayerID        uint32         `protobuf:"varint,7,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Equipped        bool           `protobuf:"varint,9,opt,name=Equipped,proto3" json:"Equipped,omitempty"`
	Defense         float32        `protobuf:"fixed32,10,opt,name=Defense,proto3" json:"Defense,omitempty"`
	Evasion         float32        `protobuf:"fixed32,11,opt,name=Evasion,proto3" json:"Evasion,omitempty"`
	Halving         float32        `protobuf:"fixed32,12,opt,name=Halving,proto3" json:"Halving,omitempty"`
	Potential       uint32         `protobuf:"varint,13,opt,name=Potential,proto3" json:"Potential,omitempty"`
}

func (x *Armor) Reset() {
	*x = Armor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Armor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Armor) ProtoMessage() {}

func (x *Armor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Armor.ProtoReflect.Descriptor instead.
func (*Armor) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{0}
}

func (x *Armor) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Armor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Armor) GetRarityID() uint32 {
	if x != nil {
		return x.RarityID
	}
	return 0
}

func (x *Armor) GetRarity() *Rarity {
	if x != nil {
		return x.Rarity
	}
	return nil
}

func (x *Armor) GetArmorCategoryID() uint32 {
	if x != nil {
		return x.ArmorCategoryID
	}
	return 0
}

func (x *Armor) GetArmorCategory() *ArmorCategory {
	if x != nil {
		return x.ArmorCategory
	}
	return nil
}

func (x *Armor) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *Armor) GetEquipped() bool {
	if x != nil {
		return x.Equipped
	}
	return false
}

func (x *Armor) GetDefense() float32 {
	if x != nil {
		return x.Defense
	}
	return 0
}

func (x *Armor) GetEvasion() float32 {
	if x != nil {
		return x.Evasion
	}
	return 0
}

func (x *Armor) GetHalving() float32 {
	if x != nil {
		return x.Halving
	}
	return 0
}

func (x *Armor) GetPotential() uint32 {
	if x != nil {
		return x.Potential
	}
	return 0
}

type GetArmorByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetArmorByIDRequest) Reset() {
	*x = GetArmorByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArmorByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArmorByIDRequest) ProtoMessage() {}

func (x *GetArmorByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArmorByIDRequest.ProtoReflect.Descriptor instead.
func (*GetArmorByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{1}
}

func (x *GetArmorByIDRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetArmorByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armor *Armor `protobuf:"bytes,1,opt,name=Armor,proto3" json:"Armor,omitempty"`
}

func (x *GetArmorByIDResponse) Reset() {
	*x = GetArmorByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArmorByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArmorByIDResponse) ProtoMessage() {}

func (x *GetArmorByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArmorByIDResponse.ProtoReflect.Descriptor instead.
func (*GetArmorByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{2}
}

func (x *GetArmorByIDResponse) GetArmor() *Armor {
	if x != nil {
		return x.Armor
	}
	return nil
}

type FindArmorByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FindArmorByNameRequest) Reset() {
	*x = FindArmorByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindArmorByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindArmorByNameRequest) ProtoMessage() {}

func (x *FindArmorByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindArmorByNameRequest.ProtoReflect.Descriptor instead.
func (*FindArmorByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{3}
}

func (x *FindArmorByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FindArmorByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armor *Armor `protobuf:"bytes,1,opt,name=Armor,proto3" json:"Armor,omitempty"`
}

func (x *FindArmorByNameResponse) Reset() {
	*x = FindArmorByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindArmorByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindArmorByNameResponse) ProtoMessage() {}

func (x *FindArmorByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindArmorByNameResponse.ProtoReflect.Descriptor instead.
func (*FindArmorByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{4}
}

func (x *FindArmorByNameResponse) GetArmor() *Armor {
	if x != nil {
		return x.Armor
	}
	return nil
}

type UpdateArmorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armor *Armor `protobuf:"bytes,1,opt,name=Armor,proto3" json:"Armor,omitempty"`
}

func (x *UpdateArmorRequest) Reset() {
	*x = UpdateArmorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArmorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArmorRequest) ProtoMessage() {}

func (x *UpdateArmorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArmorRequest.ProtoReflect.Descriptor instead.
func (*UpdateArmorRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateArmorRequest) GetArmor() *Armor {
	if x != nil {
		return x.Armor
	}
	return nil
}

type UpdateArmorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateArmorResponse) Reset() {
	*x = UpdateArmorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArmorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArmorResponse) ProtoMessage() {}

func (x *UpdateArmorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArmorResponse.ProtoReflect.Descriptor instead.
func (*UpdateArmorResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{6}
}

// GetPlayerArmors
type GetPlayerArmorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerArmorsRequest) Reset() {
	*x = GetPlayerArmorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsRequest) ProtoMessage() {}

func (x *GetPlayerArmorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerArmorsRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlayerArmorsRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerArmorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armors []*Armor `protobuf:"bytes,1,rep,name=Armors,proto3" json:"Armors,omitempty"`
}

func (x *GetPlayerArmorsResponse) Reset() {
	*x = GetPlayerArmorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsResponse) ProtoMessage() {}

func (x *GetPlayerArmorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerArmorsResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{8}
}

func (x *GetPlayerArmorsResponse) GetArmors() []*Armor {
	if x != nil {
		return x.Armors
	}
	return nil
}

// GetPlayerArmorsEquipped
type GetPlayerArmorsEquippedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerArmorsEquippedRequest) Reset() {
	*x = GetPlayerArmorsEquippedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsEquippedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsEquippedRequest) ProtoMessage() {}

func (x *GetPlayerArmorsEquippedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerArmorsEquippedRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsEquippedRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{9}
}

func (x *GetPlayerArmorsEquippedRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerArmorsEquippedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armors []*Armor `protobuf:"bytes,1,rep,name=Armors,proto3" json:"Armors,omitempty"`
}

func (x *GetPlayerArmorsEquippedResponse) Reset() {
	*x = GetPlayerArmorsEquippedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsEquippedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsEquippedResponse) ProtoMessage() {}

func (x *GetPlayerArmorsEquippedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerArmorsEquippedResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsEquippedResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{10}
}

func (x *GetPlayerArmorsEquippedResponse) GetArmors() []*Armor {
	if x != nil {
		return x.Armors
	}
	return nil
}

var File_proto_armor_proto protoreflect.FileDescriptor

var file_proto_armor_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x05, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x0d, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x72, 0x6d, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x76, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x45, 0x76, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x76, 0x69,
	0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x48, 0x61, 0x6c, 0x76, 0x69, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x05, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3d, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d,
	0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x22,
	0x38, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x52, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x34, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52,
	0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x45, 0x71, 0x75, 0x69, 0x70, 0x70,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x45, 0x71, 0x75, 0x69, 0x70, 0x70, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72,
	0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x42, 0x07,
	0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_armor_proto_rawDescOnce sync.Once
	file_proto_armor_proto_rawDescData = file_proto_armor_proto_rawDesc
)

func file_proto_armor_proto_rawDescGZIP() []byte {
	file_proto_armor_proto_rawDescOnce.Do(func() {
		file_proto_armor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_armor_proto_rawDescData)
	})
	return file_proto_armor_proto_rawDescData
}

var file_proto_armor_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_armor_proto_goTypes = []interface{}{
	(*Armor)(nil),                           // 0: armor.Armor
	(*GetArmorByIDRequest)(nil),             // 1: armor.GetArmorByIDRequest
	(*GetArmorByIDResponse)(nil),            // 2: armor.GetArmorByIDResponse
	(*FindArmorByNameRequest)(nil),          // 3: armor.FindArmorByNameRequest
	(*FindArmorByNameResponse)(nil),         // 4: armor.FindArmorByNameResponse
	(*UpdateArmorRequest)(nil),              // 5: armor.UpdateArmorRequest
	(*UpdateArmorResponse)(nil),             // 6: armor.UpdateArmorResponse
	(*GetPlayerArmorsRequest)(nil),          // 7: armor.GetPlayerArmorsRequest
	(*GetPlayerArmorsResponse)(nil),         // 8: armor.GetPlayerArmorsResponse
	(*GetPlayerArmorsEquippedRequest)(nil),  // 9: armor.GetPlayerArmorsEquippedRequest
	(*GetPlayerArmorsEquippedResponse)(nil), // 10: armor.GetPlayerArmorsEquippedResponse
	(*Rarity)(nil),                          // 11: rarity.Rarity
	(*ArmorCategory)(nil),                   // 12: armor_category.ArmorCategory
}
var file_proto_armor_proto_depIdxs = []int32{
	11, // 0: armor.Armor.Rarity:type_name -> rarity.Rarity
	12, // 1: armor.Armor.ArmorCategory:type_name -> armor_category.ArmorCategory
	0,  // 2: armor.GetArmorByIDResponse.Armor:type_name -> armor.Armor
	0,  // 3: armor.FindArmorByNameResponse.Armor:type_name -> armor.Armor
	0,  // 4: armor.UpdateArmorRequest.Armor:type_name -> armor.Armor
	0,  // 5: armor.GetPlayerArmorsResponse.Armors:type_name -> armor.Armor
	0,  // 6: armor.GetPlayerArmorsEquippedResponse.Armors:type_name -> armor.Armor
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_armor_proto_init() }
func file_proto_armor_proto_init() {
	if File_proto_armor_proto != nil {
		return
	}
	file_proto_rarity_proto_init()
	file_proto_armor_category_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_armor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Armor); i {
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
		file_proto_armor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArmorByIDRequest); i {
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
		file_proto_armor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArmorByIDResponse); i {
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
		file_proto_armor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindArmorByNameRequest); i {
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
		file_proto_armor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindArmorByNameResponse); i {
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
		file_proto_armor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArmorRequest); i {
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
		file_proto_armor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArmorResponse); i {
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
		file_proto_armor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorsRequest); i {
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
		file_proto_armor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorsResponse); i {
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
		file_proto_armor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorsEquippedRequest); i {
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
		file_proto_armor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorsEquippedResponse); i {
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
			RawDescriptor: file_proto_armor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_armor_proto_goTypes,
		DependencyIndexes: file_proto_armor_proto_depIdxs,
		MessageInfos:      file_proto_armor_proto_msgTypes,
	}.Build()
	File_proto_armor_proto = out.File
	file_proto_armor_proto_rawDesc = nil
	file_proto_armor_proto_goTypes = nil
	file_proto_armor_proto_depIdxs = nil
}