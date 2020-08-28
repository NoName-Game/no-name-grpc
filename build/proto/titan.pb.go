// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/titan.proto

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

type Titan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	PlanetID  uint32 `protobuf:"varint,3,opt,name=PlanetID,proto3" json:"PlanetID,omitempty"`
	LifeMax   uint32 `protobuf:"varint,8,opt,name=LifeMax,proto3" json:"LifeMax,omitempty"`
	LifePoint int32  `protobuf:"varint,9,opt,name=LifePoint,proto3" json:"LifePoint,omitempty"`
}

func (x *Titan) Reset() {
	*x = Titan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Titan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Titan) ProtoMessage() {}

func (x *Titan) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Titan.ProtoReflect.Descriptor instead.
func (*Titan) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{0}
}

func (x *Titan) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Titan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Titan) GetPlanetID() uint32 {
	if x != nil {
		return x.PlanetID
	}
	return 0
}

func (x *Titan) GetLifeMax() uint32 {
	if x != nil {
		return x.LifeMax
	}
	return 0
}

func (x *Titan) GetLifePoint() int32 {
	if x != nil {
		return x.LifePoint
	}
	return 0
}

// GetTitanByID
type GetTitanByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetTitanByIDRequest) Reset() {
	*x = GetTitanByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitanByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitanByIDRequest) ProtoMessage() {}

func (x *GetTitanByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitanByIDRequest.ProtoReflect.Descriptor instead.
func (*GetTitanByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{1}
}

func (x *GetTitanByIDRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetTitanByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titan *Titan `protobuf:"bytes,1,opt,name=Titan,proto3" json:"Titan,omitempty"`
}

func (x *GetTitanByIDResponse) Reset() {
	*x = GetTitanByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitanByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitanByIDResponse) ProtoMessage() {}

func (x *GetTitanByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitanByIDResponse.ProtoReflect.Descriptor instead.
func (*GetTitanByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{2}
}

func (x *GetTitanByIDResponse) GetTitan() *Titan {
	if x != nil {
		return x.Titan
	}
	return nil
}

// GetTitanByPlanetID
type GetTitanByPlanetIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanetID uint32 `protobuf:"varint,1,opt,name=PlanetID,proto3" json:"PlanetID,omitempty"`
}

func (x *GetTitanByPlanetIDRequest) Reset() {
	*x = GetTitanByPlanetIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitanByPlanetIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitanByPlanetIDRequest) ProtoMessage() {}

func (x *GetTitanByPlanetIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitanByPlanetIDRequest.ProtoReflect.Descriptor instead.
func (*GetTitanByPlanetIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{3}
}

func (x *GetTitanByPlanetIDRequest) GetPlanetID() uint32 {
	if x != nil {
		return x.PlanetID
	}
	return 0
}

type GetTitanByPlanetIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titan *Titan `protobuf:"bytes,1,opt,name=Titan,proto3" json:"Titan,omitempty"`
}

func (x *GetTitanByPlanetIDResponse) Reset() {
	*x = GetTitanByPlanetIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitanByPlanetIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitanByPlanetIDResponse) ProtoMessage() {}

func (x *GetTitanByPlanetIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitanByPlanetIDResponse.ProtoReflect.Descriptor instead.
func (*GetTitanByPlanetIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{4}
}

func (x *GetTitanByPlanetIDResponse) GetTitan() *Titan {
	if x != nil {
		return x.Titan
	}
	return nil
}

// GetTitanByName
type GetTitanByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetTitanByNameRequest) Reset() {
	*x = GetTitanByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitanByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitanByNameRequest) ProtoMessage() {}

func (x *GetTitanByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitanByNameRequest.ProtoReflect.Descriptor instead.
func (*GetTitanByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{5}
}

func (x *GetTitanByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTitanByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titan *Titan `protobuf:"bytes,1,opt,name=Titan,proto3" json:"Titan,omitempty"`
}

func (x *GetTitanByNameResponse) Reset() {
	*x = GetTitanByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitanByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitanByNameResponse) ProtoMessage() {}

func (x *GetTitanByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitanByNameResponse.ProtoReflect.Descriptor instead.
func (*GetTitanByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{6}
}

func (x *GetTitanByNameResponse) GetTitan() *Titan {
	if x != nil {
		return x.Titan
	}
	return nil
}

// HiTitan
type HitTitanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TitanID       uint32 `protobuf:"varint,1,opt,name=TitanID,proto3" json:"TitanID,omitempty"`
	PlayerID      uint32 `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	BodySelection int32  `protobuf:"varint,3,opt,name=BodySelection,proto3" json:"BodySelection,omitempty"`
}

func (x *HitTitanRequest) Reset() {
	*x = HitTitanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HitTitanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HitTitanRequest) ProtoMessage() {}

func (x *HitTitanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HitTitanRequest.ProtoReflect.Descriptor instead.
func (*HitTitanRequest) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{7}
}

func (x *HitTitanRequest) GetTitanID() uint32 {
	if x != nil {
		return x.TitanID
	}
	return 0
}

func (x *HitTitanRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *HitTitanRequest) GetBodySelection() int32 {
	if x != nil {
		return x.BodySelection
	}
	return 0
}

type HitTitanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerDie        bool  `protobuf:"varint,1,opt,name=PlayerDie,proto3" json:"PlayerDie,omitempty"`
	TitanDie         bool  `protobuf:"varint,2,opt,name=TitanDie,proto3" json:"TitanDie,omitempty"`
	DodgeAttack      bool  `protobuf:"varint,3,opt,name=DodgeAttack,proto3" json:"DodgeAttack,omitempty"`
	PlayerDamage     int32 `protobuf:"varint,4,opt,name=PlayerDamage,proto3" json:"PlayerDamage,omitempty"`
	PlayerExperience int32 `protobuf:"varint,5,opt,name=PlayerExperience,proto3" json:"PlayerExperience,omitempty"`
	TitanDamage      int32 `protobuf:"varint,6,opt,name=TitanDamage,proto3" json:"TitanDamage,omitempty"`
}

func (x *HitTitanResponse) Reset() {
	*x = HitTitanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HitTitanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HitTitanResponse) ProtoMessage() {}

func (x *HitTitanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HitTitanResponse.ProtoReflect.Descriptor instead.
func (*HitTitanResponse) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{8}
}

func (x *HitTitanResponse) GetPlayerDie() bool {
	if x != nil {
		return x.PlayerDie
	}
	return false
}

func (x *HitTitanResponse) GetTitanDie() bool {
	if x != nil {
		return x.TitanDie
	}
	return false
}

func (x *HitTitanResponse) GetDodgeAttack() bool {
	if x != nil {
		return x.DodgeAttack
	}
	return false
}

func (x *HitTitanResponse) GetPlayerDamage() int32 {
	if x != nil {
		return x.PlayerDamage
	}
	return 0
}

func (x *HitTitanResponse) GetPlayerExperience() int32 {
	if x != nil {
		return x.PlayerExperience
	}
	return 0
}

func (x *HitTitanResponse) GetTitanDamage() int32 {
	if x != nil {
		return x.TitanDamage
	}
	return 0
}

//TitanDiscovered
type TitanDiscoveredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TitanDiscoveredRequest) Reset() {
	*x = TitanDiscoveredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitanDiscoveredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitanDiscoveredRequest) ProtoMessage() {}

func (x *TitanDiscoveredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitanDiscoveredRequest.ProtoReflect.Descriptor instead.
func (*TitanDiscoveredRequest) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{9}
}

type TitanDiscoveredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Titans []*Titan `protobuf:"bytes,1,rep,name=Titans,proto3" json:"Titans,omitempty"`
}

func (x *TitanDiscoveredResponse) Reset() {
	*x = TitanDiscoveredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_titan_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitanDiscoveredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitanDiscoveredResponse) ProtoMessage() {}

func (x *TitanDiscoveredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_titan_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitanDiscoveredResponse.ProtoReflect.Descriptor instead.
func (*TitanDiscoveredResponse) Descriptor() ([]byte, []int) {
	return file_proto_titan_proto_rawDescGZIP(), []int{10}
}

func (x *TitanDiscoveredResponse) GetTitans() []*Titan {
	if x != nil {
		return x.Titans
	}
	return nil
}

var File_proto_titan_proto protoreflect.FileDescriptor

var file_proto_titan_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x74, 0x61, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x69, 0x74, 0x61, 0x6e, 0x22, 0x7f, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x66, 0x65, 0x4d, 0x61, 0x78, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4c, 0x69, 0x66, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x1c, 0x0a,
	0x09, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x3a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x74, 0x61,
	0x6e, 0x2e, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x52, 0x05, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x22, 0x37,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x74, 0x61, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x74, 0x61, 0x6e, 0x2e, 0x54, 0x69, 0x74,
	0x61, 0x6e, 0x52, 0x05, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x74, 0x61, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74,
	0x61, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x69, 0x74, 0x61, 0x6e, 0x2e, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x61, 0x6e, 0x22, 0x6d, 0x0a, 0x0f, 0x48, 0x69, 0x74, 0x54, 0x69, 0x74, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x74, 0x61, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a,
	0x0d, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x10, 0x48, 0x69, 0x74, 0x54, 0x69, 0x74, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x44,
	0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x44,
	0x69, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x6f, 0x64, 0x67, 0x65, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x44, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54, 0x69, 0x74, 0x61, 0x6e,
	0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x17, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x54,
	0x69, 0x74, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69,
	0x74, 0x61, 0x6e, 0x2e, 0x54, 0x69, 0x74, 0x61, 0x6e, 0x52, 0x06, 0x54, 0x69, 0x74, 0x61, 0x6e,
	0x73, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_titan_proto_rawDescOnce sync.Once
	file_proto_titan_proto_rawDescData = file_proto_titan_proto_rawDesc
)

func file_proto_titan_proto_rawDescGZIP() []byte {
	file_proto_titan_proto_rawDescOnce.Do(func() {
		file_proto_titan_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_titan_proto_rawDescData)
	})
	return file_proto_titan_proto_rawDescData
}

var file_proto_titan_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_titan_proto_goTypes = []interface{}{
	(*Titan)(nil),                      // 0: titan.Titan
	(*GetTitanByIDRequest)(nil),        // 1: titan.GetTitanByIDRequest
	(*GetTitanByIDResponse)(nil),       // 2: titan.GetTitanByIDResponse
	(*GetTitanByPlanetIDRequest)(nil),  // 3: titan.GetTitanByPlanetIDRequest
	(*GetTitanByPlanetIDResponse)(nil), // 4: titan.GetTitanByPlanetIDResponse
	(*GetTitanByNameRequest)(nil),      // 5: titan.GetTitanByNameRequest
	(*GetTitanByNameResponse)(nil),     // 6: titan.GetTitanByNameResponse
	(*HitTitanRequest)(nil),            // 7: titan.HitTitanRequest
	(*HitTitanResponse)(nil),           // 8: titan.HitTitanResponse
	(*TitanDiscoveredRequest)(nil),     // 9: titan.TitanDiscoveredRequest
	(*TitanDiscoveredResponse)(nil),    // 10: titan.TitanDiscoveredResponse
}
var file_proto_titan_proto_depIdxs = []int32{
	0, // 0: titan.GetTitanByIDResponse.Titan:type_name -> titan.Titan
	0, // 1: titan.GetTitanByPlanetIDResponse.Titan:type_name -> titan.Titan
	0, // 2: titan.GetTitanByNameResponse.Titan:type_name -> titan.Titan
	0, // 3: titan.TitanDiscoveredResponse.Titans:type_name -> titan.Titan
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_titan_proto_init() }
func file_proto_titan_proto_init() {
	if File_proto_titan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_titan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Titan); i {
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
		file_proto_titan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitanByIDRequest); i {
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
		file_proto_titan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitanByIDResponse); i {
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
		file_proto_titan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitanByPlanetIDRequest); i {
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
		file_proto_titan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitanByPlanetIDResponse); i {
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
		file_proto_titan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitanByNameRequest); i {
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
		file_proto_titan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitanByNameResponse); i {
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
		file_proto_titan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HitTitanRequest); i {
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
		file_proto_titan_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HitTitanResponse); i {
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
		file_proto_titan_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TitanDiscoveredRequest); i {
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
		file_proto_titan_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TitanDiscoveredResponse); i {
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
			RawDescriptor: file_proto_titan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_titan_proto_goTypes,
		DependencyIndexes: file_proto_titan_proto_depIdxs,
		MessageInfos:      file_proto_titan_proto_msgTypes,
	}.Build()
	File_proto_titan_proto = out.File
	file_proto_titan_proto_rawDesc = nil
	file_proto_titan_proto_goTypes = nil
	file_proto_titan_proto_depIdxs = nil
}