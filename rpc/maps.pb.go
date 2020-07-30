// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/maps.proto

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

type Maps struct {
	ID                   uint32     `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CellGrid             string     `protobuf:"bytes,2,opt,name=CellGrid,proto3" json:"CellGrid,omitempty"`
	StartPositionX       int32      `protobuf:"varint,3,opt,name=StartPositionX,proto3" json:"StartPositionX,omitempty"`
	StartPositionY       int32      `protobuf:"varint,4,opt,name=StartPositionY,proto3" json:"StartPositionY,omitempty"`
	Enemies              []*Enemy   `protobuf:"bytes,5,rep,name=Enemies,proto3" json:"Enemies,omitempty"`
	Tresures             []*Tresure `protobuf:"bytes,6,rep,name=Tresures,proto3" json:"Tresures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Maps) Reset()         { *m = Maps{} }
func (m *Maps) String() string { return proto.CompactTextString(m) }
func (*Maps) ProtoMessage()    {}
func (*Maps) Descriptor() ([]byte, []int) {
	return fileDescriptor_d024716cd2f0ef93, []int{0}
}

func (m *Maps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maps.Unmarshal(m, b)
}
func (m *Maps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maps.Marshal(b, m, deterministic)
}
func (m *Maps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maps.Merge(m, src)
}
func (m *Maps) XXX_Size() int {
	return xxx_messageInfo_Maps.Size(m)
}
func (m *Maps) XXX_DiscardUnknown() {
	xxx_messageInfo_Maps.DiscardUnknown(m)
}

var xxx_messageInfo_Maps proto.InternalMessageInfo

func (m *Maps) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Maps) GetCellGrid() string {
	if m != nil {
		return m.CellGrid
	}
	return ""
}

func (m *Maps) GetStartPositionX() int32 {
	if m != nil {
		return m.StartPositionX
	}
	return 0
}

func (m *Maps) GetStartPositionY() int32 {
	if m != nil {
		return m.StartPositionY
	}
	return 0
}

func (m *Maps) GetEnemies() []*Enemy {
	if m != nil {
		return m.Enemies
	}
	return nil
}

func (m *Maps) GetTresures() []*Tresure {
	if m != nil {
		return m.Tresures
	}
	return nil
}

// GetMapByID
type GetMapByIDRequest struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMapByIDRequest) Reset()         { *m = GetMapByIDRequest{} }
func (m *GetMapByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetMapByIDRequest) ProtoMessage()    {}
func (*GetMapByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d024716cd2f0ef93, []int{1}
}

func (m *GetMapByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMapByIDRequest.Unmarshal(m, b)
}
func (m *GetMapByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMapByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetMapByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMapByIDRequest.Merge(m, src)
}
func (m *GetMapByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetMapByIDRequest.Size(m)
}
func (m *GetMapByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMapByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMapByIDRequest proto.InternalMessageInfo

func (m *GetMapByIDRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetMapByIDResponse struct {
	Maps                 *Maps    `protobuf:"bytes,1,opt,name=Maps,proto3" json:"Maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMapByIDResponse) Reset()         { *m = GetMapByIDResponse{} }
func (m *GetMapByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetMapByIDResponse) ProtoMessage()    {}
func (*GetMapByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d024716cd2f0ef93, []int{2}
}

func (m *GetMapByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMapByIDResponse.Unmarshal(m, b)
}
func (m *GetMapByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMapByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetMapByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMapByIDResponse.Merge(m, src)
}
func (m *GetMapByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetMapByIDResponse.Size(m)
}
func (m *GetMapByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMapByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMapByIDResponse proto.InternalMessageInfo

func (m *GetMapByIDResponse) GetMaps() *Maps {
	if m != nil {
		return m.Maps
	}
	return nil
}

func init() {
	proto.RegisterType((*Maps)(nil), "maps.Maps")
	proto.RegisterType((*GetMapByIDRequest)(nil), "maps.GetMapByIDRequest")
	proto.RegisterType((*GetMapByIDResponse)(nil), "maps.GetMapByIDResponse")
}

func init() { proto.RegisterFile("rpc/maps.proto", fileDescriptor_d024716cd2f0ef93) }

var fileDescriptor_d024716cd2f0ef93 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x71, 0x7e, 0x4a, 0xb9, 0x85, 0x40, 0x3d, 0x59, 0x19, 0x50, 0x14, 0xa4, 0x2a, 0x03,
	0x0a, 0x52, 0xe1, 0x09, 0x4a, 0x50, 0x95, 0xa1, 0x52, 0x64, 0x18, 0x28, 0x5b, 0x28, 0x77, 0x88,
	0xd4, 0x24, 0xc6, 0x76, 0x87, 0xbc, 0x29, 0x8f, 0x83, 0x6c, 0xd3, 0x0a, 0x15, 0x36, 0xfb, 0xbb,
	0xdf, 0x70, 0xce, 0x81, 0x48, 0x8a, 0xcd, 0x5d, 0x5b, 0x0b, 0x95, 0x0b, 0xd9, 0xeb, 0x9e, 0x06,
	0xe6, 0x1d, 0x5f, 0x1a, 0x8a, 0x1d, 0xb6, 0x83, 0xc3, 0xf1, 0xd4, 0x00, 0x2d, 0x51, 0xed, 0x24,
	0x3a, 0x94, 0x7e, 0x11, 0x08, 0x56, 0xb5, 0x50, 0x34, 0x02, 0xaf, 0x2c, 0x18, 0x49, 0x48, 0x76,
	0xc1, 0xbd, 0xb2, 0xa0, 0x31, 0x8c, 0x1f, 0x71, 0xbb, 0x5d, 0xca, 0xe6, 0x83, 0x79, 0x09, 0xc9,
	0xce, 0xf8, 0xe1, 0x4f, 0x67, 0x10, 0x3d, 0xeb, 0x5a, 0xea, 0xaa, 0x57, 0x8d, 0x6e, 0xfa, 0xee,
	0x95, 0xf9, 0x09, 0xc9, 0x42, 0x7e, 0x44, 0xff, 0x78, 0x6b, 0x16, 0xfc, 0xe3, 0xad, 0xe9, 0x0c,
	0x4e, 0x9f, 0x3a, 0x6c, 0x1b, 0x54, 0x2c, 0x4c, 0xfc, 0x6c, 0x32, 0x3f, 0xcf, 0x5d, 0x6c, 0x43,
	0x07, 0xbe, 0x3f, 0xd2, 0x5b, 0x18, 0xbf, 0xb8, 0xf4, 0x8a, 0x8d, 0xac, 0x78, 0x95, 0xef, 0xeb,
	0xfc, 0x1c, 0xf8, 0xc1, 0x48, 0x6f, 0x60, 0xba, 0x44, 0xbd, 0xaa, 0xc5, 0x62, 0x28, 0x0b, 0x8e,
	0x9f, 0x3b, 0x54, 0xfa, 0xb8, 0x66, 0xfa, 0x00, 0xf4, 0xb7, 0xa4, 0x44, 0xdf, 0x29, 0xa4, 0xd7,
	0x6e, 0x14, 0xeb, 0x4d, 0xe6, 0x90, 0xdb, 0x69, 0x0d, 0xe1, 0x96, 0x2f, 0xc2, 0x37, 0x5f, 0x8a,
	0x4d, 0x75, 0x52, 0x91, 0xf7, 0x91, 0xdd, 0xf1, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x04, 0x5a,
	0xfc, 0xc1, 0x83, 0x01, 0x00, 0x00,
}
