// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/planet_biome.proto

package pb

import (
	encoding_binary "encoding/binary"
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

type PlanetBiome struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	Solid                float64  `protobuf:"fixed64,4,opt,name=Solid,proto3" json:"Solid,omitempty"`
	Liquid               float64  `protobuf:"fixed64,5,opt,name=Liquid,proto3" json:"Liquid,omitempty"`
	Gas                  float64  `protobuf:"fixed64,6,opt,name=Gas,proto3" json:"Gas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanetBiome) Reset()         { *m = PlanetBiome{} }
func (m *PlanetBiome) String() string { return proto.CompactTextString(m) }
func (*PlanetBiome) ProtoMessage()    {}
func (*PlanetBiome) Descriptor() ([]byte, []int) {
	return fileDescriptor_197204c4008784b2, []int{0}
}
func (m *PlanetBiome) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlanetBiome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlanetBiome.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlanetBiome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanetBiome.Merge(m, src)
}
func (m *PlanetBiome) XXX_Size() int {
	return m.Size()
}
func (m *PlanetBiome) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanetBiome.DiscardUnknown(m)
}

var xxx_messageInfo_PlanetBiome proto.InternalMessageInfo

func (m *PlanetBiome) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PlanetBiome) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlanetBiome) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *PlanetBiome) GetSolid() float64 {
	if m != nil {
		return m.Solid
	}
	return 0
}

func (m *PlanetBiome) GetLiquid() float64 {
	if m != nil {
		return m.Liquid
	}
	return 0
}

func (m *PlanetBiome) GetGas() float64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func init() {
	proto.RegisterType((*PlanetBiome)(nil), "planet_biome.PlanetBiome")
}

func init() { proto.RegisterFile("proto/planet_biome.proto", fileDescriptor_197204c4008784b2) }

var fileDescriptor_197204c4008784b2 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0x49, 0xcc, 0x4b, 0x2d, 0x89, 0x4f, 0xca, 0xcc, 0xcf, 0x4d, 0xd5, 0x03,
	0x0b, 0x09, 0xf1, 0x20, 0x8b, 0x29, 0xb5, 0x32, 0x72, 0x71, 0x07, 0x80, 0x05, 0x9c, 0x40, 0x7c,
	0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x26, 0x4f, 0x17,
	0x21, 0x21, 0x2e, 0x16, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30,
	0x1b, 0x24, 0x16, 0x9c, 0x53, 0x9a, 0x2e, 0xc1, 0x0c, 0x11, 0x03, 0xb1, 0x85, 0x44, 0xb8, 0x58,
	0x83, 0xf3, 0x73, 0x32, 0x53, 0x24, 0x58, 0x14, 0x18, 0x35, 0x18, 0x83, 0x20, 0x1c, 0x21, 0x31,
	0x2e, 0x36, 0x9f, 0xcc, 0xc2, 0xd2, 0xcc, 0x14, 0x09, 0x56, 0xb0, 0x30, 0x94, 0x27, 0x24, 0xc0,
	0xc5, 0xec, 0x9e, 0x58, 0x2c, 0xc1, 0x06, 0x16, 0x04, 0x31, 0x9d, 0x84, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x28, 0x16, 0x3d, 0xeb, 0x82, 0xa4, 0x24,
	0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0x85, 0x48, 0x4b, 0xcc, 0x00,
	0x00, 0x00,
}

func (m *PlanetBiome) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlanetBiome) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlanetBiome) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Gas != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Gas))))
		i--
		dAtA[i] = 0x31
	}
	if m.Liquid != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Liquid))))
		i--
		dAtA[i] = 0x29
	}
	if m.Solid != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Solid))))
		i--
		dAtA[i] = 0x21
	}
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintPlanetBiome(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPlanetBiome(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintPlanetBiome(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlanetBiome(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlanetBiome(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlanetBiome) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPlanetBiome(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlanetBiome(uint64(l))
	}
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovPlanetBiome(uint64(l))
	}
	if m.Solid != 0 {
		n += 9
	}
	if m.Liquid != 0 {
		n += 9
	}
	if m.Gas != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPlanetBiome(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlanetBiome(x uint64) (n int) {
	return sovPlanetBiome(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlanetBiome) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlanetBiome
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
			return fmt.Errorf("proto: PlanetBiome: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlanetBiome: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetBiome
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
					return ErrIntOverflowPlanetBiome
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
				return ErrInvalidLengthPlanetBiome
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlanetBiome
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
					return ErrIntOverflowPlanetBiome
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
				return ErrInvalidLengthPlanetBiome
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlanetBiome
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Solid", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Solid = float64(math.Float64frombits(v))
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquid", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Liquid = float64(math.Float64frombits(v))
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Gas = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipPlanetBiome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlanetBiome
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlanetBiome
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
func skipPlanetBiome(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlanetBiome
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
					return 0, ErrIntOverflowPlanetBiome
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
					return 0, ErrIntOverflowPlanetBiome
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
				return 0, ErrInvalidLengthPlanetBiome
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlanetBiome
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlanetBiome
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlanetBiome        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlanetBiome          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlanetBiome = fmt.Errorf("proto: unexpected end of group")
)
