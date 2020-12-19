// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/expansion.proto

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

// GetTeletrasportSafePlanetList
type GetTeletrasportSafePlanetListRequest struct {
	PlanetID uint32 `protobuf:"varint,1,opt,name=PlanetID,proto3" json:"PlanetID,omitempty"`
}

func (m *GetTeletrasportSafePlanetListRequest) Reset()         { *m = GetTeletrasportSafePlanetListRequest{} }
func (m *GetTeletrasportSafePlanetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetTeletrasportSafePlanetListRequest) ProtoMessage()    {}
func (*GetTeletrasportSafePlanetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab067e92a46b1c0b, []int{0}
}
func (m *GetTeletrasportSafePlanetListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTeletrasportSafePlanetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTeletrasportSafePlanetListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTeletrasportSafePlanetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeletrasportSafePlanetListRequest.Merge(m, src)
}
func (m *GetTeletrasportSafePlanetListRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTeletrasportSafePlanetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeletrasportSafePlanetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeletrasportSafePlanetListRequest proto.InternalMessageInfo

func (m *GetTeletrasportSafePlanetListRequest) GetPlanetID() uint32 {
	if m != nil {
		return m.PlanetID
	}
	return 0
}

type GetTeletrasportSafePlanetListResponse struct {
	SafePlanets []*GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList `protobuf:"bytes,1,rep,name=SafePlanets,proto3" json:"SafePlanets,omitempty"`
}

func (m *GetTeletrasportSafePlanetListResponse) Reset()         { *m = GetTeletrasportSafePlanetListResponse{} }
func (m *GetTeletrasportSafePlanetListResponse) String() string { return proto.CompactTextString(m) }
func (*GetTeletrasportSafePlanetListResponse) ProtoMessage()    {}
func (*GetTeletrasportSafePlanetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab067e92a46b1c0b, []int{1}
}
func (m *GetTeletrasportSafePlanetListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTeletrasportSafePlanetListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTeletrasportSafePlanetListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTeletrasportSafePlanetListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeletrasportSafePlanetListResponse.Merge(m, src)
}
func (m *GetTeletrasportSafePlanetListResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTeletrasportSafePlanetListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeletrasportSafePlanetListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeletrasportSafePlanetListResponse proto.InternalMessageInfo

func (m *GetTeletrasportSafePlanetListResponse) GetSafePlanets() []*GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList {
	if m != nil {
		return m.SafePlanets
	}
	return nil
}

type GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList struct {
	Planet *Planet `protobuf:"bytes,1,opt,name=Planet,proto3" json:"Planet,omitempty"`
	Price  uint32  `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) Reset() {
	*m = GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList{}
}
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) String() string {
	return proto.CompactTextString(m)
}
func (*GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) ProtoMessage() {}
func (*GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab067e92a46b1c0b, []int{1, 0}
}
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList.Merge(m, src)
}
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) XXX_Size() int {
	return m.Size()
}
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList proto.InternalMessageInfo

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) GetPlanet() *Planet {
	if m != nil {
		return m.Planet
	}
	return nil
}

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

// EndTeletrasportSafePlanet
type EndTeletrasportSafePlanetRequest struct {
	PlayerID uint32 `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlanetID uint32 `protobuf:"varint,2,opt,name=PlanetID,proto3" json:"PlanetID,omitempty"`
	Price    int32  `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *EndTeletrasportSafePlanetRequest) Reset()         { *m = EndTeletrasportSafePlanetRequest{} }
func (m *EndTeletrasportSafePlanetRequest) String() string { return proto.CompactTextString(m) }
func (*EndTeletrasportSafePlanetRequest) ProtoMessage()    {}
func (*EndTeletrasportSafePlanetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab067e92a46b1c0b, []int{2}
}
func (m *EndTeletrasportSafePlanetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EndTeletrasportSafePlanetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EndTeletrasportSafePlanetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndTeletrasportSafePlanetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndTeletrasportSafePlanetRequest.Merge(m, src)
}
func (m *EndTeletrasportSafePlanetRequest) XXX_Size() int {
	return m.Size()
}
func (m *EndTeletrasportSafePlanetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndTeletrasportSafePlanetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndTeletrasportSafePlanetRequest proto.InternalMessageInfo

func (m *EndTeletrasportSafePlanetRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *EndTeletrasportSafePlanetRequest) GetPlanetID() uint32 {
	if m != nil {
		return m.PlanetID
	}
	return 0
}

func (m *EndTeletrasportSafePlanetRequest) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type EndTeletrasportSafePlanetResponse struct {
}

func (m *EndTeletrasportSafePlanetResponse) Reset()         { *m = EndTeletrasportSafePlanetResponse{} }
func (m *EndTeletrasportSafePlanetResponse) String() string { return proto.CompactTextString(m) }
func (*EndTeletrasportSafePlanetResponse) ProtoMessage()    {}
func (*EndTeletrasportSafePlanetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab067e92a46b1c0b, []int{3}
}
func (m *EndTeletrasportSafePlanetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EndTeletrasportSafePlanetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EndTeletrasportSafePlanetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndTeletrasportSafePlanetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndTeletrasportSafePlanetResponse.Merge(m, src)
}
func (m *EndTeletrasportSafePlanetResponse) XXX_Size() int {
	return m.Size()
}
func (m *EndTeletrasportSafePlanetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndTeletrasportSafePlanetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndTeletrasportSafePlanetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTeletrasportSafePlanetListRequest)(nil), "expansion.GetTeletrasportSafePlanetListRequest")
	proto.RegisterType((*GetTeletrasportSafePlanetListResponse)(nil), "expansion.GetTeletrasportSafePlanetListResponse")
	proto.RegisterType((*GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList)(nil), "expansion.GetTeletrasportSafePlanetListResponse.TeletrasportSafePlanetList")
	proto.RegisterType((*EndTeletrasportSafePlanetRequest)(nil), "expansion.EndTeletrasportSafePlanetRequest")
	proto.RegisterType((*EndTeletrasportSafePlanetResponse)(nil), "expansion.EndTeletrasportSafePlanetResponse")
}

func init() { proto.RegisterFile("proto/expansion.proto", fileDescriptor_ab067e92a46b1c0b) }

var fileDescriptor_ab067e92a46b1c0b = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x28, 0x48, 0xcc, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x03, 0xf3, 0x85, 0x38,
	0xe1, 0x02, 0x52, 0x42, 0x10, 0x15, 0x05, 0x39, 0x89, 0x79, 0xa9, 0x25, 0x10, 0x69, 0x25, 0x27,
	0x2e, 0x15, 0xf7, 0xd4, 0x92, 0x90, 0xd4, 0x9c, 0xd4, 0x92, 0xa2, 0xc4, 0xe2, 0x82, 0xfc, 0xa2,
	0x92, 0xe0, 0xc4, 0xb4, 0xd4, 0x00, 0xb0, 0x12, 0x9f, 0xcc, 0xe2, 0x92, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x0e, 0x88, 0xa0, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x6f, 0x10, 0x9c, 0xaf, 0xf4, 0x89, 0x91, 0x4b, 0x95, 0x80, 0x21, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0xb9, 0x5c, 0xdc, 0x08, 0x99, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x6f,
	0x3d, 0x84, 0x9b, 0x89, 0x32, 0x46, 0x0f, 0x8f, 0x12, 0x64, 0xf3, 0xa5, 0xa2, 0xb8, 0xa4, 0x70,
	0x2b, 0x15, 0x52, 0xe3, 0x62, 0x83, 0xf0, 0xc0, 0x1e, 0xe2, 0x36, 0xe2, 0xd3, 0x83, 0x86, 0x0c,
	0x44, 0x34, 0x08, 0x2a, 0x2b, 0x24, 0xc2, 0xc5, 0x1a, 0x50, 0x94, 0x99, 0x9c, 0x2a, 0xc1, 0x04,
	0xf6, 0x37, 0x84, 0xa3, 0x54, 0xc0, 0xa5, 0xe0, 0x9a, 0x97, 0x82, 0xdd, 0x78, 0xd4, 0x40, 0xab,
	0x4c, 0x2d, 0x42, 0x09, 0x34, 0x30, 0x1f, 0x25, 0x40, 0x99, 0x50, 0x03, 0x14, 0x61, 0x23, 0xb3,
	0x02, 0xa3, 0x06, 0x2b, 0xcc, 0x46, 0x65, 0x2e, 0x45, 0x3c, 0x36, 0x42, 0x82, 0xc6, 0x49, 0xee,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x58, 0xf4, 0xac, 0x0b, 0x92, 0x92,
	0xd8, 0xc0, 0xd1, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x77, 0xb4, 0xae, 0x8d, 0x2e, 0x02,
	0x00, 0x00,
}

func (m *GetTeletrasportSafePlanetListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTeletrasportSafePlanetListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTeletrasportSafePlanetListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanetID != 0 {
		i = encodeVarintExpansion(dAtA, i, uint64(m.PlanetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetTeletrasportSafePlanetListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTeletrasportSafePlanetListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTeletrasportSafePlanetListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SafePlanets) > 0 {
		for iNdEx := len(m.SafePlanets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SafePlanets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExpansion(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintExpansion(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x10
	}
	if m.Planet != nil {
		{
			size, err := m.Planet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExpansion(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EndTeletrasportSafePlanetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndTeletrasportSafePlanetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EndTeletrasportSafePlanetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintExpansion(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x18
	}
	if m.PlanetID != 0 {
		i = encodeVarintExpansion(dAtA, i, uint64(m.PlanetID))
		i--
		dAtA[i] = 0x10
	}
	if m.PlayerID != 0 {
		i = encodeVarintExpansion(dAtA, i, uint64(m.PlayerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EndTeletrasportSafePlanetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndTeletrasportSafePlanetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EndTeletrasportSafePlanetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintExpansion(dAtA []byte, offset int, v uint64) int {
	offset -= sovExpansion(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetTeletrasportSafePlanetListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanetID != 0 {
		n += 1 + sovExpansion(uint64(m.PlanetID))
	}
	return n
}

func (m *GetTeletrasportSafePlanetListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SafePlanets) > 0 {
		for _, e := range m.SafePlanets {
			l = e.Size()
			n += 1 + l + sovExpansion(uint64(l))
		}
	}
	return n
}

func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Planet != nil {
		l = m.Planet.Size()
		n += 1 + l + sovExpansion(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovExpansion(uint64(m.Price))
	}
	return n
}

func (m *EndTeletrasportSafePlanetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlayerID != 0 {
		n += 1 + sovExpansion(uint64(m.PlayerID))
	}
	if m.PlanetID != 0 {
		n += 1 + sovExpansion(uint64(m.PlanetID))
	}
	if m.Price != 0 {
		n += 1 + sovExpansion(uint64(m.Price))
	}
	return n
}

func (m *EndTeletrasportSafePlanetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovExpansion(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExpansion(x uint64) (n int) {
	return sovExpansion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetTeletrasportSafePlanetListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpansion
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
			return fmt.Errorf("proto: GetTeletrasportSafePlanetListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTeletrasportSafePlanetListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanetID", wireType)
			}
			m.PlanetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanetID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpansion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpansion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExpansion
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
func (m *GetTeletrasportSafePlanetListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpansion
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
			return fmt.Errorf("proto: GetTeletrasportSafePlanetListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTeletrasportSafePlanetListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafePlanets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
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
				return ErrInvalidLengthExpansion
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpansion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SafePlanets = append(m.SafePlanets, &GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList{})
			if err := m.SafePlanets[len(m.SafePlanets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExpansion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpansion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExpansion
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
func (m *GetTeletrasportSafePlanetListResponse_TeletrasportSafePlanetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpansion
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
			return fmt.Errorf("proto: TeletrasportSafePlanetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TeletrasportSafePlanetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Planet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
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
				return ErrInvalidLengthExpansion
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExpansion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Planet == nil {
				m.Planet = &Planet{}
			}
			if err := m.Planet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpansion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpansion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExpansion
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
func (m *EndTeletrasportSafePlanetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpansion
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
			return fmt.Errorf("proto: EndTeletrasportSafePlanetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndTeletrasportSafePlanetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
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
				return fmt.Errorf("proto: wrong wireType = %d for field PlanetID", wireType)
			}
			m.PlanetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanetID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpansion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpansion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpansion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExpansion
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
func (m *EndTeletrasportSafePlanetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpansion
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
			return fmt.Errorf("proto: EndTeletrasportSafePlanetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndTeletrasportSafePlanetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipExpansion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpansion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExpansion
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
func skipExpansion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExpansion
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
					return 0, ErrIntOverflowExpansion
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
					return 0, ErrIntOverflowExpansion
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
				return 0, ErrInvalidLengthExpansion
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExpansion
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExpansion
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExpansion        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExpansion          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExpansion = fmt.Errorf("proto: unexpected end of group")
)