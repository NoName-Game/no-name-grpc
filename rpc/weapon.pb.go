// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/weapon.proto

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

type Weapon struct {
	ID                   uint32          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	RawDamage            float32         `protobuf:"fixed32,3,opt,name=RawDamage,proto3" json:"RawDamage,omitempty"`
	PlayerID             uint32          `protobuf:"varint,4,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Equipped             bool            `protobuf:"varint,5,opt,name=Equipped,proto3" json:"Equipped,omitempty"`
	RarityID             uint32          `protobuf:"varint,6,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity               *Rarity         `protobuf:"bytes,7,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	WeaponCategoryID     uint32          `protobuf:"varint,8,opt,name=WeaponCategoryID,proto3" json:"WeaponCategoryID,omitempty"`
	WeaponCategory       *WeaponCategory `protobuf:"bytes,9,opt,name=WeaponCategory,proto3" json:"WeaponCategory,omitempty"`
	Precision            float32         `protobuf:"fixed32,10,opt,name=Precision,proto3" json:"Precision,omitempty"`
	Penetration          float32         `protobuf:"fixed32,11,opt,name=Penetration,proto3" json:"Penetration,omitempty"`
	Potential            uint32          `protobuf:"varint,12,opt,name=Potential,proto3" json:"Potential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Weapon) Reset()         { *m = Weapon{} }
func (m *Weapon) String() string { return proto.CompactTextString(m) }
func (*Weapon) ProtoMessage()    {}
func (*Weapon) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{0}
}

func (m *Weapon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weapon.Unmarshal(m, b)
}
func (m *Weapon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weapon.Marshal(b, m, deterministic)
}
func (m *Weapon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weapon.Merge(m, src)
}
func (m *Weapon) XXX_Size() int {
	return xxx_messageInfo_Weapon.Size(m)
}
func (m *Weapon) XXX_DiscardUnknown() {
	xxx_messageInfo_Weapon.DiscardUnknown(m)
}

var xxx_messageInfo_Weapon proto.InternalMessageInfo

func (m *Weapon) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Weapon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Weapon) GetRawDamage() float32 {
	if m != nil {
		return m.RawDamage
	}
	return 0
}

func (m *Weapon) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Weapon) GetEquipped() bool {
	if m != nil {
		return m.Equipped
	}
	return false
}

func (m *Weapon) GetRarityID() uint32 {
	if m != nil {
		return m.RarityID
	}
	return 0
}

func (m *Weapon) GetRarity() *Rarity {
	if m != nil {
		return m.Rarity
	}
	return nil
}

func (m *Weapon) GetWeaponCategoryID() uint32 {
	if m != nil {
		return m.WeaponCategoryID
	}
	return 0
}

func (m *Weapon) GetWeaponCategory() *WeaponCategory {
	if m != nil {
		return m.WeaponCategory
	}
	return nil
}

func (m *Weapon) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *Weapon) GetPenetration() float32 {
	if m != nil {
		return m.Penetration
	}
	return 0
}

func (m *Weapon) GetPotential() uint32 {
	if m != nil {
		return m.Potential
	}
	return 0
}

type GetWeaponByIDRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWeaponByIDRequest) Reset()         { *m = GetWeaponByIDRequest{} }
func (m *GetWeaponByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetWeaponByIDRequest) ProtoMessage()    {}
func (*GetWeaponByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{1}
}

func (m *GetWeaponByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWeaponByIDRequest.Unmarshal(m, b)
}
func (m *GetWeaponByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWeaponByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetWeaponByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWeaponByIDRequest.Merge(m, src)
}
func (m *GetWeaponByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetWeaponByIDRequest.Size(m)
}
func (m *GetWeaponByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWeaponByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWeaponByIDRequest proto.InternalMessageInfo

func (m *GetWeaponByIDRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetWeaponByIDResponse struct {
	Weapon               *Weapon  `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWeaponByIDResponse) Reset()         { *m = GetWeaponByIDResponse{} }
func (m *GetWeaponByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetWeaponByIDResponse) ProtoMessage()    {}
func (*GetWeaponByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{2}
}

func (m *GetWeaponByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWeaponByIDResponse.Unmarshal(m, b)
}
func (m *GetWeaponByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWeaponByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetWeaponByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWeaponByIDResponse.Merge(m, src)
}
func (m *GetWeaponByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetWeaponByIDResponse.Size(m)
}
func (m *GetWeaponByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWeaponByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWeaponByIDResponse proto.InternalMessageInfo

func (m *GetWeaponByIDResponse) GetWeapon() *Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

type FindWeaponByNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWeaponByNameRequest) Reset()         { *m = FindWeaponByNameRequest{} }
func (m *FindWeaponByNameRequest) String() string { return proto.CompactTextString(m) }
func (*FindWeaponByNameRequest) ProtoMessage()    {}
func (*FindWeaponByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{3}
}

func (m *FindWeaponByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWeaponByNameRequest.Unmarshal(m, b)
}
func (m *FindWeaponByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWeaponByNameRequest.Marshal(b, m, deterministic)
}
func (m *FindWeaponByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWeaponByNameRequest.Merge(m, src)
}
func (m *FindWeaponByNameRequest) XXX_Size() int {
	return xxx_messageInfo_FindWeaponByNameRequest.Size(m)
}
func (m *FindWeaponByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWeaponByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindWeaponByNameRequest proto.InternalMessageInfo

func (m *FindWeaponByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FindWeaponByNameResponse struct {
	Weapon               *Weapon  `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWeaponByNameResponse) Reset()         { *m = FindWeaponByNameResponse{} }
func (m *FindWeaponByNameResponse) String() string { return proto.CompactTextString(m) }
func (*FindWeaponByNameResponse) ProtoMessage()    {}
func (*FindWeaponByNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{4}
}

func (m *FindWeaponByNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWeaponByNameResponse.Unmarshal(m, b)
}
func (m *FindWeaponByNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWeaponByNameResponse.Marshal(b, m, deterministic)
}
func (m *FindWeaponByNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWeaponByNameResponse.Merge(m, src)
}
func (m *FindWeaponByNameResponse) XXX_Size() int {
	return xxx_messageInfo_FindWeaponByNameResponse.Size(m)
}
func (m *FindWeaponByNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWeaponByNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindWeaponByNameResponse proto.InternalMessageInfo

func (m *FindWeaponByNameResponse) GetWeapon() *Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

type UpdateWeaponRequest struct {
	Weapon               *Weapon  `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWeaponRequest) Reset()         { *m = UpdateWeaponRequest{} }
func (m *UpdateWeaponRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWeaponRequest) ProtoMessage()    {}
func (*UpdateWeaponRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{5}
}

func (m *UpdateWeaponRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWeaponRequest.Unmarshal(m, b)
}
func (m *UpdateWeaponRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWeaponRequest.Marshal(b, m, deterministic)
}
func (m *UpdateWeaponRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWeaponRequest.Merge(m, src)
}
func (m *UpdateWeaponRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWeaponRequest.Size(m)
}
func (m *UpdateWeaponRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWeaponRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWeaponRequest proto.InternalMessageInfo

func (m *UpdateWeaponRequest) GetWeapon() *Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

type UpdateWeaponResponse struct {
	Weapon               *Weapon  `protobuf:"bytes,1,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWeaponResponse) Reset()         { *m = UpdateWeaponResponse{} }
func (m *UpdateWeaponResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateWeaponResponse) ProtoMessage()    {}
func (*UpdateWeaponResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf270cb06c620467, []int{6}
}

func (m *UpdateWeaponResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWeaponResponse.Unmarshal(m, b)
}
func (m *UpdateWeaponResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWeaponResponse.Marshal(b, m, deterministic)
}
func (m *UpdateWeaponResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWeaponResponse.Merge(m, src)
}
func (m *UpdateWeaponResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateWeaponResponse.Size(m)
}
func (m *UpdateWeaponResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWeaponResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWeaponResponse proto.InternalMessageInfo

func (m *UpdateWeaponResponse) GetWeapon() *Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

func init() {
	proto.RegisterType((*Weapon)(nil), "weapon.Weapon")
	proto.RegisterType((*GetWeaponByIDRequest)(nil), "weapon.GetWeaponByIDRequest")
	proto.RegisterType((*GetWeaponByIDResponse)(nil), "weapon.GetWeaponByIDResponse")
	proto.RegisterType((*FindWeaponByNameRequest)(nil), "weapon.FindWeaponByNameRequest")
	proto.RegisterType((*FindWeaponByNameResponse)(nil), "weapon.FindWeaponByNameResponse")
	proto.RegisterType((*UpdateWeaponRequest)(nil), "weapon.UpdateWeaponRequest")
	proto.RegisterType((*UpdateWeaponResponse)(nil), "weapon.UpdateWeaponResponse")
}

func init() { proto.RegisterFile("rpc/weapon.proto", fileDescriptor_cf270cb06c620467) }

var fileDescriptor_cf270cb06c620467 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4b, 0x6b, 0xea, 0x40,
	0x14, 0xc7, 0x49, 0xd4, 0x5c, 0x3d, 0xde, 0x2b, 0x32, 0xd7, 0xcb, 0x9d, 0x4a, 0xa1, 0x21, 0x0b,
	0x09, 0x85, 0x5a, 0x68, 0xd7, 0x6d, 0xc1, 0xa6, 0x95, 0x6c, 0x8a, 0x0c, 0x94, 0x42, 0x37, 0x65,
	0x1a, 0x07, 0x09, 0x68, 0x32, 0x4e, 0x46, 0xc4, 0x8f, 0xd2, 0x6f, 0x5b, 0xe6, 0x11, 0xe3, 0x63,
	0x53, 0x77, 0x39, 0xff, 0xdf, 0x79, 0xcc, 0x79, 0x04, 0xba, 0x82, 0x27, 0xd7, 0x6b, 0x46, 0x79,
	0x9e, 0x0d, 0xb9, 0xc8, 0x65, 0x8e, 0x3c, 0x63, 0xf5, 0x35, 0x11, 0x54, 0xa4, 0x72, 0x63, 0x48,
	0xff, 0xac, 0xf2, 0xfd, 0x48, 0xa8, 0x64, 0xb3, 0x5c, 0x58, 0x14, 0x7c, 0xd5, 0xc0, 0x7b, 0xd3,
	0x04, 0x75, 0xc0, 0x8d, 0x23, 0xec, 0xf8, 0x4e, 0xf8, 0x87, 0xb8, 0x71, 0x84, 0x10, 0xd4, 0x5f,
	0xe8, 0x82, 0x61, 0xd7, 0x77, 0xc2, 0x16, 0xd1, 0xdf, 0xe8, 0x1c, 0x5a, 0x84, 0xae, 0x23, 0xba,
	0xa0, 0x33, 0x86, 0x6b, 0xbe, 0x13, 0xba, 0xa4, 0x12, 0x50, 0x1f, 0x9a, 0x93, 0x39, 0xdd, 0x30,
	0x11, 0x47, 0xb8, 0xae, 0xf3, 0x6c, 0x6d, 0xc5, 0x9e, 0x96, 0xab, 0x94, 0x73, 0x36, 0xc5, 0x0d,
	0xdf, 0x09, 0x9b, 0x64, 0x6b, 0x2b, 0x46, 0xf4, 0x7b, 0xe3, 0x08, 0x7b, 0x26, 0xae, 0xb4, 0xd1,
	0x00, 0x3c, 0xf3, 0x8d, 0x7f, 0xf9, 0x4e, 0xd8, 0xbe, 0xe9, 0x0c, 0x6d, 0x6b, 0x46, 0x25, 0x96,
	0xa2, 0x4b, 0xe8, 0x9a, 0x3e, 0x1e, 0x6d, 0x83, 0x71, 0x84, 0x9b, 0x3a, 0xd7, 0x91, 0x8e, 0xc6,
	0xd0, 0xd9, 0xd7, 0x70, 0x4b, 0xe7, 0xbe, 0x18, 0x1e, 0x0e, 0x69, 0xdf, 0x8d, 0x1c, 0x84, 0xa9,
	0x71, 0x4c, 0x04, 0x4b, 0xd2, 0x22, 0xcd, 0x33, 0x0c, 0x66, 0x1c, 0x5b, 0x01, 0xf9, 0xd0, 0x9e,
	0xb0, 0x8c, 0x49, 0x41, 0xa5, 0xe2, 0x6d, 0xcd, 0x77, 0x25, 0x1d, 0x9f, 0x4b, 0x96, 0xc9, 0x94,
	0xce, 0xf1, 0x6f, 0xfd, 0xda, 0x4a, 0x08, 0x06, 0xd0, 0x1b, 0x33, 0x69, 0x4a, 0x8e, 0x36, 0x71,
	0x44, 0xd8, 0x72, 0xc5, 0x0a, 0xb9, 0xb3, 0xa8, 0x86, 0x5a, 0x54, 0xf0, 0x00, 0xff, 0x0e, 0xfc,
	0x0a, 0x9e, 0x67, 0x05, 0x53, 0xb3, 0x33, 0xaa, 0x76, 0x56, 0xb3, 0xb3, 0x07, 0x63, 0x54, 0x62,
	0x69, 0x70, 0x05, 0xff, 0x9f, 0xd3, 0x6c, 0x5a, 0x66, 0x50, 0x9b, 0x2e, 0x6b, 0x95, 0x47, 0xe0,
	0x54, 0x47, 0x10, 0x8c, 0x00, 0x1f, 0xbb, 0x9f, 0x58, 0xf2, 0x0e, 0xfe, 0xbe, 0xf2, 0x29, 0x95,
	0xcc, 0xea, 0xb6, 0xdc, 0x4f, 0xc3, 0xef, 0xa1, 0xb7, 0x1f, 0x7e, 0x5a, 0xf9, 0x51, 0xe3, 0xbd,
	0x26, 0x78, 0xf2, 0xe9, 0xe9, 0x9f, 0xe0, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x17, 0xf1,
	0xfc, 0x4d, 0x03, 0x00, 0x00,
}