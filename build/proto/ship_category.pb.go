// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/ship_category.proto

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

type ShipCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *ShipCategory) Reset() {
	*x = ShipCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ship_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipCategory) ProtoMessage() {}

func (x *ShipCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ship_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipCategory.ProtoReflect.Descriptor instead.
func (*ShipCategory) Descriptor() ([]byte, []int) {
	return file_proto_ship_category_proto_rawDescGZIP(), []int{0}
}

func (x *ShipCategory) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShipCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipCategory) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_proto_ship_category_proto protoreflect.FileDescriptor

var file_proto_ship_category_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x68,
	0x69, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c,
	0x75, 0x67, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_ship_category_proto_rawDescOnce sync.Once
	file_proto_ship_category_proto_rawDescData = file_proto_ship_category_proto_rawDesc
)

func file_proto_ship_category_proto_rawDescGZIP() []byte {
	file_proto_ship_category_proto_rawDescOnce.Do(func() {
		file_proto_ship_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ship_category_proto_rawDescData)
	})
	return file_proto_ship_category_proto_rawDescData
}

var file_proto_ship_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ship_category_proto_goTypes = []interface{}{
	(*ShipCategory)(nil), // 0: ship_category.ShipCategory
}
var file_proto_ship_category_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ship_category_proto_init() }
func file_proto_ship_category_proto_init() {
	if File_proto_ship_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ship_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipCategory); i {
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
			RawDescriptor: file_proto_ship_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ship_category_proto_goTypes,
		DependencyIndexes: file_proto_ship_category_proto_depIdxs,
		MessageInfos:      file_proto_ship_category_proto_msgTypes,
	}.Build()
	File_proto_ship_category_proto = out.File
	file_proto_ship_category_proto_rawDesc = nil
	file_proto_ship_category_proto_goTypes = nil
	file_proto_ship_category_proto_depIdxs = nil
}
