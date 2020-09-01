// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/trap.proto

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

type Trap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Damage    int32  `protobuf:"varint,2,opt,name=Damage,proto3" json:"Damage,omitempty"`
	PlayerDie bool   `protobuf:"varint,3,opt,name=PlayerDie,proto3" json:"PlayerDie,omitempty"`
}

func (x *Trap) Reset() {
	*x = Trap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trap) ProtoMessage() {}

func (x *Trap) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trap.ProtoReflect.Descriptor instead.
func (*Trap) Descriptor() ([]byte, []int) {
	return file_proto_trap_proto_rawDescGZIP(), []int{0}
}

func (x *Trap) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Trap) GetDamage() int32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *Trap) GetPlayerDie() bool {
	if x != nil {
		return x.PlayerDie
	}
	return false
}

var File_proto_trap_proto protoreflect.FileDescriptor

var file_proto_trap_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x72, 0x61, 0x70, 0x22, 0x4c, 0x0a, 0x04, 0x54, 0x72, 0x61, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x69, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_trap_proto_rawDescOnce sync.Once
	file_proto_trap_proto_rawDescData = file_proto_trap_proto_rawDesc
)

func file_proto_trap_proto_rawDescGZIP() []byte {
	file_proto_trap_proto_rawDescOnce.Do(func() {
		file_proto_trap_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_trap_proto_rawDescData)
	})
	return file_proto_trap_proto_rawDescData
}

var file_proto_trap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_trap_proto_goTypes = []interface{}{
	(*Trap)(nil), // 0: trap.Trap
}
var file_proto_trap_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_trap_proto_init() }
func file_proto_trap_proto_init() {
	if File_proto_trap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_trap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trap); i {
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
			RawDescriptor: file_proto_trap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_trap_proto_goTypes,
		DependencyIndexes: file_proto_trap_proto_depIdxs,
		MessageInfos:      file_proto_trap_proto_msgTypes,
	}.Build()
	File_proto_trap_proto = out.File
	file_proto_trap_proto_rawDesc = nil
	file_proto_trap_proto_goTypes = nil
	file_proto_trap_proto_depIdxs = nil
}
