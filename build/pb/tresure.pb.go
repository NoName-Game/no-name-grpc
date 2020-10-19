// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/tresure.proto

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

type Tresure struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlanetMapID          uint32   `protobuf:"varint,3,opt,name=PlanetMapID,proto3" json:"PlanetMapID,omitempty"`
	PositionX            int32    `protobuf:"varint,4,opt,name=PositionX,proto3" json:"PositionX,omitempty"`
	PositionY            int32    `protobuf:"varint,5,opt,name=PositionY,proto3" json:"PositionY,omitempty"`
	RarityID             uint32   `protobuf:"varint,6,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity               *Rarity  `protobuf:"bytes,7,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tresure) Reset()         { *m = Tresure{} }
func (m *Tresure) String() string { return proto.CompactTextString(m) }
func (*Tresure) ProtoMessage()    {}
func (*Tresure) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2123bb5e2d1ca62, []int{0}
}
func (m *Tresure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tresure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tresure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tresure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tresure.Merge(m, src)
}
func (m *Tresure) XXX_Size() int {
	return m.Size()
}
func (m *Tresure) XXX_DiscardUnknown() {
	xxx_messageInfo_Tresure.DiscardUnknown(m)
}

var xxx_messageInfo_Tresure proto.InternalMessageInfo

func (m *Tresure) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Tresure) GetPlanetMapID() uint32 {
	if m != nil {
		return m.PlanetMapID
	}
	return 0
}

func (m *Tresure) GetPositionX() int32 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *Tresure) GetPositionY() int32 {
	if m != nil {
		return m.PositionY
	}
	return 0
}

func (m *Tresure) GetRarityID() uint32 {
	if m != nil {
		return m.RarityID
	}
	return 0
}

func (m *Tresure) GetRarity() *Rarity {
	if m != nil {
		return m.Rarity
	}
	return nil
}

// DropTresure
type DropTresureRequest struct {
	TresureID            uint32   `protobuf:"varint,1,opt,name=TresureID,proto3" json:"TresureID,omitempty"`
	PlayerID             uint32   `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DropTresureRequest) Reset()         { *m = DropTresureRequest{} }
func (m *DropTresureRequest) String() string { return proto.CompactTextString(m) }
func (*DropTresureRequest) ProtoMessage()    {}
func (*DropTresureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2123bb5e2d1ca62, []int{1}
}
func (m *DropTresureRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DropTresureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DropTresureRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DropTresureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropTresureRequest.Merge(m, src)
}
func (m *DropTresureRequest) XXX_Size() int {
	return m.Size()
}
func (m *DropTresureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DropTresureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DropTresureRequest proto.InternalMessageInfo

func (m *DropTresureRequest) GetTresureID() uint32 {
	if m != nil {
		return m.TresureID
	}
	return 0
}

func (m *DropTresureRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

type DropTresureResponse struct {
	Item                 *Item          `protobuf:"bytes,1,opt,name=Item,proto3" json:"Item,omitempty"`
	Resource             *Resource      `protobuf:"bytes,2,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Transaction          *PlayerEconomy `protobuf:"bytes,3,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
	Trap                 *Trap          `protobuf:"bytes,4,opt,name=Trap,proto3" json:"Trap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DropTresureResponse) Reset()         { *m = DropTresureResponse{} }
func (m *DropTresureResponse) String() string { return proto.CompactTextString(m) }
func (*DropTresureResponse) ProtoMessage()    {}
func (*DropTresureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2123bb5e2d1ca62, []int{2}
}
func (m *DropTresureResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DropTresureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DropTresureResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DropTresureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropTresureResponse.Merge(m, src)
}
func (m *DropTresureResponse) XXX_Size() int {
	return m.Size()
}
func (m *DropTresureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DropTresureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DropTresureResponse proto.InternalMessageInfo

func (m *DropTresureResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *DropTresureResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *DropTresureResponse) GetTransaction() *PlayerEconomy {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *DropTresureResponse) GetTrap() *Trap {
	if m != nil {
		return m.Trap
	}
	return nil
}

func init() {
	proto.RegisterType((*Tresure)(nil), "tresure.Tresure")
	proto.RegisterType((*DropTresureRequest)(nil), "tresure.DropTresureRequest")
	proto.RegisterType((*DropTresureResponse)(nil), "tresure.DropTresureResponse")
}

func init() { proto.RegisterFile("proto/tresure.proto", fileDescriptor_f2123bb5e2d1ca62) }

var fileDescriptor_f2123bb5e2d1ca62 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe5, 0xd2, 0x1b, 0x8e, 0xa8, 0x90, 0xcb, 0x10, 0x45, 0x10, 0x45, 0x1d, 0x50, 0x27,
	0x23, 0x85, 0x91, 0x01, 0x09, 0x85, 0x21, 0x03, 0xa8, 0xb2, 0x3a, 0x50, 0x16, 0xe4, 0x56, 0x1e,
	0x22, 0xb5, 0xb1, 0xb1, 0xdd, 0xa1, 0xcf, 0xc6, 0x8e, 0x18, 0x79, 0x04, 0xd4, 0x27, 0x41, 0x3e,
	0x76, 0xd3, 0xcb, 0x96, 0xf3, 0xfd, 0xce, 0x7f, 0x2e, 0x3f, 0x1e, 0x2a, 0x2d, 0xad, 0xbc, 0xb3,
	0x5a, 0x98, 0xb5, 0x16, 0x14, 0x2a, 0xd2, 0x0b, 0x65, 0x42, 0xbc, 0xaa, 0xb9, 0xae, 0xec, 0xc6,
	0x8b, 0xc9, 0xa5, 0x67, 0x95, 0x15, 0xab, 0x40, 0xae, 0xc2, 0x2b, 0x61, 0xe4, 0x5a, 0x2f, 0x82,
	0x49, 0x92, 0x78, 0xaa, 0x96, 0x7c, 0x23, 0xf4, 0x87, 0x58, 0xc8, 0x5a, 0xae, 0x4e, 0x3c, 0xac,
	0xe6, 0xca, 0x93, 0xd1, 0x17, 0xc2, 0xbd, 0xa9, 0xef, 0x4a, 0x06, 0xb8, 0x55, 0x16, 0x31, 0xca,
	0xd0, 0xf8, 0x82, 0xb5, 0xca, 0x82, 0x64, 0x38, 0x9a, 0x2c, 0x79, 0x2d, 0xec, 0x0b, 0x57, 0x65,
	0x11, 0x9f, 0x81, 0x70, 0x88, 0xc8, 0x35, 0x3e, 0x9f, 0x48, 0x53, 0xd9, 0x4a, 0xd6, 0x6f, 0x71,
	0x3b, 0x43, 0xe3, 0x0e, 0xdb, 0x83, 0x43, 0x75, 0x16, 0x77, 0x8e, 0xd5, 0x19, 0x49, 0x70, 0x9f,
	0xc1, 0x7e, 0x65, 0x11, 0x77, 0xc1, 0xba, 0xa9, 0xc9, 0x2d, 0xee, 0xfa, 0xef, 0xb8, 0x97, 0xa1,
	0x71, 0x94, 0x0f, 0x68, 0x38, 0x85, 0xa7, 0x2c, 0xa8, 0xa3, 0x57, 0x4c, 0x0a, 0x2d, 0x55, 0x58,
	0x80, 0x89, 0xcf, 0xb5, 0x30, 0xd6, 0xf5, 0x0d, 0xa4, 0x59, 0x67, 0x0f, 0x5c, 0xdf, 0x09, 0xdc,
	0xa6, 0x2c, 0xe2, 0x96, 0xef, 0xbb, 0xab, 0x47, 0xdf, 0x08, 0x0f, 0x8f, 0x0c, 0x8d, 0x92, 0xb5,
	0x11, 0x24, 0xc5, 0xed, 0xd2, 0x8a, 0x15, 0x98, 0x45, 0x39, 0xa6, 0x10, 0x82, 0x23, 0x0c, 0x38,
	0xa1, 0xb8, 0xcf, 0x42, 0x0a, 0xe0, 0x19, 0xe5, 0x84, 0x36, 0xb1, 0xec, 0x14, 0xd6, 0xbc, 0x21,
	0x8f, 0x38, 0x9a, 0x6a, 0x5e, 0x1b, 0xbe, 0x70, 0xb7, 0x80, 0xcb, 0x46, 0xf9, 0x0d, 0x3d, 0xc9,
	0xcc, 0x8f, 0xf5, 0xec, 0x2b, 0x76, 0xf8, 0x87, 0x1b, 0x68, 0xaa, 0xb9, 0x82, 0x9b, 0xbb, 0x81,
	0x20, 0x51, 0x47, 0x18, 0xf0, 0x27, 0xf2, 0xb3, 0x4d, 0xd1, 0xef, 0x36, 0x45, 0x7f, 0xdb, 0x14,
	0xbd, 0xb7, 0xe9, 0x83, 0x9a, 0xcf, 0xbb, 0x90, 0xf8, 0xfd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x79, 0xbf, 0x77, 0xa7, 0x7b, 0x02, 0x00, 0x00,
}

func (m *Tresure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tresure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tresure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Rarity != nil {
		{
			size, err := m.Rarity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTresure(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.RarityID != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.RarityID))
		i--
		dAtA[i] = 0x30
	}
	if m.PositionY != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.PositionY))
		i--
		dAtA[i] = 0x28
	}
	if m.PositionX != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.PositionX))
		i--
		dAtA[i] = 0x20
	}
	if m.PlanetMapID != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.PlanetMapID))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DropTresureRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DropTresureRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DropTresureRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PlayerID != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x10
	}
	if m.TresureID != 0 {
		i = encodeVarintTresure(dAtA, i, uint64(m.TresureID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DropTresureResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DropTresureResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DropTresureResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Trap != nil {
		{
			size, err := m.Trap.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTresure(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Transaction != nil {
		{
			size, err := m.Transaction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTresure(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTresure(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Item != nil {
		{
			size, err := m.Item.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTresure(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTresure(dAtA []byte, offset int, v uint64) int {
	offset -= sovTresure(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tresure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTresure(uint64(m.ID))
	}
	if m.PlanetMapID != 0 {
		n += 1 + sovTresure(uint64(m.PlanetMapID))
	}
	if m.PositionX != 0 {
		n += 1 + sovTresure(uint64(m.PositionX))
	}
	if m.PositionY != 0 {
		n += 1 + sovTresure(uint64(m.PositionY))
	}
	if m.RarityID != 0 {
		n += 1 + sovTresure(uint64(m.RarityID))
	}
	if m.Rarity != nil {
		l = m.Rarity.Size()
		n += 1 + l + sovTresure(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DropTresureRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TresureID != 0 {
		n += 1 + sovTresure(uint64(m.TresureID))
	}
	if m.PlayerID != 0 {
		n += 1 + sovTresure(uint64(m.PlayerID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DropTresureResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Item != nil {
		l = m.Item.Size()
		n += 1 + l + sovTresure(uint64(l))
	}
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovTresure(uint64(l))
	}
	if m.Transaction != nil {
		l = m.Transaction.Size()
		n += 1 + l + sovTresure(uint64(l))
	}
	if m.Trap != nil {
		l = m.Trap.Size()
		n += 1 + l + sovTresure(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTresure(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTresure(x uint64) (n int) {
	return sovTresure(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tresure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTresure
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
			return fmt.Errorf("proto: Tresure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tresure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanetMapID", wireType)
			}
			m.PlanetMapID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanetMapID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionX", wireType)
			}
			m.PositionX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PositionX |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionY", wireType)
			}
			m.PositionY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PositionY |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RarityID", wireType)
			}
			m.RarityID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rarity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
				return ErrInvalidLengthTresure
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTresure
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
		default:
			iNdEx = preIndex
			skippy, err := skipTresure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTresure
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTresure
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
func (m *DropTresureRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTresure
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
			return fmt.Errorf("proto: DropTresureRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DropTresureRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TresureID", wireType)
			}
			m.TresureID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TresureID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlayerID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTresure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTresure
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTresure
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
func (m *DropTresureResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTresure
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
			return fmt.Errorf("proto: DropTresureResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DropTresureResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Item", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
				return ErrInvalidLengthTresure
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTresure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Item == nil {
				m.Item = &Item{}
			}
			if err := m.Item.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
				return ErrInvalidLengthTresure
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTresure
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
				return ErrInvalidLengthTresure
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTresure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Transaction == nil {
				m.Transaction = &PlayerEconomy{}
			}
			if err := m.Transaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTresure
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
				return ErrInvalidLengthTresure
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTresure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trap == nil {
				m.Trap = &Trap{}
			}
			if err := m.Trap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTresure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTresure
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTresure
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
func skipTresure(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTresure
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
					return 0, ErrIntOverflowTresure
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
					return 0, ErrIntOverflowTresure
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
				return 0, ErrInvalidLengthTresure
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTresure
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTresure
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTresure        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTresure          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTresure = fmt.Errorf("proto: unexpected end of group")
)
