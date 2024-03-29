// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/titan_damage.proto

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

type TitanDamage struct {
	ID              uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerID        uint32 `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	TitanID         uint32 `protobuf:"varint,3,opt,name=TitanID,proto3" json:"TitanID,omitempty"`
	DamageInflicted int32  `protobuf:"varint,4,opt,name=DamageInflicted,proto3" json:"DamageInflicted,omitempty"`
}

func (m *TitanDamage) Reset()         { *m = TitanDamage{} }
func (m *TitanDamage) String() string { return proto.CompactTextString(m) }
func (*TitanDamage) ProtoMessage()    {}
func (*TitanDamage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf06fac82462d5e, []int{0}
}
func (m *TitanDamage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TitanDamage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TitanDamage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TitanDamage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TitanDamage.Merge(m, src)
}
func (m *TitanDamage) XXX_Size() int {
	return m.Size()
}
func (m *TitanDamage) XXX_DiscardUnknown() {
	xxx_messageInfo_TitanDamage.DiscardUnknown(m)
}

var xxx_messageInfo_TitanDamage proto.InternalMessageInfo

func (m *TitanDamage) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TitanDamage) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *TitanDamage) GetTitanID() uint32 {
	if m != nil {
		return m.TitanID
	}
	return 0
}

func (m *TitanDamage) GetDamageInflicted() int32 {
	if m != nil {
		return m.DamageInflicted
	}
	return 0
}

// GetTitanDamageByTitanID
type GetTitanDamageByTitanIDRequest struct {
	TitanID uint32 `protobuf:"varint,1,opt,name=TitanID,proto3" json:"TitanID,omitempty"`
}

func (m *GetTitanDamageByTitanIDRequest) Reset()         { *m = GetTitanDamageByTitanIDRequest{} }
func (m *GetTitanDamageByTitanIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetTitanDamageByTitanIDRequest) ProtoMessage()    {}
func (*GetTitanDamageByTitanIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf06fac82462d5e, []int{1}
}
func (m *GetTitanDamageByTitanIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTitanDamageByTitanIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTitanDamageByTitanIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTitanDamageByTitanIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTitanDamageByTitanIDRequest.Merge(m, src)
}
func (m *GetTitanDamageByTitanIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTitanDamageByTitanIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTitanDamageByTitanIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTitanDamageByTitanIDRequest proto.InternalMessageInfo

func (m *GetTitanDamageByTitanIDRequest) GetTitanID() uint32 {
	if m != nil {
		return m.TitanID
	}
	return 0
}

type GetTitanDamageByTitanIDResponse struct {
	Damages []*TitanDamage `protobuf:"bytes,1,rep,name=damages,proto3" json:"damages,omitempty"`
}

func (m *GetTitanDamageByTitanIDResponse) Reset()         { *m = GetTitanDamageByTitanIDResponse{} }
func (m *GetTitanDamageByTitanIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetTitanDamageByTitanIDResponse) ProtoMessage()    {}
func (*GetTitanDamageByTitanIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf06fac82462d5e, []int{2}
}
func (m *GetTitanDamageByTitanIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTitanDamageByTitanIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTitanDamageByTitanIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTitanDamageByTitanIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTitanDamageByTitanIDResponse.Merge(m, src)
}
func (m *GetTitanDamageByTitanIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTitanDamageByTitanIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTitanDamageByTitanIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTitanDamageByTitanIDResponse proto.InternalMessageInfo

func (m *GetTitanDamageByTitanIDResponse) GetDamages() []*TitanDamage {
	if m != nil {
		return m.Damages
	}
	return nil
}

func init() {
	proto.RegisterType((*TitanDamage)(nil), "titan_damage.TitanDamage")
	proto.RegisterType((*GetTitanDamageByTitanIDRequest)(nil), "titan_damage.GetTitanDamageByTitanIDRequest")
	proto.RegisterType((*GetTitanDamageByTitanIDResponse)(nil), "titan_damage.GetTitanDamageByTitanIDResponse")
}

func init() { proto.RegisterFile("proto/titan_damage.proto", fileDescriptor_1cf06fac82462d5e) }

var fileDescriptor_1cf06fac82462d5e = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc9, 0x2c, 0x49, 0xcc, 0x8b, 0x4f, 0x49, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x03,
	0x0b, 0x09, 0xf1, 0x20, 0x8b, 0x29, 0xd5, 0x72, 0x71, 0x87, 0x80, 0xf8, 0x2e, 0x60, 0xae, 0x10,
	0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x93, 0xa7, 0x8b, 0x90,
	0x14, 0x17, 0x47, 0x40, 0x4e, 0x62, 0x65, 0x6a, 0x91, 0xa7, 0x8b, 0x04, 0x13, 0x58, 0x14, 0xce,
	0x17, 0x92, 0xe0, 0x62, 0x07, 0x6b, 0xf5, 0x74, 0x91, 0x60, 0x06, 0x4b, 0xc1, 0xb8, 0x42, 0x1a,
	0x5c, 0xfc, 0x10, 0xf3, 0x3c, 0xf3, 0xd2, 0x72, 0x32, 0x93, 0x4b, 0x52, 0x53, 0x24, 0x58, 0x14,
	0x18, 0x35, 0x58, 0x83, 0xd0, 0x85, 0x95, 0xac, 0xb8, 0xe4, 0xdc, 0x53, 0x4b, 0x90, 0x5c, 0xe0,
	0x54, 0x09, 0x35, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x04, 0xd9, 0x16, 0x46, 0x14, 0x5b,
	0x94, 0xc2, 0xb8, 0xe4, 0x71, 0xea, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x32, 0xe6, 0x62,
	0x87, 0xf8, 0xb3, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x52, 0x0f, 0x25, 0x44, 0x90,
	0x34, 0x07, 0xc1, 0x54, 0x3a, 0xc9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43,
	0x14, 0x8b, 0x9e, 0x75, 0x41, 0x52, 0x12, 0x1b, 0x38, 0x1c, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2d, 0x65, 0xe9, 0xc1, 0x63, 0x01, 0x00, 0x00,
}

func (m *TitanDamage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TitanDamage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TitanDamage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DamageInflicted != 0 {
		i = encodeVarintTitanDamage(dAtA, i, uint64(m.DamageInflicted))
		i--
		dAtA[i] = 0x20
	}
	if m.TitanID != 0 {
		i = encodeVarintTitanDamage(dAtA, i, uint64(m.TitanID))
		i--
		dAtA[i] = 0x18
	}
	if m.PlayerID != 0 {
		i = encodeVarintTitanDamage(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintTitanDamage(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetTitanDamageByTitanIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTitanDamageByTitanIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTitanDamageByTitanIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TitanID != 0 {
		i = encodeVarintTitanDamage(dAtA, i, uint64(m.TitanID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetTitanDamageByTitanIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTitanDamageByTitanIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTitanDamageByTitanIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Damages) > 0 {
		for iNdEx := len(m.Damages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Damages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTitanDamage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTitanDamage(dAtA []byte, offset int, v uint64) int {
	offset -= sovTitanDamage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TitanDamage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTitanDamage(uint64(m.ID))
	}
	if m.PlayerID != 0 {
		n += 1 + sovTitanDamage(uint64(m.PlayerID))
	}
	if m.TitanID != 0 {
		n += 1 + sovTitanDamage(uint64(m.TitanID))
	}
	if m.DamageInflicted != 0 {
		n += 1 + sovTitanDamage(uint64(m.DamageInflicted))
	}
	return n
}

func (m *GetTitanDamageByTitanIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TitanID != 0 {
		n += 1 + sovTitanDamage(uint64(m.TitanID))
	}
	return n
}

func (m *GetTitanDamageByTitanIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Damages) > 0 {
		for _, e := range m.Damages {
			l = e.Size()
			n += 1 + l + sovTitanDamage(uint64(l))
		}
	}
	return n
}

func sovTitanDamage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTitanDamage(x uint64) (n int) {
	return sovTitanDamage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TitanDamage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTitanDamage
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
			return fmt.Errorf("proto: TitanDamage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TitanDamage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTitanDamage
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
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTitanDamage
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TitanID", wireType)
			}
			m.TitanID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTitanDamage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TitanID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DamageInflicted", wireType)
			}
			m.DamageInflicted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTitanDamage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DamageInflicted |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTitanDamage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTitanDamage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTitanDamage
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
func (m *GetTitanDamageByTitanIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTitanDamage
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
			return fmt.Errorf("proto: GetTitanDamageByTitanIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTitanDamageByTitanIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TitanID", wireType)
			}
			m.TitanID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTitanDamage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TitanID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTitanDamage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTitanDamage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTitanDamage
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
func (m *GetTitanDamageByTitanIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTitanDamage
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
			return fmt.Errorf("proto: GetTitanDamageByTitanIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTitanDamageByTitanIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Damages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTitanDamage
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
				return ErrInvalidLengthTitanDamage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTitanDamage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Damages = append(m.Damages, &TitanDamage{})
			if err := m.Damages[len(m.Damages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTitanDamage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTitanDamage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTitanDamage
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
func skipTitanDamage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTitanDamage
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
					return 0, ErrIntOverflowTitanDamage
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
					return 0, ErrIntOverflowTitanDamage
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
				return 0, ErrInvalidLengthTitanDamage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTitanDamage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTitanDamage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTitanDamage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTitanDamage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTitanDamage = fmt.Errorf("proto: unexpected end of group")
)
