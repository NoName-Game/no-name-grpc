// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/global_message.proto

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

type Message struct {
	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Sent bool   `protobuf:"varint,3,opt,name=Sent,proto3" json:"Sent,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7889b8afc726cf83, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetSent() bool {
	if m != nil {
		return m.Sent
	}
	return false
}

type RetrieveMessageRequest struct {
}

func (m *RetrieveMessageRequest) Reset()         { *m = RetrieveMessageRequest{} }
func (m *RetrieveMessageRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveMessageRequest) ProtoMessage()    {}
func (*RetrieveMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7889b8afc726cf83, []int{1}
}
func (m *RetrieveMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetrieveMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetrieveMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetrieveMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveMessageRequest.Merge(m, src)
}
func (m *RetrieveMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *RetrieveMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveMessageRequest proto.InternalMessageInfo

type RetrieveMessageResponse struct {
	Messages []*Message `protobuf:"bytes,1,rep,name=Messages,proto3" json:"Messages,omitempty"`
}

func (m *RetrieveMessageResponse) Reset()         { *m = RetrieveMessageResponse{} }
func (m *RetrieveMessageResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveMessageResponse) ProtoMessage()    {}
func (*RetrieveMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7889b8afc726cf83, []int{2}
}
func (m *RetrieveMessageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetrieveMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetrieveMessageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetrieveMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveMessageResponse.Merge(m, src)
}
func (m *RetrieveMessageResponse) XXX_Size() int {
	return m.Size()
}
func (m *RetrieveMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveMessageResponse proto.InternalMessageInfo

func (m *RetrieveMessageResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "global_message.Message")
	proto.RegisterType((*RetrieveMessageRequest)(nil), "global_message.RetrieveMessageRequest")
	proto.RegisterType((*RetrieveMessageResponse)(nil), "global_message.RetrieveMessageResponse")
}

func init() { proto.RegisterFile("proto/global_message.proto", fileDescriptor_7889b8afc726cf83) }

var fileDescriptor_7889b8afc726cf83 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0xd5, 0x03, 0x0b, 0x0a, 0xf1, 0xa1, 0x8a, 0x2a, 0x39, 0x72, 0xb1, 0xfb, 0x42, 0x98, 0x42, 0x7c,
	0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x42,
	0x5c, 0x2c, 0x21, 0xa9, 0x15, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48,
	0x2c, 0x38, 0x35, 0xaf, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08, 0xcc, 0x56, 0x92, 0xe0,
	0x12, 0x0b, 0x4a, 0x2d, 0x29, 0xca, 0x4c, 0x2d, 0x4b, 0x85, 0x1a, 0x15, 0x94, 0x5a, 0x58, 0x9a,
	0x5a, 0x5c, 0xa2, 0xe4, 0xc7, 0x25, 0x8e, 0x21, 0x53, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64,
	0xcc, 0xc5, 0x01, 0x15, 0x2a, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd7, 0x43, 0x73,
	0x30, 0x4c, 0x0b, 0x5c, 0xa1, 0x93, 0xdc, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x44, 0xb1, 0xe8, 0x59, 0x17, 0x24, 0x25, 0xb1, 0x81, 0xfd, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xd1, 0xa0, 0xb0, 0x01, 0x01, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sent {
		i--
		if m.Sent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintGlobalMessage(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintGlobalMessage(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RetrieveMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetrieveMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetrieveMessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RetrieveMessageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetrieveMessageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetrieveMessageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGlobalMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGlobalMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovGlobalMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovGlobalMessage(uint64(m.ID))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovGlobalMessage(uint64(l))
	}
	if m.Sent {
		n += 2
	}
	return n
}

func (m *RetrieveMessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RetrieveMessageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovGlobalMessage(uint64(l))
		}
	}
	return n
}

func sovGlobalMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGlobalMessage(x uint64) (n int) {
	return sovGlobalMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalMessage
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
				return ErrInvalidLengthGlobalMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalMessage
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
			m.Sent = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalMessage
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
func (m *RetrieveMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalMessage
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
			return fmt.Errorf("proto: RetrieveMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetrieveMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalMessage
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
func (m *RetrieveMessageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalMessage
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
			return fmt.Errorf("proto: RetrieveMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetrieveMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalMessage
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
				return ErrInvalidLengthGlobalMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &Message{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalMessage
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
func skipGlobalMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGlobalMessage
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
					return 0, ErrIntOverflowGlobalMessage
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
					return 0, ErrIntOverflowGlobalMessage
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
				return 0, ErrInvalidLengthGlobalMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGlobalMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGlobalMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGlobalMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGlobalMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGlobalMessage = fmt.Errorf("proto: unexpected end of group")
)
