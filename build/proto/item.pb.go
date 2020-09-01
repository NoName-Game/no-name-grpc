// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/item.proto

package proto

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              uint32        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name            string        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug            string        `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	RecipeID        uint32        `protobuf:"varint,4,opt,name=RecipeID,proto3" json:"RecipeID,omitempty"`
	Recipe          *Recipe       `protobuf:"bytes,5,opt,name=Recipe,proto3" json:"Recipe,omitempty"`
	RarityID        uint32        `protobuf:"varint,6,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity          *Rarity       `protobuf:"bytes,7,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ItemCategoryID  uint32        `protobuf:"varint,8,opt,name=ItemCategoryID,proto3" json:"ItemCategoryID,omitempty"`
	ItemCategory    *ItemCategory `protobuf:"bytes,9,opt,name=ItemCategory,proto3" json:"ItemCategory,omitempty"`
	Value           int32         `protobuf:"varint,10,opt,name=Value,proto3" json:"Value,omitempty"`
	MinIntellectLvl uint32        `protobuf:"varint,11,opt,name=MinIntellectLvl,proto3" json:"MinIntellectLvl,omitempty"`
	Price           uint32        `protobuf:"varint,12,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Item) GetRecipeID() uint32 {
	if x != nil {
		return x.RecipeID
	}
	return 0
}

func (x *Item) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

func (x *Item) GetRarityID() uint32 {
	if x != nil {
		return x.RarityID
	}
	return 0
}

func (x *Item) GetRarity() *Rarity {
	if x != nil {
		return x.Rarity
	}
	return nil
}

func (x *Item) GetItemCategoryID() uint32 {
	if x != nil {
		return x.ItemCategoryID
	}
	return 0
}

func (x *Item) GetItemCategory() *ItemCategory {
	if x != nil {
		return x.ItemCategory
	}
	return nil
}

func (x *Item) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Item) GetMinIntellectLvl() uint32 {
	if x != nil {
		return x.MinIntellectLvl
	}
	return 0
}

func (x *Item) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

// GetAllItems
type GetAllItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllItemsRequest) Reset() {
	*x = GetAllItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllItemsRequest) ProtoMessage() {}

func (x *GetAllItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllItemsRequest.ProtoReflect.Descriptor instead.
func (*GetAllItemsRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{1}
}

type GetAllItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *GetAllItemsResponse) Reset() {
	*x = GetAllItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllItemsResponse) ProtoMessage() {}

func (x *GetAllItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllItemsResponse.ProtoReflect.Descriptor instead.
func (*GetAllItemsResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// GetItemByCategoryID
type GetItemsByCategoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID uint32 `protobuf:"varint,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
}

func (x *GetItemsByCategoryIDRequest) Reset() {
	*x = GetItemsByCategoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsByCategoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsByCategoryIDRequest) ProtoMessage() {}

func (x *GetItemsByCategoryIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsByCategoryIDRequest.ProtoReflect.Descriptor instead.
func (*GetItemsByCategoryIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemsByCategoryIDRequest) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type GetItemsByCategoryIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *GetItemsByCategoryIDResponse) Reset() {
	*x = GetItemsByCategoryIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsByCategoryIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsByCategoryIDResponse) ProtoMessage() {}

func (x *GetItemsByCategoryIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsByCategoryIDResponse.ProtoReflect.Descriptor instead.
func (*GetItemsByCategoryIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{4}
}

func (x *GetItemsByCategoryIDResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// UseItem
type UseItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	ItemID   uint32 `protobuf:"varint,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
}

func (x *UseItemRequest) Reset() {
	*x = UseItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseItemRequest) ProtoMessage() {}

func (x *UseItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseItemRequest.ProtoReflect.Descriptor instead.
func (*UseItemRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{5}
}

func (x *UseItemRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *UseItemRequest) GetItemID() uint32 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

type UseItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UseItemResponse) Reset() {
	*x = UseItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseItemResponse) ProtoMessage() {}

func (x *UseItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseItemResponse.ProtoReflect.Descriptor instead.
func (*UseItemResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{6}
}

var File_proto_item_proto protoreflect.FileDescriptor

var file_proto_item_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x03, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06,
	0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x06, 0x52, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x49, 0x74,
	0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0c,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0c, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x4c, 0x76, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x4d, 0x69,
	0x6e, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x76, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x3d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x22, 0x40, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_item_proto_rawDescOnce sync.Once
	file_proto_item_proto_rawDescData = file_proto_item_proto_rawDesc
)

func file_proto_item_proto_rawDescGZIP() []byte {
	file_proto_item_proto_rawDescOnce.Do(func() {
		file_proto_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_item_proto_rawDescData)
	})
	return file_proto_item_proto_rawDescData
}

var file_proto_item_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_item_proto_goTypes = []interface{}{
	(*Item)(nil),                         // 0: item.Item
	(*GetAllItemsRequest)(nil),           // 1: item.GetAllItemsRequest
	(*GetAllItemsResponse)(nil),          // 2: item.GetAllItemsResponse
	(*GetItemsByCategoryIDRequest)(nil),  // 3: item.GetItemsByCategoryIDRequest
	(*GetItemsByCategoryIDResponse)(nil), // 4: item.GetItemsByCategoryIDResponse
	(*UseItemRequest)(nil),               // 5: item.UseItemRequest
	(*UseItemResponse)(nil),              // 6: item.UseItemResponse
	(*Recipe)(nil),                       // 7: recipe.Recipe
	(*Rarity)(nil),                       // 8: rarity.Rarity
	(*ItemCategory)(nil),                 // 9: item_category.ItemCategory
}
var file_proto_item_proto_depIdxs = []int32{
	7, // 0: item.Item.Recipe:type_name -> recipe.Recipe
	8, // 1: item.Item.Rarity:type_name -> rarity.Rarity
	9, // 2: item.Item.ItemCategory:type_name -> item_category.ItemCategory
	0, // 3: item.GetAllItemsResponse.Items:type_name -> item.Item
	0, // 4: item.GetItemsByCategoryIDResponse.Items:type_name -> item.Item
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_item_proto_init() }
func file_proto_item_proto_init() {
	if File_proto_item_proto != nil {
		return
	}
	file_proto_item_category_proto_init()
	file_proto_rarity_proto_init()
	file_proto_recipe_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_proto_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllItemsRequest); i {
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
		file_proto_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllItemsResponse); i {
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
		file_proto_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsByCategoryIDRequest); i {
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
		file_proto_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsByCategoryIDResponse); i {
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
		file_proto_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseItemRequest); i {
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
		file_proto_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseItemResponse); i {
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
			RawDescriptor: file_proto_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_item_proto_goTypes,
		DependencyIndexes: file_proto_item_proto_depIdxs,
		MessageInfos:      file_proto_item_proto_msgTypes,
	}.Build()
	File_proto_item_proto = out.File
	file_proto_item_proto_rawDesc = nil
	file_proto_item_proto_goTypes = nil
	file_proto_item_proto_depIdxs = nil
}
