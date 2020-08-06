// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/enemy.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Enemy struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	MapID                uint32   `protobuf:"varint,3,opt,name=MapID,proto3" json:"MapID,omitempty"`
	PositionX            int32    `protobuf:"varint,4,opt,name=PositionX,proto3" json:"PositionX,omitempty"`
	PositionY            int32    `protobuf:"varint,5,opt,name=PositionY,proto3" json:"PositionY,omitempty"`
	RarityID             uint32   `protobuf:"varint,6,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity               *Rarity  `protobuf:"bytes,7,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	LifeMax              uint32   `protobuf:"varint,8,opt,name=LifeMax,proto3" json:"LifeMax,omitempty"`
	LifePoint            int32    `protobuf:"varint,9,opt,name=LifePoint,proto3" json:"LifePoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Enemy) Reset()         { *m = Enemy{} }
func (m *Enemy) String() string { return proto.CompactTextString(m) }
func (*Enemy) ProtoMessage()    {}
func (*Enemy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ae511599e1addec, []int{0}
}

func (m *Enemy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enemy.Unmarshal(m, b)
}
func (m *Enemy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enemy.Marshal(b, m, deterministic)
}
func (m *Enemy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enemy.Merge(m, src)
}
func (m *Enemy) XXX_Size() int {
	return xxx_messageInfo_Enemy.Size(m)
}
func (m *Enemy) XXX_DiscardUnknown() {
	xxx_messageInfo_Enemy.DiscardUnknown(m)
}

var xxx_messageInfo_Enemy proto.InternalMessageInfo

func (m *Enemy) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Enemy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Enemy) GetMapID() uint32 {
	if m != nil {
		return m.MapID
	}
	return 0
}

func (m *Enemy) GetPositionX() int32 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *Enemy) GetPositionY() int32 {
	if m != nil {
		return m.PositionY
	}
	return 0
}

func (m *Enemy) GetRarityID() uint32 {
	if m != nil {
		return m.RarityID
	}
	return 0
}

func (m *Enemy) GetRarity() *Rarity {
	if m != nil {
		return m.Rarity
	}
	return nil
}

func (m *Enemy) GetLifeMax() uint32 {
	if m != nil {
		return m.LifeMax
	}
	return 0
}

func (m *Enemy) GetLifePoint() int32 {
	if m != nil {
		return m.LifePoint
	}
	return 0
}

// GetEnemyByID
type GetEnemyByIDRequest struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEnemyByIDRequest) Reset()         { *m = GetEnemyByIDRequest{} }
func (m *GetEnemyByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetEnemyByIDRequest) ProtoMessage()    {}
func (*GetEnemyByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ae511599e1addec, []int{1}
}

func (m *GetEnemyByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnemyByIDRequest.Unmarshal(m, b)
}
func (m *GetEnemyByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnemyByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetEnemyByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnemyByIDRequest.Merge(m, src)
}
func (m *GetEnemyByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetEnemyByIDRequest.Size(m)
}
func (m *GetEnemyByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnemyByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnemyByIDRequest proto.InternalMessageInfo

func (m *GetEnemyByIDRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetEnemyByIDResponse struct {
	Enemy                *Enemy   `protobuf:"bytes,1,opt,name=Enemy,proto3" json:"Enemy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEnemyByIDResponse) Reset()         { *m = GetEnemyByIDResponse{} }
func (m *GetEnemyByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetEnemyByIDResponse) ProtoMessage()    {}
func (*GetEnemyByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ae511599e1addec, []int{2}
}

func (m *GetEnemyByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnemyByIDResponse.Unmarshal(m, b)
}
func (m *GetEnemyByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnemyByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetEnemyByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnemyByIDResponse.Merge(m, src)
}
func (m *GetEnemyByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetEnemyByIDResponse.Size(m)
}
func (m *GetEnemyByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnemyByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnemyByIDResponse proto.InternalMessageInfo

func (m *GetEnemyByIDResponse) GetEnemy() *Enemy {
	if m != nil {
		return m.Enemy
	}
	return nil
}

// HitEnemy
type HitEnemyRequest struct {
	EnemyID              uint32   `protobuf:"varint,1,opt,name=EnemyID,proto3" json:"EnemyID,omitempty"`
	PlayerID             uint32   `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlayerPositionX      int32    `protobuf:"varint,3,opt,name=PlayerPositionX,proto3" json:"PlayerPositionX,omitempty"`
	PlayerPositionY      int32    `protobuf:"varint,4,opt,name=PlayerPositionY,proto3" json:"PlayerPositionY,omitempty"`
	BodySelection        int32    `protobuf:"varint,5,opt,name=BodySelection,proto3" json:"BodySelection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HitEnemyRequest) Reset()         { *m = HitEnemyRequest{} }
func (m *HitEnemyRequest) String() string { return proto.CompactTextString(m) }
func (*HitEnemyRequest) ProtoMessage()    {}
func (*HitEnemyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ae511599e1addec, []int{3}
}

func (m *HitEnemyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HitEnemyRequest.Unmarshal(m, b)
}
func (m *HitEnemyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HitEnemyRequest.Marshal(b, m, deterministic)
}
func (m *HitEnemyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HitEnemyRequest.Merge(m, src)
}
func (m *HitEnemyRequest) XXX_Size() int {
	return xxx_messageInfo_HitEnemyRequest.Size(m)
}
func (m *HitEnemyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HitEnemyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HitEnemyRequest proto.InternalMessageInfo

func (m *HitEnemyRequest) GetEnemyID() uint32 {
	if m != nil {
		return m.EnemyID
	}
	return 0
}

func (m *HitEnemyRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *HitEnemyRequest) GetPlayerPositionX() int32 {
	if m != nil {
		return m.PlayerPositionX
	}
	return 0
}

func (m *HitEnemyRequest) GetPlayerPositionY() int32 {
	if m != nil {
		return m.PlayerPositionY
	}
	return 0
}

func (m *HitEnemyRequest) GetBodySelection() int32 {
	if m != nil {
		return m.BodySelection
	}
	return 0
}

type HitEnemyResponse struct {
	PlayerDie            bool                 `protobuf:"varint,1,opt,name=PlayerDie,proto3" json:"PlayerDie,omitempty"`
	EnemyDie             bool                 `protobuf:"varint,2,opt,name=EnemyDie,proto3" json:"EnemyDie,omitempty"`
	DodgeAttack          bool                 `protobuf:"varint,3,opt,name=DodgeAttack,proto3" json:"DodgeAttack,omitempty"`
	PlayerDamage         int32                `protobuf:"varint,4,opt,name=PlayerDamage,proto3" json:"PlayerDamage,omitempty"`
	PlayerExperience     int32                `protobuf:"varint,5,opt,name=PlayerExperience,proto3" json:"PlayerExperience,omitempty"`
	EnemyDamage          int32                `protobuf:"varint,6,opt,name=EnemyDamage,proto3" json:"EnemyDamage,omitempty"`
	EnemyDrop            *DropTresureResponse `protobuf:"bytes,7,opt,name=EnemyDrop,proto3" json:"EnemyDrop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HitEnemyResponse) Reset()         { *m = HitEnemyResponse{} }
func (m *HitEnemyResponse) String() string { return proto.CompactTextString(m) }
func (*HitEnemyResponse) ProtoMessage()    {}
func (*HitEnemyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ae511599e1addec, []int{4}
}

func (m *HitEnemyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HitEnemyResponse.Unmarshal(m, b)
}
func (m *HitEnemyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HitEnemyResponse.Marshal(b, m, deterministic)
}
func (m *HitEnemyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HitEnemyResponse.Merge(m, src)
}
func (m *HitEnemyResponse) XXX_Size() int {
	return xxx_messageInfo_HitEnemyResponse.Size(m)
}
func (m *HitEnemyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HitEnemyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HitEnemyResponse proto.InternalMessageInfo

func (m *HitEnemyResponse) GetPlayerDie() bool {
	if m != nil {
		return m.PlayerDie
	}
	return false
}

func (m *HitEnemyResponse) GetEnemyDie() bool {
	if m != nil {
		return m.EnemyDie
	}
	return false
}

func (m *HitEnemyResponse) GetDodgeAttack() bool {
	if m != nil {
		return m.DodgeAttack
	}
	return false
}

func (m *HitEnemyResponse) GetPlayerDamage() int32 {
	if m != nil {
		return m.PlayerDamage
	}
	return 0
}

func (m *HitEnemyResponse) GetPlayerExperience() int32 {
	if m != nil {
		return m.PlayerExperience
	}
	return 0
}

func (m *HitEnemyResponse) GetEnemyDamage() int32 {
	if m != nil {
		return m.EnemyDamage
	}
	return 0
}

func (m *HitEnemyResponse) GetEnemyDrop() *DropTresureResponse {
	if m != nil {
		return m.EnemyDrop
	}
	return nil
}

func init() {
	proto.RegisterType((*Enemy)(nil), "enemy.Enemy")
	proto.RegisterType((*GetEnemyByIDRequest)(nil), "enemy.GetEnemyByIDRequest")
	proto.RegisterType((*GetEnemyByIDResponse)(nil), "enemy.GetEnemyByIDResponse")
	proto.RegisterType((*HitEnemyRequest)(nil), "enemy.HitEnemyRequest")
	proto.RegisterType((*HitEnemyResponse)(nil), "enemy.HitEnemyResponse")
}

func init() { proto.RegisterFile("rpc/enemy.proto", fileDescriptor_4ae511599e1addec) }

var fileDescriptor_4ae511599e1addec = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0xa7, 0xce, 0x9f, 0x49, 0xd3, 0x84, 0xa5, 0x87, 0x55, 0xd4, 0x83, 0x65, 0x01,
	0xb2, 0x38, 0x04, 0xa9, 0xdc, 0x7a, 0x23, 0x72, 0x05, 0x96, 0x28, 0x8a, 0x16, 0x0e, 0x84, 0x9b,
	0x71, 0x87, 0xca, 0xa2, 0xf1, 0x9a, 0xf5, 0x56, 0x6a, 0xde, 0x85, 0x97, 0xe1, 0xb1, 0xb8, 0x55,
	0x3b, 0xbb, 0xb6, 0xe3, 0xf4, 0x36, 0xdf, 0x6f, 0x26, 0x33, 0xfb, 0xcd, 0xc4, 0x30, 0x57, 0x55,
	0xfe, 0x0e, 0x4b, 0xdc, 0xed, 0x57, 0x95, 0x92, 0x5a, 0xb2, 0x80, 0xc4, 0x72, 0x61, 0xb8, 0xca,
	0x54, 0xa1, 0x5d, 0x62, 0xf9, 0xc2, 0x10, 0xad, 0xb0, 0x7e, 0x50, 0x68, 0x51, 0xf4, 0xdf, 0x83,
	0xe0, 0xda, 0x94, 0xb3, 0x33, 0xf0, 0xd3, 0x84, 0x7b, 0xa1, 0x17, 0xcf, 0x84, 0x9f, 0x26, 0x8c,
	0xc1, 0xc9, 0x97, 0x6c, 0x87, 0xdc, 0x0f, 0xbd, 0x78, 0x22, 0x28, 0x66, 0xe7, 0x10, 0xdc, 0x64,
	0x55, 0x9a, 0xf0, 0x01, 0x95, 0x59, 0xc1, 0x2e, 0x60, 0xb2, 0x91, 0x75, 0xa1, 0x0b, 0x59, 0x7e,
	0xe7, 0x27, 0xa1, 0x17, 0x07, 0xa2, 0x03, 0x87, 0xd9, 0x2d, 0x0f, 0xfa, 0xd9, 0x2d, 0x5b, 0xc2,
	0x58, 0xd0, 0x13, 0xd3, 0x84, 0x0f, 0xa9, 0x69, 0xab, 0xd9, 0x1b, 0x18, 0xda, 0x98, 0x8f, 0x42,
	0x2f, 0x9e, 0x5e, 0x9e, 0xad, 0x9c, 0x1b, 0x4b, 0x85, 0xcb, 0x32, 0x0e, 0xa3, 0xcf, 0xc5, 0x2f,
	0xbc, 0xc9, 0x1e, 0xf9, 0x98, 0x5a, 0x34, 0xd2, 0xcc, 0x36, 0xe1, 0x46, 0x16, 0xa5, 0xe6, 0x13,
	0x3b, 0xbb, 0x05, 0xd1, 0x6b, 0x78, 0xf9, 0x11, 0x35, 0xb9, 0x5f, 0xef, 0xd3, 0x44, 0xe0, 0x9f,
	0x07, 0xac, 0xf5, 0xf1, 0x22, 0xa2, 0x2b, 0x38, 0xef, 0x97, 0xd5, 0x95, 0x2c, 0x6b, 0x64, 0x91,
	0xdb, 0x1c, 0x95, 0x4e, 0x2f, 0x4f, 0x57, 0xf6, 0x06, 0xc4, 0x84, 0x4d, 0x45, 0xff, 0x3c, 0x98,
	0x7f, 0x2a, 0xec, 0x8f, 0x9b, 0xfe, 0x1c, 0x46, 0xa4, 0xdb, 0x21, 0x8d, 0x34, 0xcb, 0xd8, 0xdc,
	0x67, 0x7b, 0x54, 0x69, 0x42, 0x6b, 0x9f, 0x89, 0x56, 0xb3, 0x18, 0xe6, 0x36, 0xee, 0x56, 0x3d,
	0x20, 0x43, 0xc7, 0xf8, 0x79, 0xe5, 0xd6, 0x1d, 0xe5, 0x18, 0xb3, 0x57, 0x30, 0x5b, 0xcb, 0xdb,
	0xfd, 0x57, 0xbc, 0xc7, 0xdc, 0x10, 0x77, 0x9e, 0x3e, 0x8c, 0xfe, 0xfa, 0xb0, 0xe8, 0x3c, 0x38,
	0xf3, 0xe6, 0xaa, 0xd4, 0x2d, 0x29, 0x90, 0x6c, 0x8c, 0x45, 0x07, 0x8c, 0x11, 0x2a, 0x37, 0x49,
	0x9f, 0x92, 0xad, 0x66, 0x21, 0x4c, 0x13, 0x79, 0x7b, 0x87, 0x1f, 0xb4, 0xce, 0xf2, 0xdf, 0x64,
	0x62, 0x2c, 0x0e, 0x11, 0x8b, 0xe0, 0xd4, 0xb5, 0xca, 0x76, 0xd9, 0x1d, 0xba, 0xd7, 0xf7, 0x18,
	0x7b, 0x0b, 0x0b, 0xab, 0xaf, 0x1f, 0x2b, 0x54, 0x05, 0x96, 0x39, 0xba, 0xd7, 0x3f, 0xe3, 0x66,
	0xa2, 0x9d, 0x6e, 0xdb, 0x0d, 0xa9, 0xec, 0x10, 0xb1, 0x2b, 0x98, 0x58, 0xa9, 0x64, 0xe5, 0xfe,
	0x6c, 0x17, 0xab, 0xe6, 0x43, 0x31, 0xf0, 0x9b, 0x8d, 0x1b, 0xfb, 0xa2, 0x2b, 0x5f, 0x07, 0x3f,
	0x06, 0xaa, 0xca, 0x7f, 0x0e, 0xe9, 0x7b, 0x7a, 0xff, 0x14, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xd3,
	0x9a, 0x7e, 0x8e, 0x03, 0x00, 0x00,
}
