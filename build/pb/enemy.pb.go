// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/enemy.proto

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

type Enemy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	MapID     uint32  `protobuf:"varint,3,opt,name=MapID,proto3" json:"MapID,omitempty"`
	PositionX int32   `protobuf:"varint,4,opt,name=PositionX,proto3" json:"PositionX,omitempty"`
	PositionY int32   `protobuf:"varint,5,opt,name=PositionY,proto3" json:"PositionY,omitempty"`
	RarityID  uint32  `protobuf:"varint,6,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity    *Rarity `protobuf:"bytes,7,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	LifeMax   uint32  `protobuf:"varint,8,opt,name=LifeMax,proto3" json:"LifeMax,omitempty"`
	LifePoint int32   `protobuf:"varint,9,opt,name=LifePoint,proto3" json:"LifePoint,omitempty"`
}

func (x *Enemy) Reset() {
	*x = Enemy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_enemy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enemy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enemy) ProtoMessage() {}

func (x *Enemy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enemy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enemy.ProtoReflect.Descriptor instead.
func (*Enemy) Descriptor() ([]byte, []int) {
	return file_proto_enemy_proto_rawDescGZIP(), []int{0}
}

func (x *Enemy) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Enemy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Enemy) GetMapID() uint32 {
	if x != nil {
		return x.MapID
	}
	return 0
}

func (x *Enemy) GetPositionX() int32 {
	if x != nil {
		return x.PositionX
	}
	return 0
}

func (x *Enemy) GetPositionY() int32 {
	if x != nil {
		return x.PositionY
	}
	return 0
}

func (x *Enemy) GetRarityID() uint32 {
	if x != nil {
		return x.RarityID
	}
	return 0
}

func (x *Enemy) GetRarity() *Rarity {
	if x != nil {
		return x.Rarity
	}
	return nil
}

func (x *Enemy) GetLifeMax() uint32 {
	if x != nil {
		return x.LifeMax
	}
	return 0
}

func (x *Enemy) GetLifePoint() int32 {
	if x != nil {
		return x.LifePoint
	}
	return 0
}

// GetEnemyByID
type GetEnemyByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnemyID uint32 `protobuf:"varint,1,opt,name=EnemyID,proto3" json:"EnemyID,omitempty"`
}

func (x *GetEnemyByIDRequest) Reset() {
	*x = GetEnemyByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_enemy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnemyByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnemyByIDRequest) ProtoMessage() {}

func (x *GetEnemyByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enemy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnemyByIDRequest.ProtoReflect.Descriptor instead.
func (*GetEnemyByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_enemy_proto_rawDescGZIP(), []int{1}
}

func (x *GetEnemyByIDRequest) GetEnemyID() uint32 {
	if x != nil {
		return x.EnemyID
	}
	return 0
}

type GetEnemyByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enemy *Enemy `protobuf:"bytes,1,opt,name=Enemy,proto3" json:"Enemy,omitempty"`
}

func (x *GetEnemyByIDResponse) Reset() {
	*x = GetEnemyByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_enemy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnemyByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnemyByIDResponse) ProtoMessage() {}

func (x *GetEnemyByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enemy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnemyByIDResponse.ProtoReflect.Descriptor instead.
func (*GetEnemyByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_enemy_proto_rawDescGZIP(), []int{2}
}

func (x *GetEnemyByIDResponse) GetEnemy() *Enemy {
	if x != nil {
		return x.Enemy
	}
	return nil
}

// HitEnemy
type HitEnemyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnemyID       uint32 `protobuf:"varint,1,opt,name=EnemyID,proto3" json:"EnemyID,omitempty"`
	PlayerID      uint32 `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	BodySelection int32  `protobuf:"varint,3,opt,name=BodySelection,proto3" json:"BodySelection,omitempty"`
}

func (x *HitEnemyRequest) Reset() {
	*x = HitEnemyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_enemy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HitEnemyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HitEnemyRequest) ProtoMessage() {}

func (x *HitEnemyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enemy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HitEnemyRequest.ProtoReflect.Descriptor instead.
func (*HitEnemyRequest) Descriptor() ([]byte, []int) {
	return file_proto_enemy_proto_rawDescGZIP(), []int{3}
}

func (x *HitEnemyRequest) GetEnemyID() uint32 {
	if x != nil {
		return x.EnemyID
	}
	return 0
}

func (x *HitEnemyRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *HitEnemyRequest) GetBodySelection() int32 {
	if x != nil {
		return x.BodySelection
	}
	return 0
}

type HitEnemyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerDie        bool                 `protobuf:"varint,1,opt,name=PlayerDie,proto3" json:"PlayerDie,omitempty"`
	EnemyDie         bool                 `protobuf:"varint,2,opt,name=EnemyDie,proto3" json:"EnemyDie,omitempty"`
	EnemyDodge       bool                 `protobuf:"varint,3,opt,name=EnemyDodge,proto3" json:"EnemyDodge,omitempty"`
	PlayerDamage     int32                `protobuf:"varint,4,opt,name=PlayerDamage,proto3" json:"PlayerDamage,omitempty"`
	PlayerExperience int32                `protobuf:"varint,5,opt,name=PlayerExperience,proto3" json:"PlayerExperience,omitempty"`
	EnemyDamage      int32                `protobuf:"varint,6,opt,name=EnemyDamage,proto3" json:"EnemyDamage,omitempty"`
	EnemyDrop        *DropTresureResponse `protobuf:"bytes,7,opt,name=EnemyDrop,proto3" json:"EnemyDrop,omitempty"`
	PlayerDodge      bool                 `protobuf:"varint,8,opt,name=PlayerDodge,proto3" json:"PlayerDodge,omitempty"`
}

func (x *HitEnemyResponse) Reset() {
	*x = HitEnemyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_enemy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HitEnemyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HitEnemyResponse) ProtoMessage() {}

func (x *HitEnemyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enemy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HitEnemyResponse.ProtoReflect.Descriptor instead.
func (*HitEnemyResponse) Descriptor() ([]byte, []int) {
	return file_proto_enemy_proto_rawDescGZIP(), []int{4}
}

func (x *HitEnemyResponse) GetPlayerDie() bool {
	if x != nil {
		return x.PlayerDie
	}
	return false
}

func (x *HitEnemyResponse) GetEnemyDie() bool {
	if x != nil {
		return x.EnemyDie
	}
	return false
}

func (x *HitEnemyResponse) GetEnemyDodge() bool {
	if x != nil {
		return x.EnemyDodge
	}
	return false
}

func (x *HitEnemyResponse) GetPlayerDamage() int32 {
	if x != nil {
		return x.PlayerDamage
	}
	return 0
}

func (x *HitEnemyResponse) GetPlayerExperience() int32 {
	if x != nil {
		return x.PlayerExperience
	}
	return 0
}

func (x *HitEnemyResponse) GetEnemyDamage() int32 {
	if x != nil {
		return x.EnemyDamage
	}
	return 0
}

func (x *HitEnemyResponse) GetEnemyDrop() *DropTresureResponse {
	if x != nil {
		return x.EnemyDrop
	}
	return nil
}

func (x *HitEnemyResponse) GetPlayerDodge() bool {
	if x != nil {
		return x.PlayerDodge
	}
	return false
}

var File_proto_enemy_proto protoreflect.FileDescriptor

var file_proto_enemy_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x65, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x70, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x4d, 0x61, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x58, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x58, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x59, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x59, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12,
	0x26, 0x0a, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52,
	0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x66, 0x65, 0x4d,
	0x61, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4c, 0x69, 0x66, 0x65, 0x4d, 0x61,
	0x78, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x49, 0x44,
	0x22, 0x3a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x45, 0x6e, 0x65, 0x6d,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x2e,
	0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x05, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x22, 0x6d, 0x0a, 0x0f,
	0x48, 0x69, 0x74, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x42, 0x6f,
	0x64, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x02, 0x0a, 0x10,
	0x48, 0x69, 0x74, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x44, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x44, 0x69, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6e,
	0x65, 0x6d, 0x79, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x45, 0x6e, 0x65, 0x6d, 0x79, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6e,
	0x65, 0x6d, 0x79, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x45, 0x6e, 0x65, 0x6d, 0x79, 0x44, 0x72, 0x6f, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x74, 0x72, 0x65, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x72,
	0x65, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x45,
	0x6e, 0x65, 0x6d, 0x79, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_enemy_proto_rawDescOnce sync.Once
	file_proto_enemy_proto_rawDescData = file_proto_enemy_proto_rawDesc
)

func file_proto_enemy_proto_rawDescGZIP() []byte {
	file_proto_enemy_proto_rawDescOnce.Do(func() {
		file_proto_enemy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_enemy_proto_rawDescData)
	})
	return file_proto_enemy_proto_rawDescData
}

var file_proto_enemy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_enemy_proto_goTypes = []interface{}{
	(*Enemy)(nil),                // 0: enemy.Enemy
	(*GetEnemyByIDRequest)(nil),  // 1: enemy.GetEnemyByIDRequest
	(*GetEnemyByIDResponse)(nil), // 2: enemy.GetEnemyByIDResponse
	(*HitEnemyRequest)(nil),      // 3: enemy.HitEnemyRequest
	(*HitEnemyResponse)(nil),     // 4: enemy.HitEnemyResponse
	(*Rarity)(nil),               // 5: rarity.Rarity
	(*DropTresureResponse)(nil),  // 6: tresure.DropTresureResponse
}
var file_proto_enemy_proto_depIdxs = []int32{
	5, // 0: enemy.Enemy.Rarity:type_name -> rarity.Rarity
	0, // 1: enemy.GetEnemyByIDResponse.Enemy:type_name -> enemy.Enemy
	6, // 2: enemy.HitEnemyResponse.EnemyDrop:type_name -> tresure.DropTresureResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_enemy_proto_init() }
func file_proto_enemy_proto_init() {
	if File_proto_enemy_proto != nil {
		return
	}
	file_proto_rarity_proto_init()
	file_proto_tresure_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_enemy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enemy); i {
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
		file_proto_enemy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnemyByIDRequest); i {
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
		file_proto_enemy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnemyByIDResponse); i {
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
		file_proto_enemy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HitEnemyRequest); i {
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
		file_proto_enemy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HitEnemyResponse); i {
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
			RawDescriptor: file_proto_enemy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enemy_proto_goTypes,
		DependencyIndexes: file_proto_enemy_proto_depIdxs,
		MessageInfos:      file_proto_enemy_proto_msgTypes,
	}.Build()
	File_proto_enemy_proto = out.File
	file_proto_enemy_proto_rawDesc = nil
	file_proto_enemy_proto_goTypes = nil
	file_proto_enemy_proto_depIdxs = nil
}
