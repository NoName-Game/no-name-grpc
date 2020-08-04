// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/armor_category.proto

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

type ArmorCategory struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArmorCategory) Reset()         { *m = ArmorCategory{} }
func (m *ArmorCategory) String() string { return proto.CompactTextString(m) }
func (*ArmorCategory) ProtoMessage()    {}
func (*ArmorCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d687c3dc474ceb, []int{0}
}

func (m *ArmorCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArmorCategory.Unmarshal(m, b)
}
func (m *ArmorCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArmorCategory.Marshal(b, m, deterministic)
}
func (m *ArmorCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArmorCategory.Merge(m, src)
}
func (m *ArmorCategory) XXX_Size() int {
	return xxx_messageInfo_ArmorCategory.Size(m)
}
func (m *ArmorCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_ArmorCategory.DiscardUnknown(m)
}

var xxx_messageInfo_ArmorCategory proto.InternalMessageInfo

func (m *ArmorCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ArmorCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArmorCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GetAllArmorCategoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllArmorCategoryRequest) Reset()         { *m = GetAllArmorCategoryRequest{} }
func (m *GetAllArmorCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllArmorCategoryRequest) ProtoMessage()    {}
func (*GetAllArmorCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d687c3dc474ceb, []int{1}
}

func (m *GetAllArmorCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllArmorCategoryRequest.Unmarshal(m, b)
}
func (m *GetAllArmorCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllArmorCategoryRequest.Marshal(b, m, deterministic)
}
func (m *GetAllArmorCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllArmorCategoryRequest.Merge(m, src)
}
func (m *GetAllArmorCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllArmorCategoryRequest.Size(m)
}
func (m *GetAllArmorCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllArmorCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllArmorCategoryRequest proto.InternalMessageInfo

type GetAllArmorCategoryResponse struct {
	ArmorCategories      []*ArmorCategory `protobuf:"bytes,1,rep,name=ArmorCategories,proto3" json:"ArmorCategories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetAllArmorCategoryResponse) Reset()         { *m = GetAllArmorCategoryResponse{} }
func (m *GetAllArmorCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllArmorCategoryResponse) ProtoMessage()    {}
func (*GetAllArmorCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d687c3dc474ceb, []int{2}
}

func (m *GetAllArmorCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllArmorCategoryResponse.Unmarshal(m, b)
}
func (m *GetAllArmorCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllArmorCategoryResponse.Marshal(b, m, deterministic)
}
func (m *GetAllArmorCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllArmorCategoryResponse.Merge(m, src)
}
func (m *GetAllArmorCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllArmorCategoryResponse.Size(m)
}
func (m *GetAllArmorCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllArmorCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllArmorCategoryResponse proto.InternalMessageInfo

func (m *GetAllArmorCategoryResponse) GetArmorCategories() []*ArmorCategory {
	if m != nil {
		return m.ArmorCategories
	}
	return nil
}

func init() {
	proto.RegisterType((*ArmorCategory)(nil), "armor_category.ArmorCategory")
	proto.RegisterType((*GetAllArmorCategoryRequest)(nil), "armor_category.GetAllArmorCategoryRequest")
	proto.RegisterType((*GetAllArmorCategoryResponse)(nil), "armor_category.GetAllArmorCategoryResponse")
}

func init() { proto.RegisterFile("rpc/armor_category.proto", fileDescriptor_21d687c3dc474ceb) }

var fileDescriptor_21d687c3dc474ceb = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2a, 0x48, 0xd6,
	0x4f, 0x2c, 0xca, 0xcd, 0x2f, 0x8a, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x43, 0x15, 0x55, 0x72, 0xe7, 0xe2, 0x75, 0x04, 0x89, 0x38,
	0x43, 0x05, 0x84, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x98,
	0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x90, 0x58, 0x70, 0x4e, 0x69, 0xba, 0x04, 0x33, 0x44, 0x0c, 0xc4, 0x56, 0x92,
	0xe1, 0x92, 0x72, 0x4f, 0x2d, 0x71, 0xcc, 0xc9, 0x41, 0x31, 0x2e, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x29, 0x8d, 0x4b, 0x1a, 0xab, 0x6c, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x3b,
	0x17, 0x3f, 0xb2, 0x44, 0x66, 0x6a, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xac, 0x1e,
	0x9a, 0x2f, 0x50, 0xf5, 0xa3, 0xeb, 0x72, 0x62, 0x8d, 0x62, 0x2e, 0x2a, 0x48, 0x4e, 0x62, 0x03,
	0x7b, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x70, 0x1c, 0xc8, 0x9c, 0x08, 0x01, 0x00, 0x00,
}
