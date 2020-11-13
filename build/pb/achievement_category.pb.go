// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/achievement_category.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AchievementCategory struct {
	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (m *AchievementCategory) Reset()         { *m = AchievementCategory{} }
func (m *AchievementCategory) String() string { return proto.CompactTextString(m) }
func (*AchievementCategory) ProtoMessage()    {}
func (*AchievementCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{0}
}
func (m *AchievementCategory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AchievementCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AchievementCategory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AchievementCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AchievementCategory.Merge(m, src)
}
func (m *AchievementCategory) XXX_Size() int {
	return m.Size()
}
func (m *AchievementCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_AchievementCategory.DiscardUnknown(m)
}

var xxx_messageInfo_AchievementCategory proto.InternalMessageInfo

func (m *AchievementCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *AchievementCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AchievementCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

// GetAllAchievementCategory
type GetAllAchievementCategoryRequest struct {
}

func (m *GetAllAchievementCategoryRequest) Reset()         { *m = GetAllAchievementCategoryRequest{} }
func (m *GetAllAchievementCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllAchievementCategoryRequest) ProtoMessage()    {}
func (*GetAllAchievementCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{1}
}
func (m *GetAllAchievementCategoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllAchievementCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllAchievementCategoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllAchievementCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllAchievementCategoryRequest.Merge(m, src)
}
func (m *GetAllAchievementCategoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAllAchievementCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllAchievementCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllAchievementCategoryRequest proto.InternalMessageInfo

type GetAllAchievementCategoryResponse struct {
	AchievementCategories []*AchievementCategory `protobuf:"bytes,1,rep,name=AchievementCategories,proto3" json:"AchievementCategories,omitempty"`
}

func (m *GetAllAchievementCategoryResponse) Reset()         { *m = GetAllAchievementCategoryResponse{} }
func (m *GetAllAchievementCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllAchievementCategoryResponse) ProtoMessage()    {}
func (*GetAllAchievementCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{2}
}
func (m *GetAllAchievementCategoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllAchievementCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllAchievementCategoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllAchievementCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllAchievementCategoryResponse.Merge(m, src)
}
func (m *GetAllAchievementCategoryResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAllAchievementCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllAchievementCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllAchievementCategoryResponse proto.InternalMessageInfo

func (m *GetAllAchievementCategoryResponse) GetAchievementCategories() []*AchievementCategory {
	if m != nil {
		return m.AchievementCategories
	}
	return nil
}

// GetAchievementCategoryByID
type GetAchievementCategoryByIDRequest struct {
	CategoryID uint32 `protobuf:"varint,1,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
}

func (m *GetAchievementCategoryByIDRequest) Reset()         { *m = GetAchievementCategoryByIDRequest{} }
func (m *GetAchievementCategoryByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAchievementCategoryByIDRequest) ProtoMessage()    {}
func (*GetAchievementCategoryByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{3}
}
func (m *GetAchievementCategoryByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAchievementCategoryByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAchievementCategoryByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAchievementCategoryByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAchievementCategoryByIDRequest.Merge(m, src)
}
func (m *GetAchievementCategoryByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAchievementCategoryByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAchievementCategoryByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAchievementCategoryByIDRequest proto.InternalMessageInfo

func (m *GetAchievementCategoryByIDRequest) GetCategoryID() uint32 {
	if m != nil {
		return m.CategoryID
	}
	return 0
}

type GetAchievementCategoryByIDResponse struct {
	AchievementCategory *AchievementCategory `protobuf:"bytes,1,opt,name=AchievementCategory,proto3" json:"AchievementCategory,omitempty"`
}

func (m *GetAchievementCategoryByIDResponse) Reset()         { *m = GetAchievementCategoryByIDResponse{} }
func (m *GetAchievementCategoryByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetAchievementCategoryByIDResponse) ProtoMessage()    {}
func (*GetAchievementCategoryByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{4}
}
func (m *GetAchievementCategoryByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAchievementCategoryByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAchievementCategoryByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAchievementCategoryByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAchievementCategoryByIDResponse.Merge(m, src)
}
func (m *GetAchievementCategoryByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAchievementCategoryByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAchievementCategoryByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAchievementCategoryByIDResponse proto.InternalMessageInfo

func (m *GetAchievementCategoryByIDResponse) GetAchievementCategory() *AchievementCategory {
	if m != nil {
		return m.AchievementCategory
	}
	return nil
}

// GetAchievementCategoryBySlug
type GetAchievementCategoryBySlugRequest struct {
	Slug string `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (m *GetAchievementCategoryBySlugRequest) Reset()         { *m = GetAchievementCategoryBySlugRequest{} }
func (m *GetAchievementCategoryBySlugRequest) String() string { return proto.CompactTextString(m) }
func (*GetAchievementCategoryBySlugRequest) ProtoMessage()    {}
func (*GetAchievementCategoryBySlugRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{5}
}
func (m *GetAchievementCategoryBySlugRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAchievementCategoryBySlugRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAchievementCategoryBySlugRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAchievementCategoryBySlugRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAchievementCategoryBySlugRequest.Merge(m, src)
}
func (m *GetAchievementCategoryBySlugRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAchievementCategoryBySlugRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAchievementCategoryBySlugRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAchievementCategoryBySlugRequest proto.InternalMessageInfo

func (m *GetAchievementCategoryBySlugRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GetAchievementCategoryBySlugResponse struct {
	AchievementCategory *AchievementCategory `protobuf:"bytes,1,opt,name=AchievementCategory,proto3" json:"AchievementCategory,omitempty"`
}

func (m *GetAchievementCategoryBySlugResponse) Reset()         { *m = GetAchievementCategoryBySlugResponse{} }
func (m *GetAchievementCategoryBySlugResponse) String() string { return proto.CompactTextString(m) }
func (*GetAchievementCategoryBySlugResponse) ProtoMessage()    {}
func (*GetAchievementCategoryBySlugResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2883501740ee9590, []int{6}
}
func (m *GetAchievementCategoryBySlugResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAchievementCategoryBySlugResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAchievementCategoryBySlugResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAchievementCategoryBySlugResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAchievementCategoryBySlugResponse.Merge(m, src)
}
func (m *GetAchievementCategoryBySlugResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAchievementCategoryBySlugResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAchievementCategoryBySlugResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAchievementCategoryBySlugResponse proto.InternalMessageInfo

func (m *GetAchievementCategoryBySlugResponse) GetAchievementCategory() *AchievementCategory {
	if m != nil {
		return m.AchievementCategory
	}
	return nil
}

func init() {
	proto.RegisterType((*AchievementCategory)(nil), "achievement_category.AchievementCategory")
	proto.RegisterType((*GetAllAchievementCategoryRequest)(nil), "achievement_category.GetAllAchievementCategoryRequest")
	proto.RegisterType((*GetAllAchievementCategoryResponse)(nil), "achievement_category.GetAllAchievementCategoryResponse")
	proto.RegisterType((*GetAchievementCategoryByIDRequest)(nil), "achievement_category.GetAchievementCategoryByIDRequest")
	proto.RegisterType((*GetAchievementCategoryByIDResponse)(nil), "achievement_category.GetAchievementCategoryByIDResponse")
	proto.RegisterType((*GetAchievementCategoryBySlugRequest)(nil), "achievement_category.GetAchievementCategoryBySlugRequest")
	proto.RegisterType((*GetAchievementCategoryBySlugResponse)(nil), "achievement_category.GetAchievementCategoryBySlugResponse")
}

func init() { proto.RegisterFile("proto/achievement_category.proto", fileDescriptor_2883501740ee9590) }

var fileDescriptor_2883501740ee9590 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0xce, 0xc8, 0x4c, 0x2d, 0x4b, 0xcd, 0x4d, 0xcd, 0x2b, 0x89, 0x4f, 0x4e,
	0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0xd4, 0x03, 0x4b, 0x09, 0x89, 0x60, 0x93, 0x53, 0xf2, 0xe5,
	0x12, 0x76, 0x44, 0x88, 0x3b, 0x43, 0x85, 0x85, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x83, 0x98, 0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x70, 0x4e, 0x69, 0xba, 0x04, 0x33,
	0x44, 0x0c, 0xc4, 0x56, 0x52, 0xe2, 0x52, 0x70, 0x4f, 0x2d, 0x71, 0xcc, 0xc9, 0xc1, 0x62, 0x68,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x52, 0x0b, 0x23, 0x97, 0x22, 0x1e, 0x45, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xf1, 0x5c, 0xa2, 0x98, 0xd2, 0x99, 0xa9, 0xc5, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x9a, 0x7a, 0x58, 0xbd, 0x8a, 0xcd, 0x44, 0xec, 0xe6, 0x28, 0x39, 0x43,
	0x5c, 0x81, 0xa9, 0xc1, 0xa9, 0xd2, 0xd3, 0x05, 0xea, 0x56, 0x21, 0x39, 0x2e, 0x2e, 0x98, 0x30,
	0x3c, 0x3c, 0x90, 0x44, 0x94, 0x1a, 0x19, 0xb9, 0x94, 0xf0, 0x99, 0x02, 0xf5, 0x4c, 0x34, 0xd6,
	0x50, 0x06, 0x9b, 0x47, 0x92, 0x57, 0xb0, 0x99, 0xa2, 0x64, 0xc9, 0xa5, 0x8c, 0xcb, 0x09, 0xa0,
	0x38, 0x81, 0x79, 0x05, 0x16, 0x5d, 0x8c, 0x48, 0xd1, 0xd5, 0xcc, 0xc8, 0xa5, 0x82, 0x5f, 0x2f,
	0x1d, 0x3c, 0xe0, 0x24, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x2c,
	0x7a, 0xd6, 0x05, 0x49, 0x49, 0x6c, 0xe0, 0x04, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0x29, 0x5d, 0x02, 0xe4, 0x02, 0x00, 0x00,
}

func (m *AchievementCategory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AchievementCategory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AchievementCategory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintAchievementCategory(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAchievementCategory(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintAchievementCategory(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetAllAchievementCategoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllAchievementCategoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllAchievementCategoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetAllAchievementCategoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllAchievementCategoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllAchievementCategoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AchievementCategories) > 0 {
		for iNdEx := len(m.AchievementCategories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AchievementCategories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAchievementCategory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetAchievementCategoryByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAchievementCategoryByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAchievementCategoryByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CategoryID != 0 {
		i = encodeVarintAchievementCategory(dAtA, i, uint64(m.CategoryID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetAchievementCategoryByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAchievementCategoryByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAchievementCategoryByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AchievementCategory != nil {
		{
			size, err := m.AchievementCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAchievementCategory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAchievementCategoryBySlugRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAchievementCategoryBySlugRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAchievementCategoryBySlugRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintAchievementCategory(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAchievementCategoryBySlugResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAchievementCategoryBySlugResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAchievementCategoryBySlugResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AchievementCategory != nil {
		{
			size, err := m.AchievementCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAchievementCategory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAchievementCategory(dAtA []byte, offset int, v uint64) int {
	offset -= sovAchievementCategory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AchievementCategory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovAchievementCategory(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAchievementCategory(uint64(l))
	}
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovAchievementCategory(uint64(l))
	}
	return n
}

func (m *GetAllAchievementCategoryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetAllAchievementCategoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AchievementCategories) > 0 {
		for _, e := range m.AchievementCategories {
			l = e.Size()
			n += 1 + l + sovAchievementCategory(uint64(l))
		}
	}
	return n
}

func (m *GetAchievementCategoryByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CategoryID != 0 {
		n += 1 + sovAchievementCategory(uint64(m.CategoryID))
	}
	return n
}

func (m *GetAchievementCategoryByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AchievementCategory != nil {
		l = m.AchievementCategory.Size()
		n += 1 + l + sovAchievementCategory(uint64(l))
	}
	return n
}

func (m *GetAchievementCategoryBySlugRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovAchievementCategory(uint64(l))
	}
	return n
}

func (m *GetAchievementCategoryBySlugResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AchievementCategory != nil {
		l = m.AchievementCategory.Size()
		n += 1 + l + sovAchievementCategory(uint64(l))
	}
	return n
}

func sovAchievementCategory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAchievementCategory(x uint64) (n int) {
	return sovAchievementCategory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AchievementCategory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AchievementCategory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AchievementCategory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAllAchievementCategoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAllAchievementCategoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllAchievementCategoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAllAchievementCategoryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAllAchievementCategoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllAchievementCategoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AchievementCategories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AchievementCategories = append(m.AchievementCategories, &AchievementCategory{})
			if err := m.AchievementCategories[len(m.AchievementCategories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAchievementCategoryByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAchievementCategoryByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAchievementCategoryByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CategoryID", wireType)
			}
			m.CategoryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CategoryID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAchievementCategoryByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAchievementCategoryByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAchievementCategoryByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AchievementCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AchievementCategory == nil {
				m.AchievementCategory = &AchievementCategory{}
			}
			if err := m.AchievementCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAchievementCategoryBySlugRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAchievementCategoryBySlugRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAchievementCategoryBySlugRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAchievementCategoryBySlugResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAchievementCategoryBySlugResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAchievementCategoryBySlugResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AchievementCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AchievementCategory == nil {
				m.AchievementCategory = &AchievementCategory{}
			}
			if err := m.AchievementCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAchievementCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAchievementCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAchievementCategory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAchievementCategory
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAchievementCategory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAchievementCategory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAchievementCategory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAchievementCategory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAchievementCategory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAchievementCategory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAchievementCategory = fmt.Errorf("proto: unexpected end of group")
)
