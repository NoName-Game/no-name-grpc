// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/resource.proto

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

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 uint32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string            `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	RarityID           uint32            `protobuf:"varint,3,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity             *Rarity           `protobuf:"bytes,4,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ResourceCategoryID uint32            `protobuf:"varint,5,opt,name=ResourceCategoryID,proto3" json:"ResourceCategoryID,omitempty"`
	ResourceCategory   *ResourceCategory `protobuf:"bytes,6,opt,name=ResourceCategory,proto3" json:"ResourceCategory,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetRarityID() uint32 {
	if x != nil {
		return x.RarityID
	}
	return 0
}

func (x *Resource) GetRarity() *Rarity {
	if x != nil {
		return x.Rarity
	}
	return nil
}

func (x *Resource) GetResourceCategoryID() uint32 {
	if x != nil {
		return x.ResourceCategoryID
	}
	return 0
}

func (x *Resource) GetResourceCategory() *ResourceCategory {
	if x != nil {
		return x.ResourceCategory
	}
	return nil
}

// GetResourceByID
type GetResourceByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetResourceByIDRequest) Reset() {
	*x = GetResourceByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByIDRequest) ProtoMessage() {}

func (x *GetResourceByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByIDRequest.ProtoReflect.Descriptor instead.
func (*GetResourceByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourceByIDRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetResourceByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=Resource,proto3" json:"Resource,omitempty"`
}

func (x *GetResourceByIDResponse) Reset() {
	*x = GetResourceByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByIDResponse) ProtoMessage() {}

func (x *GetResourceByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByIDResponse.ProtoReflect.Descriptor instead.
func (*GetResourceByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{2}
}

func (x *GetResourceByIDResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// GetResourceByName
type GetResourceByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetResourceByNameRequest) Reset() {
	*x = GetResourceByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByNameRequest) ProtoMessage() {}

func (x *GetResourceByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByNameRequest.ProtoReflect.Descriptor instead.
func (*GetResourceByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{3}
}

func (x *GetResourceByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetResourceByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=Resource,proto3" json:"Resource,omitempty"`
}

func (x *GetResourceByNameResponse) Reset() {
	*x = GetResourceByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByNameResponse) ProtoMessage() {}

func (x *GetResourceByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByNameResponse.ProtoReflect.Descriptor instead.
func (*GetResourceByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourceByNameResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// DropResource
type DropResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeExploration string `protobuf:"bytes,1,opt,name=TypeExploration,proto3" json:"TypeExploration,omitempty"`
	QtyExploration  int32  `protobuf:"varint,2,opt,name=QtyExploration,proto3" json:"QtyExploration,omitempty"`
	PlayerID        uint32 `protobuf:"varint,3,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlanetID        uint32 `protobuf:"varint,4,opt,name=PlanetID,proto3" json:"PlanetID,omitempty"`
}

func (x *DropResourceRequest) Reset() {
	*x = DropResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropResourceRequest) ProtoMessage() {}

func (x *DropResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropResourceRequest.ProtoReflect.Descriptor instead.
func (*DropResourceRequest) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{5}
}

func (x *DropResourceRequest) GetTypeExploration() string {
	if x != nil {
		return x.TypeExploration
	}
	return ""
}

func (x *DropResourceRequest) GetQtyExploration() int32 {
	if x != nil {
		return x.QtyExploration
	}
	return 0
}

func (x *DropResourceRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *DropResourceRequest) GetPlanetID() uint32 {
	if x != nil {
		return x.PlanetID
	}
	return 0
}

type DropResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Quantity int32     `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *DropResourceResponse) Reset() {
	*x = DropResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropResourceResponse) ProtoMessage() {}

func (x *DropResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropResourceResponse.ProtoReflect.Descriptor instead.
func (*DropResourceResponse) Descriptor() ([]byte, []int) {
	return file_proto_resource_proto_rawDescGZIP(), []int{6}
}

func (x *DropResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *DropResourceResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_proto_resource_proto protoreflect.FileDescriptor

var file_proto_resource_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44,
	0x12, 0x26, 0x0a, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x06, 0x52, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x4f, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2e,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x13,
	0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0e, 0x51, 0x74, 0x79, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x51, 0x74, 0x79, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x22, 0x62, 0x0a,
	0x14, 0x44, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_resource_proto_rawDescOnce sync.Once
	file_proto_resource_proto_rawDescData = file_proto_resource_proto_rawDesc
)

func file_proto_resource_proto_rawDescGZIP() []byte {
	file_proto_resource_proto_rawDescOnce.Do(func() {
		file_proto_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_resource_proto_rawDescData)
	})
	return file_proto_resource_proto_rawDescData
}

var file_proto_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_resource_proto_goTypes = []interface{}{
	(*Resource)(nil),                  // 0: resource.Resource
	(*GetResourceByIDRequest)(nil),    // 1: resource.GetResourceByIDRequest
	(*GetResourceByIDResponse)(nil),   // 2: resource.GetResourceByIDResponse
	(*GetResourceByNameRequest)(nil),  // 3: resource.GetResourceByNameRequest
	(*GetResourceByNameResponse)(nil), // 4: resource.GetResourceByNameResponse
	(*DropResourceRequest)(nil),       // 5: resource.DropResourceRequest
	(*DropResourceResponse)(nil),      // 6: resource.DropResourceResponse
	(*Rarity)(nil),                    // 7: rarity.Rarity
	(*ResourceCategory)(nil),          // 8: resource_category.ResourceCategory
}
var file_proto_resource_proto_depIdxs = []int32{
	7, // 0: resource.Resource.Rarity:type_name -> rarity.Rarity
	8, // 1: resource.Resource.ResourceCategory:type_name -> resource_category.ResourceCategory
	0, // 2: resource.GetResourceByIDResponse.Resource:type_name -> resource.Resource
	0, // 3: resource.GetResourceByNameResponse.Resource:type_name -> resource.Resource
	0, // 4: resource.DropResourceResponse.Resource:type_name -> resource.Resource
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_resource_proto_init() }
func file_proto_resource_proto_init() {
	if File_proto_resource_proto != nil {
		return
	}
	file_proto_resource_category_proto_init()
	file_proto_rarity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_proto_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByIDRequest); i {
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
		file_proto_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByIDResponse); i {
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
		file_proto_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByNameRequest); i {
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
		file_proto_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByNameResponse); i {
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
		file_proto_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropResourceRequest); i {
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
		file_proto_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropResourceResponse); i {
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
			RawDescriptor: file_proto_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_resource_proto_goTypes,
		DependencyIndexes: file_proto_resource_proto_depIdxs,
		MessageInfos:      file_proto_resource_proto_msgTypes,
	}.Build()
	File_proto_resource_proto = out.File
	file_proto_resource_proto_rawDesc = nil
	file_proto_resource_proto_goTypes = nil
	file_proto_resource_proto_depIdxs = nil
}
