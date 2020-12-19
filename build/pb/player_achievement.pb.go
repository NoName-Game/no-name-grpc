// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/player_achievement.proto

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

type PlayerAchievement struct {
	ID            uint32       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerID      uint32       `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Player        *Player      `protobuf:"bytes,3,opt,name=Player,proto3" json:"Player,omitempty"`
	AchievementID uint32       `protobuf:"varint,4,opt,name=AchievementID,proto3" json:"AchievementID,omitempty"`
	Achievement   *Achievement `protobuf:"bytes,5,opt,name=Achievement,proto3" json:"Achievement,omitempty"`
	ToNotify      bool         `protobuf:"varint,6,opt,name=ToNotify,proto3" json:"ToNotify,omitempty"`
}

func (m *PlayerAchievement) Reset()         { *m = PlayerAchievement{} }
func (m *PlayerAchievement) String() string { return proto.CompactTextString(m) }
func (*PlayerAchievement) ProtoMessage()    {}
func (*PlayerAchievement) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12157cff19983a, []int{0}
}
func (m *PlayerAchievement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerAchievement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerAchievement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlayerAchievement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAchievement.Merge(m, src)
}
func (m *PlayerAchievement) XXX_Size() int {
	return m.Size()
}
func (m *PlayerAchievement) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAchievement.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAchievement proto.InternalMessageInfo

func (m *PlayerAchievement) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PlayerAchievement) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *PlayerAchievement) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *PlayerAchievement) GetAchievementID() uint32 {
	if m != nil {
		return m.AchievementID
	}
	return 0
}

func (m *PlayerAchievement) GetAchievement() *Achievement {
	if m != nil {
		return m.Achievement
	}
	return nil
}

func (m *PlayerAchievement) GetToNotify() bool {
	if m != nil {
		return m.ToNotify
	}
	return false
}

// GetPlayerAchievementToNotify
type GetPlayerAchievementToNotifyRequest struct {
}

func (m *GetPlayerAchievementToNotifyRequest) Reset()         { *m = GetPlayerAchievementToNotifyRequest{} }
func (m *GetPlayerAchievementToNotifyRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlayerAchievementToNotifyRequest) ProtoMessage()    {}
func (*GetPlayerAchievementToNotifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12157cff19983a, []int{1}
}
func (m *GetPlayerAchievementToNotifyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlayerAchievementToNotifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlayerAchievementToNotifyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlayerAchievementToNotifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerAchievementToNotifyRequest.Merge(m, src)
}
func (m *GetPlayerAchievementToNotifyRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPlayerAchievementToNotifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerAchievementToNotifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerAchievementToNotifyRequest proto.InternalMessageInfo

type GetPlayerAchievementToNotifyResponse struct {
	PlayerAchievements []*PlayerAchievement `protobuf:"bytes,1,rep,name=PlayerAchievements,proto3" json:"PlayerAchievements,omitempty"`
}

func (m *GetPlayerAchievementToNotifyResponse) Reset()         { *m = GetPlayerAchievementToNotifyResponse{} }
func (m *GetPlayerAchievementToNotifyResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlayerAchievementToNotifyResponse) ProtoMessage()    {}
func (*GetPlayerAchievementToNotifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12157cff19983a, []int{2}
}
func (m *GetPlayerAchievementToNotifyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlayerAchievementToNotifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlayerAchievementToNotifyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlayerAchievementToNotifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerAchievementToNotifyResponse.Merge(m, src)
}
func (m *GetPlayerAchievementToNotifyResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPlayerAchievementToNotifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerAchievementToNotifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerAchievementToNotifyResponse proto.InternalMessageInfo

func (m *GetPlayerAchievementToNotifyResponse) GetPlayerAchievements() []*PlayerAchievement {
	if m != nil {
		return m.PlayerAchievements
	}
	return nil
}

// SetPlayerAchievementNotified
type SetPlayerAchievementNotifiedRequest struct {
	AchievementID uint32 `protobuf:"varint,1,opt,name=AchievementID,proto3" json:"AchievementID,omitempty"`
}

func (m *SetPlayerAchievementNotifiedRequest) Reset()         { *m = SetPlayerAchievementNotifiedRequest{} }
func (m *SetPlayerAchievementNotifiedRequest) String() string { return proto.CompactTextString(m) }
func (*SetPlayerAchievementNotifiedRequest) ProtoMessage()    {}
func (*SetPlayerAchievementNotifiedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12157cff19983a, []int{3}
}
func (m *SetPlayerAchievementNotifiedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetPlayerAchievementNotifiedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetPlayerAchievementNotifiedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetPlayerAchievementNotifiedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPlayerAchievementNotifiedRequest.Merge(m, src)
}
func (m *SetPlayerAchievementNotifiedRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetPlayerAchievementNotifiedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPlayerAchievementNotifiedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetPlayerAchievementNotifiedRequest proto.InternalMessageInfo

func (m *SetPlayerAchievementNotifiedRequest) GetAchievementID() uint32 {
	if m != nil {
		return m.AchievementID
	}
	return 0
}

type SetPlayerAchievementNotifiedResponse struct {
}

func (m *SetPlayerAchievementNotifiedResponse) Reset()         { *m = SetPlayerAchievementNotifiedResponse{} }
func (m *SetPlayerAchievementNotifiedResponse) String() string { return proto.CompactTextString(m) }
func (*SetPlayerAchievementNotifiedResponse) ProtoMessage()    {}
func (*SetPlayerAchievementNotifiedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12157cff19983a, []int{4}
}
func (m *SetPlayerAchievementNotifiedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetPlayerAchievementNotifiedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetPlayerAchievementNotifiedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetPlayerAchievementNotifiedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPlayerAchievementNotifiedResponse.Merge(m, src)
}
func (m *SetPlayerAchievementNotifiedResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetPlayerAchievementNotifiedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPlayerAchievementNotifiedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetPlayerAchievementNotifiedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PlayerAchievement)(nil), "player_achievement.PlayerAchievement")
	proto.RegisterType((*GetPlayerAchievementToNotifyRequest)(nil), "player_achievement.GetPlayerAchievementToNotifyRequest")
	proto.RegisterType((*GetPlayerAchievementToNotifyResponse)(nil), "player_achievement.GetPlayerAchievementToNotifyResponse")
	proto.RegisterType((*SetPlayerAchievementNotifiedRequest)(nil), "player_achievement.SetPlayerAchievementNotifiedRequest")
	proto.RegisterType((*SetPlayerAchievementNotifiedResponse)(nil), "player_achievement.SetPlayerAchievementNotifiedResponse")
}

func init() { proto.RegisterFile("proto/player_achievement.proto", fileDescriptor_2e12157cff19983a) }

var fileDescriptor_2e12157cff19983a = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x8a, 0x4f, 0x4c, 0xce, 0xc8, 0x4c, 0x2d, 0x4b,
	0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x03, 0x4b, 0x08, 0x09, 0x61, 0xca, 0x48, 0x09, 0x21, 0xeb, 0x81,
	0xa8, 0x93, 0x12, 0x87, 0x88, 0x61, 0x18, 0xa0, 0xf4, 0x94, 0x91, 0x4b, 0x30, 0x00, 0xac, 0xd2,
	0x11, 0x21, 0x27, 0xc4, 0xc7, 0xc5, 0xe4, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4,
	0xe4, 0xe9, 0x22, 0x24, 0xc5, 0xc5, 0x01, 0x51, 0xe4, 0xe9, 0x22, 0xc1, 0x04, 0x16, 0x85, 0xf3,
	0x85, 0xd4, 0xb8, 0xd8, 0x20, 0x6c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xa8,
	0xcd, 0x10, 0xd1, 0x20, 0xa8, 0xac, 0x90, 0x0a, 0x17, 0x2f, 0x92, 0x15, 0x9e, 0x2e, 0x12, 0x2c,
	0x60, 0x83, 0x50, 0x05, 0x85, 0xac, 0xb8, 0xb8, 0x91, 0x04, 0x24, 0x58, 0xc1, 0x46, 0x4a, 0xe8,
	0x21, 0x3b, 0x1c, 0x49, 0x3e, 0x08, 0x59, 0x31, 0xc8, 0x95, 0x21, 0xf9, 0x7e, 0xf9, 0x25, 0x99,
	0x69, 0x95, 0x12, 0x6c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x70, 0xbe, 0x92, 0x2a, 0x97, 0xb2, 0x7b,
	0x6a, 0x09, 0x86, 0x4f, 0x61, 0xf2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xb5, 0x5c,
	0x2a, 0xf8, 0x95, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x85, 0x72, 0x09, 0x61, 0x28, 0x2a,
	0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd5, 0xc3, 0x12, 0x5d, 0x18, 0xaa, 0x83, 0xb0,
	0x18, 0xa0, 0xe4, 0xcd, 0xa5, 0x1c, 0x8c, 0xc5, 0x7a, 0xb0, 0xe5, 0x99, 0xa9, 0x29, 0x50, 0x57,
	0x62, 0x06, 0x25, 0x23, 0x96, 0xa0, 0x54, 0x52, 0xe3, 0x52, 0xc1, 0x6f, 0x18, 0xc4, 0x2f, 0x4e,
	0x72, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0xa2, 0x67, 0x5d, 0x90,
	0x94, 0xc4, 0x06, 0x4e, 0x29, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xd1, 0xca, 0xca,
	0x8c, 0x02, 0x00, 0x00,
}

func (m *PlayerAchievement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerAchievement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlayerAchievement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ToNotify {
		i--
		if m.ToNotify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Achievement != nil {
		{
			size, err := m.Achievement.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlayerAchievement(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.AchievementID != 0 {
		i = encodeVarintPlayerAchievement(dAtA, i, uint64(m.AchievementID))
		i--
		dAtA[i] = 0x20
	}
	if m.Player != nil {
		{
			size, err := m.Player.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlayerAchievement(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.PlayerID != 0 {
		i = encodeVarintPlayerAchievement(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintPlayerAchievement(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlayerAchievementToNotifyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlayerAchievementToNotifyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlayerAchievementToNotifyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetPlayerAchievementToNotifyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlayerAchievementToNotifyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlayerAchievementToNotifyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlayerAchievements) > 0 {
		for iNdEx := len(m.PlayerAchievements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PlayerAchievements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlayerAchievement(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SetPlayerAchievementNotifiedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetPlayerAchievementNotifiedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetPlayerAchievementNotifiedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AchievementID != 0 {
		i = encodeVarintPlayerAchievement(dAtA, i, uint64(m.AchievementID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SetPlayerAchievementNotifiedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetPlayerAchievementNotifiedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetPlayerAchievementNotifiedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPlayerAchievement(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlayerAchievement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlayerAchievement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPlayerAchievement(uint64(m.ID))
	}
	if m.PlayerID != 0 {
		n += 1 + sovPlayerAchievement(uint64(m.PlayerID))
	}
	if m.Player != nil {
		l = m.Player.Size()
		n += 1 + l + sovPlayerAchievement(uint64(l))
	}
	if m.AchievementID != 0 {
		n += 1 + sovPlayerAchievement(uint64(m.AchievementID))
	}
	if m.Achievement != nil {
		l = m.Achievement.Size()
		n += 1 + l + sovPlayerAchievement(uint64(l))
	}
	if m.ToNotify {
		n += 2
	}
	return n
}

func (m *GetPlayerAchievementToNotifyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetPlayerAchievementToNotifyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PlayerAchievements) > 0 {
		for _, e := range m.PlayerAchievements {
			l = e.Size()
			n += 1 + l + sovPlayerAchievement(uint64(l))
		}
	}
	return n
}

func (m *SetPlayerAchievementNotifiedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AchievementID != 0 {
		n += 1 + sovPlayerAchievement(uint64(m.AchievementID))
	}
	return n
}

func (m *SetPlayerAchievementNotifiedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPlayerAchievement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlayerAchievement(x uint64) (n int) {
	return sovPlayerAchievement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerAchievement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerAchievement
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
			return fmt.Errorf("proto: PlayerAchievement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerAchievement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
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
					return ErrIntOverflowPlayerAchievement
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
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
				return ErrInvalidLengthPlayerAchievement
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Player == nil {
				m.Player = &Player{}
			}
			if err := m.Player.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AchievementID", wireType)
			}
			m.AchievementID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AchievementID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Achievement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
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
				return ErrInvalidLengthPlayerAchievement
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Achievement == nil {
				m.Achievement = &Achievement{}
			}
			if err := m.Achievement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToNotify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
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
			m.ToNotify = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerAchievement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerAchievement
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
func (m *GetPlayerAchievementToNotifyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerAchievement
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
			return fmt.Errorf("proto: GetPlayerAchievementToNotifyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlayerAchievementToNotifyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerAchievement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerAchievement
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
func (m *GetPlayerAchievementToNotifyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerAchievement
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
			return fmt.Errorf("proto: GetPlayerAchievementToNotifyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlayerAchievementToNotifyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerAchievements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
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
				return ErrInvalidLengthPlayerAchievement
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerAchievements = append(m.PlayerAchievements, &PlayerAchievement{})
			if err := m.PlayerAchievements[len(m.PlayerAchievements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerAchievement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerAchievement
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
func (m *SetPlayerAchievementNotifiedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerAchievement
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
			return fmt.Errorf("proto: SetPlayerAchievementNotifiedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetPlayerAchievementNotifiedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AchievementID", wireType)
			}
			m.AchievementID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerAchievement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AchievementID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerAchievement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerAchievement
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
func (m *SetPlayerAchievementNotifiedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerAchievement
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
			return fmt.Errorf("proto: SetPlayerAchievementNotifiedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetPlayerAchievementNotifiedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerAchievement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayerAchievement
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlayerAchievement
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
func skipPlayerAchievement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayerAchievement
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
					return 0, ErrIntOverflowPlayerAchievement
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
					return 0, ErrIntOverflowPlayerAchievement
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
				return 0, ErrInvalidLengthPlayerAchievement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlayerAchievement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlayerAchievement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlayerAchievement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayerAchievement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlayerAchievement = fmt.Errorf("proto: unexpected end of group")
)