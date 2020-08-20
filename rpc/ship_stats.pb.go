// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: rpc/ship_stats.proto

package rpc

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

type ShipStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Integrity uint32  `protobuf:"varint,2,opt,name=Integrity,proto3" json:"Integrity,omitempty"`
	Tank      float64 `protobuf:"fixed64,3,opt,name=Tank,proto3" json:"Tank,omitempty"`
	Hold      uint32  `protobuf:"varint,4,opt,name=Hold,proto3" json:"Hold,omitempty"`
	Engine    float64 `protobuf:"fixed64,5,opt,name=Engine,proto3" json:"Engine,omitempty"`
	Speed     uint32  `protobuf:"varint,6,opt,name=Speed,proto3" json:"Speed,omitempty"`
	Radar     uint32  `protobuf:"varint,7,opt,name=Radar,proto3" json:"Radar,omitempty"`
	Attack    float64 `protobuf:"fixed64,8,opt,name=Attack,proto3" json:"Attack,omitempty"`
	Shields   float64 `protobuf:"fixed64,9,opt,name=Shields,proto3" json:"Shields,omitempty"`
	Evasion   float64 `protobuf:"fixed64,10,opt,name=Evasion,proto3" json:"Evasion,omitempty"`
}

func (x *ShipStats) Reset() {
	*x = ShipStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ship_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipStats) ProtoMessage() {}

func (x *ShipStats) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ship_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipStats.ProtoReflect.Descriptor instead.
func (*ShipStats) Descriptor() ([]byte, []int) {
	return file_rpc_ship_stats_proto_rawDescGZIP(), []int{0}
}

func (x *ShipStats) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShipStats) GetIntegrity() uint32 {
	if x != nil {
		return x.Integrity
	}
	return 0
}

func (x *ShipStats) GetTank() float64 {
	if x != nil {
		return x.Tank
	}
	return 0
}

func (x *ShipStats) GetHold() uint32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *ShipStats) GetEngine() float64 {
	if x != nil {
		return x.Engine
	}
	return 0
}

func (x *ShipStats) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *ShipStats) GetRadar() uint32 {
	if x != nil {
		return x.Radar
	}
	return 0
}

func (x *ShipStats) GetAttack() float64 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *ShipStats) GetShields() float64 {
	if x != nil {
		return x.Shields
	}
	return 0
}

func (x *ShipStats) GetEvasion() float64 {
	if x != nil {
		return x.Evasion
	}
	return 0
}

var File_rpc_ship_stats_proto protoreflect.FileDescriptor

var file_rpc_ship_stats_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x54, 0x61,
	0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x64, 0x61, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x52, 0x61, 0x64, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x76, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x45,
	0x76, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x05, 0x5a, 0x03, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_ship_stats_proto_rawDescOnce sync.Once
	file_rpc_ship_stats_proto_rawDescData = file_rpc_ship_stats_proto_rawDesc
)

func file_rpc_ship_stats_proto_rawDescGZIP() []byte {
	file_rpc_ship_stats_proto_rawDescOnce.Do(func() {
		file_rpc_ship_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_ship_stats_proto_rawDescData)
	})
	return file_rpc_ship_stats_proto_rawDescData
}

var file_rpc_ship_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_ship_stats_proto_goTypes = []interface{}{
	(*ShipStats)(nil), // 0: ship_stats.ShipStats
}
var file_rpc_ship_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_ship_stats_proto_init() }
func file_rpc_ship_stats_proto_init() {
	if File_rpc_ship_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_ship_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipStats); i {
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
			RawDescriptor: file_rpc_ship_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_ship_stats_proto_goTypes,
		DependencyIndexes: file_rpc_ship_stats_proto_depIdxs,
		MessageInfos:      file_rpc_ship_stats_proto_msgTypes,
	}.Build()
	File_rpc_ship_stats_proto = out.File
	file_rpc_ship_stats_proto_rawDesc = nil
	file_rpc_ship_stats_proto_goTypes = nil
	file_rpc_ship_stats_proto_depIdxs = nil
}
