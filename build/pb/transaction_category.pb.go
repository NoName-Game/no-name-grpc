// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/transaction_category.proto

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

type TransactionCategory struct {
	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug string `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (m *TransactionCategory) Reset()         { *m = TransactionCategory{} }
func (m *TransactionCategory) String() string { return proto.CompactTextString(m) }
func (*TransactionCategory) ProtoMessage()    {}
func (*TransactionCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed3da49d8cd3625, []int{0}
}
func (m *TransactionCategory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionCategory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionCategory.Merge(m, src)
}
func (m *TransactionCategory) XXX_Size() int {
	return m.Size()
}
func (m *TransactionCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionCategory.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionCategory proto.InternalMessageInfo

func (m *TransactionCategory) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TransactionCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransactionCategory) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionCategory)(nil), "transaction_category.TransactionCategory")
}

func init() { proto.RegisterFile("proto/transaction_category.proto", fileDescriptor_fed3da49d8cd3625) }

var fileDescriptor_fed3da49d8cd3625 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x4f, 0x4e,
	0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0xd4, 0x03, 0x4b, 0x09, 0x89, 0x60, 0x93, 0x53, 0xf2, 0xe5,
	0x12, 0x0e, 0x41, 0x88, 0x3b, 0x43, 0x85, 0x85, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x83, 0x98, 0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x70, 0x4e, 0x69, 0xba, 0x04, 0x33,
	0x44, 0x0c, 0xc4, 0x76, 0x92, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0x16, 0x3d, 0xeb, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x5b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0x98, 0x17, 0xa2, 0xaf, 0x00, 0x00, 0x00,
}

func (m *TransactionCategory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionCategory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionCategory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintTransactionCategory(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTransactionCategory(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintTransactionCategory(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransactionCategory(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionCategory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionCategory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTransactionCategory(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTransactionCategory(uint64(l))
	}
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovTransactionCategory(uint64(l))
	}
	return n
}

func sovTransactionCategory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionCategory(x uint64) (n int) {
	return sovTransactionCategory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionCategory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionCategory
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
			return fmt.Errorf("proto: TransactionCategory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionCategory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionCategory
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
					return ErrIntOverflowTransactionCategory
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
				return ErrInvalidLengthTransactionCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionCategory
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
					return ErrIntOverflowTransactionCategory
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
				return ErrInvalidLengthTransactionCategory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionCategory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionCategory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransactionCategory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransactionCategory
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
func skipTransactionCategory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionCategory
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
					return 0, ErrIntOverflowTransactionCategory
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
					return 0, ErrIntOverflowTransactionCategory
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
				return 0, ErrInvalidLengthTransactionCategory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionCategory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionCategory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionCategory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionCategory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionCategory = fmt.Errorf("proto: unexpected end of group")
)
