// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/ship_stats.proto

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

type ShipStats struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Integrity            uint32   `protobuf:"varint,2,opt,name=Integrity,proto3" json:"Integrity,omitempty"`
	Tank                 float64  `protobuf:"fixed64,3,opt,name=Tank,proto3" json:"Tank,omitempty"`
	Hold                 uint32   `protobuf:"varint,4,opt,name=Hold,proto3" json:"Hold,omitempty"`
	Engine               float64  `protobuf:"fixed64,5,opt,name=Engine,proto3" json:"Engine,omitempty"`
	Speed                uint32   `protobuf:"varint,6,opt,name=Speed,proto3" json:"Speed,omitempty"`
	Radar                uint32   `protobuf:"varint,7,opt,name=Radar,proto3" json:"Radar,omitempty"`
	Attack               float64  `protobuf:"fixed64,8,opt,name=Attack,proto3" json:"Attack,omitempty"`
	Shields              float64  `protobuf:"fixed64,9,opt,name=Shields,proto3" json:"Shields,omitempty"`
	Evasion              float64  `protobuf:"fixed64,10,opt,name=Evasion,proto3" json:"Evasion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipStats) Reset()         { *m = ShipStats{} }
func (m *ShipStats) String() string { return proto.CompactTextString(m) }
func (*ShipStats) ProtoMessage()    {}
func (*ShipStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e746256f6b1b03a, []int{0}
}

func (m *ShipStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipStats.Unmarshal(m, b)
}
func (m *ShipStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipStats.Marshal(b, m, deterministic)
}
func (m *ShipStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipStats.Merge(m, src)
}
func (m *ShipStats) XXX_Size() int {
	return xxx_messageInfo_ShipStats.Size(m)
}
func (m *ShipStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipStats.DiscardUnknown(m)
}

var xxx_messageInfo_ShipStats proto.InternalMessageInfo

func (m *ShipStats) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ShipStats) GetIntegrity() uint32 {
	if m != nil {
		return m.Integrity
	}
	return 0
}

func (m *ShipStats) GetTank() float64 {
	if m != nil {
		return m.Tank
	}
	return 0
}

func (m *ShipStats) GetHold() uint32 {
	if m != nil {
		return m.Hold
	}
	return 0
}

func (m *ShipStats) GetEngine() float64 {
	if m != nil {
		return m.Engine
	}
	return 0
}

func (m *ShipStats) GetSpeed() uint32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *ShipStats) GetRadar() uint32 {
	if m != nil {
		return m.Radar
	}
	return 0
}

func (m *ShipStats) GetAttack() float64 {
	if m != nil {
		return m.Attack
	}
	return 0
}

func (m *ShipStats) GetShields() float64 {
	if m != nil {
		return m.Shields
	}
	return 0
}

func (m *ShipStats) GetEvasion() float64 {
	if m != nil {
		return m.Evasion
	}
	return 0
}

func init() {
	proto.RegisterType((*ShipStats)(nil), "ship_stats.ShipStats")
}

func init() { proto.RegisterFile("rpc/ship_stats.proto", fileDescriptor_2e746256f6b1b03a) }

var fileDescriptor_2e746256f6b1b03a = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xc9, 0xb6, 0xbb, 0x75, 0x07, 0xf4, 0x30, 0x14, 0x99, 0x83, 0x87, 0xe2, 0xa9, 0x27,
	0x3d, 0xf8, 0x04, 0x4a, 0x0b, 0xee, 0x75, 0xe3, 0xc9, 0x8b, 0xc4, 0x4d, 0xe8, 0x86, 0x96, 0x24,
	0x24, 0x41, 0xf0, 0xb1, 0x7d, 0x03, 0xc9, 0xa4, 0xa5, 0xb7, 0xf9, 0xbe, 0xff, 0xe7, 0x3f, 0x0c,
	0xac, 0x63, 0x98, 0x9e, 0xd3, 0x6c, 0xc3, 0x57, 0xca, 0x2a, 0xa7, 0xa7, 0x10, 0x7d, 0xf6, 0x08,
	0x57, 0xf3, 0xf8, 0x27, 0xa0, 0x97, 0xb3, 0x0d, 0xb2, 0x10, 0xde, 0x41, 0x33, 0xec, 0x48, 0x6c,
	0xc4, 0xf6, 0x76, 0x6c, 0x86, 0x1d, 0x3e, 0x40, 0x3f, 0xb8, 0x6c, 0x0e, 0xd1, 0xe6, 0x5f, 0x6a,
	0x58, 0x5f, 0x05, 0x22, 0x2c, 0x3f, 0x94, 0x3b, 0xd2, 0x62, 0x23, 0xb6, 0x62, 0xe4, 0xbb, 0xb8,
	0x77, 0x7f, 0xd2, 0xb4, 0xe4, 0x32, 0xdf, 0x78, 0x0f, 0xdd, 0xde, 0x1d, 0xac, 0x33, 0xd4, 0x72,
	0xf3, 0x4c, 0xb8, 0x86, 0x56, 0x06, 0x63, 0x34, 0x75, 0x5c, 0xae, 0x50, 0xec, 0xa8, 0xb4, 0x8a,
	0xb4, 0xaa, 0x96, 0xa1, 0x6c, 0xbc, 0xe6, 0xac, 0xa6, 0x23, 0xdd, 0xd4, 0x8d, 0x4a, 0x48, 0xb0,
	0x92, 0xb3, 0x35, 0x27, 0x9d, 0xa8, 0xe7, 0xe0, 0x82, 0x25, 0xd9, 0xff, 0xa8, 0x64, 0xbd, 0x23,
	0xa8, 0xc9, 0x19, 0xdf, 0xda, 0xcf, 0x45, 0x0c, 0xd3, 0x77, 0xc7, 0xdf, 0x78, 0xf9, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x0a, 0x20, 0x16, 0x80, 0x25, 0x01, 0x00, 0x00,
}