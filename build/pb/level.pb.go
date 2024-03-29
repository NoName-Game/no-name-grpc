// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/level.proto

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

type Level struct {
	ID               uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ExperienceNeeded int32  `protobuf:"varint,2,opt,name=ExperienceNeeded,proto3" json:"ExperienceNeeded,omitempty"`
	PlayerMaxLife    int64  `protobuf:"varint,3,opt,name=PlayerMaxLife,proto3" json:"PlayerMaxLife,omitempty"`
}

func (m *Level) Reset()         { *m = Level{} }
func (m *Level) String() string { return proto.CompactTextString(m) }
func (*Level) ProtoMessage()    {}
func (*Level) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a5dd38195ba9bd7, []int{0}
}
func (m *Level) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Level) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Level.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Level) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Level.Merge(m, src)
}
func (m *Level) XXX_Size() int {
	return m.Size()
}
func (m *Level) XXX_DiscardUnknown() {
	xxx_messageInfo_Level.DiscardUnknown(m)
}

var xxx_messageInfo_Level proto.InternalMessageInfo

func (m *Level) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Level) GetExperienceNeeded() int32 {
	if m != nil {
		return m.ExperienceNeeded
	}
	return 0
}

func (m *Level) GetPlayerMaxLife() int64 {
	if m != nil {
		return m.PlayerMaxLife
	}
	return 0
}

// GetLevelByID
type GetLevelByIDRequest struct {
	LevelID uint32 `protobuf:"varint,1,opt,name=LevelID,proto3" json:"LevelID,omitempty"`
}

func (m *GetLevelByIDRequest) Reset()         { *m = GetLevelByIDRequest{} }
func (m *GetLevelByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetLevelByIDRequest) ProtoMessage()    {}
func (*GetLevelByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a5dd38195ba9bd7, []int{1}
}
func (m *GetLevelByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLevelByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLevelByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLevelByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLevelByIDRequest.Merge(m, src)
}
func (m *GetLevelByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLevelByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLevelByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLevelByIDRequest proto.InternalMessageInfo

func (m *GetLevelByIDRequest) GetLevelID() uint32 {
	if m != nil {
		return m.LevelID
	}
	return 0
}

type GetLevelByIDResponse struct {
	Level *Level `protobuf:"bytes,1,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *GetLevelByIDResponse) Reset()         { *m = GetLevelByIDResponse{} }
func (m *GetLevelByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetLevelByIDResponse) ProtoMessage()    {}
func (*GetLevelByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a5dd38195ba9bd7, []int{2}
}
func (m *GetLevelByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLevelByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLevelByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLevelByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLevelByIDResponse.Merge(m, src)
}
func (m *GetLevelByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLevelByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLevelByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLevelByIDResponse proto.InternalMessageInfo

func (m *GetLevelByIDResponse) GetLevel() *Level {
	if m != nil {
		return m.Level
	}
	return nil
}

func init() {
	proto.RegisterType((*Level)(nil), "level.Level")
	proto.RegisterType((*GetLevelByIDRequest)(nil), "level.GetLevelByIDRequest")
	proto.RegisterType((*GetLevelByIDResponse)(nil), "level.GetLevelByIDResponse")
}

func init() { proto.RegisterFile("proto/level.proto", fileDescriptor_2a5dd38195ba9bd7) }

var fileDescriptor_2a5dd38195ba9bd7 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x49, 0x2d, 0x4b, 0xcd, 0xd1, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0xa5, 0x4c,
	0x2e, 0x56, 0x1f, 0x10, 0x43, 0x88, 0x8f, 0x8b, 0xc9, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x37, 0x88, 0xc9, 0xd3, 0x45, 0x48, 0x8b, 0x4b, 0xc0, 0xb5, 0xa2, 0x20, 0xb5, 0x28, 0x33, 0x35,
	0x2f, 0x39, 0xd5, 0x2f, 0x35, 0x35, 0x25, 0x35, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x08,
	0x43, 0x5c, 0x48, 0x85, 0x8b, 0x37, 0x20, 0x27, 0xb1, 0x32, 0xb5, 0xc8, 0x37, 0xb1, 0xc2, 0x27,
	0x33, 0x2d, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x08, 0x55, 0x50, 0x49, 0x9f, 0x4b, 0xd8,
	0x3d, 0xb5, 0x04, 0x6c, 0x9b, 0x53, 0xa5, 0xa7, 0x4b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x04, 0x17, 0x3b, 0x58, 0x0c, 0x6e, 0x3b, 0x8c, 0xab, 0x64, 0xc5, 0x25, 0x82, 0xaa, 0xa1,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x09, 0xea, 0x66, 0xb0, 0x7a, 0x6e, 0x23, 0x1e, 0x3d,
	0x88, 0xbf, 0xc0, 0x62, 0x41, 0x10, 0x29, 0x27, 0xb9, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0x62, 0xd1, 0xb3, 0x2e, 0x48, 0x4a, 0x62, 0x03, 0x87, 0x82, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xfe, 0x63, 0x93, 0x4b, 0x1a, 0x01, 0x00, 0x00,
}

func (m *Level) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Level) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Level) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlayerMaxLife != 0 {
		i = encodeVarintLevel(dAtA, i, uint64(m.PlayerMaxLife))
		i--
		dAtA[i] = 0x18
	}
	if m.ExperienceNeeded != 0 {
		i = encodeVarintLevel(dAtA, i, uint64(m.ExperienceNeeded))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintLevel(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetLevelByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLevelByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLevelByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LevelID != 0 {
		i = encodeVarintLevel(dAtA, i, uint64(m.LevelID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetLevelByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLevelByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLevelByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Level != nil {
		{
			size, err := m.Level.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLevel(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLevel(dAtA []byte, offset int, v uint64) int {
	offset -= sovLevel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Level) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLevel(uint64(m.ID))
	}
	if m.ExperienceNeeded != 0 {
		n += 1 + sovLevel(uint64(m.ExperienceNeeded))
	}
	if m.PlayerMaxLife != 0 {
		n += 1 + sovLevel(uint64(m.PlayerMaxLife))
	}
	return n
}

func (m *GetLevelByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LevelID != 0 {
		n += 1 + sovLevel(uint64(m.LevelID))
	}
	return n
}

func (m *GetLevelByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != nil {
		l = m.Level.Size()
		n += 1 + l + sovLevel(uint64(l))
	}
	return n
}

func sovLevel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLevel(x uint64) (n int) {
	return sovLevel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Level) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLevel
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
			return fmt.Errorf("proto: Level: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Level: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevel
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
				return fmt.Errorf("proto: wrong wireType = %d for field ExperienceNeeded", wireType)
			}
			m.ExperienceNeeded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExperienceNeeded |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerMaxLife", wireType)
			}
			m.PlayerMaxLife = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlayerMaxLife |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLevel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLevel
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
func (m *GetLevelByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLevel
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
			return fmt.Errorf("proto: GetLevelByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLevelByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LevelID", wireType)
			}
			m.LevelID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LevelID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLevel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLevel
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
func (m *GetLevelByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLevel
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
			return fmt.Errorf("proto: GetLevelByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLevelByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevel
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
				return ErrInvalidLengthLevel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLevel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Level == nil {
				m.Level = &Level{}
			}
			if err := m.Level.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLevel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLevel
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
func skipLevel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLevel
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
					return 0, ErrIntOverflowLevel
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
					return 0, ErrIntOverflowLevel
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
				return 0, ErrInvalidLengthLevel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLevel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLevel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLevel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLevel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLevel = fmt.Errorf("proto: unexpected end of group")
)
