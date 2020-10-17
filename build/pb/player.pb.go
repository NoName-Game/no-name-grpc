// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/player.proto

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username   string    `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	LevelID    uint32    `protobuf:"varint,3,opt,name=LevelID,proto3" json:"LevelID,omitempty"`
	Level      *Level    `protobuf:"bytes,4,opt,name=Level,proto3" json:"Level,omitempty"`
	LifePoint  int64     `protobuf:"varint,5,opt,name=LifePoint,proto3" json:"LifePoint,omitempty"`
	Dead       bool      `protobuf:"varint,6,opt,name=Dead,proto3" json:"Dead,omitempty"`
	Tutorial   bool      `protobuf:"varint,7,opt,name=Tutorial,proto3" json:"Tutorial,omitempty"`
	Banned     bool      `protobuf:"varint,8,opt,name=Banned,proto3" json:"Banned,omitempty"`
	ChatID     int64     `protobuf:"varint,9,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	LanguageID uint32    `protobuf:"varint,10,opt,name=LanguageID,proto3" json:"LanguageID,omitempty"`
	Language   *Language `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Player) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Player) GetLevelID() uint32 {
	if x != nil {
		return x.LevelID
	}
	return 0
}

func (x *Player) GetLevel() *Level {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *Player) GetLifePoint() int64 {
	if x != nil {
		return x.LifePoint
	}
	return 0
}

func (x *Player) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

func (x *Player) GetTutorial() bool {
	if x != nil {
		return x.Tutorial
	}
	return false
}

func (x *Player) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

func (x *Player) GetChatID() int64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

func (x *Player) GetLanguageID() uint32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *Player) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

// GetPlayerByID
type GetPlayerByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetPlayerByIDRequest) Reset() {
	*x = GetPlayerByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerByIDRequest) ProtoMessage() {}

func (x *GetPlayerByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerByIDRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlayerByIDRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetPlayerByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (x *GetPlayerByIDResponse) Reset() {
	*x = GetPlayerByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerByIDResponse) ProtoMessage() {}

func (x *GetPlayerByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerByIDResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlayerByIDResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

// FindPlayerByUsername
type GetPlayerByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *GetPlayerByUsernameRequest) Reset() {
	*x = GetPlayerByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerByUsernameRequest) ProtoMessage() {}

func (x *GetPlayerByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetPlayerByUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (x *GetPlayerByUsernameResponse) Reset() {
	*x = GetPlayerByUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerByUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerByUsernameResponse) ProtoMessage() {}

func (x *GetPlayerByUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerByUsernameResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerByUsernameResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerByUsernameResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

// SignIn
type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	ChatID     int64  `protobuf:"varint,2,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	LanguageID uint32 `protobuf:"varint,3,opt,name=LanguageID,proto3" json:"LanguageID,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{5}
}

func (x *SignInRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignInRequest) GetChatID() int64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

func (x *SignInRequest) GetLanguageID() uint32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{6}
}

func (x *SignInResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

// GetPlayerExperience
type GetPlayerExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerExperienceRequest) Reset() {
	*x = GetPlayerExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerExperienceRequest) ProtoMessage() {}

func (x *GetPlayerExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerExperienceRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerExperienceRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlayerExperienceRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerExperienceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetPlayerExperienceResponse) Reset() {
	*x = GetPlayerExperienceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerExperienceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerExperienceResponse) ProtoMessage() {}

func (x *GetPlayerExperienceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerExperienceResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerExperienceResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{8}
}

func (x *GetPlayerExperienceResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// GetPlayerLifePoint
type GetPlayerLifePointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerLifePointRequest) Reset() {
	*x = GetPlayerLifePointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerLifePointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerLifePointRequest) ProtoMessage() {}

func (x *GetPlayerLifePointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerLifePointRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerLifePointRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{9}
}

func (x *GetPlayerLifePointRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerLifePointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetPlayerLifePointResponse) Reset() {
	*x = GetPlayerLifePointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerLifePointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerLifePointResponse) ProtoMessage() {}

func (x *GetPlayerLifePointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerLifePointResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerLifePointResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{10}
}

func (x *GetPlayerLifePointResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// GetPlayerDailyReward
type GetPlayerDailyRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *GetPlayerDailyRewardRequest) Reset() {
	*x = GetPlayerDailyRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerDailyRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerDailyRewardRequest) ProtoMessage() {}

func (x *GetPlayerDailyRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerDailyRewardRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerDailyRewardRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{11}
}

func (x *GetPlayerDailyRewardRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type GetPlayerDailyRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experience int32 `protobuf:"varint,1,opt,name=Experience,proto3" json:"Experience,omitempty"`
	Money      int32 `protobuf:"varint,2,opt,name=Money,proto3" json:"Money,omitempty"`
	Diamond    int32 `protobuf:"varint,3,opt,name=Diamond,proto3" json:"Diamond,omitempty"`
}

func (x *GetPlayerDailyRewardResponse) Reset() {
	*x = GetPlayerDailyRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerDailyRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerDailyRewardResponse) ProtoMessage() {}

func (x *GetPlayerDailyRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerDailyRewardResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerDailyRewardResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{12}
}

func (x *GetPlayerDailyRewardResponse) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *GetPlayerDailyRewardResponse) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *GetPlayerDailyRewardResponse) GetDiamond() int32 {
	if x != nil {
		return x.Diamond
	}
	return 0
}

var File_proto_player_proto protoreflect.FileDescriptor

var file_proto_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x14, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69,
	0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4c,
	0x69, 0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x61, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x22, 0x38, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x63, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x22, 0x38, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x37, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x66,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x6e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_player_proto_rawDescOnce sync.Once
	file_proto_player_proto_rawDescData = file_proto_player_proto_rawDesc
)

func file_proto_player_proto_rawDescGZIP() []byte {
	file_proto_player_proto_rawDescOnce.Do(func() {
		file_proto_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_proto_rawDescData)
	})
	return file_proto_player_proto_rawDescData
}

var file_proto_player_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_player_proto_goTypes = []interface{}{
	(*Player)(nil),                       // 0: player.Player
	(*GetPlayerByIDRequest)(nil),         // 1: player.GetPlayerByIDRequest
	(*GetPlayerByIDResponse)(nil),        // 2: player.GetPlayerByIDResponse
	(*GetPlayerByUsernameRequest)(nil),   // 3: player.GetPlayerByUsernameRequest
	(*GetPlayerByUsernameResponse)(nil),  // 4: player.GetPlayerByUsernameResponse
	(*SignInRequest)(nil),                // 5: player.SignInRequest
	(*SignInResponse)(nil),               // 6: player.SignInResponse
	(*GetPlayerExperienceRequest)(nil),   // 7: player.GetPlayerExperienceRequest
	(*GetPlayerExperienceResponse)(nil),  // 8: player.GetPlayerExperienceResponse
	(*GetPlayerLifePointRequest)(nil),    // 9: player.GetPlayerLifePointRequest
	(*GetPlayerLifePointResponse)(nil),   // 10: player.GetPlayerLifePointResponse
	(*GetPlayerDailyRewardRequest)(nil),  // 11: player.GetPlayerDailyRewardRequest
	(*GetPlayerDailyRewardResponse)(nil), // 12: player.GetPlayerDailyRewardResponse
	(*Level)(nil),                        // 13: level.Level
	(*Language)(nil),                     // 14: language.Language
}
var file_proto_player_proto_depIdxs = []int32{
	13, // 0: player.Player.Level:type_name -> level.Level
	14, // 1: player.Player.Language:type_name -> language.Language
	0,  // 2: player.GetPlayerByIDResponse.Player:type_name -> player.Player
	0,  // 3: player.GetPlayerByUsernameResponse.Player:type_name -> player.Player
	0,  // 4: player.SignInResponse.Player:type_name -> player.Player
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_player_proto_init() }
func file_proto_player_proto_init() {
	if File_proto_player_proto != nil {
		return
	}
	file_proto_language_proto_init()
	file_proto_level_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerByIDRequest); i {
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
		file_proto_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerByIDResponse); i {
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
		file_proto_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerByUsernameRequest); i {
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
		file_proto_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerByUsernameResponse); i {
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
		file_proto_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_proto_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_proto_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerExperienceRequest); i {
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
		file_proto_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerExperienceResponse); i {
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
		file_proto_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerLifePointRequest); i {
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
		file_proto_player_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerLifePointResponse); i {
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
		file_proto_player_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerDailyRewardRequest); i {
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
		file_proto_player_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerDailyRewardResponse); i {
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
			RawDescriptor: file_proto_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_player_proto_goTypes,
		DependencyIndexes: file_proto_player_proto_depIdxs,
		MessageInfos:      file_proto_player_proto_msgTypes,
	}.Build()
	File_proto_player_proto = out.File
	file_proto_player_proto_rawDesc = nil
	file_proto_player_proto_goTypes = nil
	file_proto_player_proto_depIdxs = nil
}
