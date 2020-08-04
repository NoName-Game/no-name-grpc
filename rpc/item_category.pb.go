// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/item_category.proto

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

type ItemCategory struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemCategory) Reset()         { *m = ItemCategory{} }
func (m *ItemCategory) String() string { return proto.CompactTextString(m) }
func (*ItemCategory) ProtoMessage()    {}
func (*ItemCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbbf1a3ae912b0ea, []int{0}
}

func (m *ItemCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemCategory.Unmarshal(m, b)
}
func (m *ItemCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemCategory.Marshal(b, m, deterministic)
}
func (m *ItemCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemCategory.Merge(m, src)
}
func (m *ItemCategory) XXX_Size() int {
	return xxx_messageInfo_ItemCategory.Size(m)
}
func (m *ItemCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemCategory.DiscardUnknown(m)
}

var xxx_messageInfo_ItemCategory proto.InternalMessageInfo

func (m *ItemCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ItemCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ItemCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

// GetAllItemCategories
type GetAllItemCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllItemCategoriesRequest) Reset()         { *m = GetAllItemCategoriesRequest{} }
func (m *GetAllItemCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllItemCategoriesRequest) ProtoMessage()    {}
func (*GetAllItemCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbbf1a3ae912b0ea, []int{1}
}

func (m *GetAllItemCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllItemCategoriesRequest.Unmarshal(m, b)
}
func (m *GetAllItemCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllItemCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *GetAllItemCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllItemCategoriesRequest.Merge(m, src)
}
func (m *GetAllItemCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllItemCategoriesRequest.Size(m)
}
func (m *GetAllItemCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllItemCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllItemCategoriesRequest proto.InternalMessageInfo

type GetAllItemCategoriesResponse struct {
	ItemCategories       []*ItemCategory `protobuf:"bytes,1,rep,name=ItemCategories,proto3" json:"ItemCategories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllItemCategoriesResponse) Reset()         { *m = GetAllItemCategoriesResponse{} }
func (m *GetAllItemCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllItemCategoriesResponse) ProtoMessage()    {}
func (*GetAllItemCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbbf1a3ae912b0ea, []int{2}
}

func (m *GetAllItemCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllItemCategoriesResponse.Unmarshal(m, b)
}
func (m *GetAllItemCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllItemCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *GetAllItemCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllItemCategoriesResponse.Merge(m, src)
}
func (m *GetAllItemCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllItemCategoriesResponse.Size(m)
}
func (m *GetAllItemCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllItemCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllItemCategoriesResponse proto.InternalMessageInfo

func (m *GetAllItemCategoriesResponse) GetItemCategories() []*ItemCategory {
	if m != nil {
		return m.ItemCategories
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemCategory)(nil), "item_category.ItemCategory")
	proto.RegisterType((*GetAllItemCategoriesRequest)(nil), "item_category.GetAllItemCategoriesRequest")
	proto.RegisterType((*GetAllItemCategoriesResponse)(nil), "item_category.GetAllItemCategoriesResponse")
}

func init() { proto.RegisterFile("rpc/item_category.proto", fileDescriptor_cbbf1a3ae912b0ea) }

var fileDescriptor_cbbf1a3ae912b0ea = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2a, 0x48, 0xd6,
	0xcf, 0x2c, 0x49, 0xcd, 0x8d, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x45, 0x11, 0x54, 0x72, 0xe3, 0xe2, 0xf1, 0x2c, 0x49, 0xcd, 0x75,
	0x86, 0xf2, 0x85, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x98,
	0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x90, 0x58, 0x70, 0x4e, 0x69, 0xba, 0x04, 0x33, 0x44, 0x0c, 0xc4, 0x56, 0x92,
	0xe5, 0x92, 0x76, 0x4f, 0x2d, 0x71, 0xcc, 0xc9, 0x41, 0x32, 0x2d, 0x33, 0xb5, 0x38, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x29, 0x99, 0x4b, 0x06, 0xbb, 0x74, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x33, 0x17, 0x1f, 0xaa, 0x8c, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xb4, 0x1e,
	0xaa, 0x1f, 0x90, 0xdd, 0x1a, 0x84, 0xa6, 0xc5, 0x89, 0x35, 0x8a, 0xb9, 0xa8, 0x20, 0x39, 0x89,
	0x0d, 0xec, 0x51, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x04, 0x95, 0x0d, 0x03, 0x01,
	0x00, 0x00,
}
