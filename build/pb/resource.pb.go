// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/resource.proto

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

type Resource struct {
	ID                 uint32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string            `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Base               bool              `protobuf:"varint,3,opt,name=Base,proto3" json:"Base,omitempty"`
	RarityID           uint32            `protobuf:"varint,4,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity             *Rarity           `protobuf:"bytes,5,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ResourceCategoryID uint32            `protobuf:"varint,6,opt,name=ResourceCategoryID,proto3" json:"ResourceCategoryID,omitempty"`
	ResourceCategory   *ResourceCategory `protobuf:"bytes,7,opt,name=ResourceCategory,proto3" json:"ResourceCategory,omitempty"`
	PlanetSystemID     uint32            `protobuf:"varint,8,opt,name=PlanetSystemID,proto3" json:"PlanetSystemID,omitempty"`
	Enabled            bool              `protobuf:"varint,9,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetBase() bool {
	if m != nil {
		return m.Base
	}
	return false
}

func (m *Resource) GetRarityID() uint32 {
	if m != nil {
		return m.RarityID
	}
	return 0
}

func (m *Resource) GetRarity() *Rarity {
	if m != nil {
		return m.Rarity
	}
	return nil
}

func (m *Resource) GetResourceCategoryID() uint32 {
	if m != nil {
		return m.ResourceCategoryID
	}
	return 0
}

func (m *Resource) GetResourceCategory() *ResourceCategory {
	if m != nil {
		return m.ResourceCategory
	}
	return nil
}

func (m *Resource) GetPlanetSystemID() uint32 {
	if m != nil {
		return m.PlanetSystemID
	}
	return 0
}

func (m *Resource) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// GetResourceByID
type GetResourceByIDRequest struct {
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *GetResourceByIDRequest) Reset()         { *m = GetResourceByIDRequest{} }
func (m *GetResourceByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourceByIDRequest) ProtoMessage()    {}
func (*GetResourceByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{1}
}
func (m *GetResourceByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetResourceByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetResourceByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetResourceByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceByIDRequest.Merge(m, src)
}
func (m *GetResourceByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetResourceByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceByIDRequest proto.InternalMessageInfo

func (m *GetResourceByIDRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetResourceByIDResponse struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=Resource,proto3" json:"Resource,omitempty"`
}

func (m *GetResourceByIDResponse) Reset()         { *m = GetResourceByIDResponse{} }
func (m *GetResourceByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetResourceByIDResponse) ProtoMessage()    {}
func (*GetResourceByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{2}
}
func (m *GetResourceByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetResourceByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetResourceByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetResourceByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceByIDResponse.Merge(m, src)
}
func (m *GetResourceByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetResourceByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceByIDResponse proto.InternalMessageInfo

func (m *GetResourceByIDResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

// GetResourceByName
type GetResourceByNameRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *GetResourceByNameRequest) Reset()         { *m = GetResourceByNameRequest{} }
func (m *GetResourceByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourceByNameRequest) ProtoMessage()    {}
func (*GetResourceByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{3}
}
func (m *GetResourceByNameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetResourceByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetResourceByNameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetResourceByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceByNameRequest.Merge(m, src)
}
func (m *GetResourceByNameRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetResourceByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceByNameRequest proto.InternalMessageInfo

func (m *GetResourceByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResourceByNameResponse struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=Resource,proto3" json:"Resource,omitempty"`
}

func (m *GetResourceByNameResponse) Reset()         { *m = GetResourceByNameResponse{} }
func (m *GetResourceByNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetResourceByNameResponse) ProtoMessage()    {}
func (*GetResourceByNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{4}
}
func (m *GetResourceByNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetResourceByNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetResourceByNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetResourceByNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceByNameResponse.Merge(m, src)
}
func (m *GetResourceByNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetResourceByNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceByNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceByNameResponse proto.InternalMessageInfo

func (m *GetResourceByNameResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "resource.Resource")
	proto.RegisterType((*GetResourceByIDRequest)(nil), "resource.GetResourceByIDRequest")
	proto.RegisterType((*GetResourceByIDResponse)(nil), "resource.GetResourceByIDResponse")
	proto.RegisterType((*GetResourceByNameRequest)(nil), "resource.GetResourceByNameRequest")
	proto.RegisterType((*GetResourceByNameResponse)(nil), "resource.GetResourceByNameResponse")
}

func init() { proto.RegisterFile("proto/resource.proto", fileDescriptor_e41559ab98850371) }

var fileDescriptor_e41559ab98850371 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xed, 0xc6, 0xda, 0xa6, 0x53, 0x2c, 0x32, 0x88, 0xae, 0x05, 0x97, 0x10, 0xa1, 0xe4, 0xb4,
	0x42, 0x3d, 0x7a, 0xab, 0x11, 0x59, 0x04, 0x95, 0xf5, 0xe6, 0x45, 0xd2, 0xba, 0x88, 0xd0, 0x36,
	0x35, 0xd9, 0x1e, 0xfa, 0x17, 0x7e, 0x96, 0xde, 0x7a, 0xf4, 0x28, 0xed, 0x8f, 0x48, 0xb7, 0xd9,
	0x94, 0x26, 0x3d, 0x79, 0x9b, 0x79, 0xf3, 0x66, 0xde, 0xce, 0xbc, 0x85, 0xa3, 0x49, 0x12, 0xeb,
	0xf8, 0x22, 0x51, 0x69, 0x3c, 0x4d, 0x06, 0x8a, 0x9b, 0x14, 0x5d, 0x9b, 0xb7, 0xcf, 0xb6, 0xeb,
	0x2f, 0x83, 0x48, 0xab, 0xb7, 0x38, 0x99, 0xad, 0x89, 0x6d, 0xcc, 0xca, 0x51, 0xf2, 0xae, 0x33,
	0xcc, 0xff, 0x76, 0xc0, 0x95, 0x19, 0x1f, 0x5b, 0xe0, 0x88, 0x90, 0x12, 0x8f, 0x04, 0x07, 0xd2,
	0x11, 0x21, 0x22, 0x54, 0xef, 0xa3, 0x91, 0xa2, 0x8e, 0x47, 0x82, 0x86, 0x34, 0xf1, 0x0a, 0xeb,
	0x45, 0xa9, 0xa2, 0x7b, 0x1e, 0x09, 0x5c, 0x69, 0x62, 0x6c, 0x83, 0x2b, 0xcd, 0x50, 0x11, 0xd2,
	0xaa, 0xe9, 0xce, 0x73, 0xec, 0x40, 0x6d, 0x1d, 0xd3, 0x7d, 0x8f, 0x04, 0xcd, 0x6e, 0x8b, 0x67,
	0xfa, 0x6b, 0x54, 0x66, 0x55, 0xe4, 0x80, 0xf6, 0x1d, 0xd7, 0xd9, 0xb3, 0x45, 0x48, 0x6b, 0x66,
	0xda, 0x8e, 0x0a, 0x3e, 0xc0, 0x61, 0x11, 0xa5, 0x75, 0xa3, 0x70, 0xce, 0xcb, 0x07, 0x28, 0x52,
	0x65, 0xa9, 0x19, 0x3b, 0xd0, 0x7a, 0x1c, 0x46, 0x63, 0xa5, 0x9f, 0x66, 0xa9, 0x56, 0x23, 0x11,
	0x52, 0xd7, 0x88, 0x17, 0x50, 0xa4, 0x50, 0xbf, 0x19, 0x47, 0xfd, 0xa1, 0x7a, 0xa5, 0x0d, 0x73,
	0x03, 0x9b, 0xfa, 0x01, 0x1c, 0xdf, 0x2a, 0x6d, 0x07, 0xf7, 0x66, 0x22, 0x94, 0xea, 0x63, 0xaa,
	0x52, 0x5d, 0x3c, 0xac, 0x2f, 0xe0, 0xa4, 0xc4, 0x4c, 0x27, 0xf1, 0x38, 0x55, 0xc8, 0x37, 0x7e,
	0x98, 0x86, 0x66, 0x17, 0xf3, 0x7d, 0xf2, 0x35, 0x64, 0xce, 0xf1, 0x39, 0xd0, 0xad, 0x51, 0x2b,
	0x93, 0xac, 0xac, 0xf5, 0x8f, 0x6c, 0xfc, 0xf3, 0xef, 0xe0, 0x74, 0x07, 0xff, 0x7f, 0xe2, 0x3d,
	0xf6, 0xb5, 0x60, 0x64, 0xbe, 0x60, 0xe4, 0x77, 0xc1, 0xc8, 0xe7, 0x92, 0x55, 0xe6, 0x4b, 0x56,
	0xf9, 0x59, 0xb2, 0xca, 0x73, 0x95, 0x5f, 0x4d, 0xfa, 0xfd, 0x9a, 0xf9, 0x64, 0x97, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x25, 0xb9, 0xa1, 0xe4, 0xb9, 0x02, 0x00, 0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.PlanetSystemID != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.PlanetSystemID))
		i--
		dAtA[i] = 0x40
	}
	if m.ResourceCategory != nil {
		{
			size, err := m.ResourceCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.ResourceCategoryID != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.ResourceCategoryID))
		i--
		dAtA[i] = 0x30
	}
	if m.Rarity != nil {
		{
			size, err := m.Rarity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.RarityID != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.RarityID))
		i--
		dAtA[i] = 0x20
	}
	if m.Base {
		i--
		if m.Base {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetResourceByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResourceByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetResourceByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetResourceByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResourceByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetResourceByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetResourceByNameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResourceByNameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetResourceByNameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetResourceByNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResourceByNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetResourceByNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovResource(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	if m.Base {
		n += 2
	}
	if m.RarityID != 0 {
		n += 1 + sovResource(uint64(m.RarityID))
	}
	if m.Rarity != nil {
		l = m.Rarity.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	if m.ResourceCategoryID != 0 {
		n += 1 + sovResource(uint64(m.ResourceCategoryID))
	}
	if m.ResourceCategory != nil {
		l = m.ResourceCategory.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	if m.PlanetSystemID != 0 {
		n += 1 + sovResource(uint64(m.PlanetSystemID))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *GetResourceByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovResource(uint64(m.ID))
	}
	return n
}

func (m *GetResourceByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *GetResourceByNameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *GetResourceByNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Base = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RarityID", wireType)
			}
			m.RarityID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RarityID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rarity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rarity == nil {
				m.Rarity = &Rarity{}
			}
			if err := m.Rarity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceCategoryID", wireType)
			}
			m.ResourceCategoryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceCategoryID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResourceCategory == nil {
				m.ResourceCategory = &ResourceCategory{}
			}
			if err := m.ResourceCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanetSystemID", wireType)
			}
			m.PlanetSystemID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanetSystemID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResource
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
func (m *GetResourceByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: GetResourceByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResourceByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResource
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
func (m *GetResourceByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: GetResourceByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResourceByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &Resource{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResource
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
func (m *GetResourceByNameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: GetResourceByNameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResourceByNameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResource
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
func (m *GetResourceByNameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: GetResourceByNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResourceByNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &Resource{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)
