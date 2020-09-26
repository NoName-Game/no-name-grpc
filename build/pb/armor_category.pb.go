// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/armor_category.proto

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

type ArmorCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *ArmorCategory) Reset() {
	*x = ArmorCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArmorCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArmorCategory) ProtoMessage() {}

func (x *ArmorCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArmorCategory.ProtoReflect.Descriptor instead.
func (*ArmorCategory) Descriptor() ([]byte, []int) {
	return file_proto_armor_category_proto_rawDescGZIP(), []int{0}
}

func (x *ArmorCategory) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ArmorCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArmorCategory) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

// GetAllArmorCategory
type GetAllArmorCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllArmorCategoryRequest) Reset() {
	*x = GetAllArmorCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllArmorCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllArmorCategoryRequest) ProtoMessage() {}

func (x *GetAllArmorCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllArmorCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetAllArmorCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_category_proto_rawDescGZIP(), []int{1}
}

type GetAllArmorCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArmorCategories []*ArmorCategory `protobuf:"bytes,1,rep,name=ArmorCategories,proto3" json:"ArmorCategories,omitempty"`
}

func (x *GetAllArmorCategoryResponse) Reset() {
	*x = GetAllArmorCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllArmorCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllArmorCategoryResponse) ProtoMessage() {}

func (x *GetAllArmorCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllArmorCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetAllArmorCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_category_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllArmorCategoryResponse) GetArmorCategories() []*ArmorCategory {
	if x != nil {
		return x.ArmorCategories
	}
	return nil
}

// GetArmorCategoryBySlug
type GetArmorCategoryBySlugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *GetArmorCategoryBySlugRequest) Reset() {
	*x = GetArmorCategoryBySlugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArmorCategoryBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArmorCategoryBySlugRequest) ProtoMessage() {}

func (x *GetArmorCategoryBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArmorCategoryBySlugRequest.ProtoReflect.Descriptor instead.
func (*GetArmorCategoryBySlugRequest) Descriptor() ([]byte, []int) {
	return file_proto_armor_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetArmorCategoryBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetArmorCategoryBySlugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArmorCategory *ArmorCategory `protobuf:"bytes,1,opt,name=ArmorCategory,proto3" json:"ArmorCategory,omitempty"`
}

func (x *GetArmorCategoryBySlugResponse) Reset() {
	*x = GetArmorCategoryBySlugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_armor_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArmorCategoryBySlugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArmorCategoryBySlugResponse) ProtoMessage() {}

func (x *GetArmorCategoryBySlugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_armor_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArmorCategoryBySlugResponse.ProtoReflect.Descriptor instead.
func (*GetArmorCategoryBySlugResponse) Descriptor() ([]byte, []int) {
	return file_proto_armor_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetArmorCategoryBySlugResponse) GetArmorCategory() *ArmorCategory {
	if x != nil {
		return x.ArmorCategory
	}
	return nil
}

var File_proto_armor_category_proto protoreflect.FileDescriptor

var file_proto_armor_category_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x72,
	0x6d, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x47, 0x0a, 0x0d,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x53, 0x6c, 0x75, 0x67, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x72,
	0x6d, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0f, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42,
	0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75, 0x67,
	0x22, 0x65, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x72, 0x6d, 0x6f,
	0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x72, 0x6d, 0x6f, 0x72,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_armor_category_proto_rawDescOnce sync.Once
	file_proto_armor_category_proto_rawDescData = file_proto_armor_category_proto_rawDesc
)

func file_proto_armor_category_proto_rawDescGZIP() []byte {
	file_proto_armor_category_proto_rawDescOnce.Do(func() {
		file_proto_armor_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_armor_category_proto_rawDescData)
	})
	return file_proto_armor_category_proto_rawDescData
}

var file_proto_armor_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_armor_category_proto_goTypes = []interface{}{
	(*ArmorCategory)(nil),                  // 0: armor_category.ArmorCategory
	(*GetAllArmorCategoryRequest)(nil),     // 1: armor_category.GetAllArmorCategoryRequest
	(*GetAllArmorCategoryResponse)(nil),    // 2: armor_category.GetAllArmorCategoryResponse
	(*GetArmorCategoryBySlugRequest)(nil),  // 3: armor_category.GetArmorCategoryBySlugRequest
	(*GetArmorCategoryBySlugResponse)(nil), // 4: armor_category.GetArmorCategoryBySlugResponse
}
var file_proto_armor_category_proto_depIdxs = []int32{
	0, // 0: armor_category.GetAllArmorCategoryResponse.ArmorCategories:type_name -> armor_category.ArmorCategory
	0, // 1: armor_category.GetArmorCategoryBySlugResponse.ArmorCategory:type_name -> armor_category.ArmorCategory
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_armor_category_proto_init() }
func file_proto_armor_category_proto_init() {
	if File_proto_armor_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_armor_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArmorCategory); i {
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
		file_proto_armor_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllArmorCategoryRequest); i {
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
		file_proto_armor_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllArmorCategoryResponse); i {
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
		file_proto_armor_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArmorCategoryBySlugRequest); i {
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
		file_proto_armor_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArmorCategoryBySlugResponse); i {
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
			RawDescriptor: file_proto_armor_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_armor_category_proto_goTypes,
		DependencyIndexes: file_proto_armor_category_proto_depIdxs,
		MessageInfos:      file_proto_armor_category_proto_msgTypes,
	}.Build()
	File_proto_armor_category_proto = out.File
	file_proto_armor_category_proto_rawDesc = nil
	file_proto_armor_category_proto_goTypes = nil
	file_proto_armor_category_proto_depIdxs = nil
}
