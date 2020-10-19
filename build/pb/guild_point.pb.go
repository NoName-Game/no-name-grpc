// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/guild_point.proto

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

// GetGuildPoints
type GetGuildPointsRequest struct {
	GuildID              uint32   `protobuf:"varint,1,opt,name=GuildID,proto3" json:"GuildID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGuildPointsRequest) Reset()         { *m = GetGuildPointsRequest{} }
func (m *GetGuildPointsRequest) String() string { return proto.CompactTextString(m) }
func (*GetGuildPointsRequest) ProtoMessage()    {}
func (*GetGuildPointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_414ed415941276ac, []int{0}
}
func (m *GetGuildPointsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetGuildPointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetGuildPointsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetGuildPointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGuildPointsRequest.Merge(m, src)
}
func (m *GetGuildPointsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetGuildPointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGuildPointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGuildPointsRequest proto.InternalMessageInfo

func (m *GetGuildPointsRequest) GetGuildID() uint32 {
	if m != nil {
		return m.GuildID
	}
	return 0
}

type GetGuildPointsResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGuildPointsResponse) Reset()         { *m = GetGuildPointsResponse{} }
func (m *GetGuildPointsResponse) String() string { return proto.CompactTextString(m) }
func (*GetGuildPointsResponse) ProtoMessage()    {}
func (*GetGuildPointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_414ed415941276ac, []int{1}
}
func (m *GetGuildPointsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetGuildPointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetGuildPointsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetGuildPointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGuildPointsResponse.Merge(m, src)
}
func (m *GetGuildPointsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetGuildPointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGuildPointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGuildPointsResponse proto.InternalMessageInfo

func (m *GetGuildPointsResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// GetPlayerGuildPoints
type GetPlayerGuildPointsRequest struct {
	GuildID              uint32   `protobuf:"varint,1,opt,name=GuildID,proto3" json:"GuildID,omitempty"`
	PlayerID             uint32   `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerGuildPointsRequest) Reset()         { *m = GetPlayerGuildPointsRequest{} }
func (m *GetPlayerGuildPointsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlayerGuildPointsRequest) ProtoMessage()    {}
func (*GetPlayerGuildPointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_414ed415941276ac, []int{2}
}
func (m *GetPlayerGuildPointsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlayerGuildPointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlayerGuildPointsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlayerGuildPointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerGuildPointsRequest.Merge(m, src)
}
func (m *GetPlayerGuildPointsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPlayerGuildPointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerGuildPointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerGuildPointsRequest proto.InternalMessageInfo

func (m *GetPlayerGuildPointsRequest) GetGuildID() uint32 {
	if m != nil {
		return m.GuildID
	}
	return 0
}

func (m *GetPlayerGuildPointsRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

type GetPlayerGuildPointsResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerGuildPointsResponse) Reset()         { *m = GetPlayerGuildPointsResponse{} }
func (m *GetPlayerGuildPointsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlayerGuildPointsResponse) ProtoMessage()    {}
func (*GetPlayerGuildPointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_414ed415941276ac, []int{3}
}
func (m *GetPlayerGuildPointsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlayerGuildPointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlayerGuildPointsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlayerGuildPointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerGuildPointsResponse.Merge(m, src)
}
func (m *GetPlayerGuildPointsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPlayerGuildPointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerGuildPointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerGuildPointsResponse proto.InternalMessageInfo

func (m *GetPlayerGuildPointsResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*GetGuildPointsRequest)(nil), "guild_point.GetGuildPointsRequest")
	proto.RegisterType((*GetGuildPointsResponse)(nil), "guild_point.GetGuildPointsResponse")
	proto.RegisterType((*GetPlayerGuildPointsRequest)(nil), "guild_point.GetPlayerGuildPointsRequest")
	proto.RegisterType((*GetPlayerGuildPointsResponse)(nil), "guild_point.GetPlayerGuildPointsResponse")
}

func init() { proto.RegisterFile("proto/guild_point.proto", fileDescriptor_414ed415941276ac) }

var fileDescriptor_414ed415941276ac = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0xcd, 0xcc, 0x49, 0x89, 0x2f, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x03, 0x8b,
	0x08, 0x71, 0x23, 0x09, 0x29, 0x19, 0x72, 0x89, 0xba, 0xa7, 0x96, 0xb8, 0x83, 0x44, 0x02, 0x40,
	0x02, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0x60, 0x51, 0x4f,
	0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x18, 0x57, 0xc9, 0x80, 0x4b, 0x0c, 0x5d, 0x4b,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x50, 0x6a, 0x71, 0x69, 0x4e, 0x09,
	0x58, 0x0b, 0x6b, 0x10, 0x94, 0xa7, 0x14, 0xcc, 0x25, 0xed, 0x9e, 0x5a, 0x12, 0x90, 0x93, 0x58,
	0x99, 0x5a, 0x44, 0x8a, 0x55, 0x42, 0x52, 0x5c, 0x1c, 0x10, 0x5d, 0x9e, 0x2e, 0x12, 0x4c, 0x60,
	0x29, 0x38, 0x5f, 0xc9, 0x8c, 0x4b, 0x06, 0xbb, 0xa1, 0xf8, 0x1d, 0xe3, 0x24, 0x74, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0xb1, 0xe8, 0x59, 0x17, 0x24,
	0x25, 0xb1, 0x81, 0x43, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x44, 0x8e, 0x51, 0xbf, 0x34,
	0x01, 0x00, 0x00,
}

func (m *GetGuildPointsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetGuildPointsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetGuildPointsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.GuildID != 0 {
		i = encodeVarintGuildPoint(dAtA, i, uint64(m.GuildID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetGuildPointsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetGuildPointsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetGuildPointsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Result != 0 {
		i = encodeVarintGuildPoint(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlayerGuildPointsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlayerGuildPointsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlayerGuildPointsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PlayerID != 0 {
		i = encodeVarintGuildPoint(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x10
	}
	if m.GuildID != 0 {
		i = encodeVarintGuildPoint(dAtA, i, uint64(m.GuildID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlayerGuildPointsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlayerGuildPointsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlayerGuildPointsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Result != 0 {
		i = encodeVarintGuildPoint(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGuildPoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovGuildPoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetGuildPointsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GuildID != 0 {
		n += 1 + sovGuildPoint(uint64(m.GuildID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetGuildPointsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovGuildPoint(uint64(m.Result))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetPlayerGuildPointsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GuildID != 0 {
		n += 1 + sovGuildPoint(uint64(m.GuildID))
	}
	if m.PlayerID != 0 {
		n += 1 + sovGuildPoint(uint64(m.PlayerID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetPlayerGuildPointsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovGuildPoint(uint64(m.Result))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGuildPoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuildPoint(x uint64) (n int) {
	return sovGuildPoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetGuildPointsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildPoint
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
			return fmt.Errorf("proto: GetGuildPointsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetGuildPointsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuildID", wireType)
			}
			m.GuildID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GuildID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuildPoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildPoint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGuildPoint
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
func (m *GetGuildPointsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildPoint
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
			return fmt.Errorf("proto: GetGuildPointsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetGuildPointsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuildPoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildPoint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGuildPoint
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
func (m *GetPlayerGuildPointsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildPoint
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
			return fmt.Errorf("proto: GetPlayerGuildPointsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlayerGuildPointsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuildID", wireType)
			}
			m.GuildID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GuildID |= uint32(b&0x7F) << shift
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
					return ErrIntOverflowGuildPoint
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
			skippy, err := skipGuildPoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildPoint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGuildPoint
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
func (m *GetPlayerGuildPointsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuildPoint
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
			return fmt.Errorf("proto: GetPlayerGuildPointsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlayerGuildPointsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuildPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuildPoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuildPoint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGuildPoint
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
func skipGuildPoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuildPoint
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
					return 0, ErrIntOverflowGuildPoint
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
					return 0, ErrIntOverflowGuildPoint
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
				return 0, ErrInvalidLengthGuildPoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGuildPoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGuildPoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGuildPoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuildPoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGuildPoint = fmt.Errorf("proto: unexpected end of group")
)
