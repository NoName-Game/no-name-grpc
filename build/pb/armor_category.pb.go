// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/armor_category.proto

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

type ArmorCategory struct {
	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (m *ArmorCategory) Reset()         { *m = ArmorCategory{} }
func (m *ArmorCategory) String() string { return proto.CompactTextString(m) }
func (*ArmorCategory) ProtoMessage()    {}
func (*ArmorCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe9710380d9ecc4, []int{0}
}
func (m *ArmorCategory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArmorCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArmorCategory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArmorCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArmorCategory.Merge(m, src)
}
func (m *ArmorCategory) XXX_Size() int {
	return m.Size()
}
func (m *ArmorCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_ArmorCategory.DiscardUnknown(m)
}

var xxx_messageInfo_ArmorCategory proto.InternalMessageInfo

func (m *ArmorCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ArmorCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArmorCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

// GetAllArmorCategory
type GetAllArmorCategoryRequest struct {
}

func (m *GetAllArmorCategoryRequest) Reset()         { *m = GetAllArmorCategoryRequest{} }
func (m *GetAllArmorCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllArmorCategoryRequest) ProtoMessage()    {}
func (*GetAllArmorCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe9710380d9ecc4, []int{1}
}
func (m *GetAllArmorCategoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllArmorCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllArmorCategoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllArmorCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllArmorCategoryRequest.Merge(m, src)
}
func (m *GetAllArmorCategoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAllArmorCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllArmorCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllArmorCategoryRequest proto.InternalMessageInfo

type GetAllArmorCategoryResponse struct {
	ArmorCategories []*ArmorCategory `protobuf:"bytes,1,rep,name=ArmorCategories,proto3" json:"ArmorCategories,omitempty"`
}

func (m *GetAllArmorCategoryResponse) Reset()         { *m = GetAllArmorCategoryResponse{} }
func (m *GetAllArmorCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllArmorCategoryResponse) ProtoMessage()    {}
func (*GetAllArmorCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe9710380d9ecc4, []int{2}
}
func (m *GetAllArmorCategoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllArmorCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllArmorCategoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllArmorCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllArmorCategoryResponse.Merge(m, src)
}
func (m *GetAllArmorCategoryResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAllArmorCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllArmorCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllArmorCategoryResponse proto.InternalMessageInfo

func (m *GetAllArmorCategoryResponse) GetArmorCategories() []*ArmorCategory {
	if m != nil {
		return m.ArmorCategories
	}
	return nil
}

// GetArmorCategoryBySlug
type GetArmorCategoryBySlugRequest struct {
	Slug string `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (m *GetArmorCategoryBySlugRequest) Reset()         { *m = GetArmorCategoryBySlugRequest{} }
func (m *GetArmorCategoryBySlugRequest) String() string { return proto.CompactTextString(m) }
func (*GetArmorCategoryBySlugRequest) ProtoMessage()    {}
func (*GetArmorCategoryBySlugRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe9710380d9ecc4, []int{3}
}
func (m *GetArmorCategoryBySlugRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArmorCategoryBySlugRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArmorCategoryBySlugRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArmorCategoryBySlugRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArmorCategoryBySlugRequest.Merge(m, src)
}
func (m *GetArmorCategoryBySlugRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetArmorCategoryBySlugRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArmorCategoryBySlugRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArmorCategoryBySlugRequest proto.InternalMessageInfo

func (m *GetArmorCategoryBySlugRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GetArmorCategoryBySlugResponse struct {
	ArmorCategory *ArmorCategory `protobuf:"bytes,1,opt,name=ArmorCategory,proto3" json:"ArmorCategory,omitempty"`
}

func (m *GetArmorCategoryBySlugResponse) Reset()         { *m = GetArmorCategoryBySlugResponse{} }
func (m *GetArmorCategoryBySlugResponse) String() string { return proto.CompactTextString(m) }
func (*GetArmorCategoryBySlugResponse) ProtoMessage()    {}
func (*GetArmorCategoryBySlugResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe9710380d9ecc4, []int{4}
}
func (m *GetArmorCategoryBySlugResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArmorCategoryBySlugResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArmorCategoryBySlugResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArmorCategoryBySlugResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArmorCategoryBySlugResponse.Merge(m, src)
}
func (m *GetArmorCategoryBySlugResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetArmorCategoryBySlugResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArmorCategoryBySlugResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArmorCategoryBySlugResponse proto.InternalMessageInfo

func (m *GetArmorCategoryBySlugResponse) GetArmorCategory() *ArmorCategory {
	if m != nil {
		return m.ArmorCategory
	}
	return nil
}

func init() {
	proto.RegisterType((*ArmorCategory)(nil), "armor_category.ArmorCategory")
	proto.RegisterType((*GetAllArmorCategoryRequest)(nil), "armor_category.GetAllArmorCategoryRequest")
	proto.RegisterType((*GetAllArmorCategoryResponse)(nil), "armor_category.GetAllArmorCategoryResponse")
	proto.RegisterType((*GetArmorCategoryBySlugRequest)(nil), "armor_category.GetArmorCategoryBySlugRequest")
	proto.RegisterType((*GetArmorCategoryBySlugResponse)(nil), "armor_category.GetArmorCategoryBySlugResponse")
}

func init() { proto.RegisterFile("proto/armor_category.proto", fileDescriptor_abe9710380d9ecc4) }

var fileDescriptor_abe9710380d9ecc4 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0xca, 0xcd, 0x2f, 0x8a, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa,
	0xd4, 0x03, 0x0b, 0x0a, 0xf1, 0xa1, 0x8a, 0x2a, 0xb9, 0x73, 0xf1, 0x3a, 0x82, 0x44, 0x9c, 0xa1,
	0x02, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x4c, 0x9e,
	0x2e, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0x36, 0x48, 0x2c, 0x38, 0xa7, 0x34, 0x5d, 0x82, 0x19, 0x22, 0x06, 0x62, 0x2b, 0xc9, 0x70,
	0x49, 0xb9, 0xa7, 0x96, 0x38, 0xe6, 0xe4, 0xa0, 0x18, 0x17, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0xa2, 0x94, 0xc6, 0x25, 0x8d, 0x55, 0xb6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x9d, 0x8b,
	0x1f, 0x59, 0x22, 0x33, 0xb5, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x56, 0x0f, 0xcd,
	0x17, 0xa8, 0xfa, 0xd1, 0x75, 0x29, 0x19, 0x73, 0xc9, 0x82, 0xec, 0x41, 0x56, 0xe4, 0x54, 0x09,
	0x72, 0x1f, 0xd4, 0x21, 0x70, 0xa7, 0x33, 0x22, 0x39, 0x3d, 0x95, 0x4b, 0x0e, 0x97, 0x26, 0xa8,
	0xfb, 0x9c, 0xd1, 0x42, 0x09, 0xac, 0x9d, 0xa0, 0xeb, 0x50, 0xf5, 0x38, 0xc9, 0x9d, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x8b, 0x9e, 0x75, 0x41, 0x52, 0x12, 0x1b, 0x38,
	0x86, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x8e, 0xc5, 0x8c, 0xbf, 0x01, 0x00, 0x00,
}

func (m *ArmorCategory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArmorCategory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArmorCategory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintArmorCategory(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintArmorCategory(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintArmorCategory(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetAllArmorCategoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllArmorCategoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllArmorCategoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetAllArmorCategoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllArmorCategoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllArmorCategoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ArmorCategories) > 0 {
		for iNdEx := len(m.ArmorCategories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ArmorCategories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintArmorCategory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetArmorCategoryBySlugRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArmorCategoryBySlugRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArmorCategoryBySlugRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintArmorCategory(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetArmorCategoryBySlugResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArmorCategoryBySlugResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArmorCategoryBySlugResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ArmorCategory != nil {
		{
			size, err := m.ArmorCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintArmorCategory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintArmorCategory(dAtA []byte, offset int, v uint64) int {
	offset -= sovArmorCategory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ArmorCategory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovArmorCategory(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovArmorCategory(uint64(l))
	}
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovArmorCategory(uint64(l))
	}
	return n
}

func (m *GetAllArmorCategoryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetAllArmorCategoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ArmorCategories) > 0 {
		for _, e := range m.ArmorCategories {
			l = e.Size()
			n += 1 + l + sovArmorCategory(uint64(l))
		}
	}
	return n
}

func (m *GetArmorCategoryBySlugRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovArmorCategory(uint64(l))
	}
	return n
}

func (m *GetArmorCategoryBySlugResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ArmorCategory != nil {
		l = m.ArmorCategory.Size()
		n += 1 + l + sovArmorCategory(uint64(l))
	}
	return n
}

func sovArmorCategory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozArmorCategory(x uint64) (n int) {
	return sovArmorCategory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArmorCategory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArmorCategory
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
			return fmt.Errorf("proto: ArmorCategory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArmorCategory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArmorCategory
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
					return ErrIntOverflowArmorCategory
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
				return ErrInvalidLengthArmorCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArmorCategory
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
					return ErrIntOverflowArmorCategory
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
				return ErrInvalidLengthArmorCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArmorCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArmorCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArmorCategory
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
func (m *GetAllArmorCategoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArmorCategory
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
			return fmt.Errorf("proto: GetAllArmorCategoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllArmorCategoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipArmorCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArmorCategory
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
func (m *GetAllArmorCategoryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArmorCategory
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
			return fmt.Errorf("proto: GetAllArmorCategoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllArmorCategoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArmorCategories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArmorCategory
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
				return ErrInvalidLengthArmorCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArmorCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ArmorCategories = append(m.ArmorCategories, &ArmorCategory{})
			if err := m.ArmorCategories[len(m.ArmorCategories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArmorCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArmorCategory
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
func (m *GetArmorCategoryBySlugRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArmorCategory
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
			return fmt.Errorf("proto: GetArmorCategoryBySlugRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArmorCategoryBySlugRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArmorCategory
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
				return ErrInvalidLengthArmorCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArmorCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArmorCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArmorCategory
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
func (m *GetArmorCategoryBySlugResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArmorCategory
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
			return fmt.Errorf("proto: GetArmorCategoryBySlugResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArmorCategoryBySlugResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArmorCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArmorCategory
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
				return ErrInvalidLengthArmorCategory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArmorCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ArmorCategory == nil {
				m.ArmorCategory = &ArmorCategory{}
			}
			if err := m.ArmorCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArmorCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArmorCategory
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
func skipArmorCategory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArmorCategory
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
					return 0, ErrIntOverflowArmorCategory
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
					return 0, ErrIntOverflowArmorCategory
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
				return 0, ErrInvalidLengthArmorCategory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupArmorCategory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthArmorCategory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthArmorCategory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArmorCategory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupArmorCategory = fmt.Errorf("proto: unexpected end of group")
)
