// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/weapon.proto

package pb

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

type Weapon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	RawDamage     float32 `protobuf:"fixed32,3,opt,name=RawDamage,proto3" json:"RawDamage,omitempty"`
	PlayerID      uint32  `protobuf:"varint,4,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Equipped      bool    `protobuf:"varint,5,opt,name=Equipped,proto3" json:"Equipped,omitempty"`
	RarityID      uint32  `protobuf:"varint,6,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity        *Rarity `protobuf:"bytes,7,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	Precision     float32 `protobuf:"fixed32,10,opt,name=Precision,proto3" json:"Precision,omitempty"`
	Penetration   float32 `protobuf:"fixed32,11,opt,name=Penetration,proto3" json:"Penetration,omitempty"`
	Potential     uint32  `protobuf:"varint,12,opt,name=Potential,proto3" json:"Potential,omitempty"`
	Durability    int32   `protobuf:"varint,13,opt,name=Durability,proto3" json:"Durability,omitempty"`
	DurabilityCap int32   `protobuf:"varint,14,opt,name=DurabilityCap,proto3" json:"DurabilityCap,omitempty"`
}

func (x *Weapon) Reset() {
	*x = Weapon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weapon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weapon) ProtoMessage() {}

func (x *Weapon) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weapon.ProtoReflect.Descriptor instead.
func (*Weapon) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{0}
}

func (x *Weapon) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Weapon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Weapon) GetRawDamage() float32 {
	if x != nil {
		return x.RawDamage
	}
	return 0
}

func (x *Weapon) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *Weapon) GetEquipped() bool {
	if x != nil {
		return x.Equipped
	}
	return false
}

func (x *Weapon) GetRarityID() uint32 {
	if x != nil {
		return x.RarityID
	}
	return 0
}

func (x *Weapon) GetRarity() *Rarity {
	if x != nil {
		return x.Rarity
	}
	return nil
}

func (x *Weapon) GetPrecision() float32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *Weapon) GetPenetration() float32 {
	if x != nil {
		return x.Penetration
	}
	return 0
}

func (x *Weapon) GetPotential() uint32 {
	if x != nil {
		return x.Potential
	}
	return 0
}

func (x *Weapon) GetDurability() int32 {
	if x != nil {
		return x.Durability
	}
	return 0
}

func (x *Weapon) GetDurabilityCap() int32 {
	if x != nil {
		return x.DurabilityCap
	}
	return 0
}

// GetWeaponByID
type GetWeaponByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetWeaponByIDRequest) Reset() {
	*x = GetWeaponByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeaponByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeaponByIDRequest) ProtoMessage() {}

func (x *GetWeaponByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeaponByIDRequest.ProtoReflect.Descriptor instead.
func (*GetWeaponByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeaponByIDRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetWeaponByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapon *Weapon `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
}

func (x *GetWeaponByIDResponse) Reset() {
	*x = GetWeaponByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeaponByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeaponByIDResponse) ProtoMessage() {}

func (x *GetWeaponByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeaponByIDResponse.ProtoReflect.Descriptor instead.
func (*GetWeaponByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{2}
}

func (x *GetWeaponByIDResponse) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

// GetWeaponByName
type GetWeaponByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetWeaponByNameRequest) Reset() {
	*x = GetWeaponByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeaponByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeaponByNameRequest) ProtoMessage() {}

func (x *GetWeaponByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeaponByNameRequest.ProtoReflect.Descriptor instead.
func (*GetWeaponByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{3}
}

func (x *GetWeaponByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetWeaponByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapon *Weapon `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
}

func (x *GetWeaponByNameResponse) Reset() {
	*x = GetWeaponByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeaponByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeaponByNameResponse) ProtoMessage() {}

func (x *GetWeaponByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeaponByNameResponse.ProtoReflect.Descriptor instead.
func (*GetWeaponByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{4}
}

func (x *GetWeaponByNameResponse) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

// GetPlayerWeapons
type GetPlayerWeaponsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerWeaponsRequest) Reset() {
	*x = GetPlayerWeaponsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerWeaponsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerWeaponsRequest) ProtoMessage() {}

func (x *GetPlayerWeaponsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerWeaponsRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerWeaponsRequest) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlayerWeaponsRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerWeaponsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapons []*Weapon `protobuf:"bytes,1,rep,name=Weapons,proto3" json:"Weapons,omitempty"`
}

func (x *GetPlayerWeaponsResponse) Reset() {
	*x = GetPlayerWeaponsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerWeaponsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerWeaponsResponse) ProtoMessage() {}

func (x *GetPlayerWeaponsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerWeaponsResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerWeaponsResponse) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{6}
}

func (x *GetPlayerWeaponsResponse) GetWeapons() []*Weapon {
	if x != nil {
		return x.Weapons
	}
	return nil
}

// GetPlayerWeaponEquipped
type GetPlayerWeaponEquippedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerWeaponEquippedRequest) Reset() {
	*x = GetPlayerWeaponEquippedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerWeaponEquippedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerWeaponEquippedRequest) ProtoMessage() {}

func (x *GetPlayerWeaponEquippedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerWeaponEquippedRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerWeaponEquippedRequest) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlayerWeaponEquippedRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerWeaponEquippedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapon *Weapon `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
}

func (x *GetPlayerWeaponEquippedResponse) Reset() {
	*x = GetPlayerWeaponEquippedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerWeaponEquippedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerWeaponEquippedResponse) ProtoMessage() {}

func (x *GetPlayerWeaponEquippedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerWeaponEquippedResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerWeaponEquippedResponse) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{8}
}

func (x *GetPlayerWeaponEquippedResponse) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

// EquipWeapon
type EquipWeaponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	WeaponID uint32 `protobuf:"varint,2,opt,name=WeaponID,proto3" json:"WeaponID,omitempty"`
	Equip    bool   `protobuf:"varint,3,opt,name=Equip,proto3" json:"Equip,omitempty"`
}

func (x *EquipWeaponRequest) Reset() {
	*x = EquipWeaponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipWeaponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipWeaponRequest) ProtoMessage() {}

func (x *EquipWeaponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipWeaponRequest.ProtoReflect.Descriptor instead.
func (*EquipWeaponRequest) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{9}
}

func (x *EquipWeaponRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *EquipWeaponRequest) GetWeaponID() uint32 {
	if x != nil {
		return x.WeaponID
	}
	return 0
}

func (x *EquipWeaponRequest) GetEquip() bool {
	if x != nil {
		return x.Equip
	}
	return false
}

type EquipWeaponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EquipWeaponResponse) Reset() {
	*x = EquipWeaponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weapon_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipWeaponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipWeaponResponse) ProtoMessage() {}

func (x *EquipWeaponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weapon_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipWeaponResponse.ProtoReflect.Descriptor instead.
func (*EquipWeaponResponse) Descriptor() ([]byte, []int) {
	return file_proto_weapon_proto_rawDescGZIP(), []int{10}
}

var File_proto_weapon_proto protoreflect.FileDescriptor

var file_proto_weapon_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x1a, 0x12, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x02, 0x0a, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x52, 0x61, 0x77, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x52, 0x61, 0x77, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x50, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x65, 0x6e, 0x65, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x50, 0x65,
	0x6e, 0x65, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x6f,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x75, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x75, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x75, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x70, 0x22, 0x26, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52, 0x06,
	0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52,
	0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x44,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x57, 0x65,
	0x61, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x65,
	0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52, 0x07, 0x57, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x70, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x49, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x22, 0x62, 0x0a,
	0x12, 0x45, 0x71, 0x75, 0x69, 0x70, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x71, 0x75, 0x69, 0x70, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_weapon_proto_rawDescOnce sync.Once
	file_proto_weapon_proto_rawDescData = file_proto_weapon_proto_rawDesc
)

func file_proto_weapon_proto_rawDescGZIP() []byte {
	file_proto_weapon_proto_rawDescOnce.Do(func() {
		file_proto_weapon_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_weapon_proto_rawDescData)
	})
	return file_proto_weapon_proto_rawDescData
}

var file_proto_weapon_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_weapon_proto_goTypes = []interface{}{
	(*Weapon)(nil),                          // 0: weapon.Weapon
	(*GetWeaponByIDRequest)(nil),            // 1: weapon.GetWeaponByIDRequest
	(*GetWeaponByIDResponse)(nil),           // 2: weapon.GetWeaponByIDResponse
	(*GetWeaponByNameRequest)(nil),          // 3: weapon.GetWeaponByNameRequest
	(*GetWeaponByNameResponse)(nil),         // 4: weapon.GetWeaponByNameResponse
	(*GetPlayerWeaponsRequest)(nil),         // 5: weapon.GetPlayerWeaponsRequest
	(*GetPlayerWeaponsResponse)(nil),        // 6: weapon.GetPlayerWeaponsResponse
	(*GetPlayerWeaponEquippedRequest)(nil),  // 7: weapon.GetPlayerWeaponEquippedRequest
	(*GetPlayerWeaponEquippedResponse)(nil), // 8: weapon.GetPlayerWeaponEquippedResponse
	(*EquipWeaponRequest)(nil),              // 9: weapon.EquipWeaponRequest
	(*EquipWeaponResponse)(nil),             // 10: weapon.EquipWeaponResponse
	(*Rarity)(nil),                          // 11: rarity.Rarity
}
var file_proto_weapon_proto_depIdxs = []int32{
	11, // 0: weapon.Weapon.Rarity:type_name -> rarity.Rarity
	0,  // 1: weapon.GetWeaponByIDResponse.Weapon:type_name -> weapon.Weapon
	0,  // 2: weapon.GetWeaponByNameResponse.Weapon:type_name -> weapon.Weapon
	0,  // 3: weapon.GetPlayerWeaponsResponse.Weapons:type_name -> weapon.Weapon
	0,  // 4: weapon.GetPlayerWeaponEquippedResponse.Weapon:type_name -> weapon.Weapon
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_weapon_proto_init() }
func file_proto_weapon_proto_init() {
	if File_proto_weapon_proto != nil {
		return
	}
	file_proto_rarity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_weapon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weapon); i {
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
		file_proto_weapon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeaponByIDRequest); i {
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
		file_proto_weapon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeaponByIDResponse); i {
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
		file_proto_weapon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeaponByNameRequest); i {
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
		file_proto_weapon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeaponByNameResponse); i {
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
		file_proto_weapon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerWeaponsRequest); i {
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
		file_proto_weapon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerWeaponsResponse); i {
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
		file_proto_weapon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerWeaponEquippedRequest); i {
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
		file_proto_weapon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerWeaponEquippedResponse); i {
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
		file_proto_weapon_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipWeaponRequest); i {
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
		file_proto_weapon_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipWeaponResponse); i {
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
			RawDescriptor: file_proto_weapon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_weapon_proto_goTypes,
		DependencyIndexes: file_proto_weapon_proto_depIdxs,
		MessageInfos:      file_proto_weapon_proto_msgTypes,
	}.Build()
	File_proto_weapon_proto = out.File
	file_proto_weapon_proto_rawDesc = nil
	file_proto_weapon_proto_goTypes = nil
	file_proto_weapon_proto_depIdxs = nil
}
