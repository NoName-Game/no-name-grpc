// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/achievement.proto

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

type Achievement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                    uint32               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                  string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                  string               `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	Quantity              int32                `protobuf:"varint,4,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	GoldReward            int32                `protobuf:"varint,5,opt,name=GoldReward,proto3" json:"GoldReward,omitempty"`
	DiamondReward         int32                `protobuf:"varint,6,opt,name=DiamondReward,proto3" json:"DiamondReward,omitempty"`
	ExperienceReward      int32                `protobuf:"varint,7,opt,name=ExperienceReward,proto3" json:"ExperienceReward,omitempty"`
	AchievementCategoryID uint32               `protobuf:"varint,8,opt,name=AchievementCategoryID,proto3" json:"AchievementCategoryID,omitempty"`
	AchievementCategory   *AchievementCategory `protobuf:"bytes,9,opt,name=AchievementCategory,proto3" json:"AchievementCategory,omitempty"`
}

func (x *Achievement) Reset() {
	*x = Achievement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Achievement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Achievement) ProtoMessage() {}

func (x *Achievement) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Achievement.ProtoReflect.Descriptor instead.
func (*Achievement) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{0}
}

func (x *Achievement) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Achievement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Achievement) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Achievement) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Achievement) GetGoldReward() int32 {
	if x != nil {
		return x.GoldReward
	}
	return 0
}

func (x *Achievement) GetDiamondReward() int32 {
	if x != nil {
		return x.DiamondReward
	}
	return 0
}

func (x *Achievement) GetExperienceReward() int32 {
	if x != nil {
		return x.ExperienceReward
	}
	return 0
}

func (x *Achievement) GetAchievementCategoryID() uint32 {
	if x != nil {
		return x.AchievementCategoryID
	}
	return 0
}

func (x *Achievement) GetAchievementCategory() *AchievementCategory {
	if x != nil {
		return x.AchievementCategory
	}
	return nil
}

// GetAllAchievement
type GetAllAchievementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllAchievementRequest) Reset() {
	*x = GetAllAchievementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAchievementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAchievementRequest) ProtoMessage() {}

func (x *GetAllAchievementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAchievementRequest.ProtoReflect.Descriptor instead.
func (*GetAllAchievementRequest) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{1}
}

type GetAllAchievementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Achievements []*Achievement `protobuf:"bytes,1,rep,name=Achievements,proto3" json:"Achievements,omitempty"`
}

func (x *GetAllAchievementResponse) Reset() {
	*x = GetAllAchievementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAchievementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAchievementResponse) ProtoMessage() {}

func (x *GetAllAchievementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAchievementResponse.ProtoReflect.Descriptor instead.
func (*GetAllAchievementResponse) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllAchievementResponse) GetAchievements() []*Achievement {
	if x != nil {
		return x.Achievements
	}
	return nil
}

// GetAchievementForPlayerByCategory
type GetAchievementForPlayerByCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID              uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	AchievementCategoryID uint32 `protobuf:"varint,2,opt,name=AchievementCategoryID,proto3" json:"AchievementCategoryID,omitempty"`
}

func (x *GetAchievementForPlayerByCategoryRequest) Reset() {
	*x = GetAchievementForPlayerByCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAchievementForPlayerByCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAchievementForPlayerByCategoryRequest) ProtoMessage() {}

func (x *GetAchievementForPlayerByCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAchievementForPlayerByCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetAchievementForPlayerByCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{3}
}

func (x *GetAchievementForPlayerByCategoryRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *GetAchievementForPlayerByCategoryRequest) GetAchievementCategoryID() uint32 {
	if x != nil {
		return x.AchievementCategoryID
	}
	return 0
}

type GetAchievementForPlayerByCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Achievement *Achievement `protobuf:"bytes,1,opt,name=Achievement,proto3" json:"Achievement,omitempty"`
}

func (x *GetAchievementForPlayerByCategoryResponse) Reset() {
	*x = GetAchievementForPlayerByCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAchievementForPlayerByCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAchievementForPlayerByCategoryResponse) ProtoMessage() {}

func (x *GetAchievementForPlayerByCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAchievementForPlayerByCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetAchievementForPlayerByCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{4}
}

func (x *GetAchievementForPlayerByCategoryResponse) GetAchievement() *Achievement {
	if x != nil {
		return x.Achievement
	}
	return nil
}

// CheckIfPlayerHaveAchievement
type CheckIfPlayerHaveAchievementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID      uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	AchievementID uint32 `protobuf:"varint,2,opt,name=AchievementID,proto3" json:"AchievementID,omitempty"`
}

func (x *CheckIfPlayerHaveAchievementRequest) Reset() {
	*x = CheckIfPlayerHaveAchievementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfPlayerHaveAchievementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfPlayerHaveAchievementRequest) ProtoMessage() {}

func (x *CheckIfPlayerHaveAchievementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfPlayerHaveAchievementRequest.ProtoReflect.Descriptor instead.
func (*CheckIfPlayerHaveAchievementRequest) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{5}
}

func (x *CheckIfPlayerHaveAchievementRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *CheckIfPlayerHaveAchievementRequest) GetAchievementID() uint32 {
	if x != nil {
		return x.AchievementID
	}
	return 0
}

type CheckIfPlayerHaveAchievementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HaveAchievement bool         `protobuf:"varint,1,opt,name=haveAchievement,proto3" json:"haveAchievement,omitempty"`
	Achievement     *Achievement `protobuf:"bytes,2,opt,name=Achievement,proto3" json:"Achievement,omitempty"`
}

func (x *CheckIfPlayerHaveAchievementResponse) Reset() {
	*x = CheckIfPlayerHaveAchievementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_achievement_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfPlayerHaveAchievementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfPlayerHaveAchievementResponse) ProtoMessage() {}

func (x *CheckIfPlayerHaveAchievementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_achievement_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfPlayerHaveAchievementResponse.ProtoReflect.Descriptor instead.
func (*CheckIfPlayerHaveAchievementResponse) Descriptor() ([]byte, []int) {
	return file_proto_achievement_proto_rawDescGZIP(), []int{6}
}

func (x *CheckIfPlayerHaveAchievementResponse) GetHaveAchievement() bool {
	if x != nil {
		return x.HaveAchievement
	}
	return false
}

func (x *CheckIfPlayerHaveAchievementResponse) GetAchievement() *Achievement {
	if x != nil {
		return x.Achievement
	}
	return nil
}

var File_proto_achievement_proto protoreflect.FileDescriptor

var file_proto_achievement_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x0b, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x34,
	0x0a, 0x15, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x12, 0x5b, 0x0a, 0x13, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x13, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x59, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7c, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x34, 0x0a, 0x15, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x67, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0b, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x67, 0x0a, 0x23, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x48, 0x61, 0x76, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x8c, 0x01, 0x0a, 0x24, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x66, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x61, 0x76, 0x65, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61, 0x76, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x76, 0x65,
	0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_achievement_proto_rawDescOnce sync.Once
	file_proto_achievement_proto_rawDescData = file_proto_achievement_proto_rawDesc
)

func file_proto_achievement_proto_rawDescGZIP() []byte {
	file_proto_achievement_proto_rawDescOnce.Do(func() {
		file_proto_achievement_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_achievement_proto_rawDescData)
	})
	return file_proto_achievement_proto_rawDescData
}

var file_proto_achievement_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_achievement_proto_goTypes = []interface{}{
	(*Achievement)(nil),                               // 0: achievement.Achievement
	(*GetAllAchievementRequest)(nil),                  // 1: achievement.GetAllAchievementRequest
	(*GetAllAchievementResponse)(nil),                 // 2: achievement.GetAllAchievementResponse
	(*GetAchievementForPlayerByCategoryRequest)(nil),  // 3: achievement.GetAchievementForPlayerByCategoryRequest
	(*GetAchievementForPlayerByCategoryResponse)(nil), // 4: achievement.GetAchievementForPlayerByCategoryResponse
	(*CheckIfPlayerHaveAchievementRequest)(nil),       // 5: achievement.CheckIfPlayerHaveAchievementRequest
	(*CheckIfPlayerHaveAchievementResponse)(nil),      // 6: achievement.CheckIfPlayerHaveAchievementResponse
	(*AchievementCategory)(nil),                       // 7: achievement_category.AchievementCategory
}
var file_proto_achievement_proto_depIdxs = []int32{
	7, // 0: achievement.Achievement.AchievementCategory:type_name -> achievement_category.AchievementCategory
	0, // 1: achievement.GetAllAchievementResponse.Achievements:type_name -> achievement.Achievement
	0, // 2: achievement.GetAchievementForPlayerByCategoryResponse.Achievement:type_name -> achievement.Achievement
	0, // 3: achievement.CheckIfPlayerHaveAchievementResponse.Achievement:type_name -> achievement.Achievement
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_achievement_proto_init() }
func file_proto_achievement_proto_init() {
	if File_proto_achievement_proto != nil {
		return
	}
	file_proto_achievement_category_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_achievement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Achievement); i {
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
		file_proto_achievement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAchievementRequest); i {
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
		file_proto_achievement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAchievementResponse); i {
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
		file_proto_achievement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAchievementForPlayerByCategoryRequest); i {
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
		file_proto_achievement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAchievementForPlayerByCategoryResponse); i {
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
		file_proto_achievement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfPlayerHaveAchievementRequest); i {
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
		file_proto_achievement_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfPlayerHaveAchievementResponse); i {
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
			RawDescriptor: file_proto_achievement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_achievement_proto_goTypes,
		DependencyIndexes: file_proto_achievement_proto_depIdxs,
		MessageInfos:      file_proto_achievement_proto_msgTypes,
	}.Build()
	File_proto_achievement_proto = out.File
	file_proto_achievement_proto_rawDesc = nil
	file_proto_achievement_proto_goTypes = nil
	file_proto_achievement_proto_depIdxs = nil
}
