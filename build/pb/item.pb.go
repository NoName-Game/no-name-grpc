// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/item.proto

package pb

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

	ID             uint32        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name           string        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug           string        `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	Craftable      bool          `protobuf:"varint,4,opt,name=Craftable,proto3" json:"Craftable,omitempty"`
	ItemRecipeID   uint32        `protobuf:"varint,5,opt,name=ItemRecipeID,proto3" json:"ItemRecipeID,omitempty"`
	ItemRecipe     *ItemRecipe   `protobuf:"bytes,6,opt,name=ItemRecipe,proto3" json:"ItemRecipe,omitempty"`
	RarityID       uint32        `protobuf:"varint,7,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity         *Rarity       `protobuf:"bytes,8,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ItemCategoryID uint32        `protobuf:"varint,9,opt,name=ItemCategoryID,proto3" json:"ItemCategoryID,omitempty"`
	ItemCategory   *ItemCategory `protobuf:"bytes,10,opt,name=ItemCategory,proto3" json:"ItemCategory,omitempty"`
	Value          int32         `protobuf:"varint,11,opt,name=Value,proto3" json:"Value,omitempty"`
	Price          int32         `protobuf:"varint,13,opt,name=Price,proto3" json:"Price,omitempty"`
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

func (x *Item) GetCraftable() bool {
	if x != nil {
		return x.Craftable
	}
	return false
}

func (x *Item) GetItemRecipeID() uint32 {
	if x != nil {
		return x.ItemRecipeID
	}
	return 0
}

func (x *Item) GetItemRecipe() *ItemRecipe {
	if x != nil {
		return x.ItemRecipe
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

func (x *Item) GetPrice() int32 {
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

// GetShoppableItems
type GetShoppableItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetShoppableItemsRequest) Reset() {
	*x = GetShoppableItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShoppableItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShoppableItemsRequest) ProtoMessage() {}

func (x *GetShoppableItemsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetShoppableItemsRequest.ProtoReflect.Descriptor instead.
func (*GetShoppableItemsRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{3}
}

type GetShoppableItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *GetShoppableItemsResponse) Reset() {
	*x = GetShoppableItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShoppableItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShoppableItemsResponse) ProtoMessage() {}

func (x *GetShoppableItemsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetShoppableItemsResponse.ProtoReflect.Descriptor instead.
func (*GetShoppableItemsResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{4}
}

func (x *GetShoppableItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// GetItemByID
type GetItemByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID uint32 `protobuf:"varint,1,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
}

func (x *GetItemByIDRequest) Reset() {
	*x = GetItemByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemByIDRequest) ProtoMessage() {}

func (x *GetItemByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetItemByIDRequest.ProtoReflect.Descriptor instead.
func (*GetItemByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{5}
}

func (x *GetItemByIDRequest) GetItemID() uint32 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

type GetItemByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=Item,proto3" json:"Item,omitempty"`
}

func (x *GetItemByIDResponse) Reset() {
	*x = GetItemByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemByIDResponse) ProtoMessage() {}

func (x *GetItemByIDResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetItemByIDResponse.ProtoReflect.Descriptor instead.
func (*GetItemByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{6}
}

func (x *GetItemByIDResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

// GetItemsByCategoryID
type GetItemsByCategoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID uint32 `protobuf:"varint,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
}

func (x *GetItemsByCategoryIDRequest) Reset() {
	*x = GetItemsByCategoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsByCategoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsByCategoryIDRequest) ProtoMessage() {}

func (x *GetItemsByCategoryIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[7]
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
	return file_proto_item_proto_rawDescGZIP(), []int{7}
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
		mi := &file_proto_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsByCategoryIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsByCategoryIDResponse) ProtoMessage() {}

func (x *GetItemsByCategoryIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[8]
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
	return file_proto_item_proto_rawDescGZIP(), []int{8}
}

func (x *GetItemsByCategoryIDResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// GetCraftableItemsByCategoryID
type GetCraftableItemsByCategoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID uint32 `protobuf:"varint,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
}

func (x *GetCraftableItemsByCategoryIDRequest) Reset() {
	*x = GetCraftableItemsByCategoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCraftableItemsByCategoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCraftableItemsByCategoryIDRequest) ProtoMessage() {}

func (x *GetCraftableItemsByCategoryIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCraftableItemsByCategoryIDRequest.ProtoReflect.Descriptor instead.
func (*GetCraftableItemsByCategoryIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{9}
}

func (x *GetCraftableItemsByCategoryIDRequest) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type GetCraftableItemsByCategoryIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *GetCraftableItemsByCategoryIDResponse) Reset() {
	*x = GetCraftableItemsByCategoryIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCraftableItemsByCategoryIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCraftableItemsByCategoryIDResponse) ProtoMessage() {}

func (x *GetCraftableItemsByCategoryIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCraftableItemsByCategoryIDResponse.ProtoReflect.Descriptor instead.
func (*GetCraftableItemsByCategoryIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{10}
}

func (x *GetCraftableItemsByCategoryIDResponse) GetItems() []*Item {
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
		mi := &file_proto_item_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseItemRequest) ProtoMessage() {}

func (x *UseItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[11]
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
	return file_proto_item_proto_rawDescGZIP(), []int{11}
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
		mi := &file_proto_item_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseItemResponse) ProtoMessage() {}

func (x *UseItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[12]
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
	return file_proto_item_proto_rawDescGZIP(), []int{12}
}

// BuyItem
type BuyItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	ItemID   uint32 `protobuf:"varint,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *BuyItemRequest) Reset() {
	*x = BuyItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyItemRequest) ProtoMessage() {}

func (x *BuyItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyItemRequest.ProtoReflect.Descriptor instead.
func (*BuyItemRequest) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{13}
}

func (x *BuyItemRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *BuyItemRequest) GetItemID() uint32 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *BuyItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type BuyItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyItemResponse) Reset() {
	*x = BuyItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyItemResponse) ProtoMessage() {}

func (x *BuyItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyItemResponse.ProtoReflect.Descriptor instead.
func (*BuyItemResponse) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{14}
}

var File_proto_item_proto protoreflect.FileDescriptor

var file_proto_item_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x92, 0x03, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x61, 0x66, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x43, 0x72, 0x61, 0x66, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x44, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x52, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x49, 0x74, 0x65,
	0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x70,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x22, 0x35, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x3d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42,
	0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x66,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x49, 0x0a,
	0x25, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x66, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x22, 0x11,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x60, 0x0a, 0x0e, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_item_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_item_proto_goTypes = []interface{}{
	(*Item)(nil),                                  // 0: item.Item
	(*GetAllItemsRequest)(nil),                    // 1: item.GetAllItemsRequest
	(*GetAllItemsResponse)(nil),                   // 2: item.GetAllItemsResponse
	(*GetShoppableItemsRequest)(nil),              // 3: item.GetShoppableItemsRequest
	(*GetShoppableItemsResponse)(nil),             // 4: item.GetShoppableItemsResponse
	(*GetItemByIDRequest)(nil),                    // 5: item.GetItemByIDRequest
	(*GetItemByIDResponse)(nil),                   // 6: item.GetItemByIDResponse
	(*GetItemsByCategoryIDRequest)(nil),           // 7: item.GetItemsByCategoryIDRequest
	(*GetItemsByCategoryIDResponse)(nil),          // 8: item.GetItemsByCategoryIDResponse
	(*GetCraftableItemsByCategoryIDRequest)(nil),  // 9: item.GetCraftableItemsByCategoryIDRequest
	(*GetCraftableItemsByCategoryIDResponse)(nil), // 10: item.GetCraftableItemsByCategoryIDResponse
	(*UseItemRequest)(nil),                        // 11: item.UseItemRequest
	(*UseItemResponse)(nil),                       // 12: item.UseItemResponse
	(*BuyItemRequest)(nil),                        // 13: item.BuyItemRequest
	(*BuyItemResponse)(nil),                       // 14: item.BuyItemResponse
	(*ItemRecipe)(nil),                            // 15: item_recipe.ItemRecipe
	(*Rarity)(nil),                                // 16: rarity.Rarity
	(*ItemCategory)(nil),                          // 17: item_category.ItemCategory
}
var file_proto_item_proto_depIdxs = []int32{
	15, // 0: item.Item.ItemRecipe:type_name -> item_recipe.ItemRecipe
	16, // 1: item.Item.Rarity:type_name -> rarity.Rarity
	17, // 2: item.Item.ItemCategory:type_name -> item_category.ItemCategory
	0,  // 3: item.GetAllItemsResponse.Items:type_name -> item.Item
	0,  // 4: item.GetShoppableItemsResponse.Items:type_name -> item.Item
	0,  // 5: item.GetItemByIDResponse.Item:type_name -> item.Item
	0,  // 6: item.GetItemsByCategoryIDResponse.Items:type_name -> item.Item
	0,  // 7: item.GetCraftableItemsByCategoryIDResponse.Items:type_name -> item.Item
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_item_proto_init() }
func file_proto_item_proto_init() {
	if File_proto_item_proto != nil {
		return
	}
	file_proto_item_category_proto_init()
	file_proto_rarity_proto_init()
	file_proto_item_recipe_proto_init()
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
			switch v := v.(*GetShoppableItemsRequest); i {
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
			switch v := v.(*GetShoppableItemsResponse); i {
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
			switch v := v.(*GetItemByIDRequest); i {
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
			switch v := v.(*GetItemByIDResponse); i {
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
		file_proto_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_item_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCraftableItemsByCategoryIDRequest); i {
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
		file_proto_item_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCraftableItemsByCategoryIDResponse); i {
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
		file_proto_item_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_item_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_item_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyItemRequest); i {
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
		file_proto_item_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyItemResponse); i {
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
			NumMessages:   15,
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
