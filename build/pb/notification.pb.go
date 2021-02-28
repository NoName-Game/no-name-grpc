// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/notification.proto

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

type Notification struct {
	ID                     uint32                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerID               uint32                `protobuf:"varint,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Player                 *Player               `protobuf:"bytes,3,opt,name=Player,proto3" json:"Player,omitempty"`
	NotificationCategoryID uint32                `protobuf:"varint,4,opt,name=NotificationCategoryID,proto3" json:"NotificationCategoryID,omitempty"`
	NotificationCategory   *NotificationCategory `protobuf:"bytes,5,opt,name=NotificationCategory,proto3" json:"NotificationCategory,omitempty"`
	Payload                string                `protobuf:"bytes,6,opt,name=Payload,proto3" json:"Payload,omitempty"`
	ToNotify               bool                  `protobuf:"varint,7,opt,name=ToNotify,proto3" json:"ToNotify,omitempty"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return m.Size()
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Notification) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Notification) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *Notification) GetNotificationCategoryID() uint32 {
	if m != nil {
		return m.NotificationCategoryID
	}
	return 0
}

func (m *Notification) GetNotificationCategory() *NotificationCategory {
	if m != nil {
		return m.NotificationCategory
	}
	return nil
}

func (m *Notification) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Notification) GetToNotify() bool {
	if m != nil {
		return m.ToNotify
	}
	return false
}

// GetNotifications
type GetNotificationsRequest struct {
}

func (m *GetNotificationsRequest) Reset()         { *m = GetNotificationsRequest{} }
func (m *GetNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotificationsRequest) ProtoMessage()    {}
func (*GetNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{1}
}
func (m *GetNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNotificationsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationsRequest.Merge(m, src)
}
func (m *GetNotificationsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationsRequest proto.InternalMessageInfo

type GetNotificationsResponse struct {
	Notifications []*Notification `protobuf:"bytes,1,rep,name=Notifications,proto3" json:"Notifications,omitempty"`
}

func (m *GetNotificationsResponse) Reset()         { *m = GetNotificationsResponse{} }
func (m *GetNotificationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetNotificationsResponse) ProtoMessage()    {}
func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{2}
}
func (m *GetNotificationsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNotificationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNotificationsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNotificationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationsResponse.Merge(m, src)
}
func (m *GetNotificationsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetNotificationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationsResponse proto.InternalMessageInfo

func (m *GetNotificationsResponse) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

// SetNotificationNotified
type SetNotificationNotifiedRequest struct {
	NotificationID uint32 `protobuf:"varint,1,opt,name=NotificationID,proto3" json:"NotificationID,omitempty"`
}

func (m *SetNotificationNotifiedRequest) Reset()         { *m = SetNotificationNotifiedRequest{} }
func (m *SetNotificationNotifiedRequest) String() string { return proto.CompactTextString(m) }
func (*SetNotificationNotifiedRequest) ProtoMessage()    {}
func (*SetNotificationNotifiedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{3}
}
func (m *SetNotificationNotifiedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetNotificationNotifiedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetNotificationNotifiedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetNotificationNotifiedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNotificationNotifiedRequest.Merge(m, src)
}
func (m *SetNotificationNotifiedRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetNotificationNotifiedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNotificationNotifiedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetNotificationNotifiedRequest proto.InternalMessageInfo

func (m *SetNotificationNotifiedRequest) GetNotificationID() uint32 {
	if m != nil {
		return m.NotificationID
	}
	return 0
}

type SetNotificationNotifiedResponse struct {
}

func (m *SetNotificationNotifiedResponse) Reset()         { *m = SetNotificationNotifiedResponse{} }
func (m *SetNotificationNotifiedResponse) String() string { return proto.CompactTextString(m) }
func (*SetNotificationNotifiedResponse) ProtoMessage()    {}
func (*SetNotificationNotifiedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{4}
}
func (m *SetNotificationNotifiedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetNotificationNotifiedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetNotificationNotifiedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetNotificationNotifiedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNotificationNotifiedResponse.Merge(m, src)
}
func (m *SetNotificationNotifiedResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetNotificationNotifiedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNotificationNotifiedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetNotificationNotifiedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Notification)(nil), "notification.Notification")
	proto.RegisterType((*GetNotificationsRequest)(nil), "notification.GetNotificationsRequest")
	proto.RegisterType((*GetNotificationsResponse)(nil), "notification.GetNotificationsResponse")
	proto.RegisterType((*SetNotificationNotifiedRequest)(nil), "notification.SetNotificationNotifiedRequest")
	proto.RegisterType((*SetNotificationNotifiedResponse)(nil), "notification.SetNotificationNotifiedResponse")
}

func init() { proto.RegisterFile("proto/notification.proto", fileDescriptor_beb4fc010f5c4b1c) }

var fileDescriptor_beb4fc010f5c4b1c = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03,
	0x0b, 0x09, 0xf1, 0x20, 0x8b, 0x49, 0x29, 0x62, 0xaa, 0x8b, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xcf,
	0x2f, 0xaa, 0x84, 0x68, 0x90, 0x12, 0x82, 0x28, 0x29, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x82, 0x88,
	0x29, 0x6d, 0x64, 0xe2, 0xe2, 0xf1, 0x43, 0xd2, 0x23, 0xc4, 0xc7, 0xc5, 0xe4, 0xe9, 0x22, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4, 0xe4, 0xe9, 0x22, 0x24, 0xc5, 0xc5, 0x11, 0x00, 0xd6, 0xe0,
	0xe9, 0x22, 0xc1, 0x04, 0x16, 0x85, 0xf3, 0x85, 0xd4, 0xb8, 0xd8, 0x20, 0x6c, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xa8, 0xd9, 0x10, 0xd1, 0x20, 0xa8, 0xac, 0x90, 0x19, 0x97,
	0x18, 0xb2, 0x1d, 0xce, 0x50, 0x67, 0x79, 0xba, 0x48, 0xb0, 0x80, 0x4d, 0xc4, 0x21, 0x2b, 0x14,
	0xcf, 0x25, 0x82, 0x4d, 0x46, 0x82, 0x15, 0x6c, 0x9b, 0xb6, 0x1e, 0x76, 0xcf, 0x62, 0xd3, 0x12,
	0x84, 0xd5, 0x20, 0x21, 0x09, 0x2e, 0xf6, 0x80, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x36,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x17, 0xe4, 0xed, 0x90, 0x7c, 0xb0, 0x9e, 0x4a, 0x09, 0x76,
	0x05, 0x46, 0x0d, 0x8e, 0x20, 0x38, 0x5f, 0x49, 0x92, 0x4b, 0xdc, 0x3d, 0xb5, 0x04, 0xd9, 0xc0,
	0xe2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa5, 0x18, 0x2e, 0x09, 0x4c, 0xa9, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x07, 0x2e, 0x5e, 0x14, 0x09, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e,
	0x23, 0x29, 0x14, 0x6f, 0xa0, 0xb8, 0x3e, 0x08, 0x55, 0x83, 0x92, 0x07, 0x97, 0x5c, 0x30, 0xaa,
	0xe9, 0x10, 0x76, 0x6a, 0x0a, 0xd4, 0x7e, 0x21, 0x35, 0x2e, 0x3e, 0x64, 0x69, 0x78, 0x4c, 0xa2,
	0x89, 0x2a, 0x29, 0x72, 0xc9, 0xe3, 0x34, 0x09, 0xe2, 0x5c, 0x27, 0xb9, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x62, 0xd1, 0xb3, 0x2e, 0x48, 0x4a, 0x62, 0x03, 0x27, 0x20,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x99, 0xad, 0x16, 0xa1, 0x02, 0x00, 0x00,
}

func (m *Notification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Notification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Notification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x38
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintNotification(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x32
	}
	if m.NotificationCategory != nil {
		{
			size, err := m.NotificationCategory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNotification(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.NotificationCategoryID != 0 {
		i = encodeVarintNotification(dAtA, i, uint64(m.NotificationCategoryID))
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
			i = encodeVarintNotification(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.PlayerID != 0 {
		i = encodeVarintNotification(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintNotification(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetNotificationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNotificationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNotificationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetNotificationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNotificationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNotificationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Notifications) > 0 {
		for iNdEx := len(m.Notifications) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Notifications[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNotification(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SetNotificationNotifiedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetNotificationNotifiedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetNotificationNotifiedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NotificationID != 0 {
		i = encodeVarintNotification(dAtA, i, uint64(m.NotificationID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SetNotificationNotifiedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetNotificationNotifiedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetNotificationNotifiedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintNotification(dAtA []byte, offset int, v uint64) int {
	offset -= sovNotification(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Notification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovNotification(uint64(m.ID))
	}
	if m.PlayerID != 0 {
		n += 1 + sovNotification(uint64(m.PlayerID))
	}
	if m.Player != nil {
		l = m.Player.Size()
		n += 1 + l + sovNotification(uint64(l))
	}
	if m.NotificationCategoryID != 0 {
		n += 1 + sovNotification(uint64(m.NotificationCategoryID))
	}
	if m.NotificationCategory != nil {
		l = m.NotificationCategory.Size()
		n += 1 + l + sovNotification(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovNotification(uint64(l))
	}
	if m.ToNotify {
		n += 2
	}
	return n
}

func (m *GetNotificationsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetNotificationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Notifications) > 0 {
		for _, e := range m.Notifications {
			l = e.Size()
			n += 1 + l + sovNotification(uint64(l))
		}
	}
	return n
}

func (m *SetNotificationNotifiedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotificationID != 0 {
		n += 1 + sovNotification(uint64(m.NotificationID))
	}
	return n
}

func (m *SetNotificationNotifiedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovNotification(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNotification(x uint64) (n int) {
	return sovNotification(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Notification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: Notification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Notification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
					return ErrIntOverflowNotification
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
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
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
				return fmt.Errorf("proto: wrong wireType = %d for field NotificationCategoryID", wireType)
			}
			m.NotificationCategoryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotificationCategoryID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotificationCategory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NotificationCategory == nil {
				m.NotificationCategory = &NotificationCategory{}
			}
			if err := m.NotificationCategory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToNotify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotification
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotification
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
func (m *GetNotificationsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: GetNotificationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNotificationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotification
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotification
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
func (m *GetNotificationsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: GetNotificationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNotificationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notifications", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
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
				return ErrInvalidLengthNotification
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notifications = append(m.Notifications, &Notification{})
			if err := m.Notifications[len(m.Notifications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotification
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotification
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
func (m *SetNotificationNotifiedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: SetNotificationNotifiedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetNotificationNotifiedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotificationID", wireType)
			}
			m.NotificationID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotification
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotificationID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotification
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotification
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
func (m *SetNotificationNotifiedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotification
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
			return fmt.Errorf("proto: SetNotificationNotifiedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetNotificationNotifiedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNotification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotification
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotification
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
func skipNotification(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotification
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
					return 0, ErrIntOverflowNotification
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
					return 0, ErrIntOverflowNotification
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
				return 0, ErrInvalidLengthNotification
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNotification
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNotification
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNotification        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotification          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNotification = fmt.Errorf("proto: unexpected end of group")
)
