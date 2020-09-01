// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/player_inventory.proto

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

type PlayerInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerID uint32    `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Quantity int32     `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	ItemID   uint32    `protobuf:"varint,4,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	ItemType string    `protobuf:"bytes,5,opt,name=ItemType,proto3" json:"ItemType,omitempty"`
	Resource *Resource `protobuf:"bytes,6,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Item     *Item     `protobuf:"bytes,7,opt,name=Item,proto3" json:"Item,omitempty"`
}

func (x *PlayerInventory) Reset() {
	*x = PlayerInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInventory) ProtoMessage() {}

func (x *PlayerInventory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInventory.ProtoReflect.Descriptor instead.
func (*PlayerInventory) Descriptor() ([]byte, []int) {
	return file_proto_player_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInventory) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PlayerInventory) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *PlayerInventory) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PlayerInventory) GetItemID() uint32 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *PlayerInventory) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *PlayerInventory) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *PlayerInventory) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_proto_player_inventory_proto protoreflect.FileDescriptor

var file_proto_player_inventory_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_player_inventory_proto_rawDescOnce sync.Once
	file_proto_player_inventory_proto_rawDescData = file_proto_player_inventory_proto_rawDesc
)

func file_proto_player_inventory_proto_rawDescGZIP() []byte {
	file_proto_player_inventory_proto_rawDescOnce.Do(func() {
		file_proto_player_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_inventory_proto_rawDescData)
	})
	return file_proto_player_inventory_proto_rawDescData
}

var file_proto_player_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_player_inventory_proto_goTypes = []interface{}{
	(*PlayerInventory)(nil), // 0: player_inventory.PlayerInventory
	(*Resource)(nil),        // 1: resource.Resource
	(*Item)(nil),            // 2: item.Item
}
var file_proto_player_inventory_proto_depIdxs = []int32{
	1, // 0: player_inventory.PlayerInventory.Resource:type_name -> resource.Resource
	2, // 1: player_inventory.PlayerInventory.Item:type_name -> item.Item
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_player_inventory_proto_init() }
func file_proto_player_inventory_proto_init() {
	if File_proto_player_inventory_proto != nil {
		return
	}
	file_proto_item_proto_init()
	file_proto_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_player_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInventory); i {
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
			RawDescriptor: file_proto_player_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_player_inventory_proto_goTypes,
		DependencyIndexes: file_proto_player_inventory_proto_depIdxs,
		MessageInfos:      file_proto_player_inventory_proto_msgTypes,
	}.Build()
	File_proto_player_inventory_proto = out.File
	file_proto_player_inventory_proto_rawDesc = nil
	file_proto_player_inventory_proto_goTypes = nil
	file_proto_player_inventory_proto_depIdxs = nil
}
