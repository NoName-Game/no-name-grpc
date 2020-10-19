// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/ship_category.proto

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

type ShipCategory struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipCategory) Reset()         { *m = ShipCategory{} }
func (m *ShipCategory) String() string { return proto.CompactTextString(m) }
func (*ShipCategory) ProtoMessage()    {}
func (*ShipCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8989f56e1d68ed5, []int{0}
}
func (m *ShipCategory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShipCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShipCategory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShipCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipCategory.Merge(m, src)
}
func (m *ShipCategory) XXX_Size() int {
	return m.Size()
}
func (m *ShipCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipCategory.DiscardUnknown(m)
}

var xxx_messageInfo_ShipCategory proto.InternalMessageInfo

func (m *ShipCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ShipCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShipCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

// GetAllShipCategories
type GetAllShipCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllShipCategoriesRequest) Reset()         { *m = GetAllShipCategoriesRequest{} }
func (m *GetAllShipCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllShipCategoriesRequest) ProtoMessage()    {}
func (*GetAllShipCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8989f56e1d68ed5, []int{1}
}
func (m *GetAllShipCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllShipCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllShipCategoriesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllShipCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShipCategoriesRequest.Merge(m, src)
}
func (m *GetAllShipCategoriesRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAllShipCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShipCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShipCategoriesRequest proto.InternalMessageInfo

type GetAllShipCategoriesResponse struct {
	ShipCategories       []*ShipCategory `protobuf:"bytes,1,rep,name=ShipCategories,proto3" json:"ShipCategories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllShipCategoriesResponse) Reset()         { *m = GetAllShipCategoriesResponse{} }
func (m *GetAllShipCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllShipCategoriesResponse) ProtoMessage()    {}
func (*GetAllShipCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8989f56e1d68ed5, []int{2}
}
func (m *GetAllShipCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllShipCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllShipCategoriesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllShipCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShipCategoriesResponse.Merge(m, src)
}
func (m *GetAllShipCategoriesResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAllShipCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShipCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShipCategoriesResponse proto.InternalMessageInfo

func (m *GetAllShipCategoriesResponse) GetShipCategories() []*ShipCategory {
	if m != nil {
		return m.ShipCategories
	}
	return nil
}

// GetShipCategoryByID
type GetShipCategoryByIDRequest struct {
	ShipCategoryID       uint32   `protobuf:"varint,1,opt,name=ShipCategoryID,proto3" json:"ShipCategoryID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShipCategoryByIDRequest) Reset()         { *m = GetShipCategoryByIDRequest{} }
func (m *GetShipCategoryByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetShipCategoryByIDRequest) ProtoMessage()    {}
func (*GetShipCategoryByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8989f56e1d68ed5, []int{3}
}
func (m *GetShipCategoryByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetShipCategoryByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetShipCategoryByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetShipCategoryByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipCategoryByIDRequest.Merge(m, src)
}
func (m *GetShipCategoryByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetShipCategoryByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipCategoryByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipCategoryByIDRequest proto.InternalMessageInfo

func (m *GetShipCategoryByIDRequest) GetShipCategoryID() uint32 {
	if m != nil {
		return m.ShipCategoryID
	}
	return 0
}

type GetShipCategoryByIDResponse struct {
	ShipCategory         *ShipCategory `protobuf:"bytes,1,opt,name=ShipCategory,proto3" json:"ShipCategory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetShipCategoryByIDResponse) Reset()         { *m = GetShipCategoryByIDResponse{} }
func (m *GetShipCategoryByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetShipCategoryByIDResponse) ProtoMessage()    {}
func (*GetShipCategoryByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8989f56e1d68ed5, []int{4}
}
func (m *GetShipCategoryByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetShipCategoryByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetShipCategoryByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetShipCategoryByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipCategoryByIDResponse.Merge(m, src)
}
func (m *GetShipCategoryByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetShipCategoryByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipCategoryByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipCategoryByIDResponse proto.InternalMessageInfo

func (m *GetShipCategoryByIDResponse) GetShipCategory() *ShipCategory {
	if m != nil {
		return m.ShipCategory
	}
	return nil
}

func init() {
	proto.RegisterType((*ShipCategory)(nil), "ship_category.ShipCategory")
	proto.RegisterType((*GetAllShipCategoriesRequest)(nil), "ship_category.GetAllShipCategoriesRequest")
	proto.RegisterType((*GetAllShipCategoriesResponse)(nil), "ship_category.GetAllShipCategoriesResponse")
	proto.RegisterType((*GetShipCategoryByIDRequest)(nil), "ship_category.GetShipCategoryByIDRequest")
	proto.RegisterType((*GetShipCategoryByIDResponse)(nil), "ship_category.GetShipCategoryByIDResponse")
}

func init() { proto.RegisterFile("proto/ship_category.proto", fileDescriptor_a8989f56e1d68ed5) }

var fileDescriptor_a8989f56e1d68ed5 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xc8, 0x2c, 0x88, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0xd4,
	0x03, 0x8b, 0x09, 0xf1, 0xa2, 0x08, 0x2a, 0xb9, 0x71, 0xf1, 0x04, 0x67, 0x64, 0x16, 0x38, 0x43,
	0xf9, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x4c, 0x9e,
	0x2e, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0x36, 0x48, 0x2c, 0x38, 0xa7, 0x34, 0x5d, 0x82, 0x19, 0x22, 0x06, 0x62, 0x2b, 0xc9, 0x72,
	0x49, 0xbb, 0xa7, 0x96, 0x38, 0xe6, 0xe4, 0x20, 0x99, 0x96, 0x99, 0x5a, 0x1c, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0xa2, 0x94, 0xcc, 0x25, 0x83, 0x5d, 0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0xc8, 0x99, 0x8b, 0x0f, 0x55, 0x46, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5a, 0x0f, 0xd5,
	0x0f, 0xc8, 0x6e, 0x0d, 0x42, 0xd3, 0xa2, 0xe4, 0xc2, 0x25, 0xe5, 0x9e, 0x5a, 0x82, 0xac, 0xc4,
	0xa9, 0xd2, 0xd3, 0x05, 0xea, 0x04, 0x21, 0x35, 0x14, 0x2b, 0x2a, 0xe1, 0xbe, 0x44, 0x13, 0x55,
	0x8a, 0x03, 0xfb, 0x04, 0xd3, 0x14, 0xa8, 0x4b, 0xed, 0x51, 0x03, 0x0c, 0x6c, 0x08, 0x01, 0x77,
	0xa2, 0x68, 0x70, 0x12, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0xa3, 0x58, 0xf4, 0xac, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x71, 0x63, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0x3e, 0x5e, 0x2e, 0xb8, 0x01, 0x00, 0x00,
}

func (m *ShipCategory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShipCategory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShipCategory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintShipCategory(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintShipCategory(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintShipCategory(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetAllShipCategoriesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllShipCategoriesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllShipCategoriesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *GetAllShipCategoriesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllShipCategoriesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllShipCategoriesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ShipCategories) > 0 {
		for iNdEx := len(m.ShipCategories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShipCategories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShipCategory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetShipCategoryByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetShipCategoryByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetShipCategoryByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ShipCategoryID != 0 {
		i = encodeVarintShipCategory(dAtA, i, uint64(m.ShipCategoryID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetShipCategoryByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetShipCategoryByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetShipCategoryByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ShipCategory != nil {
		{
			size, err := m.ShipCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintShipCategory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintShipCategory(dAtA []byte, offset int, v uint64) int {
	offset -= sovShipCategory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ShipCategory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovShipCategory(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovShipCategory(uint64(l))
	}
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovShipCategory(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAllShipCategoriesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAllShipCategoriesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ShipCategories) > 0 {
		for _, e := range m.ShipCategories {
			l = e.Size()
			n += 1 + l + sovShipCategory(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetShipCategoryByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShipCategoryID != 0 {
		n += 1 + sovShipCategory(uint64(m.ShipCategoryID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetShipCategoryByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShipCategory != nil {
		l = m.ShipCategory.Size()
		n += 1 + l + sovShipCategory(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShipCategory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShipCategory(x uint64) (n int) {
	return sovShipCategory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShipCategory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShipCategory
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
			return fmt.Errorf("proto: ShipCategory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShipCategory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShipCategory
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
					return ErrIntOverflowShipCategory
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
				return ErrInvalidLengthShipCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShipCategory
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
					return ErrIntOverflowShipCategory
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
				return ErrInvalidLengthShipCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShipCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShipCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAllShipCategoriesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShipCategory
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
			return fmt.Errorf("proto: GetAllShipCategoriesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllShipCategoriesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShipCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAllShipCategoriesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShipCategory
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
			return fmt.Errorf("proto: GetAllShipCategoriesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllShipCategoriesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShipCategories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShipCategory
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
				return ErrInvalidLengthShipCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShipCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShipCategories = append(m.ShipCategories, &ShipCategory{})
			if err := m.ShipCategories[len(m.ShipCategories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShipCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetShipCategoryByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShipCategory
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
			return fmt.Errorf("proto: GetShipCategoryByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetShipCategoryByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShipCategoryID", wireType)
			}
			m.ShipCategoryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShipCategory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShipCategoryID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShipCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetShipCategoryByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShipCategory
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
			return fmt.Errorf("proto: GetShipCategoryByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetShipCategoryByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShipCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShipCategory
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
				return ErrInvalidLengthShipCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShipCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShipCategory == nil {
				m.ShipCategory = &ShipCategory{}
			}
			if err := m.ShipCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShipCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShipCategory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipShipCategory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShipCategory
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
					return 0, ErrIntOverflowShipCategory
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
					return 0, ErrIntOverflowShipCategory
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
				return 0, ErrInvalidLengthShipCategory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShipCategory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShipCategory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShipCategory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShipCategory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShipCategory = fmt.Errorf("proto: unexpected end of group")
)
