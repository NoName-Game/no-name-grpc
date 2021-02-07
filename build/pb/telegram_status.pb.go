// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/telegram_status.proto

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

// CreateTelegramStatus
type CreateTelegramStatusRequest struct {
	PlayerID   uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Stage      int32  `protobuf:"varint,2,opt,name=Stage,proto3" json:"Stage,omitempty"`
	Controller string `protobuf:"bytes,3,opt,name=Controller,proto3" json:"Controller,omitempty"`
	Payload    string `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (m *CreateTelegramStatusRequest) Reset()         { *m = CreateTelegramStatusRequest{} }
func (m *CreateTelegramStatusRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTelegramStatusRequest) ProtoMessage()    {}
func (*CreateTelegramStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5615c81e0c4e524d, []int{0}
}
func (m *CreateTelegramStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateTelegramStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateTelegramStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateTelegramStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTelegramStatusRequest.Merge(m, src)
}
func (m *CreateTelegramStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreateTelegramStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTelegramStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTelegramStatusRequest proto.InternalMessageInfo

func (m *CreateTelegramStatusRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *CreateTelegramStatusRequest) GetStage() int32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *CreateTelegramStatusRequest) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *CreateTelegramStatusRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type CreateTelegramStatusResponse struct {
}

func (m *CreateTelegramStatusResponse) Reset()         { *m = CreateTelegramStatusResponse{} }
func (m *CreateTelegramStatusResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTelegramStatusResponse) ProtoMessage()    {}
func (*CreateTelegramStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5615c81e0c4e524d, []int{1}
}
func (m *CreateTelegramStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateTelegramStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateTelegramStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateTelegramStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTelegramStatusResponse.Merge(m, src)
}
func (m *CreateTelegramStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreateTelegramStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTelegramStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTelegramStatusResponse proto.InternalMessageInfo

// GetTelegramStatus
type GetTelegramStatusRequest struct {
	PlayerID   uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Controller string `protobuf:"bytes,3,opt,name=Controller,proto3" json:"Controller,omitempty"`
}

func (m *GetTelegramStatusRequest) Reset()         { *m = GetTelegramStatusRequest{} }
func (m *GetTelegramStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetTelegramStatusRequest) ProtoMessage()    {}
func (*GetTelegramStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5615c81e0c4e524d, []int{2}
}
func (m *GetTelegramStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTelegramStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTelegramStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTelegramStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTelegramStatusRequest.Merge(m, src)
}
func (m *GetTelegramStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTelegramStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTelegramStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTelegramStatusRequest proto.InternalMessageInfo

func (m *GetTelegramStatusRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *GetTelegramStatusRequest) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

type GetTelegramStatusResponse struct {
	PlayerID   uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Stage      int32  `protobuf:"varint,2,opt,name=Stage,proto3" json:"Stage,omitempty"`
	Controller string `protobuf:"bytes,3,opt,name=Controller,proto3" json:"Controller,omitempty"`
	Payload    string `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (m *GetTelegramStatusResponse) Reset()         { *m = GetTelegramStatusResponse{} }
func (m *GetTelegramStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetTelegramStatusResponse) ProtoMessage()    {}
func (*GetTelegramStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5615c81e0c4e524d, []int{3}
}
func (m *GetTelegramStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTelegramStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTelegramStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTelegramStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTelegramStatusResponse.Merge(m, src)
}
func (m *GetTelegramStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTelegramStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTelegramStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTelegramStatusResponse proto.InternalMessageInfo

func (m *GetTelegramStatusResponse) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *GetTelegramStatusResponse) GetStage() int32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *GetTelegramStatusResponse) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *GetTelegramStatusResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

// DeleteTelegramStatus
type DeleteTelegramStatusRequest struct {
	PlayerID   uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Controller string `protobuf:"bytes,3,opt,name=Controller,proto3" json:"Controller,omitempty"`
}

func (m *DeleteTelegramStatusRequest) Reset()         { *m = DeleteTelegramStatusRequest{} }
func (m *DeleteTelegramStatusRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTelegramStatusRequest) ProtoMessage()    {}
func (*DeleteTelegramStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5615c81e0c4e524d, []int{4}
}
func (m *DeleteTelegramStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteTelegramStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteTelegramStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeleteTelegramStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTelegramStatusRequest.Merge(m, src)
}
func (m *DeleteTelegramStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeleteTelegramStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTelegramStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTelegramStatusRequest proto.InternalMessageInfo

func (m *DeleteTelegramStatusRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *DeleteTelegramStatusRequest) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

type DeleteTelegramStatusResponse struct {
}

func (m *DeleteTelegramStatusResponse) Reset()         { *m = DeleteTelegramStatusResponse{} }
func (m *DeleteTelegramStatusResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTelegramStatusResponse) ProtoMessage()    {}
func (*DeleteTelegramStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5615c81e0c4e524d, []int{5}
}
func (m *DeleteTelegramStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteTelegramStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteTelegramStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeleteTelegramStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTelegramStatusResponse.Merge(m, src)
}
func (m *DeleteTelegramStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeleteTelegramStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTelegramStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTelegramStatusResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTelegramStatusRequest)(nil), "telegram_status.CreateTelegramStatusRequest")
	proto.RegisterType((*CreateTelegramStatusResponse)(nil), "telegram_status.CreateTelegramStatusResponse")
	proto.RegisterType((*GetTelegramStatusRequest)(nil), "telegram_status.GetTelegramStatusRequest")
	proto.RegisterType((*GetTelegramStatusResponse)(nil), "telegram_status.GetTelegramStatusResponse")
	proto.RegisterType((*DeleteTelegramStatusRequest)(nil), "telegram_status.DeleteTelegramStatusRequest")
	proto.RegisterType((*DeleteTelegramStatusResponse)(nil), "telegram_status.DeleteTelegramStatusResponse")
}

func init() { proto.RegisterFile("proto/telegram_status.proto", fileDescriptor_5615c81e0c4e524d) }

var fileDescriptor_5615c81e0c4e524d = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0xcd, 0x49, 0x4d, 0x2f, 0x4a, 0xcc, 0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x29,
	0x2d, 0xd6, 0x03, 0x8b, 0x0a, 0xf1, 0xa3, 0x09, 0x2b, 0x75, 0x32, 0x72, 0x49, 0x3b, 0x17, 0xa5,
	0x26, 0x96, 0xa4, 0x86, 0x40, 0x65, 0x82, 0xc1, 0x12, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x52, 0x5c, 0x1c, 0x01, 0x39, 0x89, 0x95, 0xa9, 0x45, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xbc, 0x41, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x6b, 0x70, 0x49, 0x62, 0x7a, 0xaa, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0x23, 0x24, 0xc7, 0xc5, 0xe5, 0x9c, 0x9f, 0x57, 0x52, 0x94,
	0x9f, 0x93, 0x93, 0x5a, 0x24, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x24, 0x22, 0x24, 0xc1,
	0xc5, 0x1e, 0x90, 0x58, 0x99, 0x93, 0x9f, 0x98, 0x22, 0xc1, 0x02, 0x96, 0x84, 0x71, 0x95, 0xe4,
	0xb8, 0x64, 0xb0, 0x3b, 0xa5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x29, 0x8c, 0x4b, 0xc2, 0x3d,
	0xb5, 0x84, 0x74, 0x77, 0x12, 0x70, 0x91, 0x52, 0x3b, 0x23, 0x97, 0x24, 0x16, 0x83, 0x21, 0xb6,
	0xd2, 0x35, 0x04, 0x22, 0xb9, 0xa4, 0x5d, 0x52, 0x73, 0x52, 0xc9, 0x89, 0x0c, 0x42, 0x9e, 0x94,
	0xe3, 0x92, 0xc1, 0x6e, 0x34, 0xc4, 0x9b, 0x4e, 0x72, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xc5, 0xa2, 0x67, 0x5d, 0x90, 0x94, 0xc4, 0x06, 0x4e, 0x40, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x23, 0x06, 0x74, 0x5f, 0x02, 0x00, 0x00,
}

func (m *CreateTelegramStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateTelegramStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateTelegramStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTelegramStatus(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTelegramStatus(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Stage != 0 {
		i = encodeVarintTelegramStatus(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x10
	}
	if m.PlayerID != 0 {
		i = encodeVarintTelegramStatus(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateTelegramStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateTelegramStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateTelegramStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetTelegramStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTelegramStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTelegramStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTelegramStatus(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PlayerID != 0 {
		i = encodeVarintTelegramStatus(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetTelegramStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTelegramStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTelegramStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTelegramStatus(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTelegramStatus(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Stage != 0 {
		i = encodeVarintTelegramStatus(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x10
	}
	if m.PlayerID != 0 {
		i = encodeVarintTelegramStatus(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DeleteTelegramStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteTelegramStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeleteTelegramStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTelegramStatus(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PlayerID != 0 {
		i = encodeVarintTelegramStatus(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DeleteTelegramStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteTelegramStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeleteTelegramStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTelegramStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovTelegramStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateTelegramStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlayerID != 0 {
		n += 1 + sovTelegramStatus(uint64(m.PlayerID))
	}
	if m.Stage != 0 {
		n += 1 + sovTelegramStatus(uint64(m.Stage))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTelegramStatus(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTelegramStatus(uint64(l))
	}
	return n
}

func (m *CreateTelegramStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetTelegramStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlayerID != 0 {
		n += 1 + sovTelegramStatus(uint64(m.PlayerID))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTelegramStatus(uint64(l))
	}
	return n
}

func (m *GetTelegramStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlayerID != 0 {
		n += 1 + sovTelegramStatus(uint64(m.PlayerID))
	}
	if m.Stage != 0 {
		n += 1 + sovTelegramStatus(uint64(m.Stage))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTelegramStatus(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTelegramStatus(uint64(l))
	}
	return n
}

func (m *DeleteTelegramStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlayerID != 0 {
		n += 1 + sovTelegramStatus(uint64(m.PlayerID))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTelegramStatus(uint64(l))
	}
	return n
}

func (m *DeleteTelegramStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTelegramStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTelegramStatus(x uint64) (n int) {
	return sovTelegramStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateTelegramStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelegramStatus
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
			return fmt.Errorf("proto: CreateTelegramStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateTelegramStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return ErrInvalidLengthTelegramStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return ErrInvalidLengthTelegramStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTelegramStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTelegramStatus
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
func (m *CreateTelegramStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelegramStatus
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
			return fmt.Errorf("proto: CreateTelegramStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateTelegramStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTelegramStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTelegramStatus
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
func (m *GetTelegramStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelegramStatus
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
			return fmt.Errorf("proto: GetTelegramStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTelegramStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return ErrInvalidLengthTelegramStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTelegramStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTelegramStatus
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
func (m *GetTelegramStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelegramStatus
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
			return fmt.Errorf("proto: GetTelegramStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTelegramStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return ErrInvalidLengthTelegramStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return ErrInvalidLengthTelegramStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTelegramStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTelegramStatus
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
func (m *DeleteTelegramStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelegramStatus
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
			return fmt.Errorf("proto: DeleteTelegramStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteTelegramStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelegramStatus
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
				return ErrInvalidLengthTelegramStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTelegramStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTelegramStatus
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
func (m *DeleteTelegramStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelegramStatus
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
			return fmt.Errorf("proto: DeleteTelegramStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteTelegramStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTelegramStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTelegramStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTelegramStatus
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
func skipTelegramStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTelegramStatus
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
					return 0, ErrIntOverflowTelegramStatus
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
					return 0, ErrIntOverflowTelegramStatus
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
				return 0, ErrInvalidLengthTelegramStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTelegramStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTelegramStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTelegramStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTelegramStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTelegramStatus = fmt.Errorf("proto: unexpected end of group")
)
