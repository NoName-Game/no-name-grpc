// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/armor.proto

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
	Defense         float64        `protobuf:"fixed64,10,opt,name=Defense,proto3" json:"Defense,omitempty"`
	Evasion         float64        `protobuf:"fixed64,11,opt,name=Evasion,proto3" json:"Evasion,omitempty"`
	Halving         float64        `protobuf:"fixed64,12,opt,name=Halving,proto3" json:"Halving,omitempty"`
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

func (x *Armor) GetDefense() float64 {
	if x != nil {
		return x.Defense
	}
	return 0
}

func (x *Armor) GetEvasion() float64 {
	if x != nil {
		return x.Evasion
	}
	return 0
}

func (x *Armor) GetHalving() float64 {
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

// GetArmorByID
type GetArmorByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArmorID uint32 `protobuf:"varint,1,opt,name=ArmorID,proto3" json:"ArmorID,omitempty"`
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

func (x *GetArmorByIDRequest) GetArmorID() uint32 {
	if x != nil {
		return x.ArmorID
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

// GetArmorByName
type GetArmorByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetArmorByNameRequest) Reset() {
	*x = GetArmorByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArmorByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArmorByNameRequest) ProtoMessage() {}

func (x *GetArmorByNameRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetArmorByNameRequest.ProtoReflect.Descriptor instead.
func (*GetArmorByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{3}
}

func (x *GetArmorByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetArmorByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armor *Armor `protobuf:"bytes,1,opt,name=Armor,proto3" json:"Armor,omitempty"`
}

func (x *GetArmorByNameResponse) Reset() {
	*x = GetArmorByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArmorByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArmorByNameResponse) ProtoMessage() {}

func (x *GetArmorByNameResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetArmorByNameResponse.ProtoReflect.Descriptor instead.
func (*GetArmorByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{4}
}

func (x *GetArmorByNameResponse) GetArmor() *Armor {
	if x != nil {
		return x.Armor
	}
	return nil
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
		mi := &file_proto_armor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsRequest) ProtoMessage() {}

func (x *GetPlayerArmorsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPlayerArmorsRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{5}
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
		mi := &file_proto_armor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsResponse) ProtoMessage() {}

func (x *GetPlayerArmorsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPlayerArmorsResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{6}
}

func (x *GetPlayerArmorsResponse) GetArmors() []*Armor {
	if x != nil {
		return x.Armors
	}
	return nil
}

// GetPlayerArmorsByCategoryID
type GetPlayerArmorsByCategoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	CategoryID uint32 `protobuf:"varint,2,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
}

func (x *GetPlayerArmorsByCategoryIDRequest) Reset() {
	*x = GetPlayerArmorsByCategoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsByCategoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsByCategoryIDRequest) ProtoMessage() {}

func (x *GetPlayerArmorsByCategoryIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPlayerArmorsByCategoryIDRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsByCategoryIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlayerArmorsByCategoryIDRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *GetPlayerArmorsByCategoryIDRequest) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type GetPlayerArmorsByCategoryIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armors []*Armor `protobuf:"bytes,1,rep,name=Armors,proto3" json:"Armors,omitempty"`
}

func (x *GetPlayerArmorsByCategoryIDResponse) Reset() {
	*x = GetPlayerArmorsByCategoryIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorsByCategoryIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorsByCategoryIDResponse) ProtoMessage() {}

func (x *GetPlayerArmorsByCategoryIDResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPlayerArmorsByCategoryIDResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorsByCategoryIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{8}
}

func (x *GetPlayerArmorsByCategoryIDResponse) GetArmors() []*Armor {
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

// GetPlayerArmorsEquippedByCategoryID
type GetPlayerArmorEquippedByCategoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	CategoryID uint32 `protobuf:"varint,2,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
}

func (x *GetPlayerArmorEquippedByCategoryIDRequest) Reset() {
	*x = GetPlayerArmorEquippedByCategoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorEquippedByCategoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorEquippedByCategoryIDRequest) ProtoMessage() {}

func (x *GetPlayerArmorEquippedByCategoryIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerArmorEquippedByCategoryIDRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorEquippedByCategoryIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{11}
}

func (x *GetPlayerArmorEquippedByCategoryIDRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *GetPlayerArmorEquippedByCategoryIDRequest) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type GetPlayerArmorEquippedByCategoryIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Armor *Armor `protobuf:"bytes,1,opt,name=Armor,proto3" json:"Armor,omitempty"`
}

func (x *GetPlayerArmorEquippedByCategoryIDResponse) Reset() {
	*x = GetPlayerArmorEquippedByCategoryIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerArmorEquippedByCategoryIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerArmorEquippedByCategoryIDResponse) ProtoMessage() {}

func (x *GetPlayerArmorEquippedByCategoryIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerArmorEquippedByCategoryIDResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerArmorEquippedByCategoryIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{12}
}

func (x *GetPlayerArmorEquippedByCategoryIDResponse) GetArmor() *Armor {
	if x != nil {
		return x.Armor
	}
	return nil
}

// EquipArmor
type EquipArmorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	ArmorID  uint32 `protobuf:"varint,2,opt,name=ArmorID,proto3" json:"ArmorID,omitempty"`
	Equip    bool   `protobuf:"varint,3,opt,name=Equip,proto3" json:"Equip,omitempty"`
}

func (x *EquipArmorRequest) Reset() {
	*x = EquipArmorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipArmorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipArmorRequest) ProtoMessage() {}

func (x *EquipArmorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipArmorRequest.ProtoReflect.Descriptor instead.
func (*EquipArmorRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{13}
}

func (x *EquipArmorRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *EquipArmorRequest) GetArmorID() uint32 {
	if x != nil {
		return x.ArmorID
	}
	return 0
}

func (x *EquipArmorRequest) GetEquip() bool {
	if x != nil {
		return x.Equip
	}
	return false
}

type EquipArmorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EquipArmorResponse) Reset() {
	*x = EquipArmorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipArmorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipArmorResponse) ProtoMessage() {}

func (x *EquipArmorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipArmorResponse.ProtoReflect.Descriptor instead.
func (*EquipArmorResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_proto_rawDescGZIP(), []int{14}
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
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x76, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x45, 0x76, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x76, 0x69,
	0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x48, 0x61, 0x6c, 0x76, 0x69, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22,
	0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x49, 0x44,
	0x22, 0x3a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72,
	0x52, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3f, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72,
	0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x22, 0x60,
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72,
	0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x22, 0x4b, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x22, 0x3c, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x06, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x73, 0x22, 0x67, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x70, 0x65, 0x64, 0x42, 0x79,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x50, 0x0a,
	0x2a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x70, 0x65, 0x64, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6d,
	0x6f, 0x72, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x05, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x22,
	0x5f, 0x0a, 0x11, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x22, 0x14, 0x0a, 0x12, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_armor_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_armor_proto_goTypes = []interface{}{
	(*Armor)(nil),                                      // 0: armor.Armor
	(*GetArmorByIDRequest)(nil),                        // 1: armor.GetArmorByIDRequest
	(*GetArmorByIDResponse)(nil),                       // 2: armor.GetArmorByIDResponse
	(*GetArmorByNameRequest)(nil),                      // 3: armor.GetArmorByNameRequest
	(*GetArmorByNameResponse)(nil),                     // 4: armor.GetArmorByNameResponse
	(*GetPlayerArmorsRequest)(nil),                     // 5: armor.GetPlayerArmorsRequest
	(*GetPlayerArmorsResponse)(nil),                    // 6: armor.GetPlayerArmorsResponse
	(*GetPlayerArmorsByCategoryIDRequest)(nil),         // 7: armor.GetPlayerArmorsByCategoryIDRequest
	(*GetPlayerArmorsByCategoryIDResponse)(nil),        // 8: armor.GetPlayerArmorsByCategoryIDResponse
	(*GetPlayerArmorsEquippedRequest)(nil),             // 9: armor.GetPlayerArmorsEquippedRequest
	(*GetPlayerArmorsEquippedResponse)(nil),            // 10: armor.GetPlayerArmorsEquippedResponse
	(*GetPlayerArmorEquippedByCategoryIDRequest)(nil),  // 11: armor.GetPlayerArmorEquippedByCategoryIDRequest
	(*GetPlayerArmorEquippedByCategoryIDResponse)(nil), // 12: armor.GetPlayerArmorEquippedByCategoryIDResponse
	(*EquipArmorRequest)(nil),                          // 13: armor.EquipArmorRequest
	(*EquipArmorResponse)(nil),                         // 14: armor.EquipArmorResponse
	(*Rarity)(nil),                                     // 15: rarity.Rarity
	(*ArmorCategory)(nil),                              // 16: armor_category.ArmorCategory
}
var file_proto_armor_proto_depIdxs = []int32{
	15, // 0: armor.Armor.Rarity:type_name -> rarity.Rarity
	16, // 1: armor.Armor.ArmorCategory:type_name -> armor_category.ArmorCategory
	0,  // 2: armor.GetArmorByIDResponse.Armor:type_name -> armor.Armor
	0,  // 3: armor.GetArmorByNameResponse.Armor:type_name -> armor.Armor
	0,  // 4: armor.GetPlayerArmorsResponse.Armors:type_name -> armor.Armor
	0,  // 5: armor.GetPlayerArmorsByCategoryIDResponse.Armors:type_name -> armor.Armor
	0,  // 6: armor.GetPlayerArmorsEquippedResponse.Armors:type_name -> armor.Armor
	0,  // 7: armor.GetPlayerArmorEquippedByCategoryIDResponse.Armor:type_name -> armor.Armor
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
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
			switch v := v.(*GetArmorByNameRequest); i {
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
			switch v := v.(*GetArmorByNameResponse); i {
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
		file_proto_armor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_armor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorsByCategoryIDRequest); i {
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
			switch v := v.(*GetPlayerArmorsByCategoryIDResponse); i {
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
		file_proto_armor_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorEquippedByCategoryIDRequest); i {
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
		file_proto_armor_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerArmorEquippedByCategoryIDResponse); i {
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
		file_proto_armor_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipArmorRequest); i {
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
		file_proto_armor_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipArmorResponse); i {
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
			NumMessages:   15,
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
