// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/player_stats.proto

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

type PlayerStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerID  uint32 `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Dead      bool   `protobuf:"varint,3,opt,name=Dead,proto3" json:"Dead,omitempty"`
	LifePoint int32  `protobuf:"varint,4,opt,name=LifePoint,proto3" json:"LifePoint,omitempty"`
	Level     uint32 `protobuf:"varint,5,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *PlayerStats) Reset() {
	*x = PlayerStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStats) ProtoMessage() {}

func (x *PlayerStats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStats.ProtoReflect.Descriptor instead.
func (*PlayerStats) Descriptor() ([]byte, []int) {
	return file_proto_player_stats_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerStats) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PlayerStats) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *PlayerStats) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

func (x *PlayerStats) GetLifePoint() int32 {
	if x != nil {
		return x.LifePoint
	}
	return 0
}

func (x *PlayerStats) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_proto_player_stats_proto protoreflect.FileDescriptor

var file_proto_player_stats_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x44, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69, 0x66, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4c, 0x69, 0x66,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_player_stats_proto_rawDescOnce sync.Once
	file_proto_player_stats_proto_rawDescData = file_proto_player_stats_proto_rawDesc
)

func file_proto_player_stats_proto_rawDescGZIP() []byte {
	file_proto_player_stats_proto_rawDescOnce.Do(func() {
		file_proto_player_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_stats_proto_rawDescData)
	})
	return file_proto_player_stats_proto_rawDescData
}

var file_proto_player_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_player_stats_proto_goTypes = []interface{}{
	(*PlayerStats)(nil), // 0: player_stats.PlayerStats
}
var file_proto_player_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_player_stats_proto_init() }
func file_proto_player_stats_proto_init() {
	if File_proto_player_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_player_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerStats); i {
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
			RawDescriptor: file_proto_player_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_player_stats_proto_goTypes,
		DependencyIndexes: file_proto_player_stats_proto_depIdxs,
		MessageInfos:      file_proto_player_stats_proto_msgTypes,
	}.Build()
	File_proto_player_stats_proto = out.File
	file_proto_player_stats_proto_rawDesc = nil
	file_proto_player_stats_proto_goTypes = nil
	file_proto_player_stats_proto_depIdxs = nil
}
