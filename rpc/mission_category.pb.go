// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/mission_category.proto

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

type MissionCategory struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissionCategory) Reset()         { *m = MissionCategory{} }
func (m *MissionCategory) String() string { return proto.CompactTextString(m) }
func (*MissionCategory) ProtoMessage()    {}
func (*MissionCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59d103818b3694d, []int{0}
}

func (m *MissionCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissionCategory.Unmarshal(m, b)
}
func (m *MissionCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissionCategory.Marshal(b, m, deterministic)
}
func (m *MissionCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissionCategory.Merge(m, src)
}
func (m *MissionCategory) XXX_Size() int {
	return xxx_messageInfo_MissionCategory.Size(m)
}
func (m *MissionCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_MissionCategory.DiscardUnknown(m)
}

var xxx_messageInfo_MissionCategory proto.InternalMessageInfo

func (m *MissionCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *MissionCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MissionCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

// Struct per la tipologia di mission
type MissionResourcesFinding struct {
	ResourceQty          uint32   `protobuf:"varint,1,opt,name=ResourceQty,proto3" json:"ResourceQty,omitempty"`
	ResourceID           uint32   `protobuf:"varint,2,opt,name=ResourceID,proto3" json:"ResourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissionResourcesFinding) Reset()         { *m = MissionResourcesFinding{} }
func (m *MissionResourcesFinding) String() string { return proto.CompactTextString(m) }
func (*MissionResourcesFinding) ProtoMessage()    {}
func (*MissionResourcesFinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59d103818b3694d, []int{1}
}

func (m *MissionResourcesFinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissionResourcesFinding.Unmarshal(m, b)
}
func (m *MissionResourcesFinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissionResourcesFinding.Marshal(b, m, deterministic)
}
func (m *MissionResourcesFinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissionResourcesFinding.Merge(m, src)
}
func (m *MissionResourcesFinding) XXX_Size() int {
	return xxx_messageInfo_MissionResourcesFinding.Size(m)
}
func (m *MissionResourcesFinding) XXX_DiscardUnknown() {
	xxx_messageInfo_MissionResourcesFinding.DiscardUnknown(m)
}

var xxx_messageInfo_MissionResourcesFinding proto.InternalMessageInfo

func (m *MissionResourcesFinding) GetResourceQty() uint32 {
	if m != nil {
		return m.ResourceQty
	}
	return 0
}

func (m *MissionResourcesFinding) GetResourceID() uint32 {
	if m != nil {
		return m.ResourceID
	}
	return 0
}

type MissionKillMob struct {
	EnemyID              uint32   `protobuf:"varint,1,opt,name=EnemyID,proto3" json:"EnemyID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissionKillMob) Reset()         { *m = MissionKillMob{} }
func (m *MissionKillMob) String() string { return proto.CompactTextString(m) }
func (*MissionKillMob) ProtoMessage()    {}
func (*MissionKillMob) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59d103818b3694d, []int{2}
}

func (m *MissionKillMob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissionKillMob.Unmarshal(m, b)
}
func (m *MissionKillMob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissionKillMob.Marshal(b, m, deterministic)
}
func (m *MissionKillMob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissionKillMob.Merge(m, src)
}
func (m *MissionKillMob) XXX_Size() int {
	return xxx_messageInfo_MissionKillMob.Size(m)
}
func (m *MissionKillMob) XXX_DiscardUnknown() {
	xxx_messageInfo_MissionKillMob.DiscardUnknown(m)
}

var xxx_messageInfo_MissionKillMob proto.InternalMessageInfo

func (m *MissionKillMob) GetEnemyID() uint32 {
	if m != nil {
		return m.EnemyID
	}
	return 0
}

type MissionPlanetFinding struct {
	PlanetID             uint32   `protobuf:"varint,1,opt,name=PlanetID,proto3" json:"PlanetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissionPlanetFinding) Reset()         { *m = MissionPlanetFinding{} }
func (m *MissionPlanetFinding) String() string { return proto.CompactTextString(m) }
func (*MissionPlanetFinding) ProtoMessage()    {}
func (*MissionPlanetFinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59d103818b3694d, []int{3}
}

func (m *MissionPlanetFinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissionPlanetFinding.Unmarshal(m, b)
}
func (m *MissionPlanetFinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissionPlanetFinding.Marshal(b, m, deterministic)
}
func (m *MissionPlanetFinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissionPlanetFinding.Merge(m, src)
}
func (m *MissionPlanetFinding) XXX_Size() int {
	return xxx_messageInfo_MissionPlanetFinding.Size(m)
}
func (m *MissionPlanetFinding) XXX_DiscardUnknown() {
	xxx_messageInfo_MissionPlanetFinding.DiscardUnknown(m)
}

var xxx_messageInfo_MissionPlanetFinding proto.InternalMessageInfo

func (m *MissionPlanetFinding) GetPlanetID() uint32 {
	if m != nil {
		return m.PlanetID
	}
	return 0
}

// GetAllMissionCategories
type GetAllMissionCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllMissionCategoriesRequest) Reset()         { *m = GetAllMissionCategoriesRequest{} }
func (m *GetAllMissionCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllMissionCategoriesRequest) ProtoMessage()    {}
func (*GetAllMissionCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59d103818b3694d, []int{4}
}

func (m *GetAllMissionCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMissionCategoriesRequest.Unmarshal(m, b)
}
func (m *GetAllMissionCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMissionCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *GetAllMissionCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMissionCategoriesRequest.Merge(m, src)
}
func (m *GetAllMissionCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllMissionCategoriesRequest.Size(m)
}
func (m *GetAllMissionCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMissionCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMissionCategoriesRequest proto.InternalMessageInfo

type GetAllMissionCategoriesResponse struct {
	MissionCategories    []*MissionCategory `protobuf:"bytes,1,rep,name=MissionCategories,proto3" json:"MissionCategories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAllMissionCategoriesResponse) Reset()         { *m = GetAllMissionCategoriesResponse{} }
func (m *GetAllMissionCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllMissionCategoriesResponse) ProtoMessage()    {}
func (*GetAllMissionCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59d103818b3694d, []int{5}
}

func (m *GetAllMissionCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMissionCategoriesResponse.Unmarshal(m, b)
}
func (m *GetAllMissionCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMissionCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *GetAllMissionCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMissionCategoriesResponse.Merge(m, src)
}
func (m *GetAllMissionCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllMissionCategoriesResponse.Size(m)
}
func (m *GetAllMissionCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMissionCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMissionCategoriesResponse proto.InternalMessageInfo

func (m *GetAllMissionCategoriesResponse) GetMissionCategories() []*MissionCategory {
	if m != nil {
		return m.MissionCategories
	}
	return nil
}

func init() {
	proto.RegisterType((*MissionCategory)(nil), "mission_category.MissionCategory")
	proto.RegisterType((*MissionResourcesFinding)(nil), "mission_category.MissionResourcesFinding")
	proto.RegisterType((*MissionKillMob)(nil), "mission_category.MissionKillMob")
	proto.RegisterType((*MissionPlanetFinding)(nil), "mission_category.MissionPlanetFinding")
	proto.RegisterType((*GetAllMissionCategoriesRequest)(nil), "mission_category.GetAllMissionCategoriesRequest")
	proto.RegisterType((*GetAllMissionCategoriesResponse)(nil), "mission_category.GetAllMissionCategoriesResponse")
}

func init() { proto.RegisterFile("rpc/mission_category.proto", fileDescriptor_b59d103818b3694d) }

var fileDescriptor_b59d103818b3694d = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xe2, 0x73, 0x4a, 0xab, 0x2e, 0x82, 0xa1, 0x87, 0x1a, 0xf7, 0x14, 0x3c, 0x54,
	0xa8, 0x7f, 0x81, 0x1a, 0x95, 0x20, 0xf5, 0xb1, 0xde, 0xf4, 0x20, 0x69, 0x1c, 0xc2, 0xc2, 0x66,
	0x37, 0xee, 0x6e, 0x0e, 0xf9, 0xef, 0xa5, 0x79, 0x51, 0x1b, 0xbc, 0xcd, 0xf7, 0xdb, 0x99, 0x6f,
	0x87, 0xf9, 0x60, 0xaa, 0x8b, 0xf4, 0x2a, 0xe7, 0xc6, 0x70, 0x25, 0xbf, 0xd2, 0xc4, 0x62, 0xa6,
	0x74, 0x35, 0x2f, 0xb4, 0xb2, 0x8a, 0x1c, 0x6f, 0x73, 0x1a, 0xc3, 0xd1, 0xb2, 0x61, 0x77, 0x2d,
	0x22, 0x13, 0x70, 0xe3, 0xc8, 0x77, 0x02, 0x27, 0x1c, 0x33, 0x37, 0x8e, 0x08, 0x81, 0x9d, 0xe7,
	0x24, 0x47, 0xdf, 0x0d, 0x9c, 0xf0, 0x90, 0xd5, 0xf5, 0x9a, 0xbd, 0x8b, 0x32, 0xf3, 0xbd, 0x86,
	0xad, 0x6b, 0xfa, 0x09, 0x67, 0xad, 0x15, 0x43, 0xa3, 0x4a, 0x9d, 0xa2, 0x79, 0xe0, 0xf2, 0x9b,
	0xcb, 0x8c, 0x04, 0x30, 0xea, 0xd8, 0x9b, 0xad, 0x5a, 0xef, 0x4d, 0x44, 0x66, 0x00, 0x9d, 0x8c,
	0xa3, 0xfa, 0xab, 0x31, 0xdb, 0x20, 0xf4, 0x12, 0x26, 0xad, 0xf9, 0x13, 0x17, 0x62, 0xa9, 0x56,
	0xc4, 0x87, 0xfd, 0x7b, 0x89, 0x79, 0xd5, 0xef, 0xda, 0x49, 0xba, 0x80, 0xd3, 0xb6, 0xf7, 0x55,
	0x24, 0x12, 0x6d, 0xb7, 0xc5, 0x14, 0x0e, 0x1a, 0xd0, 0x8f, 0xf4, 0x9a, 0x06, 0x30, 0x7b, 0x44,
	0x7b, 0x23, 0xc4, 0xdf, 0x6b, 0x70, 0x34, 0x0c, 0x7f, 0x4a, 0x34, 0x96, 0x6a, 0x38, 0xff, 0xb7,
	0xc3, 0x14, 0x4a, 0x1a, 0x24, 0x2f, 0x70, 0x32, 0x78, 0xf4, 0x9d, 0xc0, 0x0b, 0x47, 0x8b, 0x8b,
	0xf9, 0x20, 0x92, 0xad, 0xbb, 0xb3, 0xe1, 0xec, 0xed, 0xee, 0x87, 0xa7, 0x8b, 0x74, 0xb5, 0x57,
	0xa7, 0x77, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x09, 0xc5, 0xa5, 0xdb, 0x01, 0x00, 0x00,
}