// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/item.proto

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

type Item struct {
	ID       uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug     string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	RecipeID uint32 `protobuf:"varint,4,opt,name=RecipeID,proto3" json:"RecipeID,omitempty"`
	// TODO: Recipe Recipe
	RarityID             uint32        `protobuf:"varint,5,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity               *Rarity       `protobuf:"bytes,6,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ItemCategoryID       uint32        `protobuf:"varint,7,opt,name=ItemCategoryID,proto3" json:"ItemCategoryID,omitempty"`
	ItemCategory         *ItemCategory `protobuf:"bytes,8,opt,name=ItemCategory,proto3" json:"ItemCategory,omitempty"`
	Value                int32         `protobuf:"varint,9,opt,name=Value,proto3" json:"Value,omitempty"`
	MinIntellectLvl      uint32        `protobuf:"varint,10,opt,name=MinIntellectLvl,proto3" json:"MinIntellectLvl,omitempty"`
	Price                uint32        `protobuf:"varint,11,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Item) GetRecipeID() uint32 {
	if m != nil {
		return m.RecipeID
	}
	return 0
}

func (m *Item) GetRarityID() uint32 {
	if m != nil {
		return m.RarityID
	}
	return 0
}

func (m *Item) GetRarity() *Rarity {
	if m != nil {
		return m.Rarity
	}
	return nil
}

func (m *Item) GetItemCategoryID() uint32 {
	if m != nil {
		return m.ItemCategoryID
	}
	return 0
}

func (m *Item) GetItemCategory() *ItemCategory {
	if m != nil {
		return m.ItemCategory
	}
	return nil
}

func (m *Item) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Item) GetMinIntellectLvl() uint32 {
	if m != nil {
		return m.MinIntellectLvl
	}
	return 0
}

func (m *Item) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

// GetAllItems
type GetAllItemsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllItemsRequest) Reset()         { *m = GetAllItemsRequest{} }
func (m *GetAllItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllItemsRequest) ProtoMessage()    {}
func (*GetAllItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{1}
}

func (m *GetAllItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllItemsRequest.Unmarshal(m, b)
}
func (m *GetAllItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllItemsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllItemsRequest.Merge(m, src)
}
func (m *GetAllItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllItemsRequest.Size(m)
}
func (m *GetAllItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllItemsRequest proto.InternalMessageInfo

type GetAllItemsResponse struct {
	Item                 []*Item  `protobuf:"bytes,1,rep,name=Item,proto3" json:"Item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllItemsResponse) Reset()         { *m = GetAllItemsResponse{} }
func (m *GetAllItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllItemsResponse) ProtoMessage()    {}
func (*GetAllItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{2}
}

func (m *GetAllItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllItemsResponse.Unmarshal(m, b)
}
func (m *GetAllItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllItemsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllItemsResponse.Merge(m, src)
}
func (m *GetAllItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllItemsResponse.Size(m)
}
func (m *GetAllItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllItemsResponse proto.InternalMessageInfo

func (m *GetAllItemsResponse) GetItem() []*Item {
	if m != nil {
		return m.Item
	}
	return nil
}

// GetItemByCategoryID
type GetItemByCategoryIDRequest struct {
	CategoryID           uint32   `protobuf:"varint,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemByCategoryIDRequest) Reset()         { *m = GetItemByCategoryIDRequest{} }
func (m *GetItemByCategoryIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemByCategoryIDRequest) ProtoMessage()    {}
func (*GetItemByCategoryIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{3}
}

func (m *GetItemByCategoryIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemByCategoryIDRequest.Unmarshal(m, b)
}
func (m *GetItemByCategoryIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemByCategoryIDRequest.Marshal(b, m, deterministic)
}
func (m *GetItemByCategoryIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemByCategoryIDRequest.Merge(m, src)
}
func (m *GetItemByCategoryIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemByCategoryIDRequest.Size(m)
}
func (m *GetItemByCategoryIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemByCategoryIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemByCategoryIDRequest proto.InternalMessageInfo

func (m *GetItemByCategoryIDRequest) GetCategoryID() uint32 {
	if m != nil {
		return m.CategoryID
	}
	return 0
}

type GetItemByCategoryIDResponse struct {
	Item                 []*Item  `protobuf:"bytes,1,rep,name=Item,proto3" json:"Item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemByCategoryIDResponse) Reset()         { *m = GetItemByCategoryIDResponse{} }
func (m *GetItemByCategoryIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetItemByCategoryIDResponse) ProtoMessage()    {}
func (*GetItemByCategoryIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{4}
}

func (m *GetItemByCategoryIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemByCategoryIDResponse.Unmarshal(m, b)
}
func (m *GetItemByCategoryIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemByCategoryIDResponse.Marshal(b, m, deterministic)
}
func (m *GetItemByCategoryIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemByCategoryIDResponse.Merge(m, src)
}
func (m *GetItemByCategoryIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetItemByCategoryIDResponse.Size(m)
}
func (m *GetItemByCategoryIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemByCategoryIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemByCategoryIDResponse proto.InternalMessageInfo

func (m *GetItemByCategoryIDResponse) GetItem() []*Item {
	if m != nil {
		return m.Item
	}
	return nil
}

// UseItem
type UseItemRequest struct {
	PlayerID             uint32   `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	ItemID               uint32   `protobuf:"varint,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseItemRequest) Reset()         { *m = UseItemRequest{} }
func (m *UseItemRequest) String() string { return proto.CompactTextString(m) }
func (*UseItemRequest) ProtoMessage()    {}
func (*UseItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{5}
}

func (m *UseItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseItemRequest.Unmarshal(m, b)
}
func (m *UseItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseItemRequest.Marshal(b, m, deterministic)
}
func (m *UseItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseItemRequest.Merge(m, src)
}
func (m *UseItemRequest) XXX_Size() int {
	return xxx_messageInfo_UseItemRequest.Size(m)
}
func (m *UseItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UseItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UseItemRequest proto.InternalMessageInfo

func (m *UseItemRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *UseItemRequest) GetItemID() uint32 {
	if m != nil {
		return m.ItemID
	}
	return 0
}

type UseItemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseItemResponse) Reset()         { *m = UseItemResponse{} }
func (m *UseItemResponse) String() string { return proto.CompactTextString(m) }
func (*UseItemResponse) ProtoMessage()    {}
func (*UseItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da201d68c428a6ca, []int{6}
}

func (m *UseItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseItemResponse.Unmarshal(m, b)
}
func (m *UseItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseItemResponse.Marshal(b, m, deterministic)
}
func (m *UseItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseItemResponse.Merge(m, src)
}
func (m *UseItemResponse) XXX_Size() int {
	return xxx_messageInfo_UseItemResponse.Size(m)
}
func (m *UseItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UseItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UseItemResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Item)(nil), "item.Item")
	proto.RegisterType((*GetAllItemsRequest)(nil), "item.GetAllItemsRequest")
	proto.RegisterType((*GetAllItemsResponse)(nil), "item.GetAllItemsResponse")
	proto.RegisterType((*GetItemByCategoryIDRequest)(nil), "item.GetItemByCategoryIDRequest")
	proto.RegisterType((*GetItemByCategoryIDResponse)(nil), "item.GetItemByCategoryIDResponse")
	proto.RegisterType((*UseItemRequest)(nil), "item.UseItemRequest")
	proto.RegisterType((*UseItemResponse)(nil), "item.UseItemResponse")
}

func init() { proto.RegisterFile("rpc/item.proto", fileDescriptor_da201d68c428a6ca) }

var fileDescriptor_da201d68c428a6ca = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0xeb, 0xd3, 0x40,
	0x10, 0xc6, 0x4d, 0x9a, 0xc4, 0xfe, 0xa7, 0xfe, 0x53, 0x5d, 0x8b, 0x2e, 0x29, 0x94, 0x90, 0x43,
	0xc9, 0xa9, 0x42, 0xc5, 0x9b, 0x22, 0xd6, 0x40, 0x59, 0x50, 0x09, 0x2b, 0x7a, 0xf0, 0x22, 0x31,
	0x0c, 0x25, 0xb0, 0x6d, 0xe2, 0x66, 0x2b, 0xf4, 0xc3, 0xfa, 0x5d, 0x64, 0x5f, 0x4c, 0x5f, 0xf0,
	0xe0, 0x6d, 0x9e, 0xdf, 0x93, 0x99, 0x67, 0xb2, 0xbb, 0x10, 0xcb, 0xae, 0x7e, 0xd1, 0x28, 0xdc,
	0xaf, 0x3a, 0xd9, 0xaa, 0x96, 0x04, 0xba, 0x4e, 0x9e, 0xff, 0xa5, 0xdf, 0xeb, 0x4a, 0xe1, 0xae,
	0x95, 0x27, 0x6b, 0x27, 0x8f, 0xb5, 0x21, 0x2b, 0xd9, 0x28, 0x47, 0xb2, 0xdf, 0x3e, 0x04, 0x4c,
	0xe1, 0x9e, 0xc4, 0xe0, 0xb3, 0x82, 0x7a, 0xa9, 0x97, 0xdf, 0x73, 0x9f, 0x15, 0x84, 0x40, 0xf0,
	0xa9, 0xda, 0x23, 0xf5, 0x53, 0x2f, 0xbf, 0xe3, 0xa6, 0xd6, 0xec, 0xb3, 0x38, 0xee, 0xe8, 0xc8,
	0x32, 0x5d, 0x93, 0x04, 0xc6, 0x1c, 0xeb, 0xa6, 0x43, 0x56, 0xd0, 0xc0, 0x74, 0x0f, 0xda, 0x78,
	0x26, 0x8c, 0x15, 0x34, 0x74, 0x9e, 0xd3, 0x64, 0x09, 0x91, 0xad, 0x69, 0x94, 0x7a, 0xf9, 0x64,
	0x1d, 0xaf, 0xdc, 0x5e, 0x96, 0x72, 0xe7, 0x92, 0x25, 0xc4, 0x7a, 0xbf, 0xf7, 0xee, 0x47, 0x58,
	0x41, 0x1f, 0x9a, 0x49, 0x37, 0x94, 0xbc, 0x85, 0x47, 0x97, 0x84, 0x8e, 0xcd, 0xd4, 0xf9, 0xea,
	0xfa, 0x18, 0x2e, 0x3f, 0xe1, 0x57, 0x0d, 0x64, 0x06, 0xe1, 0xd7, 0x4a, 0x1c, 0x91, 0xde, 0xa5,
	0x5e, 0x1e, 0x72, 0x2b, 0x48, 0x0e, 0xd3, 0x8f, 0xcd, 0x81, 0x1d, 0x14, 0x0a, 0x81, 0xb5, 0xfa,
	0xf0, 0x4b, 0x50, 0x30, 0xf9, 0xb7, 0x58, 0xf7, 0x97, 0xb2, 0xa9, 0x91, 0x4e, 0x8c, 0x6f, 0x45,
	0x36, 0x03, 0xb2, 0x45, 0xf5, 0x4e, 0x08, 0x9d, 0xd5, 0x73, 0xfc, 0x79, 0xc4, 0x5e, 0x65, 0xaf,
	0xe0, 0xe9, 0x15, 0xed, 0xbb, 0xf6, 0xd0, 0x23, 0x59, 0xd8, 0xbb, 0xa0, 0x5e, 0x3a, 0xca, 0x27,
	0x6b, 0x30, 0xbb, 0x9b, 0x95, 0xb9, 0xe1, 0xd9, 0x6b, 0x48, 0xb6, 0xa8, 0x74, 0xb9, 0x39, 0x9d,
	0x7f, 0xdd, 0x0d, 0x25, 0x0b, 0x80, 0xfa, 0x7c, 0x4a, 0xf6, 0x26, 0x2f, 0x48, 0xf6, 0x06, 0xe6,
	0xff, 0xec, 0xfe, 0xcf, 0xf0, 0x02, 0xe2, 0x2f, 0x3d, 0x1a, 0xe0, 0x02, 0x13, 0x18, 0x97, 0xa2,
	0x3a, 0xa1, 0x1c, 0xe2, 0x06, 0x4d, 0x9e, 0x41, 0xa4, 0x3f, 0x65, 0x85, 0x79, 0x40, 0xf7, 0xdc,
	0xa9, 0xec, 0x09, 0x4c, 0x87, 0x29, 0x36, 0x78, 0x13, 0x7e, 0x1b, 0xc9, 0xae, 0x2e, 0x1f, 0x94,
	0xde, 0x8f, 0xc8, 0x3c, 0xca, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x30, 0xf5, 0xc7,
	0xd7, 0x02, 0x00, 0x00,
}
