// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/npc.proto

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

type NPC struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	OpeningTime          int32    `protobuf:"varint,4,opt,name=OpeningTime,proto3" json:"OpeningTime,omitempty"`
	ClosingTime          int32    `protobuf:"varint,5,opt,name=ClosingTime,proto3" json:"ClosingTime,omitempty"`
	Enabled              bool     `protobuf:"varint,6,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NPC) Reset()         { *m = NPC{} }
func (m *NPC) String() string { return proto.CompactTextString(m) }
func (*NPC) ProtoMessage()    {}
func (*NPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{0}
}

func (m *NPC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NPC.Unmarshal(m, b)
}
func (m *NPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NPC.Marshal(b, m, deterministic)
}
func (m *NPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NPC.Merge(m, src)
}
func (m *NPC) XXX_Size() int {
	return xxx_messageInfo_NPC.Size(m)
}
func (m *NPC) XXX_DiscardUnknown() {
	xxx_messageInfo_NPC.DiscardUnknown(m)
}

var xxx_messageInfo_NPC proto.InternalMessageInfo

func (m *NPC) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *NPC) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NPC) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *NPC) GetOpeningTime() int32 {
	if m != nil {
		return m.OpeningTime
	}
	return 0
}

func (m *NPC) GetClosingTime() int32 {
	if m != nil {
		return m.ClosingTime
	}
	return 0
}

func (m *NPC) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// GetMapByID
type GetAllRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{1}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func (m *GetAllRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetAllResponse struct {
	NPCs                 []*NPC   `protobuf:"bytes,1,rep,name=NPCs,proto3" json:"NPCs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{2}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetNPCs() []*NPC {
	if m != nil {
		return m.NPCs
	}
	return nil
}

// Bank
type BankRequest struct {
	OperationType        uint32   `protobuf:"varint,1,opt,name=OperationType,proto3" json:"OperationType,omitempty"`
	PlayerID             uint32   `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Amount               int32    `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankRequest) Reset()         { *m = BankRequest{} }
func (m *BankRequest) String() string { return proto.CompactTextString(m) }
func (*BankRequest) ProtoMessage()    {}
func (*BankRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{3}
}

func (m *BankRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankRequest.Unmarshal(m, b)
}
func (m *BankRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankRequest.Marshal(b, m, deterministic)
}
func (m *BankRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankRequest.Merge(m, src)
}
func (m *BankRequest) XXX_Size() int {
	return xxx_messageInfo_BankRequest.Size(m)
}
func (m *BankRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BankRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BankRequest proto.InternalMessageInfo

func (m *BankRequest) GetOperationType() uint32 {
	if m != nil {
		return m.OperationType
	}
	return 0
}

func (m *BankRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *BankRequest) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type BankResponse struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankResponse) Reset()         { *m = BankResponse{} }
func (m *BankResponse) String() string { return proto.CompactTextString(m) }
func (*BankResponse) ProtoMessage()    {}
func (*BankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{4}
}

func (m *BankResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankResponse.Unmarshal(m, b)
}
func (m *BankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankResponse.Marshal(b, m, deterministic)
}
func (m *BankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankResponse.Merge(m, src)
}
func (m *BankResponse) XXX_Size() int {
	return xxx_messageInfo_BankResponse.Size(m)
}
func (m *BankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BankResponse proto.InternalMessageInfo

func (m *BankResponse) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Craft
type CraftRequest struct {
	CraftType            string   `protobuf:"bytes,1,opt,name=CraftType,proto3" json:"CraftType,omitempty"`
	Category             string   `protobuf:"bytes,2,opt,name=Category,proto3" json:"Category,omitempty"`
	Items                string   `protobuf:"bytes,3,opt,name=Items,proto3" json:"Items,omitempty"`
	PlayerID             uint32   `protobuf:"varint,4,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CraftRequest) Reset()         { *m = CraftRequest{} }
func (m *CraftRequest) String() string { return proto.CompactTextString(m) }
func (*CraftRequest) ProtoMessage()    {}
func (*CraftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{5}
}

func (m *CraftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CraftRequest.Unmarshal(m, b)
}
func (m *CraftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CraftRequest.Marshal(b, m, deterministic)
}
func (m *CraftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CraftRequest.Merge(m, src)
}
func (m *CraftRequest) XXX_Size() int {
	return xxx_messageInfo_CraftRequest.Size(m)
}
func (m *CraftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CraftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CraftRequest proto.InternalMessageInfo

func (m *CraftRequest) GetCraftType() string {
	if m != nil {
		return m.CraftType
	}
	return ""
}

func (m *CraftRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *CraftRequest) GetItems() string {
	if m != nil {
		return m.Items
	}
	return ""
}

func (m *CraftRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

type CraftResponse struct {
	Armor                *Armor   `protobuf:"bytes,1,opt,name=Armor,proto3" json:"Armor,omitempty"`
	Weapon               *Weapon  `protobuf:"bytes,2,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CraftResponse) Reset()         { *m = CraftResponse{} }
func (m *CraftResponse) String() string { return proto.CompactTextString(m) }
func (*CraftResponse) ProtoMessage()    {}
func (*CraftResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d06bd5f0abd0ff1, []int{6}
}

func (m *CraftResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CraftResponse.Unmarshal(m, b)
}
func (m *CraftResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CraftResponse.Marshal(b, m, deterministic)
}
func (m *CraftResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CraftResponse.Merge(m, src)
}
func (m *CraftResponse) XXX_Size() int {
	return xxx_messageInfo_CraftResponse.Size(m)
}
func (m *CraftResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CraftResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CraftResponse proto.InternalMessageInfo

func (m *CraftResponse) GetArmor() *Armor {
	if m != nil {
		return m.Armor
	}
	return nil
}

func (m *CraftResponse) GetWeapon() *Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

func init() {
	proto.RegisterType((*NPC)(nil), "npc.NPC")
	proto.RegisterType((*GetAllRequest)(nil), "npc.GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "npc.GetAllResponse")
	proto.RegisterType((*BankRequest)(nil), "npc.BankRequest")
	proto.RegisterType((*BankResponse)(nil), "npc.BankResponse")
	proto.RegisterType((*CraftRequest)(nil), "npc.CraftRequest")
	proto.RegisterType((*CraftResponse)(nil), "npc.CraftResponse")
}

func init() { proto.RegisterFile("rpc/npc.proto", fileDescriptor_6d06bd5f0abd0ff1) }

var fileDescriptor_6d06bd5f0abd0ff1 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0x87, 0x4d, 0xd3, 0xd4, 0xf6, 0xa4, 0xa9, 0x3a, 0x88, 0x84, 0xb2, 0x60, 0x08, 0x8b, 0xe4,
	0x2a, 0x42, 0x7d, 0x82, 0x6e, 0x2a, 0xd2, 0x9b, 0x6e, 0x18, 0x17, 0x05, 0xbd, 0x9a, 0xad, 0xc7,
	0x52, 0x4c, 0x66, 0xc6, 0xc9, 0x04, 0xa9, 0xaf, 0xe2, 0xcb, 0xca, 0xfc, 0x49, 0xb7, 0xbd, 0x9b,
	0xdf, 0x37, 0x87, 0x39, 0x1f, 0x67, 0x0e, 0x24, 0x4a, 0xee, 0xdf, 0x73, 0xb9, 0x2f, 0xa5, 0x12,
	0x5a, 0x90, 0x90, 0xcb, 0xfd, 0xf2, 0x85, 0x61, 0xc8, 0xb1, 0x3d, 0x39, 0xba, 0x7c, 0x65, 0x80,
	0x56, 0xd8, 0xf5, 0x0a, 0x3d, 0xb2, 0x35, 0x4c, 0xb5, 0x42, 0x79, 0xf0, 0xd2, 0x80, 0x3f, 0xc8,
	0xa4, 0xe0, 0x8e, 0xe4, 0xff, 0x02, 0x08, 0x77, 0x75, 0x45, 0x16, 0x30, 0xda, 0x6e, 0xd2, 0x20,
	0x0b, 0x8a, 0x84, 0x8e, 0xb6, 0x1b, 0x42, 0x60, 0xbc, 0x63, 0x2d, 0xa6, 0xa3, 0x2c, 0x28, 0x66,
	0xd4, 0x9e, 0x0d, 0xfb, 0xdc, 0xf4, 0x87, 0x34, 0x74, 0xcc, 0x9c, 0x49, 0x06, 0xf1, 0xbd, 0x44,
	0x7e, 0xe4, 0x87, 0x87, 0x63, 0x8b, 0xe9, 0x38, 0x0b, 0x8a, 0x88, 0x5e, 0x22, 0x53, 0x51, 0x35,
	0xa2, 0x1b, 0x2a, 0x22, 0x57, 0x71, 0x81, 0x48, 0x0a, 0xcf, 0x3f, 0x72, 0xf6, 0xd8, 0xe0, 0x8f,
	0x74, 0x92, 0x05, 0xc5, 0x94, 0x0e, 0x31, 0x7f, 0x0b, 0xc9, 0x27, 0xd4, 0xeb, 0xa6, 0xa1, 0xf8,
	0xbb, 0xc7, 0x4e, 0x5f, 0x68, 0x46, 0x46, 0x33, 0x2f, 0x61, 0x31, 0x14, 0x74, 0x52, 0xf0, 0x0e,
	0xc9, 0x0d, 0x8c, 0x77, 0x75, 0xd5, 0xa5, 0x41, 0x16, 0x16, 0xf1, 0x6a, 0x5a, 0x9a, 0xb1, 0xed,
	0xea, 0x8a, 0x5a, 0x9a, 0x1f, 0x20, 0xbe, 0x63, 0xfc, 0xd7, 0xf0, 0xdc, 0x2d, 0x24, 0xf7, 0x12,
	0x15, 0xd3, 0x47, 0xc1, 0x1f, 0x4e, 0x12, 0xfd, 0x00, 0xae, 0x21, 0x59, 0xc2, 0xb4, 0x6e, 0xd8,
	0x09, 0xd5, 0x76, 0x63, 0xe7, 0x91, 0xd0, 0x73, 0x26, 0x6f, 0x60, 0xb2, 0x6e, 0x45, 0xcf, 0xb5,
	0x9d, 0x4a, 0x44, 0x7d, 0xca, 0x6f, 0x61, 0xee, 0x1a, 0x79, 0xad, 0xd7, 0x10, 0x7d, 0x61, 0x4d,
	0x8f, 0xde, 0xdd, 0x85, 0xfc, 0x2f, 0xcc, 0x2b, 0xc5, 0x7e, 0xea, 0xc1, 0xe7, 0x06, 0x66, 0x36,
	0x9f, 0x5d, 0x66, 0xf4, 0x09, 0x18, 0x8f, 0x8a, 0x69, 0x3c, 0x08, 0x75, 0xf2, 0xff, 0x72, 0xce,
	0xe6, 0xfd, 0xad, 0xc6, 0xb6, 0xf3, 0x9f, 0xe3, 0xc2, 0x95, 0xf9, 0xf8, 0xda, 0x3c, 0xff, 0x0e,
	0x89, 0xef, 0xed, 0x15, 0x73, 0x88, 0xd6, 0x66, 0x57, 0x6c, 0xe3, 0x78, 0x35, 0x2f, 0xdd, 0xe6,
	0x58, 0x46, 0xdd, 0x15, 0x79, 0x07, 0x93, 0xaf, 0x76, 0x7d, 0xac, 0x40, 0xbc, 0x5a, 0x94, 0x7e,
	0x9b, 0x1c, 0xa5, 0xfe, 0xf6, 0x2e, 0xfa, 0x16, 0x2a, 0xb9, 0xaf, 0x9f, 0xd5, 0x41, 0x3d, 0xaa,
	0xc3, 0xc7, 0x89, 0x5d, 0xb6, 0x0f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x36, 0x8f, 0x94,
	0xc9, 0x02, 0x00, 0x00,
}
