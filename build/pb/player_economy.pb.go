// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/player_economy.proto

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

type GetPlayerEconomyRequest_EconomyTypeEnum int32

const (
	GetPlayerEconomyRequest_MONEY   GetPlayerEconomyRequest_EconomyTypeEnum = 0
	GetPlayerEconomyRequest_DIAMOND GetPlayerEconomyRequest_EconomyTypeEnum = 1
	GetPlayerEconomyRequest_BANK    GetPlayerEconomyRequest_EconomyTypeEnum = 2
)

var GetPlayerEconomyRequest_EconomyTypeEnum_name = map[int32]string{
	0: "MONEY",
	1: "DIAMOND",
	2: "BANK",
}

var GetPlayerEconomyRequest_EconomyTypeEnum_value = map[string]int32{
	"MONEY":   0,
	"DIAMOND": 1,
	"BANK":    2,
}

func (x GetPlayerEconomyRequest_EconomyTypeEnum) String() string {
	return proto.EnumName(GetPlayerEconomyRequest_EconomyTypeEnum_name, int32(x))
}

func (GetPlayerEconomyRequest_EconomyTypeEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fabfea28d0d5694, []int{1, 0}
}

type PlayerEconomy struct {
	ID                    uint32               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Value                 int32                `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	TransactionTypeID     uint32               `protobuf:"varint,3,opt,name=TransactionTypeID,proto3" json:"TransactionTypeID,omitempty"`
	TransactionType       *TransactionType     `protobuf:"bytes,4,opt,name=TransactionType,proto3" json:"TransactionType,omitempty"`
	TransactionCategoryID uint32               `protobuf:"varint,5,opt,name=TransactionCategoryID,proto3" json:"TransactionCategoryID,omitempty"`
	TransactionCategory   *TransactionCategory `protobuf:"bytes,6,opt,name=TransactionCategory,proto3" json:"TransactionCategory,omitempty"`
	PlayerID              uint32               `protobuf:"varint,7,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (m *PlayerEconomy) Reset()         { *m = PlayerEconomy{} }
func (m *PlayerEconomy) String() string { return proto.CompactTextString(m) }
func (*PlayerEconomy) ProtoMessage()    {}
func (*PlayerEconomy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fabfea28d0d5694, []int{0}
}
func (m *PlayerEconomy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerEconomy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerEconomy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlayerEconomy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerEconomy.Merge(m, src)
}
func (m *PlayerEconomy) XXX_Size() int {
	return m.Size()
}
func (m *PlayerEconomy) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerEconomy.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerEconomy proto.InternalMessageInfo

func (m *PlayerEconomy) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PlayerEconomy) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PlayerEconomy) GetTransactionTypeID() uint32 {
	if m != nil {
		return m.TransactionTypeID
	}
	return 0
}

func (m *PlayerEconomy) GetTransactionType() *TransactionType {
	if m != nil {
		return m.TransactionType
	}
	return nil
}

func (m *PlayerEconomy) GetTransactionCategoryID() uint32 {
	if m != nil {
		return m.TransactionCategoryID
	}
	return 0
}

func (m *PlayerEconomy) GetTransactionCategory() *TransactionCategory {
	if m != nil {
		return m.TransactionCategory
	}
	return nil
}

func (m *PlayerEconomy) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

// GetPlayerEconomy
type GetPlayerEconomyRequest struct {
	PlayerID    uint32                                  `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	EconomyType GetPlayerEconomyRequest_EconomyTypeEnum `protobuf:"varint,2,opt,name=EconomyType,proto3,enum=player_economy.GetPlayerEconomyRequest_EconomyTypeEnum" json:"EconomyType,omitempty"`
}

func (m *GetPlayerEconomyRequest) Reset()         { *m = GetPlayerEconomyRequest{} }
func (m *GetPlayerEconomyRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlayerEconomyRequest) ProtoMessage()    {}
func (*GetPlayerEconomyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fabfea28d0d5694, []int{1}
}
func (m *GetPlayerEconomyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlayerEconomyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlayerEconomyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlayerEconomyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerEconomyRequest.Merge(m, src)
}
func (m *GetPlayerEconomyRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPlayerEconomyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerEconomyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerEconomyRequest proto.InternalMessageInfo

func (m *GetPlayerEconomyRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *GetPlayerEconomyRequest) GetEconomyType() GetPlayerEconomyRequest_EconomyTypeEnum {
	if m != nil {
		return m.EconomyType
	}
	return GetPlayerEconomyRequest_MONEY
}

type GetPlayerEconomyResponse struct {
	Value int32 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *GetPlayerEconomyResponse) Reset()         { *m = GetPlayerEconomyResponse{} }
func (m *GetPlayerEconomyResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlayerEconomyResponse) ProtoMessage()    {}
func (*GetPlayerEconomyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fabfea28d0d5694, []int{2}
}
func (m *GetPlayerEconomyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlayerEconomyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlayerEconomyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlayerEconomyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerEconomyResponse.Merge(m, src)
}
func (m *GetPlayerEconomyResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPlayerEconomyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerEconomyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerEconomyResponse proto.InternalMessageInfo

func (m *GetPlayerEconomyResponse) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("player_economy.GetPlayerEconomyRequest_EconomyTypeEnum", GetPlayerEconomyRequest_EconomyTypeEnum_name, GetPlayerEconomyRequest_EconomyTypeEnum_value)
	proto.RegisterType((*PlayerEconomy)(nil), "player_economy.PlayerEconomy")
	proto.RegisterType((*GetPlayerEconomyRequest)(nil), "player_economy.GetPlayerEconomyRequest")
	proto.RegisterType((*GetPlayerEconomyResponse)(nil), "player_economy.GetPlayerEconomyResponse")
}

func init() { proto.RegisterFile("proto/player_economy.proto", fileDescriptor_8fabfea28d0d5694) }

var fileDescriptor_8fabfea28d0d5694 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xce, 0xe4, 0x36, 0x6d, 0xef, 0x29, 0xfd, 0xb9, 0x73, 0xef, 0xc5, 0x21, 0x48, 0x88, 0x59,
	0x55, 0x90, 0x28, 0xad, 0xe0, 0xc2, 0x55, 0x6b, 0x8a, 0x84, 0xd2, 0x56, 0x42, 0x11, 0xaa, 0x8b,
	0x92, 0x96, 0x41, 0x84, 0x36, 0x89, 0xc9, 0x74, 0x91, 0xb7, 0xf0, 0xa5, 0x04, 0x97, 0x5d, 0xba,
	0x53, 0xda, 0x17, 0x11, 0x27, 0xc5, 0xe6, 0xcf, 0xe5, 0x9c, 0xef, 0xe7, 0x9c, 0xf9, 0xf8, 0x40,
	0xf6, 0x7c, 0x97, 0xb9, 0xa7, 0xde, 0xc2, 0x0e, 0xa9, 0x3f, 0xa5, 0x73, 0xd7, 0x71, 0x97, 0xa1,
	0xce, 0x87, 0xb8, 0x96, 0x9c, 0xca, 0x87, 0x11, 0x97, 0xf9, 0xb6, 0x13, 0xd8, 0x73, 0xf6, 0xe8,
	0x3a, 0x53, 0x16, 0x7a, 0x34, 0x62, 0xcb, 0x6a, 0x16, 0x9d, 0xdb, 0x8c, 0x3e, 0xb8, 0xfe, 0xce,
	0x4f, 0x7b, 0x17, 0xa1, 0x7a, 0xc3, 0x2d, 0x7b, 0x91, 0x23, 0xae, 0x81, 0x68, 0x1a, 0x04, 0xa9,
	0xa8, 0x59, 0xb5, 0x44, 0xd3, 0xc0, 0xff, 0x40, 0xba, 0xb5, 0x17, 0x2b, 0x4a, 0x44, 0x15, 0x35,
	0x25, 0x2b, 0x7a, 0xe0, 0x13, 0xf8, 0x33, 0xde, 0xbb, 0x8e, 0x43, 0x8f, 0x9a, 0x06, 0xf9, 0xc5,
	0x45, 0x59, 0x00, 0xf7, 0xa1, 0x9e, 0x1a, 0x92, 0x82, 0x8a, 0x9a, 0x95, 0xd6, 0x91, 0x9e, 0xb9,
	0x3c, 0x45, 0xb4, 0xd2, 0x4a, 0x7c, 0x0e, 0xff, 0x63, 0xa3, 0xab, 0xdd, 0x7f, 0x4c, 0x83, 0x48,
	0x7c, 0x7d, 0x3e, 0x88, 0xef, 0xe1, 0x6f, 0x0e, 0x40, 0x8a, 0xfc, 0x8c, 0x63, 0x3d, 0x37, 0xa2,
	0x1c, 0x81, 0x95, 0xe7, 0x82, 0x65, 0x28, 0x47, 0x21, 0x9a, 0x06, 0x29, 0xf1, 0x2b, 0xbe, 0xdf,
	0xda, 0x0b, 0x82, 0x83, 0x6b, 0xca, 0x12, 0x21, 0x5b, 0xf4, 0x69, 0x45, 0x03, 0x96, 0xd0, 0xa1,
	0xa4, 0x0e, 0x4f, 0xa0, 0xb2, 0x63, 0xf3, 0xbc, 0xbe, 0xd2, 0xaf, 0xb5, 0x2e, 0xf4, 0x54, 0x2b,
	0x7e, 0x70, 0xd6, 0x63, 0xd2, 0x9e, 0xb3, 0x5a, 0x5a, 0x71, 0x2f, 0xad, 0x0d, 0xf5, 0x14, 0x8e,
	0x7f, 0x83, 0x34, 0x18, 0x0d, 0x7b, 0x93, 0x86, 0x80, 0x2b, 0x50, 0x32, 0xcc, 0xce, 0x60, 0x34,
	0x34, 0x1a, 0x08, 0x97, 0xa1, 0xd0, 0xed, 0x0c, 0xfb, 0x0d, 0x51, 0x3b, 0x03, 0x92, 0x5d, 0x16,
	0x78, 0xae, 0x13, 0xd0, 0x7d, 0x47, 0x50, 0xac, 0x23, 0x5d, 0xe5, 0x75, 0xa3, 0xa0, 0xf5, 0x46,
	0x41, 0x1f, 0x1b, 0x05, 0x3d, 0x6f, 0x15, 0x61, 0xbd, 0x55, 0x84, 0xb7, 0xad, 0x22, 0xdc, 0x15,
	0xf4, 0x4b, 0x6f, 0x36, 0x2b, 0xf2, 0x0a, 0xb6, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xec, 0x72,
	0x84, 0xc1, 0xf0, 0x02, 0x00, 0x00,
}

func (m *PlayerEconomy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerEconomy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlayerEconomy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlayerID != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x38
	}
	if m.TransactionCategory != nil {
		{
			size, err := m.TransactionCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlayerEconomy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.TransactionCategoryID != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.TransactionCategoryID))
		i--
		dAtA[i] = 0x28
	}
	if m.TransactionType != nil {
		{
			size, err := m.TransactionType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlayerEconomy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TransactionTypeID != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.TransactionTypeID))
		i--
		dAtA[i] = 0x18
	}
	if m.Value != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlayerEconomyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlayerEconomyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlayerEconomyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EconomyType != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.EconomyType))
		i--
		dAtA[i] = 0x10
	}
	if m.PlayerID != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlayerEconomyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlayerEconomyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlayerEconomyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintPlayerEconomy(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlayerEconomy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlayerEconomy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlayerEconomy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.ID))
	}
	if m.Value != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.Value))
	}
	if m.TransactionTypeID != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.TransactionTypeID))
	}
	if m.TransactionType != nil {
		l = m.TransactionType.Size()
		n += 1 + l + sovPlayerEconomy(uint64(l))
	}
	if m.TransactionCategoryID != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.TransactionCategoryID))
	}
	if m.TransactionCategory != nil {
		l = m.TransactionCategory.Size()
		n += 1 + l + sovPlayerEconomy(uint64(l))
	}
	if m.PlayerID != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.PlayerID))
	}
	return n
}

func (m *GetPlayerEconomyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlayerID != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.PlayerID))
	}
	if m.EconomyType != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.EconomyType))
	}
	return n
}

func (m *GetPlayerEconomyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovPlayerEconomy(uint64(m.Value))
	}
	return n
}

func sovPlayerEconomy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlayerEconomy(x uint64) (n int) {
	return sovPlayerEconomy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerEconomy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerEconomy
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
			return fmt.Errorf("proto: PlayerEconomy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerEconomy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionTypeID", wireType)
			}
			m.TransactionTypeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionTypeID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
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
				return ErrInvalidLengthPlayerEconomy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerEconomy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransactionType == nil {
				m.TransactionType = &TransactionType{}
			}
			if err := m.TransactionType.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionCategoryID", wireType)
			}
			m.TransactionCategoryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionCategoryID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
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
				return ErrInvalidLengthPlayerEconomy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerEconomy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransactionCategory == nil {
				m.TransactionCategory = &TransactionCategory{}
			}
			if err := m.TransactionCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
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
			skippy, err := skipPlayerEconomy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerEconomy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerEconomy
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
func (m *GetPlayerEconomyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerEconomy
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
			return fmt.Errorf("proto: GetPlayerEconomyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlayerEconomyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EconomyType", wireType)
			}
			m.EconomyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EconomyType |= GetPlayerEconomyRequest_EconomyTypeEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerEconomy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerEconomy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerEconomy
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
func (m *GetPlayerEconomyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerEconomy
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
			return fmt.Errorf("proto: GetPlayerEconomyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlayerEconomyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerEconomy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerEconomy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerEconomy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerEconomy
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
func skipPlayerEconomy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayerEconomy
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
					return 0, ErrIntOverflowPlayerEconomy
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
					return 0, ErrIntOverflowPlayerEconomy
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
				return 0, ErrInvalidLengthPlayerEconomy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlayerEconomy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlayerEconomy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlayerEconomy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayerEconomy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlayerEconomy = fmt.Errorf("proto: unexpected end of group")
)
