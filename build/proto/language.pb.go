// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/language.proto

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

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{0}
}

func (x *Language) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Language) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

// FindLanguageBySlug
type FindLanguageBySlugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *FindLanguageBySlugRequest) Reset() {
	*x = FindLanguageBySlugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLanguageBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLanguageBySlugRequest) ProtoMessage() {}

func (x *FindLanguageBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLanguageBySlugRequest.ProtoReflect.Descriptor instead.
func (*FindLanguageBySlugRequest) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{1}
}

func (x *FindLanguageBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type FindLanguageBySlugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *Language `protobuf:"bytes,1,opt,name=Language,proto3" json:"Language,omitempty"`
}

func (x *FindLanguageBySlugResponse) Reset() {
	*x = FindLanguageBySlugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLanguageBySlugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLanguageBySlugResponse) ProtoMessage() {}

func (x *FindLanguageBySlugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLanguageBySlugResponse.ProtoReflect.Descriptor instead.
func (*FindLanguageBySlugResponse) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{2}
}

func (x *FindLanguageBySlugResponse) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

// FindLanguageBySlug
type FindLanguageByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FindLanguageByNameRequest) Reset() {
	*x = FindLanguageByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLanguageByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLanguageByNameRequest) ProtoMessage() {}

func (x *FindLanguageByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLanguageByNameRequest.ProtoReflect.Descriptor instead.
func (*FindLanguageByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{3}
}

func (x *FindLanguageByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FindLanguageByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *Language `protobuf:"bytes,1,opt,name=Language,proto3" json:"Language,omitempty"`
}

func (x *FindLanguageByNameResponse) Reset() {
	*x = FindLanguageByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLanguageByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLanguageByNameResponse) ProtoMessage() {}

func (x *FindLanguageByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLanguageByNameResponse.ProtoReflect.Descriptor instead.
func (*FindLanguageByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{4}
}

func (x *FindLanguageByNameResponse) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

// GetLanguages
type GetAllLanguagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllLanguagesRequest) Reset() {
	*x = GetAllLanguagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLanguagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLanguagesRequest) ProtoMessage() {}

func (x *GetAllLanguagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLanguagesRequest.ProtoReflect.Descriptor instead.
func (*GetAllLanguagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{5}
}

type GetAllLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []*Language `protobuf:"bytes,1,rep,name=Languages,proto3" json:"Languages,omitempty"`
}

func (x *GetAllLanguagesResponse) Reset() {
	*x = GetAllLanguagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_language_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLanguagesResponse) ProtoMessage() {}

func (x *GetAllLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_language_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLanguagesResponse.ProtoReflect.Descriptor instead.
func (*GetAllLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_language_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllLanguagesResponse) GetLanguages() []*Language {
	if x != nil {
		return x.Languages
	}
	return nil
}

var File_proto_language_proto protoreflect.FileDescriptor

var file_proto_language_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x22, 0x42, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x53, 0x6c, 0x75, 0x67, 0x22, 0x2f, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x53, 0x6c, 0x75, 0x67, 0x22, 0x4c, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x09,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_language_proto_rawDescOnce sync.Once
	file_proto_language_proto_rawDescData = file_proto_language_proto_rawDesc
)

func file_proto_language_proto_rawDescGZIP() []byte {
	file_proto_language_proto_rawDescOnce.Do(func() {
		file_proto_language_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_language_proto_rawDescData)
	})
	return file_proto_language_proto_rawDescData
}

var file_proto_language_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_language_proto_goTypes = []interface{}{
	(*Language)(nil),                   // 0: language.Language
	(*FindLanguageBySlugRequest)(nil),  // 1: language.FindLanguageBySlugRequest
	(*FindLanguageBySlugResponse)(nil), // 2: language.FindLanguageBySlugResponse
	(*FindLanguageByNameRequest)(nil),  // 3: language.FindLanguageByNameRequest
	(*FindLanguageByNameResponse)(nil), // 4: language.FindLanguageByNameResponse
	(*GetAllLanguagesRequest)(nil),     // 5: language.GetAllLanguagesRequest
	(*GetAllLanguagesResponse)(nil),    // 6: language.GetAllLanguagesResponse
}
var file_proto_language_proto_depIdxs = []int32{
	0, // 0: language.FindLanguageBySlugResponse.Language:type_name -> language.Language
	0, // 1: language.FindLanguageByNameResponse.Language:type_name -> language.Language
	0, // 2: language.GetAllLanguagesResponse.Languages:type_name -> language.Language
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_language_proto_init() }
func file_proto_language_proto_init() {
	if File_proto_language_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_language_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_proto_language_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLanguageBySlugRequest); i {
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
		file_proto_language_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLanguageBySlugResponse); i {
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
		file_proto_language_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLanguageByNameRequest); i {
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
		file_proto_language_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLanguageByNameResponse); i {
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
		file_proto_language_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllLanguagesRequest); i {
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
		file_proto_language_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllLanguagesResponse); i {
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
			RawDescriptor: file_proto_language_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_language_proto_goTypes,
		DependencyIndexes: file_proto_language_proto_depIdxs,
		MessageInfos:      file_proto_language_proto_msgTypes,
	}.Build()
	File_proto_language_proto = out.File
	file_proto_language_proto_rawDesc = nil
	file_proto_language_proto_goTypes = nil
	file_proto_language_proto_depIdxs = nil
}